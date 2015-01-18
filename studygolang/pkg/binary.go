package pkg

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

func BinaryEncoding() {
	b := []byte{0x00, 0x00, 0x03, 0xe8}
	buf := bytes.NewBuffer(b)
	var x int32
	binary.Read(buf, binary.BigEndian, &x)
	fmt.Println(x)

	fmt.Println(strings.Repeat("-", 10))

	x = 1000
	buf = bytes.NewBuffer([]byte{})
	binary.Write(buf, binary.BigEndian, x)
	fmt.Println(buf.Bytes())
}
