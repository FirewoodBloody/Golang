package main

import "fmt"

//学习重在理解,讲复杂的结构进行细化解剖,方可理解其中的逻辑

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int { //匿名函数
		base += i
		return base
	}
	sub := func(i int) int { //匿名函数
		base -= i
		return base
	}
	return add, sub
}

func calc1() (func(int) int, func(int) int) {
	var base int                              //间接等于函数定义时传入参数,区别是一个可以在外面传入参数且可变,一个不可变的局部变量
	var add func(int) int = func(i int) int { //add函数是一个匿名变量,type是func(int)int
		base += i
		return base
	}
	var sub func(int) int = func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	a, b := calc(0)
	c, d := calc1()
	//这两个函数在base为0 的情况下结果一样
	fmt.Println(a(1), b(1))
	fmt.Println(c(1), d(1))

}
