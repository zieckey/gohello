package main


import (
	"github.com/mitchellh/cli"
	"fmt"
)

type ProxyCommand struct {}
type StandaloneCommand struct {}

func (c *ProxyCommand) Help() string {
	return "ProxyCommand help xxx"
}

func (c *ProxyCommand) Run(args []string) int {
	fmt.Printf("ProxyCommand Running ...\n")
	return 0
}

func (c *ProxyCommand) Synopsis() string {
	return "ProxyCommand Synopsis Text"
}


func main() {
	command := new(ProxyCommand)
	c := &cli.CLI{
		Args: []string{"foo", "-bar", "-baz"},
		Commands: map[string]cli.CommandFactory{
			"foo": func() (cli.Command, error) {
				return command, nil
			},
		},
	}

	c.Run()
}
