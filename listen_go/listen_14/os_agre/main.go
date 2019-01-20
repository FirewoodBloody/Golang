package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
	if len(os.Args) > 0 {
		for index, _ := range os.Args {
			if index == 0 {
				continue
			}
			fmt.Println(index)
		}
	}
}
