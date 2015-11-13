package main

import (
    "os/exec"
    "fmt"
    "os"
)

func main() {
    message := "hello, this is the email body"
    c := exec.Command("echo", message)
    buf, err := c.Output()
    if err != nil {
        fmt.Fprintf(os.Stderr, "The command failed to perform: %s (Command: %s, Arguments: %s)", err, c.Args)
        return
    }
    fmt.Fprintf(os.Stdout, "Result: %s", buf)
}