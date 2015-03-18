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
	Eat()
}

func (this *Person) Test() {
	fmt.Println("\tthis =", &this, "Person.Test")
}

func (this *Person) Eat() {
	fmt.Println("\tthis =", &this, "Person.Eat")
}

// Employee 从Person继承, 直接继承了 Eat 方法，并且将 Test 方法覆盖了。
type Employee struct {
	Person
}

func (this *Employee) Test() {
	fmt.Println("\tthis =", &this, "Employee.Test")
	this.Person.Test() // 调用父类的方法，且该方法被子类覆盖了
}

func main() {
	fmt.Println("An Employee instance :")
	var nu Employee
	nu.Id = 2
	nu.Name = "NTom"
	nu.Test()
	nu.Eat()
	fmt.Println()
	
	fmt.Println("A Tester interface to Employee instance :")
	var t Tester
	t = &nu
	t.Test()
	t.Eat()
	fmt.Println()
	
	fmt.Println("A Tester interface to Person instance :")
	t = &nu.Person
	t.Test()
	t.Eat()
	/* Output:
An Employee instance :
	this = 0xc082024020 Employee.Test
	this = 0xc082024028 Person.Test
	this = 0xc082024030 Person.Eat

A Tester interface to Employee instance :
	this = 0xc082024038 Employee.Test
	this = 0xc082024040 Person.Test
	this = 0xc082024048 Person.Eat

A Tester interface to Person instance :
	this = 0xc082024050 Person.Test
	this = 0xc082024058 Person.Eat
	*/
}