package dpdk

/*
#cgo CFLAGS: -m64 -pthread -O3 -march=native -I/usr/local/include/dpdk
#cgo LDFLAGS: -L/usr/local/lib -ldpdk -lz -lrt -lm -ldl -lfuse

#include <rte_config.h>
#include <rte_mbuf.h>
*/
import "C"
import "unsafe"

const (
	PKT_RX_VLAN_PKT                     = uint64(C.PKT_RX_VLAN_PKT)
	PKT_RX_RSS_HASH                     = uint64(C.PKT_RX_RSS_HASH)
	PKT_RX_FDIR                         = uint64(C.PKT_RX_FDIR)
	PKT_RX_L4_CKSUM_BAD                 = uint64(C.PKT_RX_L4_CKSUM_BAD)
	PKT_RX_IP_CKSUM_BAD                 = uint64(C.PKT_RX_IP_CKSUM_BAD)
	PKT_RX_EIP_CKSUM_BAD                = uint64(C.PKT_RX_EIP_CKSUM_BAD)
	PKT_RX_OVERSIZE                     = uint64(C.PKT_RX_OVERSIZE)
	PKT_RX_HBUF_OVERFLOW                = uint64(C.PKT_RX_HBUF_OVERFLOW)
	PKT_RX_RECIP_ERR                    = uint64(C.PKT_RX_RECIP_ERR)
	PKT_RX_MAC_ERR                      = uint64(C.PKT_RX_MAC_ERR)
	PKT_RX_IEEE1588_PTP                 = uint64(C.PKT_RX_IEEE1588_PTP)
	PKT_RX_IEEE1588_TMST                = uint64(C.PKT_RX_IEEE1588_TMST)
	PKT_RX_FDIR_ID                      = uint64(C.PKT_RX_FDIR_ID)
	PKT_RX_FDIR_FLX                     = uint64(C.PKT_RX_FDIR_FLX)
	PKT_RX_QINQ_PKT                     = uint64(C.PKT_RX_QINQ_PKT)
	PKT_TX_QINQ_PKT                     = uint64(C.PKT_TX_QINQ_PKT)
	PKT_TX_TCP_SEG                      = uint64(C.PKT_TX_TCP_SEG)
	PKT_TX_IEEE1588_TMST                = uint64(C.PKT_TX_IEEE1588_TMST)
	PKT_TX_L4_NO_CKSUM                  = uint64(C.PKT_TX_L4_NO_CKSUM)
	PKT_TX_TCP_CKSUM                    = uint64(C.PKT_TX_TCP_CKSUM)
	PKT_TX_SCTP_CKSUM                   = uint64(C.PKT_TX_SCTP_CKSUM)
	PKT_TX_UDP_CKSUM                    = uint64(C.PKT_TX_UDP_CKSUM)
	PKT_TX_L4_MASK                      = uint64(C.PKT_TX_L4_MASK)
	PKT_TX_IP_CKSUM                     = uint64(C.PKT_TX_IP_CKSUM)
	PKT_TX_IPV4                         = uint64(C.PKT_TX_IPV4)
	PKT_TX_IPV6                         = uint64(C.PKT_TX_IPV6)
	PKT_TX_VLAN_PKT                     = uint64(C.PKT_TX_VLAN_PKT)
	PKT_TX_OUTER_IP_CKSUM               = uint64(C.PKT_TX_OUTER_IP_CKSUM)
	PKT_TX_OUTER_IPV4                   = uint64(C.PKT_TX_OUTER_IPV4)
	PKT_TX_OUTER_IPV6                   = uint64(C.PKT_TX_OUTER_IPV6)
	__RESERVED                          = uint64(C.__RESERVED)
	IND_ATTACHED_MBUF                   = uint64(C.IND_ATTACHED_MBUF)
	RTE_PTYPE_L2_ETHER                  = uint64(C.RTE_PTYPE_L2_ETHER)
	RTE_PTYPE_L2_ETHER_TIMESYNC         = uint64(C.RTE_PTYPE_L2_ETHER_TIMESYNC)
	RTE_PTYPE_L2_ETHER_ARP              = uint64(C.RTE_PTYPE_L2_ETHER_ARP)
	RTE_PTYPE_L2_ETHER_LLDP             = uint64(C.RTE_PTYPE_L2_ETHER_LLDP)
	RTE_PTYPE_L2_MASK                   = uint64(C.RTE_PTYPE_L2_MASK)
	RTE_PTYPE_L3_IPV4                   = uint64(C.RTE_PTYPE_L3_IPV4)
	RTE_PTYPE_L3_IPV4_EXT               = uint64(C.RTE_PTYPE_L3_IPV4_EXT)
	RTE_PTYPE_L3_IPV6                   = uint64(C.RTE_PTYPE_L3_IPV6)
	RTE_PTYPE_L3_IPV4_EXT_UNKNOWN       = uint64(C.RTE_PTYPE_L3_IPV4_EXT_UNKNOWN)
	RTE_PTYPE_L3_IPV6_EXT               = uint64(C.RTE_PTYPE_L3_IPV6_EXT)
	RTE_PTYPE_L3_IPV6_EXT_UNKNOWN       = uint64(C.RTE_PTYPE_L3_IPV6_EXT_UNKNOWN)
	RTE_PTYPE_L3_MASK                   = uint64(C.RTE_PTYPE_L3_MASK)
	RTE_PTYPE_L4_TCP                    = uint64(C.RTE_PTYPE_L4_TCP)
	RTE_PTYPE_L4_UDP                    = uint64(C.RTE_PTYPE_L4_UDP)
	RTE_PTYPE_L4_FRAG                   = uint64(C.RTE_PTYPE_L4_FRAG)
	RTE_PTYPE_L4_SCTP                   = uint64(C.RTE_PTYPE_L4_SCTP)
	RTE_PTYPE_L4_ICMP                   = uint64(C.RTE_PTYPE_L4_ICMP)
	RTE_PTYPE_L4_NONFRAG                = uint64(C.RTE_PTYPE_L4_NONFRAG)
	RTE_PTYPE_L4_MASK                   = uint64(C.RTE_PTYPE_L4_MASK)
	RTE_PTYPE_TUNNEL_IP                 = uint64(C.RTE_PTYPE_TUNNEL_IP)
	RTE_PTYPE_TUNNEL_GRE                = uint64(C.RTE_PTYPE_TUNNEL_GRE)
	RTE_PTYPE_TUNNEL_VXLAN              = uint64(C.RTE_PTYPE_TUNNEL_VXLAN)
	RTE_PTYPE_TUNNEL_NVGRE              = uint64(C.RTE_PTYPE_TUNNEL_NVGRE)
	RTE_PTYPE_TUNNEL_GENEVE             = uint64(C.RTE_PTYPE_TUNNEL_GENEVE)
	RTE_PTYPE_TUNNEL_GRENAT             = uint64(C.RTE_PTYPE_TUNNEL_GRENAT)
	RTE_PTYPE_TUNNEL_MASK               = uint64(C.RTE_PTYPE_TUNNEL_MASK)
	RTE_PTYPE_INNER_L2_ETHER            = uint64(C.RTE_PTYPE_INNER_L2_ETHER)
	RTE_PTYPE_INNER_L2_ETHER_VLAN       = uint64(C.RTE_PTYPE_INNER_L2_ETHER_VLAN)
	RTE_PTYPE_INNER_L2_MASK             = uint64(C.RTE_PTYPE_INNER_L2_MASK)
	RTE_PTYPE_INNER_L3_IPV4             = uint64(C.RTE_PTYPE_INNER_L3_IPV4)
	RTE_PTYPE_INNER_L3_IPV4_EXT         = uint64(C.RTE_PTYPE_INNER_L3_IPV4_EXT)
	RTE_PTYPE_INNER_L3_IPV6             = uint64(C.RTE_PTYPE_INNER_L3_IPV6)
	RTE_PTYPE_INNER_L3_IPV4_EXT_UNKNOWN = uint64(C.RTE_PTYPE_INNER_L3_IPV4_EXT_UNKNOWN)
	RTE_PTYPE_INNER_L3_IPV6_EXT         = uint64(C.RTE_PTYPE_INNER_L3_IPV6_EXT)
	RTE_PTYPE_INNER_L3_IPV6_EXT_UNKNOWN = uint64(C.RTE_PTYPE_INNER_L3_IPV6_EXT_UNKNOWN)
	RTE_PTYPE_INNER_L3_MASK             = uint64(C.RTE_PTYPE_INNER_L3_MASK)
	RTE_PTYPE_INNER_L4_TCP              = uint64(C.RTE_PTYPE_INNER_L4_TCP)
	RTE_PTYPE_INNER_L4_UDP              = uint64(C.RTE_PTYPE_INNER_L4_UDP)
	RTE_PTYPE_INNER_L4_FRAG             = uint64(C.RTE_PTYPE_INNER_L4_FRAG)
	RTE_PTYPE_INNER_L4_SCTP             = uint64(C.RTE_PTYPE_INNER_L4_SCTP)
	RTE_PTYPE_INNER_L4_ICMP             = uint64(C.RTE_PTYPE_INNER_L4_ICMP)
	RTE_PTYPE_INNER_L4_NONFRAG          = uint64(C.RTE_PTYPE_INNER_L4_NONFRAG)
	RTE_PTYPE_INNER_L4_MASK             = uint64(C.RTE_PTYPE_INNER_L4_MASK)
	RTE_MBUF_PRIV_ALIGN                 = uint64(C.RTE_MBUF_PRIV_ALIGN)
	RTE_MBUF_DEFAULT_DATAROOM           = uint64(C.RTE_MBUF_DEFAULT_DATAROOM)
	RTE_MBUF_DEFAULT_BUF_SIZE           = uint(C.RTE_MBUF_DEFAULT_BUF_SIZE)
)

// RteMbuf is generic rte_mbuf, containing a packet mbuf
type RteMbuf C.struct_rte_mbuf
type RtePktMbufPoolPrivate C.struct_rte_pktmbuf_pool_private

func RtePktMbufPoolCreate(name string, n, cache_size, priv_size,
	data_room_size uint, socket_id int) *RteMemPool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return (*RteMemPool)(C.rte_pktmbuf_pool_create(cName,
		C.unsigned(n), C.unsigned(cache_size), C.uint16_t(priv_size),
		C.uint16_t(data_room_size), C.int(socket_id)))
}

// RtePktMbufAlloc allocates a new mbuf from a mempool
func RtePktMbufAlloc(mp *RteMemPool) *RteMbuf {
	return (*RteMbuf)(unsafe.Pointer(C.rte_pktmbuf_alloc((*C.struct_rte_mempool)(unsafe.Pointer(mp)))))
}

// RtePktMbufFree frees a packet mbuf back into its original mempool
func RtePktMbufFree(buf *RteMbuf) {
	C.rte_pktmbuf_free((*C.struct_rte_mbuf)(unsafe.Pointer(buf)))
}

// RtePktMbufMtoD returns a pointer pointing to the start of the data in mbuf
func RtePktMbufMtoD(buf *RteMbuf) unsafe.Pointer {
	mbuf := (*C.struct_rte_mbuf)(unsafe.Pointer(buf))
	return unsafe.Pointer(uintptr(unsafe.Pointer((*C.char)(mbuf.buf_addr))) + uintptr(mbuf.data_off))
}

// RtePktMbufDataLen returns the length of the segment.
func RtePktMbufDataLen(buf *RteMbuf) uint16 {
	mbuf := (*C.struct_rte_mbuf)(unsafe.Pointer(buf))
	return uint16(mbuf.data_len)
}
