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
		log.Printf("recv %v byte, sum=%v, err=%v\n", n, sum, err.Error())
		if sum >= int(len) {
			break
		}
		if err != nil {
			log.Printf("recv %v byte, sum=%v, err=%v\n", n, sum, err.Error())
			break
		}
	}

	w.Write([]byte("OK"))
}

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	http.HandleFunc("/echo", handler)
	log.Fatal(http.ListenAndServe(":8091", nil))
}
