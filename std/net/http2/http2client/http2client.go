package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    client := http.Client{
        // InsecureTLSDial is temporary and will likely be
        // replaced by a different API later.
        // Transport: &http2.Transport{InsecureTLSDial: true},
        Transport: http.,
    }

    resp, err := client.Get("https://localhost:8000/")
    if err != nil {
        log.Fatal(err)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(body))
}