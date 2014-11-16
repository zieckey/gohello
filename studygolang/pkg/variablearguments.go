package pkg

import (
	"path/filepath"
	"fmt"
	"os"
)

func TestVariableNumberArguments() {
	path := joinVariableNumberArguments("a","b","c")
	fmt.Printf("\n\n\n\npath=%v\n", path)
	if path != string("a") + string(os.PathSeparator) + "b" + string(os.PathSeparator) + "c" {
		fmt.Printf("Error !!!!\n")
	}
}


func joinVariableNumberArguments(branch ...string) string {
	var path string
	for _, p := range branch {
		if len(path) == 0 {
			path = p
		} else {
			path = filepath.Join(path, p)
		}
	}
	return path
}

