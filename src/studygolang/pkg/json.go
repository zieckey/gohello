package pkg

import (
	"encoding/json"
	"fmt"
)

func TestJSON() {
	//key必须是大写字母才会解析出来
	fmt.Printf("\n\n")
	var s = make(map[string]interface{})
	s["Tame"] = "zieckey"
	s["Id"] = 123
	s["Iime"] = "2014-11-11 17:46:00"
	s["Vip"] = true
	result, _ := json.Marshal(s)
	fmt.Println("json.Marshal ==> ", string(result))
	
	type User struct {
		Name string
		Id int
		Time string
		Vip bool
	}
	
	u := &User{}
	_ = json.Unmarshal(result, u)
	fmt.Printf("json.Unmarshal ==> %v %v %v %v\n", u.Name, u.Id, u.Time, u.Vip)
}
