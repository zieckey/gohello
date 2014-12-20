package goid

import (
	"fmt"
	"strconv"
	"runtime"
	"bytes"
)

// Take it from : http://camlistore.org/pkg/syncutil/lock.go?s=821:1134

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

func Get() int64 {
	b := getBuf()
	defer putBuf(b)
	b = b[:runtime.Stack(b, false)]
	// Parse the 4707 out of "goroutine 4707 [running]:"
	fmt.Printf("STACK BEGIN:\n[%s]\nSTACK END\n", string(b))
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
