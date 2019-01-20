package main

import "fmt"

func main() {
	a := [...]string{"a", "b", "c", "d", "e"}
	b := a[0:3]
	fmt.Printf("b slice is %s , b len is %d , b cap is %d", b, len(b), cap(b))
}

func len_cap() {
	var a []int
	fmt.Printf("a len is %d , a cap is %d ", len(a), cap(a))
	if a == nil {
		fmt.Println("a is nil")
	}
}
