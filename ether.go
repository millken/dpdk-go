package dpdk

/*
#cgo CFLAGS: -m64 -pthread -O3 -march=native -I/usr/local/include/dpdk
#cgo LDFLAGS: -L/usr/lib -ldpdk -lz -lrt -lm -ldl -lfuse

#include <rte_config.h>
#include <rte_ether.h>
*/
import "C"
import "unsafe"

const (
	ETHER_ADDR_LEN            = int(C.ETHER_ADDR_LEN)
	ETHER_TYPE_LEN            = int(C.ETHER_TYPE_LEN)
	ETHER_CRC_LEN             = int(C.ETHER_CRC_LEN)
	ETHER_HDR_LEN             = int(C.ETHER_HDR_LEN)
	ETHER_MIN_LEN             = int(C.ETHER_MIN_LEN)
	ETHER_MAX_LEN             = int(C.ETHER_MAX_LEN)
	ETHER_MTU                 = int(C.ETHER_MTU)
	ETHER_MAX_VLAN_FRAME_LEN  = int(C.ETHER_MAX_VLAN_FRAME_LEN)
	ETHER_MAX_JUMBO_FRAME_LEN = int(C.ETHER_MAX_JUMBO_FRAME_LEN)
	ETHER_MAX_VLAN_ID         = int(C.ETHER_MAX_VLAN_ID)
	ETHER_MIN_MTU             = int(C.ETHER_MIN_MTU)
	ETHER_LOCAL_ADMIN_ADDR    = int(C.ETHER_LOCAL_ADMIN_ADDR)
	ETHER_GROUP_ADDR          = int(C.ETHER_GROUP_ADDR)
	ETHER_TYPE_IPv4           = int(C.ETHER_TYPE_IPv4)
	ETHER_TYPE_IPv6           = int(C.ETHER_TYPE_IPv6)
	ETHER_TYPE_ARP            = int(C.ETHER_TYPE_ARP)
	ETHER_TYPE_RARP           = int(C.ETHER_TYPE_RARP)
	ETHER_TYPE_VLAN           = int(C.ETHER_TYPE_VLAN)
	ETHER_TYPE_1588           = int(C.ETHER_TYPE_1588)
	ETHER_TYPE_SLOW           = int(C.ETHER_TYPE_SLOW)
	ETHER_TYPE_TEB            = int(C.ETHER_TYPE_TEB)
)

type EtherAddr C.struct_ether_addr
type EtherHdr C.struct_ether_hdr

func (a *EtherAddr) IsSameEtherAddr(b *EtherAddr) bool {
	return int(C.is_same_ether_addr((*C.struct_ether_addr)(a), (*C.struct_ether_addr)(b))) == 1
}

func (addr *EtherAddr) IsZeroEtherAddr() bool {
	return int(C.is_zero_ether_addr((*C.struct_ether_addr)(addr))) == 1
}

func (addr *EtherAddr) IsUnicastEtherAddr() bool {
	return int(C.is_unicast_ether_addr((*C.struct_ether_addr)(addr))) == 1
}

func (addr *EtherAddr) IsMulticastEtherAddr() bool {
	return int(C.is_multicast_ether_addr((*C.struct_ether_addr)(addr))) == 1
}

func (addr *EtherAddr) IsBroadcastEtherAddr() bool {
	return int(C.is_broadcast_ether_addr((*C.struct_ether_addr)(addr))) == 1
}

func (addr *EtherAddr) IsUniversalEtherAddr() bool {
	return int(C.is_universal_ether_addr((*C.struct_ether_addr)(addr))) == 1
}

func (addr *EtherAddr) IsLocalAdminEtherAddr() bool {
	return int(C.is_local_admin_ether_addr((*C.struct_ether_addr)(addr))) == 1
}

func EthRandomAddr() uint {
	var addr uint
	C.eth_random_addr((*C.uint8_t)(unsafe.Pointer((&addr))))
	return addr
}

func (addr *EtherAddr) Copy() *EtherAddr {
	var b EtherAddr
	C.ether_addr_copy((*C.struct_ether_addr)(addr), (*C.struct_ether_addr)(&b))
	return &b
}

func (addr *EtherAddr) String() string {
	var out *C.char
	C.ether_format_addr(out, C.uint16_t(6),
		(*C.struct_ether_addr)(unsafe.Pointer(addr)))
	return C.GoString(out)
}
