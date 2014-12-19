package main

import (
	"fmt"
	"strconv"
	"runtime"
	"bytes"
)

const stackBufSize = 16 << 20

var stackBuf = make(chan []byte, 8)

func getBuf() []byte {
	select {
	case b := <-stackBuf:
		return b[:stackBufSize]
	default:
		return make([]byte, stackBufSize)
	}
}

func putBuf(b []byte) {
	select {
	case stackBuf <- b:
	default:
	}
}

var goroutineSpace = []byte("goroutine ")

func GoroutineID() int64 {
	b := getBuf()
	defer putBuf(b)
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 otu of "goroutine 4707 ["
	b = bytes.TrimPrefix(b, goroutineSpace)
	i := bytes.IndexByte(b, ' ')
	if i < 0 {
		panic(fmt.Sprintf("No space found in %q", b))
	}
	b = b[:i]
	n, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse goroutine ID out of %q: %v", b, err))
	}
	return int64(n)
}

func main() {

    ch := make(chan int)
    fmt.Printf("begin\n")
    for j := 0; j < 1000; j++ {
	    task("A" + strconv.Itoa(j), ch)
    }
    <-ch

}

func task(name string, ch chan int) {
    go func() {
        i := 1
        for {
            //fmt.Printf("%s %s\n", name, strconv.Itoa(i))
            for j := 0; j < 20; j++ {
                strconv.Itoa(i) // cost some CPU time
            }
            print("goid=" + strconv.Itoa(int(GoroutineID())) + " " + name + " " + strconv.Itoa(i) + "\n")
            i++
        }
        ch <- 1
    }()
}
