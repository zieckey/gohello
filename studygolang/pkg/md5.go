package pkg

import (
	"crypto/md5"
	"fmt"
)

func TestMd5() {
	fmt.Printf("\n\n\n ==================> TestMd5\n")
	md5test1()
	md5test2()
	md5test3()
}

func md5test1() {
	fmt.Printf("\n ==================> md5test1:\n")
	d := []byte("a")
	m := md5.Sum(d)
	fmt.Printf("%x %v\n", m, m) // 0cc175b9c0f1b6a831c399e269772661
	hex := fmt.Sprintf("%x", m)
	fmt.Printf("%v", hex)
}

func md5test2() {
	fmt.Printf("\n==================> md5test2:\n")
	d1 := []byte("a")
	mh := md5.New()
	m := mh.Sum(d1)
	fmt.Printf("%x %v\n", m, m) // MUST BE : 0cc175b9c0f1b6a831c399e269772661
	hex := fmt.Sprintf("%x", m)
	fmt.Printf("%v", hex)
}

func md5test3() {
	fmt.Printf("\n==================> md5test3:\n")
	d1 := []byte("a")
	d2 := []byte("a")
	mh := md5.New()
	mh.Write(d1)	
	m := mh.Sum(d2)
	fmt.Printf("%x %v\n", m, m) // 0cc175b9c0f1b6a831c399e269772661
	hex := fmt.Sprintf("%x", m)
	fmt.Printf("%v", hex)
}