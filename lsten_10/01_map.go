package main

import "fmt"

func main() {
	var a map[string]int
	//a[str] = 100
	if a == nil {
		a = make(map[string]int, 16)
		fmt.Printf("a type is %T\n", a)
		a["str"] = 100
		a["st"] = 200
		a["s"] = 300
		fmt.Printf("a = %#v\n", a) //%# 可以输出map的key类型和value的类型
	}
}
