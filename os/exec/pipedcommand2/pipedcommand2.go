package main

import (
	"os/exec"
	"fmt"
	"os"
)

func main() {
	wcCommand := exec.Command("wc", "-l")
	mailStdin, err := wcCommand.StdinPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "StdinPipe Get failed: %s (Command: %s, Arguments: %s)", err, wcCommand.Path, wcCommand.Args)
		return
	}
	mailStdin.Write([]byte("1111\n22222\n3333\n4444\n"))
	mailStdin.Close()
	buf, err := wcCommand.Output()
	if err == nil {
		fmt.Printf("%v execute OK, the result is [%v]\n", wcCommand.Path, string(buf))
	} else {
		fmt.Fprintf(os.Stderr, "The command failed to perform: %s (Command: %s, Arguments: %s)", err, wcCommand.Path, wcCommand.Args)
	}
}