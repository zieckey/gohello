package main

import (
	"github.com/zieckey/gohello/gogit/gogit"
	"os"
	"os/exec"
	"path"
)

func usage() {
	println("Usage : ", os.Args[0], " <the github project clone url>")
	println("  e.g.: ", os.Args[0], " https://github.com/zieckey/goini")
}

/*
ERROR !!!!!!

NOT FINISHED!!!

$ ./gogit.exe https://github.com/ant0ine/go-urlrouter
cmd 1:  git clone https://github.com/ant0ine/go-urlrouter
cmd 2:  mkdir -p e:/goworkspace/go/src/github.com/ant0ine
cmd 3:  mv go-urlrouter e:/goworkspace/go/src/github.com/ant0ine
exec command failed :  git clone https://github.com/ant0ine/go-urlrouter exec: "git clone https://github.com/ant0ine/go-urlrouter": file does not exist
*/

func main() {
	if len(os.Args) != 2 {
		usage()
		return
	}

	url := os.Args[1]
	dir := gogit.GetProjectParentPath(url)
	project := path.Base(url)
	c1 := "git clone " + url
	c2 := "mkdir -p " + dir
	c3 := "mv " + project + " " + dir
	println("cmd 1: ", c1)
	println("cmd 2: ", c2)
	println("cmd 3: ", c3)
	runCmd(c1, c2, c3)

}

func runCmd(cmd ...string) {
	
	for _, c := range cmd {
		cmd := exec.Command(c)
		err := cmd.Run()
		if err != nil {
			println("exec command failed : ", c, err.Error())
			break
		}
	}
}
