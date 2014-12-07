package main

import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	buf := make([]byte, 1024)
	sum := 0
	for {
		n, err := r.Body.Read(buf)
		sum += n
		if sum >= int(len) {
			break
		}
		if err != nil {
			log.Printf("recv %v byte, ERROR %v\n", n, err.Error())
			break
		}
		log.Printf("recv %v byte\n", n)
	}

	w.Write([]byte("OK"))
}

func main() {
	http.HandleFunc("/echo", handler)
	log.Fatal(http.ListenAndServe(":8091", nil))
}
