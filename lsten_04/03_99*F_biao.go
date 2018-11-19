package main

import "fmt"

//打印99乘法表
func testMulti() { //定义函数
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}
}
func main() {
	testMulti()
}
