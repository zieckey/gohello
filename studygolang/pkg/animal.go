package pkg

import (
	"fmt"
	"studygolang/util"
)

type Animal interface {
	Eat(s string)
	Name() string
	SetName(s string)
}

//////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////
type Cat struct {
	name string
}

func (m Cat) Eat(s string) {
	fmt.Printf("%v - %s\n", util.CallerFuncInfo(), s)
}

func (m Cat) Name() string {
	fmt.Printf("%v name=%s\n", util.CallerFuncInfo(), m.name)
	return m.name
}

func (m Cat) SetName(n string) {
	fmt.Printf("%v - %s\n", util.CallerFuncInfo(), n)
	m.name = n
}


//////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////
type Dog struct {
	name string
}

func (m Dog) Eat(s string) {
	fmt.Printf("%v - %s\n", util.CallerFuncInfo(), s)
}

func (m Dog) Name() string {
	fmt.Printf("%v name=%s\n", util.CallerFuncInfo(), m.name)
	return m.name
}

func (m Dog) SetName(n string) {
	fmt.Printf("%v - %s\n", util.CallerFuncInfo(), n)
	m.name = n
}


//////////////////////////////////////////////////////////
//////////////////////////////////////////////////////////
func TestAnimalInterface() {
	d := Dog{"wangcai"}
	c := Cat{"kitty"}
	
	var a Animal = d
	a.Eat("rice")
	a.SetName("new-name-for" + a.Name())
	
	
	a = c
	a.Eat("rice")
	a.SetName("new-name-for" + a.Name())
}

