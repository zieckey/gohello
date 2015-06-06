package main

type HandlerFunc func(msg string)

type Module struct {
	name string
}

func (m* Module) Print(msg string) {
	println(msg)
}

func foo(f HandlerFunc, msg string) {
	f(msg)
	// 问题：这里如何从 f 中获取到 Module 对象的实例，然后可以获取其成员变量  name 的值。
	// m := xxxgetfrom(f)
	// println(m.name)
}

func main() {
	m := &Module{}
	m.name = "demo"
	foo(m.Print, "hello")
}