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
	flag.Parse()
	var istream io.Reader
	if len(*input) > 0 {
		istream = strings.NewReader(*input)
	} else {
		istream = bufio.NewReader(os.Stdin)
	}
	body, err := ioutil.ReadAll(istream)
	if err != nil && err != io.EOF {
		fmt.Printf("ERROR : %v\ninput body:\n%v\n", err.Error(), string(body))
		return
	}
	
    js, err := simplejson.NewJson(body)
    if err != nil {
        fmt.Printf("Parse JSON ERROR : %v\ninput body:\n%v\n", err.Error(), string(body))
        return
    }

    pretty, err := js.EncodePretty()
    fmt.Printf("%v\n", string(pretty))
}

