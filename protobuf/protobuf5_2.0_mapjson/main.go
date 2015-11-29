package main

import (
	"fmt"
	proto "github.com/golang/protobuf/proto"
	"encoding/base64"
	"github.com/bitly/go-simplejson"
	"encoding/json"
)

func main() {
	regMessage := &RegMessage{
		Id:       proto.Int32(1001),
		Username: proto.String("Jane"),
		Password: proto.String("123456"),
		Email:    proto.String("jane@163.com"),
	}
	regMessage.Exts = make(map[string]string)
	
	regMessage.Exts["key0"] = string("value0")
	regMessage.Exts["key1"] = string("value1")
	regMessage.Exts["key2"] = string("value2")
	regMessage.Exts["key3"] = string("value3")

	buffer, err := json.Marshal(regMessage)
	sjson, err := simplejson.NewJson(buffer)
	prettyJSON, err := sjson.EncodePretty()
	fmt.Printf("\n\n\n\n%v\n========================\n", string(prettyJSON))

	buffer, err = proto.Marshal(regMessage)
	if err != nil {
		fmt.Printf("failed: %s\n", err.Error())
		return
	}

	fmt.Printf("%v\n\n%v\n\n========================\n", base64.StdEncoding.EncodeToString(buffer), regMessage)

	m := &RegMessage{}
	err = proto.Unmarshal(buffer, m)
	fmt.Printf("name=[%v] email=[%v] id=[%v]\nm=%v\n", m.Username, m.Email, m.Id, m)

	cpp, err := base64.StdEncoding.DecodeString("CHsSB2MrK3Rlc3QaBmMrK3B3ZCIOdGVzdGNwcEAzNjAuY24qFAoKY3BwX3Rlc2V0MRIGMjIyMjIyKhMKCmNwcF90ZXNldDISBXZhbGUyKhMKCmNwcF90ZXNldDMSBXRlc3Qy")

	m = &RegMessage{}
	err = proto.Unmarshal(cpp, m)
	fmt.Printf("\n\nname=[%v] \nemail=[%v] \nid=[%v]\n %v\n", m.Username, m.Email, m.Id, m.GetExts())

}
