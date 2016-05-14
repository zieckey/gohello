package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"encoding/base64"
)

// Protobuf各个整型类型之间转换之后的序列化结果是一样的
// 因此一般来说将类型定义为int64比较好。备注：Java中没有uint64，int64也足够大。

func main() {
	t := IntType1 {
		T1:1,
		T2:2,
		T3:3,
		T4:4,
		T5:5,
		T6:6,
		T7:7,
		T8:8,
		T9:9,
	}

	tbuf, err := proto.Marshal(&t)
	if err != nil {
		fmt.Printf("failed: %s\n", err.Error())
		return
	}

	fmt.Printf("%v %v\n", base64.StdEncoding.EncodeToString(tbuf), tbuf)

	m := IntType2{}
	err = proto.Unmarshal(tbuf, &m)
	if m.T1 != int64(t.T1) || m.T2 != int64(t.T2) ||
	   m.T3 != int64(t.T3) || m.T4 != int64(t.T4) ||
	   m.T5 != int64(t.T5) || m.T6 != int64(t.T6) ||
	   m.T7 != int64(t.T7) || m.T8 != int64(t.T8) ||
	   m.T9 != int64(t.T9) {
		fmt.Printf("ERROR!!!!!!!\n")
	}

	mbuf, err := proto.Marshal(&m)
	if err != nil {
		fmt.Printf("failed: %s\n", err.Error())
		return
	}

	if base64.StdEncoding.EncodeToString(tbuf) == base64.StdEncoding.EncodeToString(mbuf) {
		fmt.Printf("Transform OK\n")
	}

	fmt.Printf("t=%v\n", t)
	fmt.Printf("m=%v\n", m)

	// 整体输出
	/*
	CAEQAhgDIAQoBTAGOAdACEgJ [8 1 16 2 24 3 32 4 40 5 48 6 56 7 64 8 72 9]
	Transform OK
	t={1 2 3 4 5 6 7 8 9}
	m={1 2 3 4 5 6 7 8 9}
	 */
}
