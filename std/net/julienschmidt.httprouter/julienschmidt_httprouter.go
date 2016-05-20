package main

import (
    "fmt"
    "github.com/julienschmidt/httprouter"
    "net/http"
    "log"
)

func Hello(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
    fmt.Printf("%v\n", ps) // 这里的ps参数并没有解析出url上的参数。只是按照httprouter自己的规则解析URI的path上的参数。
    w.Write([]byte("OK"))
}

func main() {
    router := httprouter.New()
    router.Handle("GET", "/hello/:name", Hello)
    log.Fatal(http.ListenAndServe(":8080", router))
}