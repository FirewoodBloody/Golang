package main

import "fmt"

func main() {
	var a = make([]int, 5, 10)
	a[0] = 10
	a = append(a, 10) //在切片长度之外增加一位
	fmt.Printf("a slice is %d", a)
	for i := 0; i < 8; i++ {
		a = append(a, i)
		fmt.Printf("a = %d , a len is %d , a cap is %d \n", a, len(a), cap(a))
	}
}
