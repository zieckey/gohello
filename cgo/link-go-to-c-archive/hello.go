// package name: nautilus
package main

import "C"
import "fmt"

//export SayHello
func SayHello(name string) {
	fmt.Printf("func in Golang SayHello says: Hello, %s!\n", name)
}

//export SayHelloByte
func SayHelloByte(name []byte) {
	fmt.Printf("func in Golang SayHelloByte says: Hello, %s!\n", string(name))
}

//export SayBye
func SayBye() {
	fmt.Println("func in Golang SayBye says: Bye!\n")
}

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}
