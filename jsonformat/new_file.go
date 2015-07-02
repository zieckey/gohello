package main 

import (
	"fmt"
	"github.com/bitly/go-simplejson"
	"os"
	"bufio"
	"io/ioutil"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	body, err := ioutil.ReadAll(stdin)
	if err != nil {
		fmt.Printf("ERROR : %v \n", err.Error())
		return
	}
	
    js, err := simplejson.NewJson(body)
    if err != nil {
        fmt.Printf("ERROR : %v \n", err.Error())
        return
    }

    pretty, err := js.EncodePretty()
    fmt.Printf("%v\n", string(pretty))
}

