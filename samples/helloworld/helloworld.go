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

func helloworld(arg unsafe.Pointer) int {
	lcoreID := dpdk.RteLcoreID()
	log.Printf("Hello from core %d", lcoreID)
	return 0
}

func main() {
	ret := dpdk.RteEalInit(os.Args)
	if ret < 0 {
		log.Fatalf("rte_eal_init failed: %s", dpdk.StrError(ret))
		return
	}

	log.Println("rte inited.")

	// call helloworld() on every slave lcore
	dpdk.RteEalLaunchAllSlave(helloworld, unsafe.Pointer(nil))

	// call helloworld on master lcore too
	helloworld(unsafe.Pointer(nil))

	dpdk.RteEalMpWaitLCore()
}
