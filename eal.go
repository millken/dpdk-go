package dpdk

/*
#cgo CFLAGS: -m64 -pthread -O3 -march=native -I/usr/local/include/dpdk
#cgo LDFLAGS: -L/usr/local/lib -ldpdk -lz -lrt -lm -ldl -lfuse

extern void go_usage_hook(char *prg);

#include <rte_config.h>
#include <rte_common.h>
#include <rte_eal.h>

typedef const char* const_char_ptr;

inline void rte_eal_exit(int exit_code)
{
	rte_exit(exit_code, "Error with EAL initialization");
}
*/
import "C"
import "unsafe"

const (
	RTE_PROC_AUTO    = int(C.RTE_PROC_AUTO)
	RTE_PROC_PRIMARY = int(C.RTE_PROC_PRIMARY)
	RTE_PROC_SECONDARY
	RTE_PROC_INVALID
)

const (
	ROLE_RTE = C.ROLE_RTE
	ROLE_OFF = C.ROLE_OFF
)

const (
	RTE_MAGIC = 19820526
)

type RteConfig C.struct_rte_config

func RteEalGetConfiguration() *RteConfig {
	return (*RteConfig)(C.rte_eal_get_configuration())
}

func RteEalInit(argv []string) int {
	var args [][]byte
	for _, arg := range argv {
		args = append(args, []byte(arg))
	}

	var b *C.char
	ptrSize := unsafe.Sizeof(b)

	ptr := C.malloc(C.size_t(len(args)) * C.size_t(ptrSize))
	defer C.free(ptr)

	for i := 0; i < len(args); i++ {
		element := (**C.char)(unsafe.Pointer(uintptr(ptr) + uintptr(i)*ptrSize))
		*element = (*C.char)(unsafe.Pointer(&args[i][0]))
	}

	return int(C.rte_eal_init(C.int(len(args)), (**C.char)(ptr)))
}

func RteEalProcessType() int {
	return int(C.rte_eal_process_type())
}

func RteEalLcoreRole(lcoreId uint) int {
	return int(C.rte_eal_lcore_role(C.unsigned(lcoreId)))
}

func RteEalIoplInit() int {
	return int(C.rte_eal_iopl_init())
}

func RteEalHasHugePages() int {
	return int(C.rte_eal_has_hugepages())
}

func RteSysGetTid() int {
	return int(C.rte_sys_gettid())
}

func RteGetTid() int {
	return int(C.rte_gettid())
}

func RteExit(exitCode int) {
	C.rte_eal_exit(C.int(exitCode))
}

var applicationUsageHook func(string)

//export go_usage_hook
func go_usage_hook(prg *C.char) {
	applicationUsageHook(C.GoString(prg))
}

func RteSetApplicationUsageHook(hook func(string)) {
	applicationUsageHook = hook
	C.rte_set_application_usage_hook((C.rte_usage_hook_t)(C.go_usage_hook))
}
