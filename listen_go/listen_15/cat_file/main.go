package main

import (
	"bufio"
	"fmt"
	"github.com/urfave/cli"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	var File string
	var num bool
	cat := cli.NewApp()
	cat.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "r",
			Value:       "defaules",
			Usage:       "open file name",
			Destination: &File,
		},
		cli.BoolFlag{
			Name:        "n",
			Usage:       "sort num file",
			Destination: &num,
		},
	}
	fmt.Println(File)
	cat.Action = func(c *cli.Context) error {
		var cmd string
		if c.NArg() > 0 {
			cmd = c.Args()[0]
			fmt.Println(cmd)
		}
		if len(File) != 0 {
			if num == false {
				cats, err := ioutil.ReadFile(File)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				} else {
					fmt.Println(string(cats))
					os.Exit(0)
				}
			} else {
				f, err := os.Open(File)
				cat := bufio.NewReader(f)
				for {
					cats, _ := cat.ReadString('\n')
					if err == io.EOF {
						break
					}
					fmt.Fprintf(os.Stdout, "%s", cats)
				}
			}
		} else {
			fmt.Println("then is cat on file name!")
			os.Exit(2)
		}
		return nil
	}
	cat.Run(os.Args)
}
