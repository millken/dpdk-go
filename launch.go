package dpdk

/*
#cgo CFLAGS: -m64 -pthread -O3 -march=native -I/usr/local/include/dpdk
#cgo LDFLAGS: -L/usr/local/lib -ldpdk -lz -lrt -lm -ldl -lfuse

extern int go_remote_launch(void *);
extern int go_mp_remote_launch(void *);
#include <stdlib.h>
#include <limits.h>
#include <rte_config.h>
#include <rte_common.h>
#include <rte_launch.h>
#include <rte_lcore.h>
#include <rte_per_lcore.h>
*/
import "C"
import "unsafe"

const (
	WAIT = iota
	RUNNING
	FINISHED
)

const (
	SKIP_MASTER = iota
	CALL_MASTER
)

var lCoreFuncs [](func(unsafe.Pointer) int)
var lCoreFunc func(unsafe.Pointer) int
var lCoreFuncMp func(unsafe.Pointer) int

//export go_remote_launch
func go_remote_launch(arg unsafe.Pointer) C.int {
	return C.int(lCoreFunc(arg))
}

//export go_mp_remote_launch
func go_mp_remote_launch(arg unsafe.Pointer) C.int {
	return C.int(lCoreFuncMp(arg))
}

func RteEalRemoteLaunch(fn func(unsafe.Pointer) int, arg unsafe.Pointer, slave_id uint) int {
	if len(lCoreFuncs) == 0 {
		lCoreFuncs = make([](func(unsafe.Pointer) int), 128)
	}
	lCoreFuncs[int(slave_id)] = fn
	lCoreFunc = fn
	return int(C.rte_eal_remote_launch((*C.lcore_function_t)(C.go_remote_launch),
		arg, C.unsigned(slave_id)))
}

func RteEalLaunchAllSlave(fn func(unsafe.Pointer) int, arg unsafe.Pointer) {
	for lcoreID := C.rte_get_next_lcore(C.UINT_MAX, 1, 0); lcoreID < C.RTE_MAX_LCORE; lcoreID = C.rte_get_next_lcore(C.uint(lcoreID), 1, 0) {
		RteEalRemoteLaunch(fn, arg, uint(lcoreID))
	}
}

func RteEalMpRemoteLaunch(fn func(unsafe.Pointer) int, arg unsafe.Pointer, call_master int) int {
	lCoreFuncMp = fn
	return int(C.rte_eal_mp_remote_launch((*C.lcore_function_t)(C.go_mp_remote_launch),
		arg, C.enum_rte_rmt_call_master_t(call_master)))
}

func RteEalGetLCoreState(slave_id uint) int {
	return int(C.rte_eal_get_lcore_state(C.unsigned(slave_id)))
}

func RteEalWaitLCore(slave_id uint) int {
	return int(C.rte_eal_wait_lcore(C.unsigned(slave_id)))
}

func RteSocketID() int {
	return int(C.rte_socket_id())
}

func RteEalMpWaitLCore() {
	C.rte_eal_mp_wait_lcore()
}

func RteLcoreID() uint {
	return uint(C.rte_lcore_id())
}
