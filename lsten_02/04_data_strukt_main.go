package main

import (
	"fmt"
)

func textbool() {
	var a bool
	fmt.Println(a)
	a = true
	fmt.Println(a)
	a = !a
	fmt.Println(a)

	var b bool
	if a == true && b == true {
		fmt.Println("right")
	} else {
		fmt.Println("not right")
	}
	b = true
	if a == true || b == true {
		fmt.Println("|| right")
	} else {
		fmt.Println("|| not right")
	}
}

func teatInt() {
	var (
		a int8 = 14
		b int  = 123123123
	)
	fmt.Println("a = ", a, ",b = ", b)
	a, b = -11, -123123
	fmt.Println("a = ", a, ",b = ", b)
	b = int(a)
	fmt.Printf("b type is %T,b = ", b, b)

	var c = 4.3
	fmt.Println(c)

	fmt.Printf("a = %d , b = %x , c = %f ", a, b, c)
}

func main() {
	teatInt()
}
