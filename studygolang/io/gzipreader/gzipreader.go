package main

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Printf("++++%v++++", "1123")
	pattern := "./*"
	if len(os.Args) == 2 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "-help" {
			fmt.Printf("usage : %v <the file pattern>\n", os.Args[0])
			return
		}
		pattern = os.Args[1]
	}

	files, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Printf("error happened : %v\n", err.Error())
		return
	}
	for _, f := range files {
		r, err := OpenFile(f)
		if err != nil {
			fmt.Printf("read file error : %v\n", err.Error())
			continue
		}

		fmt.Printf("=====> read file %v \n", f)
		for {
			line, err := r.ReadString('\n')
			line = strings.Trim(line, "\n\r\t ")
			fmt.Printf("+++ %s +++\n", line)
			if err != nil {
				break
			}	
		}
	}

}

func OpenFile(name string) (*bufio.Reader, error) {
	buf, err := ioutil.ReadFile(name)
	if err != nil {
		fmt.Printf("read file error: %v\n", err.Error())
		return nil, err
	}

	r := bytes.NewReader(buf)
	if strings.HasSuffix(name, ".gz") {
		gr, err := gzip.NewReader(r)
		if err != nil {
			return nil, err
		}
		return bufio.NewReader(gr), nil
	}

	return bufio.NewReader(r), nil
}
