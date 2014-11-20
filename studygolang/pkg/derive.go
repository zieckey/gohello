package pkg

import (
	"fmt"
)

//see : testinterface.go
//
//type Person struct {
//	name string
//	age  int
//}

type Women struct {
	Person
	shopping int
}

func (this Women)eat(s string) {
	fmt.Printf("Women %s is eating %s ...\n", this.name, s)
}

func TestDerive() {
	w := Women{
		shopping:100,
		}
	w.name = "jane"
	w.age = 30
	w.eat("apple")
	w.drink("water")
}