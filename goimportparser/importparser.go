package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"path/filepath"
)

var EnvPathSeperator = ":"

func init() {
	if runtime.GOOS == "windows" {
		EnvPathSeperator = ";"
	}
}

func usage() {
	fmt.Printf("Usage : %v GOPATH1 GOPATH2 ...\n", os.Args[0])
}

func main() {
	var paths []string

	if len(os.Args) >= 2 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" || os.Args[1] == "-help" {
			usage()
			return
		}

		paths = os.Args[1:]
	}

	gopath := strings.Split(os.Getenv("GOPATH"), EnvPathSeperator)
	for _, s := range gopath {
		paths = append(paths, string(s))
	}

	fmt.Printf("%v\n", paths)

	var imports map[string]string // the imported libraries
	for _, s := range paths {
		walkdir(s, &imports)
	}
}

func walkdir(dir string, imports *map[string]string) {
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}
		l := len(path)
		if l > 3 && path[l-3:] == ".go" {
			parse(path, imports) // parse one go source file
		}
		return nil
	})
	
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}

func parse(path string, imports *map[string]string) {
	//fmt.Printf("%v\n", path)
//	contents, err := ioutil.ReadFile(filename)
//	if err != nil {
//		return err
//	}
}
