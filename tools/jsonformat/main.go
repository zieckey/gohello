package main 

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"os"
	"bufio"
	"io"
	"flag"
	"io/ioutil"
	"strings"
)

func main() {
	var input *string = flag.String("s", "", "The input json string")
	var compact *bool = flag.Bool("c", false, "Compact output")
	flag.Parse()
	var in io.Reader
	if len(*input) > 0 {
		in = strings.NewReader(*input)
	} else {
		in = bufio.NewReader(os.Stdin)
	}
	body, err := ioutil.ReadAll(in)
	if err != nil && err != io.EOF {
		fmt.Printf("ERROR : %v\ninput body:\n%v\n", err.Error(), string(body))
		return
	}
	
    js, err := simplejson.NewJson(body)
    if err != nil {
        fmt.Printf("Parse JSON ERROR : %v\ninput body:\n%v\n", err.Error(), string(body))
        return
    }

	var json []byte
	if *compact {
		json, err = js.Encode()
	} else {
		json, err = js.EncodePretty()
	}

    fmt.Printf("%v\n", string(json))
}

