package main

import "fmt"

func testAyyay() {
	var a []int = []int{1, 2, 3}
	var b []int = []int{8, 9, 0}
	a = append(a, 4, 5, 6, 7)
	fmt.Printf("a then is %d \n", a)
	a = append(a, b...)
	fmt.Printf("a then is %d \n", a)
}

func main() {
	testAyyay()
}
