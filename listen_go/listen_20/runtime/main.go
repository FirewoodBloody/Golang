package main

import (
	"Golang/ranges"
	"fmt"
)

func main() {
	var a int = 8527536549510
	var sum int
	b, err := ranges.DivisionInt(a)
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range b[:] {
		sum += v
	}
	fmt.Println(sum)
}
