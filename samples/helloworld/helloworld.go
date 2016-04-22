package main

import (
	"log"
	"os"
	"unsafe"

	"github.com/feiskyer/dpdk-go"
)

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
