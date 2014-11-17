package pkg

import (
	"log"
	"fmt"
	"encoding/base64"
	"github.com/bitly/go-simplejson"
)

func TestGoSimpleJSON() {
	fmt.Println()
	
	body := `
{
    "employees": [
        {
            "firstName": "Bill",
            "lastName": "Gates"
        },
        {
            "firstName": "George",
            "lastName": "Bush"
        },
        {
            "firstName": "Thomas",
            "lastName": "Carter"
        }
    ],
    "total": 123,
    "dept": "server_dev",
    "obj": {
        "a": 1,
        "B": "bbb",
        "c" : {
        	"cname":"ccccc"
        }
    }
}
	`

    // fmt.Printf("%s\n", string(body))

    js, err := simplejson.NewJson([]byte(body))
    if err != nil {
        log.Fatalln(err)
    }

    total := js.Get("total").MustInt()
	dept := js.Get("dept").MustString()
	obja := js.Get("obj").Get("a").MustInt()
	objb := js.GetPath("obj", "B").MustString()
	objc := js.GetPath("obj", "c").MustString()
	cname := js.GetPath("obj", "c", "cname").MustString()
	
	//获取JSONArray中的元素
	employees := js.Get("employees")
	obj1 := employees.GetIndex(0)
	obj1firstName := obj1.Get("firstName").MustString()
	
    fmt.Printf("total=[%v] dept=[%v] a=[%v] b=[%v] c=[%v] cname=[%v] obj1firstName=[%v]", total, dept, obja, objb, objc, cname, obj1firstName)
    
    buf, err := js.EncodePretty()
    fmt.Printf("=====================\n%v\n==================\n", string(buf))
}

func TestGoSimpleJSONUnicode() {
	fmt.Println()
	
	body := `
{
    "employees": [
        {
            "firstName": "Bill",
            "lastName": "Gates"
        },
        {
            "firstName": "George",
            "lastName": "Bush"
        },
        {
            "firstName": "Thomas",
            "lastName": "Carter"
        }
    ],
    "total": 123,
    "dept": "server_dev",
    "addr": "洛杉矶",
    "obj": {
        "a": 1,
        "B": "bbb",
        "c" : {
        	"cname":"ccccc"
        }
    }
}
	`

    // fmt.Printf("%s\n", string(body))

    js, err := simplejson.NewJson([]byte(body))
    if err != nil {
        log.Fatalln(err)
    }

	addr := js.Get("addr").MustString()
	addrBase64 := base64.StdEncoding.EncodeToString([]byte(addr))
    total := js.Get("total").MustInt()
	dept := js.Get("dept").MustString()
	obja := js.Get("obj").Get("a").MustInt()
	objb := js.GetPath("obj", "B").MustString()
	objc := js.GetPath("obj", "c").MustString()
	cname := js.GetPath("obj", "c", "cname").MustString()
	
	//获取JSONArray中的元素
	employees := js.Get("employees")
	obj1 := employees.GetIndex(0)
	obj1firstName := obj1.Get("firstName").MustString()
	
    fmt.Printf("addr=[%v] [%v] total=[%v] dept=[%v] a=[%v] b=[%v] c=[%v] cname=[%v] obj1firstName=[%v]", addr, addrBase64, total, dept, obja, objb, objc, cname, obj1firstName)
    
    str, err := js.Encode()
    pretty, err := js.EncodePretty()
    fmt.Printf("\n=====================%v\n%v\n==================\n", string(str), string(pretty))
}




