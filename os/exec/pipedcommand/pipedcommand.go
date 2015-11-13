package main

import (
	"os/exec"
	"fmt"
	"os"
)

func main() {
	title := "hello, mail title"
	message := "hello, this is the email body"
	mailCommand := exec.Command("mail", "-s", title, "weizili@360.cn")
	mailStdin, err := mailCommand.StdinPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "StdinPipe failed to perform: %s (Command: %s, Arguments: %s)", err, mailCommand.Path, mailCommand.Args)
		return
	}
	mailStdin.Write([]byte(message))
	mailStdin.Close()
	mailCommand.Output()
	fmt.Println(mailCommand.ProcessState.String())
}