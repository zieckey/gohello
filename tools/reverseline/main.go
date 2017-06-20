package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"flag"
	"io/ioutil"
	"strings"
)

/*
将字符串按特定分隔符分割后逆序存放
 */
func main() {
	var sep *string = flag.String("sep", ",", "The new seperator when joins the multi lines to one line.")
	flag.Parse()
	in := bufio.NewReader(os.Stdin)
	body, err := ioutil.ReadAll(in)
	if err != nil && err != io.EOF {
		fmt.Printf("ERROR : %v\ninput body:\n%v\n", err.Error(), string(body))
		return
	}

	s := string(body)

	ss := strings.Split(s, *sep)
	l := len(ss)
	for i, _ := range ss {
		if i != 0 {
			os.Stdout.WriteString(*sep)
		}
		os.Stdout.WriteString(ss[l - i - 1])
	}
}

