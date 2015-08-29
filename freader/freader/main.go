package main 

import (
	"github.com/zieckey/gohello/freader"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	freader.Run()
}

