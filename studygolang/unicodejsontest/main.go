package main

import "fmt"
import "encoding/json"

/* The php code:

<?php
$s = '{"Name":"魏"}';
echo json_encode(json_decode($s));
// It output : {"Name":"\u9b4f"}
*/

// Go Playground : http://play.golang.org/p/D9uQlFgPJo

func main() {
	fmt.Printf("\n\n")
	var s = make(map[string]interface{})
	s["Name"] = "魏"
	result, _ := json.Marshal(s)
	fmt.Println("json.Marshal ==> ", string(result))

	type User struct {
		Name string
	}

	u := &User{}
	_ = json.Unmarshal(result, u)
	fmt.Printf("json.Unmarshal ==> Name=%v\n", u.Name)

}
