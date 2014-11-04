package main

import "fmt"
import "newmath"

func main() {
    fmt.Printf("Hello, world. sqrt(2)=%v\n", newmath.Sqrt(2.0))
    
    for pos, char := range "a你x" {
		fmt.Printf("character '%c' starts at byte position %d\n", char, pos)
	}
}


