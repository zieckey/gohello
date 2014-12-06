package pkg

import (
	"fmt"
)

type tree struct {
	name string
	id   int
}

func testStructCopyParameter(t1 tree) {
	//这里传参调用过程中进行了数据拷贝，因此下面的修改不会影响外面的值
	t1.id = 2
	t1.name = "tree2"
	fmt.Printf("testStructCopyParameter:\nt1=%v &t1=%p\n", t1, &t1)
}

func testStructCopy2() {
	var t1 tree
	t1.name = "tree1"
	t1.id = 1
	
	testStructCopyParameter(t1)
	fmt.Printf("testStructCopy2:\nt1=%v &t1=%p\n", t1, &t1)
}

func testStructCopy1() {
	var t1 tree
	t1.name = "tree1"
	t1.id = 1

	t2 := t1

	fmt.Printf("testStructCopy1:\nt1=%v t2=%v &t1=%p &t2=%p\n", t1, t2, &t1, &t2)

	t2.id = 2
	t2.name = "tree2"

	fmt.Printf("t1=%v t2=%v\n", t1, t2)
}

/* Output:
testStructCopy1:
t1={tree1 1} t2={tree1 1} &t1=0xc08207f7e0 &t2=0xc08207f800
t1={tree1 1} t2={tree2 2}
testStructCopyParameter:
t1={tree2 2} &t1=0xc08207f8e0
testStructCopy2:
t1={tree1 1} &t1=0xc08207f8c0
*/
func TestStructCopy() {
	fmt.Printf("\n\n")
	testStructCopy1()
	testStructCopy2()
}
