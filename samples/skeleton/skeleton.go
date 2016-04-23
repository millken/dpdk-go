package main

import (
	"log"
	"os"
	"unsafe"

	"github.com/feiskyer/dpdk-go"
)

// #cgo CFLAGS: -m64 -pthread -O3 -march=native -I/usr/local/include/dpdk
// #cgo LDFLAGS: -L/usr/local/lib -Wl,--whole-archive -ldpdk -lz -Wl,--start-group -lrt -lm -ldl -lfuse -Wl,--end-group -Wl,--no-whole-archive
import "C"

const (
	RX_RING_SIZE    = 128
	TX_RING_SIZE    = 512
	NUM_MBUFS       = 8191
	MBUF_CACHE_SIZE = 250
	BURST_SIZE      = 32
)

type Params struct {
	portId uint
}

/*
 * Receive packets on a port and forward them on the paired
 * port. The mapping is 0 -> 1, 1 -> 0, 2 -> 3, 3 -> 2, etc.
 */
func loop(arg unsafe.Pointer) int {
	params := (*Params)(arg)
	bufs := dpdk.GetCArray(uint(BURST_SIZE))

	for {
		// Get burst of RX packets
		nb_rx := dpdk.RteEthRxBurst(params.portId, 0, bufs, BURST_SIZE)
		// Send burst of TX packets
		nb_tx := dpdk.RteEthTxBurst(params.portId^1, 0, bufs, nb_rx)
		// Free any unsent packets
		if nb_tx < nb_rx {
			for buf := nb_tx; buf < nb_rx; buf++ {
				dpdk.RtePktMbufFree(((*dpdk.RteMbuf)(dpdk.ElemFromCArray(bufs, buf))))
			}
		}
	}
	return 0
}

func main() {
	ret := dpdk.RteEalInit(os.Args)
	if ret < 0 {
		log.Fatalln("Failed to init EAL: ", dpdk.StrError(ret))
	}

	nbPorts := dpdk.RteEthDevCount()
	log.Println("Got dev count: ", nbPorts)
	if nbPorts < 2 || (nbPorts&1) != 0 {
		log.Fatalln("Error: number of ports must be even")
	}

	mbufPool := dpdk.RtePktMbufPoolCreate(
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
		ret := dpdk.RteEthDevConfigure(portId, 1, 1, &dpdk.RteEthConf{})
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

		// Start the Ethernet port
		ret = dpdk.RteEthDevStart(portId)
		if ret < 0 {
			log.Fatalln("Failed to start eth dev")
		}

		// Enable RX in promiscuous mode for the Ethernet device
		dpdk.RteEthPromiscuousEnable(portId)

		// Launch loop on lcore
		dpdk.RteEalRemoteLaunch(loop, unsafe.Pointer(&Params{portId: portId}), portId)
	}

	dpdk.RteEalMpWaitLCore()
}
