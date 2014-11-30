package gogit

import (
    "testing"
    "os"
)

func TestgetSrcPath(t *testing.T) {
	println(getSrcPath())
}


func Test1(t *testing.T) {
	path := GetProjectParentPath("https://github.com/zieckey111/11")
	err := os.MkdirAll(path, 0755)
	if err != nil {
		println(err.Error())
	} else {
		println("GetProjectParentPath: ", path)
	}
}

