package main 

import (
	"github.com/zieckey/gohello/freader"
	"log"
)

// ./freader.exe -file_path="e:\1\1" -file_pattern="ddd*" -stderrthreshold=0 -logtostderr=true
func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	freader.Run()
}

