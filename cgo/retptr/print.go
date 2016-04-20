package main

/*
#include <stdlib.h>
#include <string.h>
char* retmalloc(int len, int *retlen)
{
    static const char* s = "0123456789";
    char* p = malloc(len);
    if (len <= strlen(s)) {
        memcpy(p, s, len);
    } else {
        memset(p, 'a', len);
    }
    *retlen = len;
    return p;
}
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
	len := 10
	cstr := C.retmalloc(C.int(len), &retlen)
	defer C.free(unsafe.Pointer(cstr))
	gostr := C.GoStringN(cstr, retlen)
	fmt.Printf("retlen=%v\n", retlen)
	println(gostr)
}
