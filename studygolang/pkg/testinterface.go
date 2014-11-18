package pkg

import (
	"fmt"
	"io"
	"os"
)

type Person struct {
	name string
	age  int
}

func (p Person) printMsg() {
	fmt.Printf("I am %s, and my age is %d.\n", p.name, p.age)
}

func (p Person) eat(s string) {
	fmt.Printf("%s is eating %s ...\n", p.name, s)
}

func (p Person) drink(s string) {
	fmt.Printf("%s is drinking %s ...\n", p.name, s)
}

type People interface {
	printMsg()
	PeopleEat //组合
	PeopleDrink
	//eat() //不能出现重复的方法
}

/*
//与上面等价
type People interface {
	printMsg()
	eat()
	drink()
}
*/
type PeopleDrink interface {
	drink(s string)
}

type PeopleEat interface {
	eat(s string)
}

type PeopleEatDrink interface {
	eat(s string)
	drink(s string)
}

//以上 Person 类[型]就实现了 People/PeopleDrink/PeopleEat/PeopleEatDrink interface 类型

type Foodie struct {
	name string
}

func (f Foodie) eat(s string) {
	fmt.Printf("I am foodie, %s. My favorite food is the %s.\n", f.name, s)
}

//Foodie 类实现了 PeopleEat interface 类型

func TestInterface() {
	//定义一个 People interface 类型的变量p1
	var p1 People
	p1 = Person{"Rain", 23}
	p1.printMsg()            //I am Rain, and my age is 23.
	p1.drink("orange juice") //print result: Rain is drinking orange juice

	//同一类可以属于多个 interface, 只要这个类实现了这个 interface中的方法
	var p2 PeopleEat
	p2 = Person{"Sun", 24}
	p2.eat("chaffy dish") //print result: Sun is eating chaffy dish ...

	//不同类也可以实现同一个 interface
	var p3 PeopleEat
	p3 = Foodie{"James"}
	p3.eat("noodle") //print result: I am foodie, James. My favorite food is the noodle

	//interface 赋值
	p3 = p1 //p3 中的方法会被 p1 中的覆盖
	p3.eat("noodle")
	/************************************/
	/*print result                      */
	/*Rain is eating noodle ...         */
	/************************************/

	// interface 查询
	// 将(子集) PeopleEat 转为 People 类型
	if p4, ok := p2.(People); ok {
		p4.drink("water") //调用 People interface 中有而 PeopleEat 中没有的方法
		fmt.Println(p4)
	}
	/************************************/
	/* print result                      */
	/* Sun is drink water ...            */
	/* {Sun 24}                          */
	/************************************/

	//查询 p2 是否为 Person 类型变量
	if p5, ok := p2.(Person); ok {
		fmt.Println(p5, "type is Person")
		p5.drink("***") //此时也可以调用 Person 所有的方法
	}
	/************************************/
	/*print result                      */
	/*{Sun 24} type is Person           */
	/*Sun is drink *** ...              */
	/************************************/

	var p6 PeopleEat = Foodie{"Tom"}

	if p7, ok := p6.(People); ok {
		fmt.Println(p7)
	} else {
		fmt.Println("Error: can not convert")
	}
	//result: Error: can not convert

	if p8, ok := p6.(Foodie); ok {
		fmt.Println(p8, "type is Foodie")
	}
	//result: {Tom} type is Foodie
}

//////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////////////////

type UUser struct {
	Id   int
	Name string
}

type Tester interface {
	Test()
}

func (this *UUser) Test() {
	fmt.Println(&this)
}
func t11() Tester {
	var x *UUser = nil
	return x
}

func TestInterface1() {
	var t Tester = &UUser{1, "Tom"}
	if x := t11(); x == nil {
		fmt.Println("t11() return nil") // 此值不输出
	} else {
		fmt.Printf("t11() return not nil, it is [%v]\n", x) // TODO 这里是为什么，返回的难度是interface nil？？？？？？
	}
	fmt.Println(t)
}

//////////////////////////////////////////

func TestEmptyInterface2() {
	var r io.Reader
	path := "tmp/test.create.file.exe"
	tty, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC, 0)
	defer func() {
		tty.Close()
		os.Remove(path)
	}()
	if err != nil {
		fmt.Printf("open file failed [%v] error=[%v]\n", path, err)
		return
	}
	r = tty

	var w, w2 io.Writer
	w = r.(io.Writer)
	w2 = tty

	var empty interface{}
	empty = w

	fmt.Printf("interface empty=[%#v] w=[%#v] w2=[%#v] r=[%#v]\n", empty, w, w2, r)
}

func TestEmptyInterface3() {
	type Any interface{}
	var i = 5
	var str = "abc"
	var val Any
	val = i
	fmt.Printf("val has the value: %v\n", val)
	val = str
	fmt.Printf("val has the value: %v\n", val)
	pers1 := new(Person)
	pers1.name = "Rob Pike"
	pers1.age = 55
	val = pers1
	fmt.Printf("val has the value: %v\n", val)
	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person:
		fmt.Printf("Type pointer to Person %T\n", *t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}

	/*
		Output: val has the value: 5
		val has the value: ABC
		val has the value: &{Rob Pike 55}
		Type pointer to Person main.Person
	*/
}
