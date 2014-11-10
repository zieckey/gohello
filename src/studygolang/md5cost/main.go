package main

import (
	"crypto/md5"
	"time"
	"fmt"
)

func main() {
	data := []byte("ja;slfja;lsdjfpu-2341293ursadklf;asnvzaxcvklahpsdfp-98ur4-23hf;asdfja")
	loop := 100000
	begin := time.Now()
	for i := 0; i < loop; i++ {
		md5.Sum(data)
	}
	end := time.Now()
	d := end.Sub(begin)
	fmt.Printf("cost %v", d.Seconds())
}
