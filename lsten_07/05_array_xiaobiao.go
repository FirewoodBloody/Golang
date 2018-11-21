package main

import "fmt"

//求一个数组元素给定条件的下标
func sunArray(a [10]int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i]+a[j] == 10 {
				fmt.Printf("a[%d] + a [%d] = %d\n", i, j, a[i]+a[j])
			}
		}
	}
}

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	sunArray(a)

}
