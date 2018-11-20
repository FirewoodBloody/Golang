package main

import "fmt"

var a = 100 //全局变量,在当前程序内均可以使用--初始化

func testgovar() {
	fmt.Printf("%d\n", a)
}

func testlocalvar() {
	var b = 200 //局部变量
	fmt.Printf("%d\n", a, b)
	if a == 100 {
		var c = 150
		fmt.Printf("%d\n", c)
	}
	//var i int
	//fmt.Printf("%d\n", c)
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\n", i)
	}
	//fmt.Printf("%d\n", i)
}

func main() {
	testgovar()
	testlocalvar()
}
