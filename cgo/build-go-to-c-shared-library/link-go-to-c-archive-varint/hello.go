package main

import "C"
import "fmt"

import (
	"bytes"
	"encoding/binary"
)

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

type DocId struct {
	Id  uint32
	Row uint32
}

func (d *DocId) Encode() []byte {
	buf := make([]byte, 8)
	n := binary.PutUvarint(buf, uint64(d.Id))
	n = binary.PutUvarint(buf[n:], uint64(d.Row)) + n
	return buf[0:n]
}

func (d *DocId) Decode(buf []byte) error {
	r := bytes.NewReader(buf)
	u, err := binary.ReadUvarint(r)
	if err != nil {
		return err
	}
	d.Id = uint32(u)
	u, err = binary.ReadUvarint(r)
	d.Row = uint32(u)
	return err
}

//export DocIdEncode
func DocIdEncode(in interface{}) []byte {
	d := in.(DocId)
	return d.Encode()
}
