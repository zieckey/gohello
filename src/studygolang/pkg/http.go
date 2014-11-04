package pkg

import (
	"io/ioutil"
	"net/http"
	"fmt"
)

func HttpGet(url string) string {
	r, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s\n", err.Error()) 
		return ""
	}
	
	b, err := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if err != nil { 
		fmt.Printf("%s\n", err.Error())
		return "" 
	}
	
	return string(b)
}

func TestHttpGet() {
	url := "http://www.so.com/robots.txt"
	s := HttpGet(url)
	fmt.Printf("HttpGet from %s :\n%s\n", url, s)
}

