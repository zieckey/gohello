package pkg

import (
	"fmt"
)

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf(" i=%v ", i)
	}

	j := 1
	{
		defer fmt.Printf(" j=%v ", j) // output j=1
		j++
	}
	j++
}

func TestDefer() {
	fmt.Printf("\nTestDefer:\n")
	f()
}
