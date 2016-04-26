package dpdk

/*
#cgo CFLAGS: -m64 -pthread -O3 -march=native -I/usr/local/include/dpdk
#cgo LDFLAGS: -L/usr/local/lib -ldpdk -lz -lrt -lm -ldl -lfuse

#include <rte_config.h>
#include <rte_ethdev.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

/* Macros */
const (
	ETH_LINK_SPEED_AUTONEG          = int(C.ETH_LINK_SPEED_AUTONEG)
	ETH_LINK_SPEED_10               = int(C.ETH_LINK_SPEED_10)
	ETH_LINK_SPEED_100              = int(C.ETH_LINK_SPEED_100)
	ETH_LINK_SPEED_1000             = int(C.ETH_LINK_SPEED_1000)
	ETH_LINK_SPEED_10000            = int(C.ETH_LINK_SPEED_10000)
	ETH_LINK_SPEED_10G              = int(C.ETH_LINK_SPEED_10G)
	ETH_LINK_SPEED_20G              = int(C.ETH_LINK_SPEED_20G)
	ETH_LINK_SPEED_40G              = int(C.ETH_LINK_SPEED_40G)
	ETH_LINK_AUTONEG_DUPLEX         = int(C.ETH_LINK_AUTONEG_DUPLEX)
	ETH_LINK_HALF_DUPLEX            = int(C.ETH_LINK_HALF_DUPLEX)
	ETH_LINK_FULL_DUPLEX            = int(C.ETH_LINK_FULL_DUPLEX)
	ETH_MQ_RX_RSS_FLAG              = int(C.ETH_MQ_RX_RSS_FLAG)
	ETH_RSS                         = int(C.ETH_RSS)
	ETH_DCB_NONE                    = int(C.ETH_DCB_NONE)
	ETH_RSS_SCTP                    = int(C.ETH_RSS_SCTP)
	ETH_VMDQ_MAX_VLAN_FILTERS       = int(C.ETH_VMDQ_MAX_VLAN_FILTERS)
	ETH_DCB_NUM_USER_PRIORITIES     = int(C.ETH_DCB_NUM_USER_PRIORITIES)
	ETH_VMDQ_DCB_NUM_QUEUES         = int(C.ETH_VMDQ_DCB_NUM_QUEUES)
	ETH_DCB_NUM_QUEUES              = int(C.ETH_DCB_NUM_QUEUES)
	ETH_DCB_PG_SUPPORT              = int(C.ETH_DCB_PG_SUPPORT)
	ETH_DCB_PFC_SUPPORT             = int(C.ETH_DCB_PFC_SUPPORT)
	ETH_VLAN_STRIP_OFFLOAD          = int(C.ETH_VLAN_STRIP_OFFLOAD)
	ETH_VLAN_FILTER_OFFLOAD         = int(C.ETH_VLAN_FILTER_OFFLOAD)
	ETH_VLAN_EXTEND_OFFLOAD         = int(C.ETH_VLAN_EXTEND_OFFLOAD)
	ETH_VLAN_STRIP_MASK             = int(C.ETH_VLAN_STRIP_MASK)
	ETH_VLAN_FILTER_MASK            = int(C.ETH_VLAN_FILTER_MASK)
	ETH_VLAN_EXTEND_MASK            = int(C.ETH_VLAN_EXTEND_MASK)
	ETH_VLAN_ID_MAX                 = int(C.ETH_VLAN_ID_MAX)
	ETH_NUM_RECEIVE_MAC_ADDR        = int(C.ETH_NUM_RECEIVE_MAC_ADDR)
	ETH_VMDQ_NUM_UC_HASH_ARRAY      = int(C.ETH_VMDQ_NUM_UC_HASH_ARRAY)
	ETH_VMDQ_ACCEPT_UNTAG           = int(C.ETH_VMDQ_ACCEPT_UNTAG)
	ETH_VMDQ_ACCEPT_HASH_MC         = int(C.ETH_VMDQ_ACCEPT_HASH_MC)
	ETH_VMDQ_ACCEPT_HASH_UC         = int(C.ETH_VMDQ_ACCEPT_HASH_UC)
	ETH_VMDQ_ACCEPT_BROADCAST       = int(C.ETH_VMDQ_ACCEPT_BROADCAST)
	ETH_VMDQ_ACCEPT_MULTICAST       = int(C.ETH_VMDQ_ACCEPT_MULTICAST)
	ETH_MIRROR_MAX_VLANS            = int(C.ETH_MIRROR_MAX_VLANS)
	ETH_MIRROR_VIRTUAL_POOL_UP      = int(C.ETH_MIRROR_VIRTUAL_POOL_UP)
	ETH_MIRROR_UPLINK_PORT          = int(C.ETH_MIRROR_UPLINK_PORT)
	ETH_MIRROR_DOWNLINK_PORT        = int(C.ETH_MIRROR_DOWNLINK_PORT)
	ETH_MIRROR_VLAN                 = int(C.ETH_MIRROR_VLAN)
	ETH_MIRROR_VIRTUAL_POOL_DOWN    = int(C.ETH_MIRROR_VIRTUAL_POOL_DOWN)
	ETH_TXQ_FLAGS_NOMULTSEGS        = int(C.ETH_TXQ_FLAGS_NOMULTSEGS)
	ETH_TXQ_FLAGS_NOREFCOUNT        = int(C.ETH_TXQ_FLAGS_NOREFCOUNT)
	ETH_TXQ_FLAGS_NOMULTMEMP        = int(C.ETH_TXQ_FLAGS_NOMULTMEMP)
	ETH_TXQ_FLAGS_NOVLANOFFL        = int(C.ETH_TXQ_FLAGS_NOVLANOFFL)
	ETH_TXQ_FLAGS_NOXSUMSCTP        = int(C.ETH_TXQ_FLAGS_NOXSUMSCTP)
	ETH_TXQ_FLAGS_NOXSUMUDP         = int(C.ETH_TXQ_FLAGS_NOXSUMUDP)
	ETH_TXQ_FLAGS_NOXSUMTCP         = int(C.ETH_TXQ_FLAGS_NOXSUMTCP)
	DEV_RX_OFFLOAD_VLAN_STRIP       = int(C.DEV_RX_OFFLOAD_VLAN_STRIP)
	DEV_TX_OFFLOAD_VLAN_INSERT      = int(C.DEV_TX_OFFLOAD_VLAN_INSERT)
	DEV_TX_OFFLOAD_OUTER_IPV4_CKSUM = int(C.DEV_TX_OFFLOAD_OUTER_IPV4_CKSUM)
	RTE_ETH_XSTATS_NAME_SIZE        = int(C.RTE_ETH_XSTATS_NAME_SIZE)
	RTE_ETH_QUEUE_STATE_STOPPED     = int(C.RTE_ETH_QUEUE_STATE_STOPPED)
	RTE_ETH_DEV_DETACHABLE          = int(C.RTE_ETH_DEV_DETACHABLE)
	RTE_ETH_DEV_INTR_LSC            = int(C.RTE_ETH_DEV_INTR_LSC)
)

/* enum rte_eth_rx_mq_mode */
const (
	ETH_MQ_RX_NONE         = int(C.ETH_MQ_RX_NONE)
	ETH_MQ_RX_RSS          = int(C.ETH_MQ_RX_RSS)
	ETH_MQ_RX_DCB          = int(C.ETH_MQ_RX_DCB)
	ETH_MQ_RX_DCB_RSS      = int(C.ETH_MQ_RX_DCB_RSS)
	ETH_MQ_RX_VMDQ_ONLY    = int(C.ETH_MQ_RX_VMDQ_ONLY)
	ETH_MQ_RX_VMDQ_RSS     = int(C.ETH_MQ_RX_VMDQ_RSS)
	ETH_MQ_RX_VMDQ_DCB     = int(C.ETH_MQ_RX_VMDQ_DCB)
	ETH_MQ_RX_VMDQ_DCB_RSS = int(C.ETH_MQ_RX_VMDQ_DCB_RSS)
)

/* enum rte_eth_tx_mq_mode */
const (
	ETH_MQ_TX_NONE      = int(C.ETH_MQ_TX_NONE)
	ETH_MQ_TX_DCB       = int(C.ETH_MQ_TX_DCB)
	ETH_MQ_TX_VMDQ_DCB  = int(C.ETH_MQ_TX_VMDQ_DCB)
	ETH_MQ_TX_VMDQ_ONLY = int(C.ETH_MQ_TX_VMDQ_ONLY)
)

/* enum rte_eth_nb_tcs */
const (
	ETH_4_TCS = int(C.ETH_4_TCS)
	ETH_8_TCS = int(C.ETH_8_TCS)
)

/* enum rte_eth_nb_pools */
const (
	ETH_8_POOLS  = int(C.ETH_8_POOLS)
	ETH_16_POOLS = int(C.ETH_16_POOLS)
	ETH_32_POOLS = int(C.ETH_32_POOLS)
	ETH_64_POOLS = int(C.ETH_64_POOLS)
)

/* enum rte_eth_fc_mode */
const (
	RTE_FC_NONE     = int(C.RTE_FC_NONE)
	RTE_FC_RX_PAUSE = int(C.RTE_FC_RX_PAUSE)
	RTE_FC_TX_PAUSE = int(C.RTE_FC_TX_PAUSE)
	RTE_FC_FULL     = int(C.RTE_FC_FULL)
)

/* enum rte_fdir_pballoc_type */
const (
	RTE_FDIR_PBALLOC_64K  = int(C.RTE_FDIR_PBALLOC_64K)
	RTE_FDIR_PBALLOC_128K = int(C.RTE_FDIR_PBALLOC_128K)
	RTE_FDIR_PBALLOC_256K = int(C.RTE_FDIR_PBALLOC_256K)
)

/* enum rte_fdir_status_mode */
const (
	RTE_FDIR_NO_REPORT_STATUS     = int(C.RTE_FDIR_NO_REPORT_STATUS)
	RTE_FDIR_REPORT_STATUS        = int(C.RTE_FDIR_REPORT_STATUS)
	RTE_FDIR_REPORT_STATUS_ALWAYS = int(C.RTE_FDIR_REPORT_STATUS_ALWAYS)
)

/* enum rte_eth_dev_type */
const (
	RTE_ETH_DEV_UNKNOWN = int(C.RTE_ETH_DEV_UNKNOWN)
	RTE_ETH_DEV_PCI     = int(C.RTE_ETH_DEV_PCI)
	RTE_ETH_DEV_VIRTUAL = int(C.RTE_ETH_DEV_VIRTUAL)
	RTE_ETH_DEV_MAX     = int(C.RTE_ETH_DEV_MAX)
)

/* enum rte_eth_event_type */
const (
	RTE_ETH_EVENT_UNKNOWN  = int(C.RTE_ETH_EVENT_UNKNOWN)
	RTE_ETH_EVENT_INTR_LSC = int(C.RTE_ETH_EVENT_INTR_LSC)
	RTE_ETH_EVENT_MAX      = int(C.RTE_ETH_EVENT_MAX)
)

type RteEthStats C.struct_rte_eth_stats
type RteEthLink C.struct_rte_eth_link
type RteEthThresh C.struct_rte_eth_thresh
type RteEthRxMode C.struct_rte_eth_rxmode
type RteEthRssConf C.struct_rte_eth_rss_conf
type RteEthVlanMirror C.struct_rte_eth_vlan_mirror
type RteEthMirrorConf C.struct_rte_eth_mirror_conf
type RteEthRssRetaEntry64 C.struct_rte_eth_rss_reta_entry64
type RteEthVmdqDcbConf C.struct_rte_eth_vmdq_dcb_conf
type RteEthTxmode C.struct_rte_eth_txmode
type RteEthRxConf C.struct_rte_eth_rxconf
type RteEthTxConf C.struct_rte_eth_txconf
type RteEthDescLim C.struct_rte_eth_desc_lim
type RteEthFcConf C.struct_rte_eth_fc_conf
type RteEthPfcConf C.struct_rte_eth_pfc_conf
type RteFdirConf C.struct_rte_fdir_conf
type RteEthUdpTunnel C.struct_rte_eth_udp_tunnel
type RteIntrConf C.struct_rte_intr_conf
type RteEthConf C.struct_rte_eth_conf
type RteEthRxqInfo C.struct_rte_eth_rxq_info
type RteEthTxqInfo C.struct_rte_eth_txq_info
type RteEthXStats C.struct_rte_eth_xstats
type RteEthDcbTcQueueMapping C.struct_rte_eth_dcb_tc_queue_mapping
type RteEthDcbInfo C.struct_rte_eth_dcb_info
type RteEthAddr C.struct_ether_addr

//
func RteEthDevCount() uint {
	return uint(C.rte_eth_dev_count())
}

func RteEthDevAttach(devargs string, port_id *uint) int {
	return int(C.rte_eth_dev_attach(C.CString(devargs),
		(*C.uint8_t)(unsafe.Pointer((port_id)))))
}

func RteEthDevDetach(port_id uint, devname string) int {
	return int(C.rte_eth_dev_detach(C.uint8_t(port_id), C.CString(devname)))
}

func RteEthDevConfigure(port_id, nb_rx_queue, nb_tx_queue uint, eth_conf *RteEthConf) int {
	return int(C.rte_eth_dev_configure(C.uint8_t(port_id),
		C.uint16_t(nb_rx_queue), C.uint16_t(nb_tx_queue),
		(*C.struct_rte_eth_conf)(eth_conf)))
}

func RteEthRxQueueSetup(port_id, rx_queue_id, nb_rx_desc, socket_id uint,
	rx_conf *RteEthRxConf, mb_pool *RteMemPool) int {
	return int(C.rte_eth_rx_queue_setup(C.uint8_t(port_id),
		C.uint16_t(rx_queue_id), C.uint16_t(nb_rx_desc),
		C.unsigned(socket_id), (*C.struct_rte_eth_rxconf)(rx_conf),
		(*C.struct_rte_mempool)(mb_pool)))
}

func RteEthTxQueueSetup(port_id, tx_queue_id, nb_tx_desc, socket_id uint,
	tx_conf *RteEthTxConf) int {
	return int(C.rte_eth_tx_queue_setup(C.uint8_t(port_id),
		C.uint16_t(tx_queue_id), C.uint16_t(nb_tx_desc),
		C.unsigned(socket_id), (*C.struct_rte_eth_txconf)(tx_conf)))
}

func RteEthDevStart(port_id uint) int {
	return int(C.rte_eth_dev_start(C.uint8_t(port_id)))
}

func RteEthDevStop(port_id uint) {
	C.rte_eth_dev_stop(C.uint8_t(port_id))
}

func RteEthDevSetLinkUp(port_id uint) int {
	return int(C.rte_eth_dev_set_link_up(C.uint8_t(port_id)))
}

func RteEthDevSetLinkDown(port_id uint) int {
	return int(C.rte_eth_dev_set_link_down(C.uint8_t(port_id)))
}

func RteEthDevClose(port_id uint) {
	C.rte_eth_dev_close(C.uint8_t(C.uint8_t(port_id)))
}

func RteEthPromiscuousEnable(port_id uint) {
	C.rte_eth_promiscuous_enable(C.uint8_t(port_id))
}

func RteEthPromiscuousDisable(port_id uint) {
	C.rte_eth_promiscuous_disable(C.uint8_t(port_id))
}

func RteEthPromiscuousGet(port_id uint) int {
	return int(C.rte_eth_promiscuous_get(C.uint8_t(port_id)))
}

func RteEthRxBurst(port_id, queue_id uint, rx_pkts *unsafe.Pointer, nb_pkts uint) uint {
	return uint(C.rte_eth_rx_burst(C.uint8_t(port_id), C.uint16_t(queue_id),
		(**C.struct_rte_mbuf)(unsafe.Pointer(rx_pkts)), C.uint16_t(nb_pkts)))
}

func RteEthRxQueueCount(port_id, queue_id uint) uint {
	return uint(C.rte_eth_rx_queue_count(C.uint8_t(port_id),
		C.uint16_t(queue_id)))
}

func RteEthRxQueueDescriptorDone(port_id, queue_id, offset uint) uint {
	return uint(C.rte_eth_rx_descriptor_done(C.uint8_t(port_id),
		C.uint16_t(queue_id), C.uint16_t(offset)))
}

func RteEthTxBurst(port_id, queue_id uint, tx_pkts *unsafe.Pointer, nb_pkts uint) uint {
	return uint(C.rte_eth_tx_burst(C.uint8_t(port_id), C.uint16_t(queue_id),
		(**C.struct_rte_mbuf)(unsafe.Pointer(tx_pkts)), C.uint16_t(nb_pkts)))
}

func RteEthDevSocketID(port_id uint) uint {
	return uint(C.rte_eth_dev_socket_id(C.uint8_t(port_id)))
}

func RteEthMacAddr(port_id uint) string {
	var addr C.struct_ether_addr
	C.rte_eth_macaddr_get(C.uint8_t(port_id), &addr)
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", addr.addr_bytes[0], addr.addr_bytes[1], addr.addr_bytes[2], addr.addr_bytes[3], addr.addr_bytes[4], addr.addr_bytes[5])
}
