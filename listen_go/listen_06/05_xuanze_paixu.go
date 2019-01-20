package main

import "fmt"

func select_sort(a [8]int) [8]int {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[i] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}

func main() {
	var a = [8]int{3, 5, 8, 9, 1, 2, 7, 9}
	b := select_sort(a)
	fmt.Println(a)
	fmt.Println(b)

}
