package main

import (
	"fmt"
)

func bubbl_cort(a [8]int) [8]int {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return a
}

func main() {
	var a = [8]int{2, 5, 7, 8, 1, 2, 5, 8}
	b := bubbl_cort(a)
	fmt.Println(a)
	fmt.Println(b)
}
