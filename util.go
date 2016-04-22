package dpdk

/*
#include <rte_config.h>
#include <rte_malloc.h>
#include <rte_errno.h>
*/
import "C"
import "unsafe"

func GetCArray(n uint) *unsafe.Pointer {
	var p *unsafe.Pointer
	arr := C.rte_malloc((*C.char)(nil), C.size_t(n)*C.size_t(unsafe.Sizeof(p)), C.unsigned(0))
	return (*unsafe.Pointer)(arr)
}

func SliceFromCArray(arr *unsafe.Pointer, n uint) []unsafe.Pointer {
	return (*[1 << 30](unsafe.Pointer))(unsafe.Pointer(arr))[:n:n]
}

func StrError(errno int) string {
    return C.GoString(C.rte_strerror(C.int(errno)))
}
