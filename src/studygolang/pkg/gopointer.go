package pkg

import (
	"fmt"
)

type User struct {
	id   int
	name string
}

func New() *User {
	up := new(User)
	up.id = 1
	up.name = "Jack"
	return up
}

// 

func TestGoPointer() {
	u := New()
	fmt.Printf("The original value, u_ptr=%p value=%v\n", u, u)
	fmt.Printf("pointer assignment to p:\n")
	p := u
	fmt.Printf("The original value, p_ptr=%p value=%v\n", p, p)
	fmt.Printf("The address of pointer p_addr=%p u_addr=%p\n", &p, &u)
	pp := &p
	fmt.Printf("pp_ptr=%p *pp=%p *pp=%v\n", pp, *pp, *pp)
	fmt.Printf("value assignment to u2 and change the value of u2:\n")
	u2 := *u
	u2.name = "Tom"
	fmt.Printf("The original value, u_ptr=%p value=%v\n", u, u)
	fmt.Printf("The original value, u2=%p value=%v\n", &u2, u2)
}

func Teststring() {
	name := "zieckey"
	
}