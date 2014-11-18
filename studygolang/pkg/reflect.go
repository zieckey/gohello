package pkg

import (
	"encoding/xml"
	"fmt"
	"reflect"
)

func TestReflect() {
	TestReflect1()
	TestReflect2()
	TestReflect3()
	TestReflect4()
	TestReflect5()
	TestReflect6()
}

func TestReflect1() {

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

type st struct {
}

func (this *st) Echo() {
	fmt.Println("echo()")
}

func (this *st) Echo2() {
	fmt.Println("echo--------------------()")
}

var xmlstr string = `<root>
<func>Echo</func>
<func>Echo2</func>
</root>`

type st2 struct {
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

//////////////////////////////////////////////////////////
type MyStruct struct {
	name string
}

func (this *MyStruct) GetName() string {
	return this.name
}


func TestReflect3() {
	fmt.Println("\n")
	s := "this is string"
	fmt.Println(reflect.TypeOf(s))
	fmt.Println("-------------------")

	fmt.Println(reflect.ValueOf(s))
	var x float64 = 3.4
	fmt.Println(reflect.ValueOf(x))
	fmt.Println("-------------------")

	a := new(MyStruct)
	a.name = "yejianfeng"
	typ := reflect.TypeOf(a)

	fmt.Println(typ.NumMethod())
	fmt.Println("-------------------")

	b := reflect.ValueOf(a).MethodByName("GetName").Call([]reflect.Value{})
	fmt.Println(b[0])
}

////////////////////////////////////////////////////
func TestReflect4() {
	fmt.Printf("\n\n")
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))   //这个程序打印    type: float64
	fmt.Println("value:", reflect.ValueOf(x)) //这个程序打印    value: <float64 Value>
}

func TestReflect5() {
	fmt.Printf("\n\n")
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())
}

func TestReflect6() {
	fmt.Printf("\n\n\nTestReflect6:\n")
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	t := reflect.TypeOf(x)
	fmt.Println("reflect.ValueOf(x).type:", v.Type())
	fmt.Println("reflect.ValueOf(x).kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("reflect.ValueOf(x).value:", v.Float())
	fmt.Println("reflect.TypeOf(x):", t)
}
