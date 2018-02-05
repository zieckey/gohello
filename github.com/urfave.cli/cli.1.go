package main

import (
    "fmt"
    "os"

    "github.com/urfave/cli"
)


/*
➜  urfave.cli git:(master) ✗ ./urfave.cli.1
Hello friend!
➜  urfave.cli git:(master) ✗ ./urfave.cli.1 --help
NAME:
   greet - fight the loneliness!

USAGE:
   urfave.cli.1 [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
➜  urfave.cli git:(master) ✗ ./urfave.cli.1 greet
Hello friend!
 */

func main() {
    app := cli.NewApp()
    app.Name = "greet"
    app.Usage = "fight the loneliness!"
    app.Action = func(c *cli.Context) error {
        fmt.Println("Hello friend!")
        return nil
    }

    app.Run(os.Args)
}