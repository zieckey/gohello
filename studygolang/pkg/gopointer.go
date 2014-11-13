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
	//name := "zieckey"
	
}

func TestArray() {
	TestArray1()
	TestArray2()
}

func modifyArray1( args [4]int ){
    args[1] = 100;
}
 
func TestArray1() {
	fmt.Println()
	var args = [4]int{1, 2, 3, 4};
    modifyArray1(args);
    fmt.Println(args);//输出結果是 [1 2 3 4], modifyArray1产生的效果没有保存到结果中，表面数组是按值传递的，也就是说函数参数传递时要复制一份 
}

func modifyArray2( args * [4]int ){
    args[1] = 100;
}
 
func TestArray2() {
	fmt.Println()
	var args = [4]int{1, 2, 3, 4};
    modifyArray2(&args);
    fmt.Println(args);//输出結果是 [1 100 3 4], modifyArray2产生了效果。表面必须按照数组的指针传递 
}

