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
	fmt.Println("func in Golang SayBye says: Bye!")
}

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library

	/*
		-buildmode=c-archive
			Build the listed main package, plus all packages it imports,
			into a C archive file. The only callable symbols will be those
			functions exported using a cgo //export comment. Requires
			exactly one main package to be listed.

		-buildmode=c-shared
			Build the listed main packages, plus all packages that they
			import, into C shared libraries. The only callable symbols will
			be those functions exported using a cgo //export comment.
			Non-main packages are ignored.
	*/
}
