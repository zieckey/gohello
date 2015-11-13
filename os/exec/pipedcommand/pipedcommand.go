package main

import (
	"os/exec"
	"fmt"
	"os"
)

func main() {
	message := "hello, this is the email body"
	echoCommand := exec.Command("echo", message)
	buf, err := echoCommand.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "The command failed to perform: %s (Command: %s, Arguments: %s)", err, echoCommand.Path, echoCommand.Args)
		return
	}

	mailCommand := exec.Command("mail", "-s", message, "weizili@360.cn")
	mailStdin, err := mailCommand.StdinPipe()
	mailStdin.Write(buf)
	err = mailCommand.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "The command failed to perform: %s (Command: %s, Arguments: %s)", err, mailCommand.Path, mailCommand.Args)
		return
	}
	mailCommand.Wait()
	mailStdin.Close()
}