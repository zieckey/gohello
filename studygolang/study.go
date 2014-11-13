package main

import "fmt"
import "math"
import "github.com/zieckey/gohello/studygolang/pkg"
import "github.com/zieckey/gohello/studygolang/util"


func main() {
    StudyForLoop()
    FizzBuzzConstIfElse()
    StringFunc()
    StringRune()
    TestDefer()
    TestPanicRecover()
    pkg.TestInterface()
    pkg.TestAnimalInterface()
	util.TestCallerName()
	pkg.TestHttpGet()
	pkg.TestGoPointer()
	pkg.TestArray()
	pkg.TestGobBinary()
	pkg.TestCustomEncodeDecode()
	pkg.TestInterfaceSerialize()
	pkg.TestInterface1()
	pkg.TestEmptyInterface2()
	pkg.TestFileReadAndWrite()
	pkg.TestReflect()
	pkg.TestJSON()
	pkg.TestGoSimpleJSON()
}





func StudyForLoop() {
	fmt.Printf("for loop:")
	for i := 0; i < 10; i++ {
		fmt.Printf("i=%v ", i)
	}
	
	fmt.Printf("\n\ngoto loop:")
	i := 0
GotoLoop:
	fmt.Printf("i=%d", i)
	fmt.Printf(" ")
	i++
	if i < 10 {
		goto GotoLoop
	}
	
	fmt.Printf("\n\narray iterator:")
	a := [...]int{10, 11, 22, 33, 44, 55, 66, 77, 88}
	for index, v := range a {
		fmt.Printf("index=%v,v=%v ", index, v)
	}
}

func FizzBuzzConstIfElse() {
	const (
		FIZZ = 3
		BUZZ = 5
	)
	fmt.Printf("\n\nFizzBuzz: ")
	for i := 0; i < 100; i++ {
		if i % FIZZ == 0 && i % BUZZ == 0 {
			fmt.Printf(" fizzbuzz")
		} else if i % FIZZ == 0 {
			fmt.Printf(" fizz")
		} else if i % BUZZ == 0 {
			fmt.Printf(" buzz")
		} else {
			fmt.Printf(" %d", i)
		}
	}
	fmt.Println("")
}

func StringFunc() {
	fmt.Printf("\n\nStringFunc:  ")
	s := "asSASA"
	//b := []byte(s)
	l := 0
	for i,v := range s {
		l++
		fmt.Printf("%d=%c ", i, v)
		
	}
	fmt.Printf("   len(s)=%d    buildin-len(s)=%d", l, len(s))
}

func StringRune() {
	fmt.Printf("\n\nStringRune:\n")
	
	str :="weigo老魏"
	fmt.Printf("len=%v\n", len(str))
	for  i:=0;i<len(str);i++ {
	    fmt.Println(str[i])
	}
	
	for  i,s :=  range str {
	    fmt.Println(i,"Unicode(",s,") string=",string(s))
	}
	
	r := []rune(str)
	fmt.Println("rune=",r, " len(rune)=", len(r))
	for i:=0;i<len(r) ; i++ {
	 	fmt.Println("r[",i,"]=",r[i],"string=",string(r[i]))
	}

	b := []byte(str)
	fmt.Println("byte=",r)
	for i:=0;i<len(b) ; i++ {
	 	fmt.Println("b[",i,"]=",b[i],"string=",string(b[i]))
	}
}

func TestDefer() {
	fmt.Printf("\n\nTestDefer")
	for i := 0; i < 5; i++ {
		defer fmt.Printf(" i=%v", i)
	}
}

func ConvertInt64ToInt(x int64) int{
	if math.MinInt32<=x && x<=math.MaxInt32 {
		return int(x)
	}
	panic(fmt.Sprintf("%d is out of the int32 range",x))
}

func IntFromInt64(x int64)(i int,err error) {
	defer func(){
		if e := recover(); e != nil {
			//err=fmt.Errorf("%v",e)
			fmt.Printf("%v",e)
		}
	}()
	i=ConvertInt64ToInt(x)
	return i,nil
}

func TestPanicRecover() {
	fmt.Printf("\n\nTestPanicRecover:\n")
	IntFromInt64(1)
	IntFromInt64(1099511627775)
	IntFromInt64(2)
}
