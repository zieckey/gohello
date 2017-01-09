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
		fmt.Printf("renameall will rename the file name and the content of file with the specified string.\n", os.Args[0])
		fmt.Printf("Usage : %v <old-string> <new-string>\n", os.Args[0])
		fmt.Printf("  e.g.: %v wcpp simcc\n", os.Args[0])
		return
	}

	// 当碰到文件夹的名称也符合rename条件时， filepath.Walk 一次性可能不能搞定，这里就多重试几次。性能没关系，反正都很快，只要把事情搞定就好。
	for i := 0; i < 10; i++ {
		LookupFiles(".", os.Args[1], os.Args[2])
	}
}

func LookupFiles(dir string, old, new string) error {
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if strings.Contains(path, ".git") {
			return nil
		}

		if strings.Contains(path, ".svn") {
			return nil
		}

		if !f.IsDir() {
			buf, err := ioutil.ReadFile(path)
			if err == nil {
				if bytes.Contains(buf, []byte(old)) {
					buf = bytes.Replace(buf, []byte(old), []byte(new), -1)
					err = ioutil.WriteFile(path, buf, f.Mode())
					if err == nil {
						fmt.Printf("replace the file content %v OK\n", path)
					} else {
						fmt.Printf("replace file content %v failed : %v\n", path, err.Error())
					}
				}
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
