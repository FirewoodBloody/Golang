package main

import "fmt"

func array_Duo(a [3][2]int) {
	for index1, value1 := range a {
		for index, value := range value1 {
			fmt.Printf("a[%d][%d] = %d ", index1, index, value)
		}
		fmt.Printf("\n")
	}
}

func main() {
	a := [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	array_Duo(a)
}
