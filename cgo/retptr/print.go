package main

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L .
#include "myprint.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"
import "fmt"

func Prints() {

	/*
	   see : http://golang.org/cmd/cgo/

	   // Go string to C string
	   // The C string is allocated in the C heap using malloc.
	   // It is the caller's responsibility to arrange for it to be
	   // freed, such as by calling C.free (be sure to include stdlib.h
	   // if C.free is needed).
	   func C.CString(string) *C.char

	   // C string to Go string
	   func C.GoString(*C.char) string

	   // C string, length to Go string
	   func C.GoStringN(*C.char, C.int) string

	   // C pointer, length to Go []byte
	   func C.GoBytes(unsafe.Pointer, C.int) []byte

	*/
	retlen := C.int(0)
	len := 5
	cstr := C.retmalloc(C.int(len), &retlen)
	gostr := C.GoStringN(cstr, retlen)
	fmt.Printf("retlen=%v\n", retlen)
	println(gostr)
	C.free(unsafe.Pointer(cstr))
}
