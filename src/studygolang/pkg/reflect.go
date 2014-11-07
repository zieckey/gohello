package pkg

import (
	"fmt"
	"reflect"
	"encoding/xml"
)

func TestReflect() {

	type T struct {
		A int
		B string
	}
	t := T{12, "hello"}
	reflectedType := reflect.ValueOf(&t).Elem()
	types := reflectedType.Type()
	for i := 0; i < reflectedType.NumField(); i++ {
		f := reflectedType.Field(i)
		fmt.Printf("Reflecting index=%d : %s %s = %v\r\n", i, types.Field(i).Name, f.Type(), f.Interface())
	}
}


/////////////////////////////////////////

type st struct{
}

func (this *st)Echo(){
    fmt.Println("echo()")
}

func (this *st)Echo2(){
    fmt.Println("echo--------------------()")
}

var xmlstr string=`<root>
<func>Echo</func>
<func>Echo2</func>
</root>`

type st2 struct{
    E []string `xml:"func"`
}

//TODO 不太懂？？？？
//利用golang的反射包，实现根据函数名自动调用函数。  
func TestReflect2() {
    s2 := st2{}
    xml.Unmarshal([]byte(xmlstr), &s2)

    s := &st{}
    v := reflect.ValueOf(s)
  
    v.MethodByName(s2.E[1]).Call(nil)
 }

