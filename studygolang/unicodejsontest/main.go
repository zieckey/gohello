package main

import "log"
//import "encoding/json"
import "github.com/zieckey/gohello/studygolang/unicodejsontest/json"

/* The php code:

<?php
$s = '{"Name":"魏"}';
echo json_encode(json_decode($s));
// It output : {"Name":"\u9b4f"}
*/

// Go Playground : http://play.golang.org/p/D9uQlFgPJo

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	log.Printf("\n\n")
	var s = make(map[string]interface{})
	s["Name"] = "\n\rabc\n魏\n国d\nef"
	result, _ := json.Marshal(s)
	log.Println("json.Marshal ==> ", string(result))

	type User struct {
		Name string
	}

	u := &User{}
	_ = json.Unmarshal(result, u)
	log.Printf("json.Unmarshal ==> Name=%v\n", u.Name)

}
