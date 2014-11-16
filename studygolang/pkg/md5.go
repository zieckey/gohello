package pkg

import (
	"crypto/md5"
	"fmt"
)

func TestMd5() {
	fmt.Printf("\n\n\n")
	d := []byte("a")
	m := md5.Sum(d)
	fmt.Printf("%x %v\n", m, m) // 0cc175b9c0f1b6a831c399e269772661
	hex := fmt.Sprintf("%x", m)
	fmt.Printf("%v", hex)
}
