package dpdk

/*
extern int go_remote_launch(void *);
extern int go_mp_remote_launch(void *);

#include <rte_config.h>
#include <rte_launch.h>
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

func RteEalMpWaitLCore() {
	C.rte_eal_mp_wait_lcore()
}
