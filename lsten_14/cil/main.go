package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "俄罗斯"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {
		fmt.Println("Hello World")
		return nil
	}
	app.Run(os.Args)
}
