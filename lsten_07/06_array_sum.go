package main

import "fmt"

//求数组元素之和
func sumArray(a [10]int) int {
	var sum int
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	b := sumArray(a)
	fmt.Printf("a[:] sum = %d \n", b)

}
