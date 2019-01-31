package main

import (
	"fmt"
	"github.com/FirewoodBloody/PacketProup/ranges"
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
