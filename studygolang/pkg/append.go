package pkg

import (
	"fmt"
)

func Testappend() {
	in := []byte("abc")
	hash := []byte("xyz")
	sum := append(in, hash[:]...)
	
	fmt.Printf("\n\n====> Testappend sum=[%s]\n", sum)
}