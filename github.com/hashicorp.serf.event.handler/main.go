package main

import (
	"os"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime)
	name := os.Getenv("SERF_USER_EVENT")
	buf, err := ioutil.ReadAll(os.Stdin)
	log.Printf("name=[%v] stdin=[%v] err=[%v]\n", name, strings.TrimSpace(string(buf)), err)
}

