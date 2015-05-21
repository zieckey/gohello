package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func makeBuffer() []byte {
	return make([]byte, rand.Intn(5000000)+5000000)
}

func main() {
	plot := `===
{ "Name" : "line", "Height" : 600, "Width" : 1900, "ItemName" : ["HeapSys", "bytes", "HeapAlloc", "HeapIdle", "HeapReleased", "NumGC", "HeapObjects"] }
---`
	fmt.Printf("%v\n", plot)
	
	pool := make([][]byte, 20)

	var m runtime.MemStats
	makes := 0
	for j := 1; j < 10000000; j++ {
		b := makeBuffer()
		makes += 1
		i := rand.Intn(len(pool))
		pool[i] = b

		time.Sleep(time.Second)

		bytes := 0

		for i := 0; i < len(pool); i++ {
			if pool[i] != nil {
				bytes += len(pool[i])
			}
		}

		runtime.ReadMemStats(&m)
		fmt.Printf("%d %d %d %d %d %d %d %d\n", j, m.HeapSys, bytes, m.HeapAlloc,
			m.HeapIdle, m.HeapReleased, m.NumGC, m.HeapObjects)
	}
}
