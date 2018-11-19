package main

import "fmt"

func chello() { //定义函数,无输入参数,无返回值
	fmt.Println("Hello World")
}

func add(a, b int) int {
	sum := a + b
	return sum
}

func main() {
	chello() //函数调用
	s := add(100, 200)
	fmt.Println(s)
}
