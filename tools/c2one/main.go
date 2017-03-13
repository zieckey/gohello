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
change a multi-line text input into one line.
 */
func main() {
	var new_sep *string = flag.String("sep", ",", "The new seperator when joins the multi lines to one line.")
	flag.Parse()
	in := bufio.NewReader(os.Stdin)
	body, err := ioutil.ReadAll(in)
	if err != nil && err != io.EOF {
		fmt.Printf("ERROR : %v\ninput body:\n%v\n", err.Error(), string(body))
		return
	}

	s := string(body)
	s = strings.Replace(s, "\n", *new_sep, -1)

	println(s)
}

