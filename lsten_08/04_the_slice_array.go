package main

import "fmt"

func sliceArray(a []int) int {
	var Sum int
	for _, value := range a {
		Sum += value
	}
	return Sum
}

func sunArray(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == 3 {
				fmt.Printf("a[%d] + a [%d] = %d\n", i, j, a[i]+a[j])
			}
		}
	}
}

func main() {
	var a = [4]int{1, 2, 3, 4}
	var alias []int
	alias = a[:]
	b := sliceArray(alias)
	fmt.Printf("sum = %d \n", b)
	sunArray(alias)
}
