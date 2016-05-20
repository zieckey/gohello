package main

import (
    "fmt"
    "github.com/nbio/httpcontext"
    "net/http"
    "log"
)

func Hello(w http.ResponseWriter, r *http.Request) {
    httpcontext.Set(r, "key1", "value1") // Set a context with this request r
    val := httpcontext.Get(r, "key1")    // Get the context
    v, _ := val.(string)
    fmt.Printf("Got a value associated with key1 : %v\n", v)
    w.Write([]byte("OK"))
}

func main() {
    http.HandleFunc("/hello", Hello)
    log.Fatal(http.ListenAndServe(":8080", nil))
}