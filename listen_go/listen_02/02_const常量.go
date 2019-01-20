package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
	)
	fmt.Println(a, b, c)

	const (
		d = 1 << iota
		e
		f
	)
	fmt.Println(d, e, f)
}
