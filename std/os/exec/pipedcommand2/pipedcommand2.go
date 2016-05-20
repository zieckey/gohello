package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	wcCommand := exec.Command("wc", "-l")
	stdin, err := wcCommand.StdinPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "StdinPipe Get failed: %s (Command: %s, Arguments: %s)", err, wcCommand.Path, wcCommand.Args)
		return
	}
	stdin.Write([]byte("1111\n22222\n3333\n4444\n"))
	stdin.Close()
	buf, err := wcCommand.Output()
	if err == nil {
		fmt.Printf("%v execute OK, the result is [%v]\n", wcCommand.Path, string(buf))
	} else {
		fmt.Fprintf(os.Stderr, "The command failed to perform: %s (Command: %s, Arguments: %s)", err, wcCommand.Path, wcCommand.Args)
	}
}
