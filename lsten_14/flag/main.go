package main

import (
	"flag"
	"fmt"
)

var (
	str string
	shu int
	bb  bool
)

func init() {
	flag.StringVar(&str, "s", "defaule", "ssssssss")
	flag.IntVar(&shu, "t", 1, "sssssss")
	flag.BoolVar(&bb, "b", false, "sssssss")
}

func main() {
	fmt.Println(str)
	fmt.Println(shu)
	fmt.Println(bb)
}
