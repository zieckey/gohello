package main

import (
	"fmt"
)


type Person struct {
	Id   int
	Name string
}

type Tester interface {
	Test()
}

func (this *Person) Test() {
	fmt.Println("\tthis =", &this, "Person.Test")
}

// Employee 从Person继承
type Employee struct {
	Person
}

func (this *Employee) Test() {
	fmt.Println("\tthis =", &this, "Employee.Test")
	this.Person.Test() // 调用父类的被覆盖的方法
}

func main() {
	fmt.Println("An Employee instance :")
	var nu Employee
	nu.Id = 2
	nu.Name = "NTom"
	nu.Test()
	fmt.Println()
	
	fmt.Println("A Tester interface to Employee instance :")
	var t Tester
	t = &nu
	t.Test()
	fmt.Println()
	
	fmt.Println("A Tester interface to Person instance :")
	t = &nu.Person
	t.Test()
	/* Output:
An Employee instance :
	this = 0xc082024020 Employee.Test
	this = 0xc082024028 Person.Test

A Tester interface to Employee instance :
	this = 0xc082024030 Employee.Test
	this = 0xc082024038 Person.Test

A Tester interface to Person instance :
	this = 0xc082024040 Person.Test
	*/
}