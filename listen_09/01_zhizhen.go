package main

import "fmt"

func main() {
	var a int = 100
	fmt.Println(a)
	fmt.Println(&a)
	var b *int
	fmt.Printf("%p %v", &b, b)
}
