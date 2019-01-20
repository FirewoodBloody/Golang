package main

import "fmt"

func testArray() {
	var a [6]int //定义数组 格式:var 变量 [数组个数] 数组类型
	//var a [6]int = [6]int{12, 3, 5, 5, 6, 66}
	//定于变量并初始化格式
	b := [...]int{1, 1, 1, 1, 1} //定义数组是初始化
	c := [3]int{1, 1, 1}
	d := [3]int{10}
	e := [3]int{2: 10}

	fmt.Println(a)
	a[0] = 100
	a[1] = 200
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

func main() {
	testArray()
}
