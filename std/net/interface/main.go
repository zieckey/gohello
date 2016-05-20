package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

func test1() {
	println(time.Now().String())
}

func test2() {
	info, _ := net.InterfaceAddrs()
	for _, addr := range info {
		fmt.Println(strings.Split(addr.String(), "/")[0])
	}
}

func main() {
	//test1()
	test2()
}
