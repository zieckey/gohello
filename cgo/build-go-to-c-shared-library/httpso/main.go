package main

import (
	//"C"

	"net/http"
	"log"
	"os"
)

// main is required to build a shared library, but does nothing
func main() {}

// export StartHTTP
func StartHTTP() {
	http.HandleFunc("/echo", handler)
	hostname, _ := os.Hostname()
	log.Printf("start http://%s:8091/echo", hostname)
	log.Fatal(http.ListenAndServe(":8091", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	buf := make([]byte, 1024*128)
	n, _ := r.Body.Read(buf) //Read the http body
	log.Printf("[%v]", string(buf[0:n]))
	w.Write(buf[0:n])
}



