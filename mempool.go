package dpdk

/*
#cgo CFLAGS: -m64 -pthread -O3 -march=native -I/usr/local/include/dpdk
#cgo LDFLAGS: -L/usr/lib -ldpdk -lz -lrt -lm -ldl -lfuse

#include <rte_config.h>
#include <rte_mempool.h>

typedef void(*go_walk_t)(const struct rte_mempool *, void *);
extern void go_mempool_walk(struct rte_mempool *mp, void *arg);;
*/
import "C"
import "unsafe"

type RteMemPool C.struct_rte_mempool
type RteMemPoolObjSz C.struct_rte_mempool_objsz
type RteMemPoolObjHdr C.struct_rte_mempool_objsz
type RteMemPoolObjTlr C.struct_rte_mempool_objsz

const (
	RTE_MEMPOOL_HEADER_COOKIE1 = 0xbadbadbadadd2e55
	RTE_MEMPOOL_HEADER_COOKIE2 = 0xf2eef2eedadd2e55
	RTE_MEMPOOL_TRAILER_COOKIE = 0xadd2e55badbadbad
	RTE_MEMPOOL_NAMESIZE       = 32
	MEMPOOL_PG_NUM_DEFAULT     = 1
	MEMPOOL_F_NO_SPREAD        = 0x0001
	MEMPOOL_F_NO_CACHE_ALIGN   = 0x0002
	MEMPOOL_F_SP_PUT           = 0x0004
	MEMPOOL_F_SC_GET           = 0x0008
)

func (mp *RteMemPool) Name() string {
	var arr []byte
	for _, c := range (*C.struct_rte_mempool)(mp).name {
		arr = append(arr, byte(c))
	}
	return string(arr)
}

func RteMemPoolCreate(name string, n, elt_size, cache_size, priv_data_size uint,
	socket_id int, flags uint) *RteMemPool {
	return (*RteMemPool)(C.rte_mempool_create(C.CString(name),
		C.unsigned(n),
		C.unsigned(elt_size),
		C.unsigned(cache_size),
		C.unsigned(priv_data_size),
		nil, nil, nil, nil,
		C.int(socket_id),
		C.unsigned(flags)))
}

func RteMemPoolLookup(name string) *RteMemPool {
	return (*RteMemPool)(C.rte_mempool_lookup(C.CString(name)))
}

func (mp *RteMemPool) PutBulk(tbl *unsafe.Pointer, n uint) {
	C.rte_mempool_put_bulk((*C.struct_rte_mempool)(mp), tbl, C.unsigned(n))
}

func (mp *RteMemPool) Put(obj unsafe.Pointer) {
	C.rte_mempool_put((*C.struct_rte_mempool)(mp), obj)
}

func (mp *RteMemPool) GetBulk(tbl *unsafe.Pointer, n uint) int {
	return int(C.rte_mempool_get_bulk((*C.struct_rte_mempool)(mp), tbl, C.unsigned(n)))
}

func (mp *RteMemPool) Get(obj *unsafe.Pointer) int {
	return int(C.rte_mempool_get((*C.struct_rte_mempool)(mp), obj))
}

func (mp *RteMemPool) Count() uint {
	return uint(C.rte_mempool_count((*C.struct_rte_mempool)(mp)))
}

func (mp *RteMemPool) FreeCount() uint {
	return uint(C.rte_mempool_free_count((*C.struct_rte_mempool)(mp)))
}

func (mp *RteMemPool) Full() bool {
	return int(C.rte_mempool_full((*C.struct_rte_mempool)(mp))) == 1
}

func (mp *RteMemPool) Empty() bool {
	return int(C.rte_mempool_empty((*C.struct_rte_mempool)(mp))) == 1
}

func (mp *RteMemPool) CalcObjSize(elt_size, flags uint32, sz *RteMemPoolObjSz) uint32 {
	return uint32(C.rte_mempool_calc_obj_size(C.uint32_t(elt_size), C.uint32_t(flags),
		(*C.struct_rte_mempool_objsz)(sz)))
}

var mempoolWalk func(*RteMemPool, unsafe.Pointer)

//export go_mempool_walk
func go_mempool_walk(mp *C.struct_rte_mempool, arg unsafe.Pointer) {
	mempoolWalk((*RteMemPool)(mp), arg)
}

func RteMemPoolWalk(fn func(*RteMemPool, unsafe.Pointer), arg unsafe.Pointer) {
	mempoolWalk = fn
	C.rte_mempool_walk((C.go_walk_t)(C.go_mempool_walk), arg)
}
