package main 

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"flag"
	"io/ioutil"
	"strings"
	"encoding/base64"
)


/*
base64 decode PE info string
 */
func main() {
	var input *string = flag.String("s", "", "The input string")
	//var encode *bool = flag.Bool("e", false, "Base64 encode")
	flag.Parse()
	var in io.Reader
	if len(*input) > 0 {
		in = strings.NewReader(*input)
	} else if len(os.Args) == 2 {
        in = strings.NewReader(os.Args[1])
    } else if len(os.Args) == 1 {
		in = bufio.NewReader(os.Stdin)
	} else {
        fmt.Printf("ERROR input parameter.\n")
        fmt.Printf("Usage: %v <The-base64-string>\n", os.Args[0])
        return
    }
	body, err := ioutil.ReadAll(in)
	if err != nil && err != io.EOF {
		fmt.Printf("ERROR : %v\ninput body:\n%v\n", err.Error(), string(body))
		return
	}
	
	body, err = base64.StdEncoding.DecodeString(string(body))
	if err != nil {
		fmt.Printf("ERROR : %v\n", err)
		return
	}

	ss := strings.Split(string(body), "\t")
	for _, s := range ss {
        body, err = base64.StdEncoding.DecodeString(strings.TrimSpace(s))
        if err == nil {
            fmt.Printf("%v\n", string(body))
        } else {
            fmt.Printf("ERROR %v\n", err)
        }
    }
}

