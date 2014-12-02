package main

import (
	"fmt"
	"os"	
	"github.com/zieckey/gohello/studygolang/goid"
)

func main() {
	for i := 0; i < 20; i++ {
		go func() {
			fmt.Printf("goid=%v\n", goid.Get())
			for {
				b := make([]byte, 10)
				os.Stdin.Read(b) // will block
			}
		}()
	}

	select {}
}
