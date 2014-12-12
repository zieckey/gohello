package main

import (
	"fmt"
)

type Handler interface {
	ServeHTTP(buf string)
}

type HandlerFunc func(buf string)

func (f HandlerFunc) ServeHTTP(buf string) {
	f(buf)
}

func Handle(pattern string, handler Handler) {
	fmt.Printf("Handle --> pattern=%v\n", pattern)
	handler.ServeHTTP(pattern)
}

func HandleFunc(pattern string, handler func(buf string)) {
	fmt.Printf("HandleFunc --> pattern=%v\n", pattern)
	Handle(pattern, HandlerFunc(handler))
}

func mycallback1(buf string) {
	fmt.Printf("mycallback1 called : %v\n", buf)
}

type Mytype int

func (t Mytype) ServeHTTP(buf string) {
	fmt.Printf("Mytype.ServeHTTP called : t=%v %v\n", int(t), buf)
}

func main() {
	HandleFunc("c-like func", mycallback1)
	
	fmt.Println()
	var t Mytype
	t = 1
	Handle("cpp-like func", t)
}
