package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"

	"github.com/feiskyer/dpdk-go"
)

/*
#cgo CFLAGS: -m64 -pthread -O3 -march=native -I/usr/local/include/dpdk
#cgo LDFLAGS: -L/usr/local/lib -Wl,--whole-archive -ldpdk -lz -Wl,--start-group -lrt -lm -ldl -lfuse -Wl,--end-group -Wl,--no-whole-archive
#include <rte_common.h>
#include <rte_config.h>
#include <rte_ethdev.h>
#include <rte_mbuf.h>

#include <sys/ioctl.h>
#include <sys/socket.h>
#include <linux/if.h>
#include <linux/if_tun.h>

#define IFREQ_SIZE sizeof(struct ifreq)

const struct rte_eth_conf port_conf = {
	.rxmode = {
		.header_split = 0,      // Header Split disabled
		.hw_ip_checksum = 0,    // IP checksum offload disabled
		.hw_vlan_filter = 0,    // VLAN filtering disabled
		.jumbo_frame = 0,       // Jumbo Frame Support disabled
		.hw_strip_crc = 0,      // CRC stripped by hardware
	},
	.txmode = {
		.mq_mode = ETH_MQ_TX_NONE,
	},
};
*/
import "C"

const (
	RX_RING_SIZE    = 128
	TX_RING_SIZE    = 512
	NUM_MBUFS       = 8192
	MBUF_CACHE_SIZE = 32
	BURST_SIZE      = 32

	flagTruncated = C.TUN_PKT_STRIP
	iffTun        = C.IFF_TUN
	iffTap        = C.IFF_TAP
	iffOneQueue   = C.IFF_ONE_QUEUE
	iffNoPI       = C.IFF_NO_PI
)

var mbufPool *dpdk.RteMemPool

type Params struct {
	InputMask  uint
	OutputMask uint
}

type ifReq struct {
	Name  [C.IFNAMSIZ]byte
	Flags uint16
	pad   [C.IFREQ_SIZE - C.IFNAMSIZ - 2]byte
}

func createTap(devname string) (*os.File, error) {
	file, err := os.OpenFile("/dev/net/tun", os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	var req ifReq
	copy(req.Name[:15], []byte(devname))
	req.Flags = iffOneQueue | iffTap
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, file.Fd(), uintptr(syscall.TUNSETIFF), uintptr(unsafe.Pointer(&req)))
	if int(errno) != 0 {
		file.Close()
		return nil, errno
	}

	return file, nil
}

func loop(arg unsafe.Pointer) int {
	var bufs [BURST_SIZE]*dpdk.RteMbuf

	params := (*Params)(arg)
	lcoreID := dpdk.RteLcoreID()
	tapName := fmt.Sprintf("tap_dpdk_%.2d", lcoreID)
	tapFile, err := createTap(tapName)
	if err != nil {
		log.Fatalln("CreateTap failed:", err)
	}
	defer tapFile.Close()

	if ((uint(1) << lcoreID) & params.InputMask) != 0 {
		fmt.Println("Lcore", lcoreID, "is reading from port 0 and writing to", tapName)

		for {
			// Get burst of RX packets
			numRx := dpdk.RteEthRxBurst(0, 0, (*unsafe.Pointer)(unsafe.Pointer(&bufs[0])), BURST_SIZE)
			// Print the buf content, just for debug
			for i := uint(0); i < numRx; i++ {
				// m := (*C.struct_rte_mbuf)(unsafe.Pointer(bufs[i]))
				// data := C.get_pkt_data(m)
				data := dpdk.RtePktMbufMtoD(bufs[i])
				// dataLen := int(C.get_pkt_data_len(m))
				dataLen := dpdk.RtePktMbufDataLen(bufs[i])

				wnb, err := tapFile.Write((*[1 << 30]byte)(data)[:dataLen:dataLen])
				if err != nil {
					log.Println("Write data failed: ", err)
				}

				wn := uint16(wnb)
				if wn < dataLen {
					for i := wn; i < dataLen; i++ {
						dpdk.RtePktMbufFree(bufs[i])
					}
				}
			}
		}
	} else if ((uint(1) << lcoreID) & params.OutputMask) != 0 {
		fmt.Println("Lcore", lcoreID, "is reading from", tapName, "and writing to port 0")

		for {
			m := dpdk.RtePktMbufAlloc(mbufPool)
			data := dpdk.RtePktMbufMtoD(m)
			if m == nil {
				continue
			}

			bytes := (*[1 << 30]byte)(data)[:2048:2048]
			rnb, err := tapFile.Read(bytes)
			if err != nil {
				log.Println("Read from tap failed:", err)
				continue
			}

			mbuf := (*C.struct_rte_mbuf)(unsafe.Pointer(m))
			mbuf.nb_segs = 1
			mbuf.next = nil
			mbuf.pkt_len = C.uint32_t(rnb)
			mbuf.data_len = C.uint16_t(rnb)

			numTx := dpdk.RteEthTxBurst(0, 0, (*unsafe.Pointer)(unsafe.Pointer(&m)), 1)
			if numTx < 1 {
				dpdk.RtePktMbufFree(m)
			}
		}
	} else {
		log.Println("Lcore", lcoreID, "has nothing to do")
	}

	return 0
}

func main() {
	inputMask := flag.Uint("input-core-mask", 1, "input core mask")
	outputMask := flag.Uint("output-core-mask", 2, "output core mask")

	flag.Parse()

	ret := dpdk.RteEalInit(os.Args)
	if ret < 0 {
		log.Fatalln("Failed to init EAL: ", dpdk.StrError(ret))
	}

	nbPorts := dpdk.RteEthDevCount()
	log.Println("Got dev count: ", nbPorts)
	if nbPorts < 2 {
		log.Fatalln("Error: number of ports must be > 2")
	}

	mbufPool = dpdk.RtePktMbufPoolCreate(
		"MBUF_POOL",
		nbPorts*NUM_MBUFS,
		MBUF_CACHE_SIZE,
		0,
		dpdk.RTE_MBUF_DEFAULT_BUF_SIZE,
		dpdk.RteSocketID(),
	)
	if mbufPool == nil {
		log.Fatalln("Cannot create mbuf pool")
	}

	// Initialize all ports
	for portId := uint(0); portId < nbPorts; portId++ {
		// Configure the Ethernet device
		ret := dpdk.RteEthDevConfigure(portId, 1, 1, (*dpdk.RteEthConf)(unsafe.Pointer(&C.port_conf)))
		if ret < 0 {
			log.Fatalln("Failed to setup eth dev: ", dpdk.StrError(ret))
		}

		// Allocate and set up 1 RX queue per Ethernet port
		ret = dpdk.RteEthRxQueueSetup(
			portId,
			0,
			RX_RING_SIZE,
			dpdk.RteEthDevSocketID(portId),
			nil,
			mbufPool,
		)
		if ret < 0 {
			log.Fatalln("Failed to setup rx queue")
		}

		// Allocate and set up 1 TX queue per Ethernet port
		ret = dpdk.RteEthTxQueueSetup(
			portId,
			0,
			TX_RING_SIZE,
			dpdk.RteEthDevSocketID(portId),
			nil,
		)
		if ret < 0 {
			log.Fatalln("Failed to setup tx queue")
		}

		macAddr := dpdk.RteEthMacAddr(portId)
		log.Println("Port ", portId, " 's mac address is ", macAddr)

		// Start the Ethernet port
		ret = dpdk.RteEthDevStart(portId)
		if ret < 0 {
			log.Fatalln("Failed to start eth dev")
		}

		// Enable RX in promiscuous mode for the Ethernet device
		dpdk.RteEthPromiscuousEnable(portId)

		// Launch loop on lcore
		// dpdk.RteEalRemoteLaunch(loop, unsafe.Pointer(&Params{portId: portId}), portId)
	}

	dpdk.RteEalMpRemoteLaunch(loop, unsafe.Pointer(&Params{
		InputMask:  *inputMask,
		OutputMask: *outputMask,
	}), dpdk.CALL_MASTER)
	if dpdk.RteEalWaitAllSlave() < 0 {
		log.Fatalln("RteEalWaitAllSlave failed.")
	}
}
