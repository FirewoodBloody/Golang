package main

import "fmt"

func main() {
	var a map[int](map[string](map[int]int))
	a=make(map[int](map[string](map[int]int)) ,1000)
	a[]=
	a[0]["张三"][18] = 89
	fmt.Println(a[01])
}
