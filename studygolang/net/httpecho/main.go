package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	buf := make([]byte, 1024)
	n, _ := r.Body.Read(buf) //Read the http body
	log.Printf("[%v]", string(buf[0:n]))
	w.Write(buf[0:n])
}

func main() {
	http.HandleFunc("/echo", handler)
	log.Fatal(http.ListenAndServe(":8091", nil))
}
