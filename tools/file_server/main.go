package main

import (
    "net/http"
    "os"
)

func main() {
    dir := "."
    if len(os.Args) == 2 {
        dir = os.Args[1]
    }
    http.Handle("/", http.FileServer(http.Dir(dir)))
    http.ListenAndServe(":8000", nil)
}


