package dpdk

/*
#cgo CFLAGS: -m64 -pthread -O3 -march=native -I/usr/local/include/dpdk
#cgo LDFLAGS: -L/usr/lib -ldpdk -lz -lrt -lm -ldl -lfuse

#include <rte_config.h>
#include <rte_ring.h>
*/
import "C"
import "unsafe"

type RteRing C.struct_rte_ring

const (
	RTE_RING_NAMESIZE        = 32
	RTE_RING_PAUSE_REP_COUNT = 0
	RING_F_SP_ENQ            = 0x0001
	RING_F_SC_DEQ            = 0x0002
	RTE_RING_QUOT_EXCEED     = 1 << 31
	RTE_RING_SZ_MASK         = 0x0fffffff
)

func (r *RteRing) Name() string {
	var arr []byte
	for _, c := range (*C.struct_rte_ring)(r).name {
		arr = append(arr, byte(c))
	}
	return string(arr)
}

func RteRingGetMemsize(count uint) uint {
	return uint(C.rte_ring_get_memsize(C.unsigned(count)))
}

func (r *RteRing) RteRingInit(name string, count, flags uint) int {
	return int(C.rte_ring_init((*C.struct_rte_ring)(r), C.CString(name),
		C.unsigned(count), C.unsigned(flags)))
}

func RteRingCreate(name string, count uint, socket_id int, flags uint) *RteRing {
	return (*RteRing)(C.rte_ring_create(C.CString(name), C.unsigned(count),
		C.int(socket_id), C.unsigned(flags)))
}

func RteRingLookup(name string) *RteRing {
	return (*RteRing)(C.rte_ring_lookup(C.CString(name)))
}

func (r *RteRing) Free() {
	C.rte_ring_free((*C.struct_rte_ring)(r))
}

func (r *RteRing) SetWaterMark(count uint) int {
	return int(C.rte_ring_set_water_mark((*C.struct_rte_ring)(r), C.unsigned(count)))
}

func (r *RteRing) EnqueueBurst(tbl *unsafe.Pointer, n uint) uint {
	return uint(C.rte_ring_enqueue_burst((*C.struct_rte_ring)(r), tbl, C.unsigned(n)))
}

func (r *RteRing) EnqueueBulk(tbl *unsafe.Pointer, n uint) int {
	return int(C.rte_ring_enqueue_bulk((*C.struct_rte_ring)(r), tbl, C.unsigned(n)))
}

func (r *RteRing) Enqueue(obj unsafe.Pointer) int {
	return int(C.rte_ring_enqueue((*C.struct_rte_ring)(r), obj))
}

func (r *RteRing) DequeueBurst(tbl *unsafe.Pointer, n uint) uint {
	return uint(C.rte_ring_dequeue_burst((*C.struct_rte_ring)(r), tbl, C.unsigned(n)))
}

func (r *RteRing) DequeueBulk(tbl *unsafe.Pointer, n uint) int {
	return int(C.rte_ring_dequeue_bulk((*C.struct_rte_ring)(r), tbl, C.unsigned(n)))
}

func (r *RteRing) Dequeue(obj *unsafe.Pointer) int {
	return int(C.rte_ring_dequeue((*C.struct_rte_ring)(r), obj))
}

func (r *RteRing) Full() bool {
	return int(C.rte_ring_full((*C.struct_rte_ring)(r))) == 1
}

func (r *RteRing) Empty() bool {
	return int(C.rte_ring_empty((*C.struct_rte_ring)(r))) == 1
}

func (r *RteRing) Count() uint {
	return uint(C.rte_ring_count((*C.struct_rte_ring)(r)))
}

func (r *RteRing) FreeCount() uint {
	return uint(C.rte_ring_free_count((*C.struct_rte_ring)(r)))
}
