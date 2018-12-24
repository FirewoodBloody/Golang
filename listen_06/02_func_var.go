package main

import (
	"fmt"
	"strings"
)

func add(a, b int) int {
	return a + b
}

func testFunc() {
	f1 := add
	fmt.Printf("f1 type is %T\n", f1)
	sum := f1(2, 2)
	fmt.Printf("%d\n", sum)
}

func testFunc1() {
	f2 := func(a, b int) int { //匿名函数定义
		return a + b
	}
	fmt.Printf("f2 type is %T\n", f2)
	sum := f2(2, 3)
	fmt.Printf("%d\n", sum)
}

func testFunc2() {
	var i = 0
	defer func() {
		fmt.Printf("%d\n", i)
	}()
	i = 100
}

func add1(a, b int) int {
	return a + b
}

func sub1(a, b int) int {
	return a - b
}

func clac(a, b int, op func(int, int) int) int {
	return op(a, b) //clac的第三个参数为一个op函数类型的变量
}

func testFunc5() {
	sum := clac(100, 300, add1)
	sub := clac(100, 300, sub1)
	fmt.Printf("sum = %d , sub = %d \n", sum, sub)
}

func makeSuffixFunc(suffi string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffi) {
			return name + suffi
		}
		return name
	}
}

func testFunc3() {
	func1 := makeSuffixFunc(".bmp")
	func2 := makeSuffixFunc(".jpg")
	fmt.Println(func1("test1"))
	fmt.Println(func2("test2"))
}

func main() {
	defer testFunc()
	defer testFunc1() //defer
	testFunc2()
	testFunc3()
	testFunc5()
}
