package main

import (
	"os/exec"
	"fmt"
	"os"
)

func main() {
	bcCommand := exec.Command("bc", "-l")
	mailStdin, err := bcCommand.StdinPipe()
	err = bcCommand.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "The command failed to perform: %s (Command: %s, Arguments: %s)", err, bcCommand.Path, bcCommand.Args)
		return
	}
	mailStdin.Write("13/3")
	mailStdin.Close()
	bcCommand.Wait()
	buf, err := bcCommand.Output()
	if err == nil {
		fmt.Printf("bc execute OK, the result is [%v]\n", buf)
	} else {
		fmt.Fprintf(os.Stderr, "The command failed to perform: %s (Command: %s, Arguments: %s)", err, bcCommand.Path, bcCommand.Args)
	}
}