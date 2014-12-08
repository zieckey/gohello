package pkg

import (
	"fmt"
)

func StringRune() {
	fmt.Printf("\n\nStringRune:\n")

	str := "weigo老魏"
	fmt.Printf("len=%v\n", len(str))
	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	for i, s := range str {
		fmt.Println(i, "Unicode(", s, ") string=", string(s))
	}

	/*Output:
	rune= [119 101 105 103 111 32769 39759]  len(rune)= 7
	r[ 0 ]= 119 string= w
	r[ 1 ]= 101 string= e
	r[ 2 ]= 105 string= i
	r[ 3 ]= 103 string= g
	r[ 4 ]= 111 string= o
	r[ 5 ]= 32769 string= 老
	r[ 6 ]= 39759 string= 魏
	*/
	r := []rune(str)
	fmt.Println("rune=", r, " len(rune)=", len(r))
	for i := 0; i < len(r); i++ {
		fmt.Println("r[", i, "]=", r[i], "string=", string(r[i]))
	}

	b := []byte(str)
	fmt.Println("byte=", r)
	for i := 0; i < len(b); i++ {
		fmt.Println("b[", i, "]=", b[i], "string=", string(b[i]))
	}
}
