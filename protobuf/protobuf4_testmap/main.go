package main

import (
	"fmt"
	proto "github.com/golang/protobuf/proto"
	//"flag"
	"encoding/base64"
)

func main() {
	regMessage := &RegMessage{
		Id:       proto.Int32(10001),
		Username: proto.String("vicky"),
		Password: proto.String("123456"),
		Email:    proto.String("eclipser@163.com"),
	}
	regMessage.Exts = make([]*RegMessage_ExtsEntry, 2)
	ext := &RegMessage_ExtsEntry{
		Key:proto.String("key0"),
		Value:proto.String("value0"), 
		}
	regMessage.Exts[0] = ext
		ext = &RegMessage_ExtsEntry{
		Key:proto.String("key1"),
		Value:proto.String("value1"), 
		}
	regMessage.Exts[1] = ext
	buffer, err := proto.Marshal(regMessage)
	if err != nil {
		fmt.Printf("failed: %s\n", err.Error())
		return
	}
	
	fmt.Printf("%v\n\n%v\n\n========================\n", base64.StdEncoding.EncodeToString(buffer), regMessage)

	m := &RegMessage{}
	err = proto.Unmarshal(buffer, m)
	fmt.Printf("name=[%v] email=[%v] id=[%v]\nm=%v\n", m.GetUsername(), m.GetEmail(), m.GetId(), m)
	
	cpp, err := base64.StdEncoding.DecodeString("CHsSB2MrK3Rlc3QaBmMrK3B3ZCIOdGVzdGNwcEAzNjAuY24qFAoKY3BwX3Rlc2V0MRIGMjIyMjIyKhMKCmNwcF90ZXNldDISBXZhbGUyKhMKCmNwcF90ZXNldDMSBXRlc3Qy")
	
	m = &RegMessage{}
	err = proto.Unmarshal(cpp, m)
	fmt.Printf("\n\nname=[%v] \nemail=[%v] \nid=[%v]\n %v\n", m.GetUsername(), m.GetEmail(), m.GetId(), m.GetExts())

}
