package main

type HandlerFunc func(msg string)

type Module struct {
	name string
}

func (m *Module) Print(msg string) {
	println(msg)
}

func foo(f HandlerFunc, msg string) {
	f(msg)
	// 问题：这里如何从 f 中获取到 Module 对象的实例，然后可以获取其成员变量  name 的值。
	// m := xxxgetfrom(f)
	// println(m.name)
}

func test1() {
	m := &Module{}
	m.name = "demo"
	foo(m.Print, "hello")
}

func testInterfaceConvertToConcreteType() {
	var a interface{}

	var s = "hello"
	var i = 123
	var buf = []byte(s)

	a = s
	rs := a.(string)
	println(rs, s)

	a = i
	ri := a.(int)
	println(ri, i)

	a = buf
	rbuf := a.([]byte)
	println(rbuf, buf)
}

func main() {
	test1()
	testInterfaceConvertToConcreteType()
}
