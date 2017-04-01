package main
import (
	"path/filepath"
	"os"
	"fmt"
	"strings"
)

func main() {

	file_name_prefix := "lib"

	files, err := LookupFiles("d:/cppworkspace/boost_1_58_0-bak/output/lib", "*.lib")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	for _, f := range files {
		fmt.Printf("processing %v\n", f)
		fname := filepath.Base(f)
		dir := filepath.Dir(f)
		if strings.Index(fname, file_name_prefix) == 0 {
			fname = strings.Replace(fname, file_name_prefix, "", 1)
			nf := filepath.Join(dir, fname)
			err = os.Rename(f, nf)
			if err == nil {
				fmt.Printf("rename <%v> to <%v> OK\n", f, nf)
			} else {
				fmt.Printf("rename <%v> to <%v> failed : %v\n", f, nf, err.Error())
			}
		}
	}
}


func LookupFiles(dir string, pattern string) ([]string, error) {
	var files []string = make([]string, 0, 5)

	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}

		if f.IsDir() {
			return nil
		}

		if ok, err := filepath.Match(pattern, f.Name()); err != nil {
			return err
		} else if ok {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}
