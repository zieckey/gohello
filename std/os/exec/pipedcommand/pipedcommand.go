package main

import (
	"fmt"
	"os"
	"os/exec"
)

// It works.
func main() {
	title := "title"
	message := "body 1111111111111111111"
	mailCommand := exec.Command("mail", "-s", title, "weizili@360.cn")
	stdin, err := mailCommand.StdinPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "StdinPipe failed to perform: %s (Command: %s, Arguments: %s)", err, mailCommand.Path, mailCommand.Args)
		return
	}
	stdin.Write([]byte(message))
	stdin.Close()
	mailCommand.Output()
	fmt.Println(mailCommand.ProcessState.String())
}
