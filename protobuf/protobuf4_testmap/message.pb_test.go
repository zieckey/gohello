package main

import (
	"testing"
	"fmt"
	proto "github.com/golang/protobuf/proto"
)

/*

$ go test -v -bench=".*"
testing: warning: no tests to run
PASS
Benchmark_Marshal          30000             52471 ns/op
Benchmark_Unmarshal        50000             24482 ns/op
ok      github.com/zieckey/gohello/protobuf/protobuf4_testmap   7.936s

*/

func Benchmark_Marshal(b *testing.B) {
	regMessage := &RegMessage{
		Id:       int32(10001),
		Username: string("vicky"),
		Password: string("123456"),
		Email:    string("eclipser@163.com"),
	}
	regMessage.Exts = make(map[string]string)

	regMessage.Exts["key0"] = string("value0")
	regMessage.Exts["key1"] = string("value1")
	regMessage.Exts["key2"] = string("value2")
	regMessage.Exts["key3"] = string("value3")
	regMessage.Exts["key00"] = string("value0")
	regMessage.Exts["key10"] = string("value1")
	regMessage.Exts["key20"] = string("value2")
	regMessage.Exts["key30"] = string("value3")
	regMessage.Exts["key01"] = string("value0")
	regMessage.Exts["key11"] = string("value1")
	regMessage.Exts["key21"] = string("value2")
	regMessage.Exts["key31"] = string("value3")
	regMessage.Exts["key02"] = string("value0")
	regMessage.Exts["key12"] = string("value1")
	regMessage.Exts["key22"] = string("value2")
	regMessage.Exts["key32"] = string("value3")
	regMessage.Exts["key03"] = string("value0")
	regMessage.Exts["key13"] = string("value1")
	regMessage.Exts["key23"] = string("value2")
	regMessage.Exts["key33"] = string("value3")

	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(regMessage)
		if err != nil {
			fmt.Printf("failed: %s\n", err.Error())
			return
		}
	}
}

func Benchmark_Unmarshal(b *testing.B) {
	regMessage := &RegMessage{
		Id:       int32(10001),
		Username: string("vicky"),
		Password: string("123456"),
		Email:    string("eclipser@163.com"),
	}
	regMessage.Exts = make(map[string]string)

	regMessage.Exts["key0"] = string("value0")
	regMessage.Exts["key1"] = string("value1")
	regMessage.Exts["key2"] = string("value2")
	regMessage.Exts["key3"] = string("value3")
	regMessage.Exts["key00"] = string("value0")
	regMessage.Exts["key10"] = string("value1")
	regMessage.Exts["key20"] = string("value2")
	regMessage.Exts["key30"] = string("value3")
	regMessage.Exts["key01"] = string("value0")
	regMessage.Exts["key11"] = string("value1")
	regMessage.Exts["key21"] = string("value2")
	regMessage.Exts["key31"] = string("value3")
	regMessage.Exts["key02"] = string("value0")
	regMessage.Exts["key12"] = string("value1")
	regMessage.Exts["key22"] = string("value2")
	regMessage.Exts["key32"] = string("value3")
	regMessage.Exts["key03"] = string("value0")
	regMessage.Exts["key13"] = string("value1")
	regMessage.Exts["key23"] = string("value2")
	regMessage.Exts["key33"] = string("value3")

	buffer, err := proto.Marshal(regMessage)
	if err != nil {
		fmt.Printf("failed: %s\n", err.Error())
		return
	}

	for i := 0; i < b.N; i++ {
		m := &RegMessage{}
		err = proto.Unmarshal(buffer, m)
	}
}
