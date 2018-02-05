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

    t := time.Now()
    t.Unix()
}

