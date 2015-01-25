package main

import (
	"fmt"
	"path/filepath"
	"os"
)

func main() {
	
	if len(os.Args) == 2 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "-help" {
			fmt.Printf("usage : %v <the file pattern>\n", os.Args[0])
			return
		}
		//pattern = os.Args[1]
	}

	GlobTest()
}

func GlobTest() ([]string, error) {
	pattern := "E:/goworkspace/go/*/*/*/*.go"
	files, err := filepath.Glob(pattern)
	fmt.Printf("%v\n", files)
	for i, f := range files {
		fmt.Printf("%v\t%v\n", i, f)
	}
	
	return files, err
}
