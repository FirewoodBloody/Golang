package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	var language string
	var recusive bool
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "lang,l",
			Value:       "englosh",
			Usage:       "select language",
			Destination: &language,
		},
		cli.BoolFlag{
			Name:        "recusive ,r",
			Usage:       "recrsive for the greeting",
			Destination: &recusive,
		},
	}

	app.Action = func(c *cli.Context) error {
		var cmd string
		if c.NArg() > 0 {
			cmd = c.Args()[0]
			fmt.Println(cmd)
		}
		fmt.Println(language)
		fmt.Println(recusive)
		return nil
	}
	app.Run(os.Args)
}
