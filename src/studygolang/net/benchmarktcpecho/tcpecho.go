package main

import (
	"flag"
	"log"
	"net"
	"time"
	"sync/atomic"
)

type Stat struct {
	connections      int32
	receivedMessages int32
	transferredBytes int32
}

var port = flag.String("port", "2007", "input ur name")

func main() {
	stat := new(Stat)
	ln, err := net.Listen("tcp", string(":")+*port)
	if err != nil {
		log.Println(err)
		return
	}
	go print(stat)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go echoFunc(conn, stat)
	}
}

func echoFunc(c net.Conn, stat *Stat) {
	buf := make([]byte, 1024)

	for {
		n, err := c.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}

		c.Write(buf[:n])
	}
}

func print(stat *Stat) {
	for {
		time.Sleep(1000 * time.Millisecond)
		bytes := float64(atomic.SwapInt32(&stat.transferredBytes, 0)) / 1024.0
		msgs := float64(atomic.SwapInt32(&stat.receivedMessages, 0))
		bytesPerMsg := 0.0
		if msgs > 0 {
			bytesPerMsg = bytes / msgs
		}
		log.Printf("qps=%v conns=%v transfer=%v bytesPerMsg=%v\n", msgs, stat.connections, bytes, bytesPerMsg)
	}
}
