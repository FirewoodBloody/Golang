package main

import (
	"fmt"
	"sort"
)

var A int

func arraysort(a []int) []int {
	sort.Ints(a)
	return a
}

func main() {
	var a [5]int = [5]int{7, 3, 8, 1, 4}
	var sora []int
	sora = a[:]
	aa := arraysort(sora)
	fmt.Printf("%d\n", aa)
}
