package main

import (
	"fmt"
)

func funcA() int {
	x := 6
	defer func() {
		x++
		fmt.Println("1", x)
	}()
	fmt.Println("2", x)
	return x
}

func funcB() (x int) {
	fmt.Println("1", x)
	defer func() {
		x += 1
		fmt.Println("2", x)
	}()
	fmt.Println("3", x)
	return 5
}

func funcA() (y int) {
	x := 6
	defer func() {
		x++
		fmt.Println("1", x)
	}()
	fmt.Println("2", x)
	return x
}

func main() {
	//	a := funcA()
	a := funcB()
	fmt.Println("3", a)
}
