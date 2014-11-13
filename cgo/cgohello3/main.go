package main

/*
#cgo CFLAGS: -I .
#cgo LDFLAGS: -L .
#include "myprint.h"
#include <stdlib.h>
*/
import "C"
import "unsafe"

func Prints(s string) {

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
    p := C.CString(s);
    defer C.free(unsafe.Pointer(p))
    C.myprint(p);
}

//export FunctionExportedFromGo
func FunctionExportedFromGo() {
    println("this is a string printed from go")
}

//export GoFuncPrintxx 
func GoFuncPrintxx() {
    println("this is a string printed from go. GoFuncPrintxx")
}

func main() {
    s := "hello string in golang";
    Prints(s)
}
