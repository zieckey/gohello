package main

import (
    "fmt"
    "log"
    "time"
)

func main() {
    fmt.Printf("hello world\n")
    fmt.Printf("xx\n")
    log.Printf("xxx")
    var n int
    var s string
    fmt.Scanf("%d %s", &n, &s)
    fmt.Printf("%d %s\n", n, s)
    fmt.Printf("%d\n", len(s))
    fmt.Printf("%s\n", s[0:2])
    fmt.Printf("%s\n", s[2:])
    t := time.Now()
    t.Unix()
}

