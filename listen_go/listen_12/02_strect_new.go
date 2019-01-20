package main

import "fmt"

type Test struct {
	A int32
	B int32
	C int32
	D int32
}

func main() {
	var a Test
	fmt.Printf("a addre :%p\n", &a.A)
	fmt.Printf("a addre :%p\n", &a.B)
	fmt.Printf("a addre :%p\n", &a.C)
	fmt.Printf("a addre :%p\n", &a.D)
}
