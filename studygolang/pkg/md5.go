package pkg

import (
	"crypto/md5"
	"fmt"
	"encoding/hex"
)

func TestMd5() {
	fmt.Printf("\n\n\n ==================> TestMd5\n")
	md5test1()
	md5test2()
	md5test3()
	md5test4()
	md5test5()
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
	/* WRONG CODE
	fmt.Printf("\n==================> md5test2:\n")
	d1 := []byte("a")
	mh := md5.New()
	m := mh.Sum(d1)
	fmt.Printf("%x %v\n", m, m) // MUST BE : 0cc175b9c0f1b6a831c399e269772661
	hex := fmt.Sprintf("%x", m)
	fmt.Printf("%v", hex)
	*/
}

func md5test3() {
	/* WRONG CODE
	// MD5 ("aa") = 4124bc0a9335c27f086f24ba207a4912	
	fmt.Printf("\n==================> md5test3:\n")
	d1 := []byte("a")
	d2 := []byte("a")
	mh := md5.New()
	mh.Write(d1)	
	m := mh.Sum(d2)
	fmt.Printf("%x %v\n", m, m) // 0cc175b9c0f1b6a831c399e269772661
	hex := fmt.Sprintf("%x", m)
	fmt.Printf("%v", hex)
	*/
}

func md5test4() {
	// MD5 ("aa") = 4124bc0a9335c27f086f24ba207a4912	
	fmt.Printf("\n==================> md5test4:\n")
	d1 := []byte("a")
	d2 := []byte("a")
	mh := md5.New()
	mh.Write(d1)
	mh.Write(d2)
	m := mh.Sum(nil)
	fmt.Printf("%x %v\n", m, m) // 0cc175b9c0f1b6a831c399e269772661
	hex := fmt.Sprintf("%x\t", m)
	fmt.Printf("%v", hex)
}

func md5test5() {
	// MD5 ("abcxyz") = 70fb874a43097a25234382390c0baeb3
	fmt.Printf("\n==================> md5test5:\n")	
	hasher := md5.New()
    hasher.Write([]byte("abc"))
    hasher.Write([]byte("xyz"))
    sum := hasher.Sum(nil)
    fmt.Printf("%x %v", sum, hex.EncodeToString(sum))
}