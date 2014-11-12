package pkg

import (
	"log"
	"fmt"
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
        "B": "bbb"
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
	
	//获取JSONArray中的元素
	employees := js.Get("employees")
	obj1 := employees.GetIndex(0)
	obj1firstName := obj1.Get("firstName").MustString()
	
    fmt.Printf("total=[%v] dept=[%v] a=[%v] obj1firstName=[%v]", total, dept, obja, obj1firstName)
}


