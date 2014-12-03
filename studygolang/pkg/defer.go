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
	fmt.Printf("\n\n")
	funcB()
	fmt.Printf("\n\n")
}

// Output : acdb
// defer的执行点是函数退出时顺序执行，不是离开代码块的时候执行
func funcB() {
	fmt.Printf("a")
	{
		defer fmt.Printf("b")
	}
	fmt.Printf("c")
	defer fmt.Printf("d")
}
