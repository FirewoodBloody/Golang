package main

import "fmt"

func switchif() {
	var a int
	fmt.Scan(&a)

	if a == 1 {
		fmt.Println("a = ", a)
	} else if a == 2 {
		fmt.Println("a = ", a)
	} else {
		fmt.Println("a = ", a)
	}
}

func testswitch() {
	switch a := 1; a {
	case 1:
		fmt.Println("a = ", a)
	case 2:
		fmt.Println(a)
	default:
		fmt.Println(a)

	}
}
func main() {
	switchif()
}
