package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage : %v <old-string> <new-string>\n", os.Args[0])
		fmt.Printf("  e.g.: %v wcpp simcc\n", os.Args[0])
		return
	}

	LookupFiles(".", os.Args[1], os.Args[2])
}

func LookupFiles(dir string, old, new string) error {
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if strings.Contains(path, ".git") {
			return nil
		}

		if !f.IsDir() {
			buf, err := ioutil.ReadFile(path)
			if err == nil {
				buf = bytes.Replace(buf, []byte(old), []byte(new), -1)
				ioutil.WriteFile(path, buf, f.Mode())
			} else {
				fmt.Printf("replace file content %v failed : %v\n", path, err.Error())
			}
		}

		if strings.Contains(path, old) {
			f1 := strings.Replace(path, old, new, -1)
			err = os.Rename(path, f1)
			if err == nil {
				fmt.Printf("rename <%v> to <%v> OK\n", path, f1)
			} else {
				fmt.Printf("rename <%v> to <%v> failed : %v\n", path, f1, err.Error())
			}
		}

		return nil
	})

	return err
}
