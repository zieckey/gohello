package main

import (
	"fmt"
	"strconv"
	"runtime"
)


func main() {
    //runtime.GOMAXPROCS(runtime.NumCPU())
    runtime.GOMAXPROCS(2)
    ch := make(chan int)
    fmt.Printf("begin\n")
    loop := 10
    for j := 0; j < loop; j++ {
	    task("A" + strconv.Itoa(j), ch)
    }
    <-ch

}

func task(name string, ch chan int) {
    go func() {
        i := 1
        for {
            //fmt.Printf("%s %s\n", name, strconv.Itoa(i))
            for j := 0; j < 20; j++ {
                strconv.Itoa(i) // cost some CPU time
            }
            fmt.Printf("goid=" + strconv.Itoa(int(GoroutineID())) + " " + name + " " + strconv.Itoa(i) + "\n")
            i++
        }
        ch <- 1
    }()
}
