package main

import "fmt"

func array_copy() {
	a := [3]int{1, 2, 3}
	b := a
	b[0] = 100
	fmt.Printf("a == %d \n", a)
	fmt.Printf("b == %d \n", b)
}

func tset_Array(b [3]int) {
	b[0] = 10
}

func array_test() {
	var a = [3]int{1, 2, 3}
	tset_Array(a)
	fmt.Println(a)
}

func main() {
	array_copy()
	array_test()
}
