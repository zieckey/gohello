package main

import (
	"fmt"
	proto "github.com/golang/protobuf/proto"
	"encoding/base64"
)

func main() {
	i := 1
	m1 := &M1{
		I:       proto.Int32(int32(i)),
	}

	m2 := &M2{
		I:       proto.Uint32(uint32(i)),
	}

	buf1, _ := proto.Marshal(m1) // len(buf1)=2  两个字节：0x08 0x01
	buf2, _ := proto.Marshal(m2) // len(buf1)=2  两个字节：0x08 0x01
	fmt.Printf("protobuf1 i=%v [%v]: %v\n", i, base64.StdEncoding.EncodeToString(buf1), buf1)
	fmt.Printf("protobuf2 i=%v [%v]: %v\n", i, base64.StdEncoding.EncodeToString(buf2), buf2)

	i = 127
	m1.I = proto.Int32(int32(i))
	m2.I = proto.Uint32(uint32(i))
	buf1, _ = proto.Marshal(m1) // len(buf1)=2  两个字节：0x08 0x7F
	buf2, _ = proto.Marshal(m2) // len(buf1)=2  两个字节：0x08 0x7F
	fmt.Printf("protobuf1 i=%v [%v]: %v\n", i, base64.StdEncoding.EncodeToString(buf1), buf1)
	fmt.Printf("protobuf2 i=%v [%v]: %v\n", i, base64.StdEncoding.EncodeToString(buf2), buf2)
}
