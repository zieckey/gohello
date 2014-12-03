package pkg

import (
	"fmt"
)

func testBytesParameter(buf []byte) {
	fmt.Printf("len(buf)=%v content:%v\n", len(buf), string(buf))
}

func testBytesParameter2(buf []byte) {
	fmt.Printf("input    -> len(buf)=%v content:%v\n", len(buf), string(buf))
	buf[0] = 'a'
	fmt.Printf("modified -> len(buf)=%v content:%v\n", len(buf), string(buf))
}

func testBytesCap() {
	buf := make([]byte, 32)
	copy(buf, []byte("123"))
	fmt.Println("len=%v cap=%v\n", len(buf), cap(buf))
}

/*All output:

TestBytes:
len(buf)=1 content:1
len(buf)=2 content:12
len(buf)=3 content:123
len(buf)=4 content:1234
input -> len(buf)=4 content:1234
modified -> len(buf)=4 content:a234
output -> len(buf)=4 content:a234

*/
func TestBytes() {
	fmt.Printf("TestBytes:\n")
	testBytesParameter([]byte("1"))
	testBytesParameter([]byte("12"))
	testBytesParameter([]byte("123"))
	testBytesParameter([]byte("1234"))
	
	buf := []byte("1234")
	testBytesParameter2(buf)
	fmt.Printf("output   -> len(buf)=%v content:%v\n", len(buf), string(buf))

	testBytesCap()
}