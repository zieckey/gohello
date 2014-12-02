package main

import (
	"fmt"
	"github.com/zieckey/gohello/studygolang/goid"
	"runtime"
	"strconv"
)

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.GOMAXPROCS(2)
	ch := make(chan int)
	task("A", ch)
	task("B", ch)
	task("C", ch)
	task("D", ch)
	fmt.Printf("begin\n")
	<-ch
	<-ch
	<-ch
	<-ch
}

func task(name string, ch chan int) {
	go func() {
		i := 1
		for {
			//fmt.Printf("%s %s\n", name, strconv.Itoa(i))
			for j := 0; j < 2000000; j++ {
				strconv.Itoa(i) // cost some CPU time
			}
			print("goid=" + strconv.Itoa(int(goid.Get())) + " " + name + " " + strconv.Itoa(i) + "\n")
			i++
		}
		ch <- 1
	}()
}
