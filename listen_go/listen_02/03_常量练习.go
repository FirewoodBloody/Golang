package main

import "fmt"

func main() {
	const (
		a = iota
		c
		d
		e = 8
		f = iota
		g
	)
	const (
		a1 = iota
		a2
	)
	fmt.Println(a, c, d, e, f, g)
	fmt.Println(a1, a2)
}
