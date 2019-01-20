package main

import "fmt"

func testslice() { //切片的定义,未初始化的切边无法进行赋值
	var a []int
	if a == nil {
		fmt.Printf("a is nil\n")
	} else {
		fmt.Printf("a is %d \n", a)
	}
}

func testslice1() {
	a := [5]int{1, 2, 3, 4, 5}
	var b []int
	b = a[:]
	fmt.Printf("slice b: %d \n", b)
	fmt.Printf("slice b[0] is %d\n", b[0])
	fmt.Printf("slice b[5] is %d\n", b[4])
}

func main() {
	testslice()
	testslice1()
}
