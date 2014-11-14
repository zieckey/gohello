package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmds := [...]string{"go get github.com/bitly/go-simplejson"}

	for i := 0; i < len(cmds); i++ {
		cmd := cmds[i]
		_, err := exec.Command(cmd).Output()
		if err != nil {
			fmt.Printf("Error exec [%v] [%v]\n", cmds, err.Error())
		}
	}
}
