package main

import "fmt"

func main() {
	var a int
	var b bool
	var c string
	var d float64
	fmt.Printf("a=%d,b=%t,c=%s,d=%f\n", a, b, c, d)

	a, b, c, d = 1, true, "hello", 2.12
	fmt.Printf("a=%d,b=%t,c=%s,d=%f\n", a, b, c, d)

}
