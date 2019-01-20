package main

import "fmt"

func main() {
	var a map[string]int = map[string]int{ //定义map是初始化赋值
		"str": 100,
		"st":  200,
		"s":   300,
	}
	fmt.Printf("%#v\n", a)
	a["s"] = 111
	a["a"] = 444
	fmt.Printf("%#v\n", a)
	key := "s"
	fmt.Printf("a value of key[%s] is : %d\n", key, a[key])
}
