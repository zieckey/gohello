package main

/*
#cgo CFLAGS: -I ./pcap/include
#cgo linux LDFLAGS: -L ./pcap -lpcap
#cgo windows,386 LDFLAGS: -L ./pcap/msvc/bin/debug -lpcap
#cgo windows,amd64 LDFLAGS: -L ./pcap/msvc/bin/debug -lpcap
#include <stdlib.h>
#include "pcap.h"
*/
import "C"
import (
	"unsafe"
)

func main() {
	s := "Hello Cgo"
	cs := C.CString(s)
	C.pcap_print(cs)
	C.free(unsafe.Pointer(cs))
}
