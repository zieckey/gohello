package main

import (
    "net/http"
    "os"
)

func main() {
    dir := "."
    if len(os.Args) == 2 {
        if (os.Args[1] == "--help" || os.Args[1] == "-help" || os.Args[1] == "-h") {
            println("Usage: %v <the-path-to-dir>", os.Args[0])
            return
        }
        dir = os.Args[1]
    }
    println("Start file server at http://127.0.0.1:8000/")
    http.Handle("/", http.FileServer(http.Dir(dir)))
    http.ListenAndServe(":8000", nil)
}


