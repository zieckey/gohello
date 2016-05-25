package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	buf := make([]byte, 1024)
	n, _ := r.Body.Read(buf) //Read the http body
	log.Printf("[%v]", string(buf[0:n]))
	w.Write(buf[0:n])
}

func main() {
	port := 8091
	if len(os.Args) == 2 {
		port, _ = strconv.Atoi(os.Args[1])
	}
	http.HandleFunc("/echo", handler)
	hostname, _ := os.Hostname()
	log.Printf("start http://%s:%v/echo", hostname, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
