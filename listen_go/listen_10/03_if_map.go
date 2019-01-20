package main

import "fmt"

func main() {
	a := map[int]int{
		01: 111,
		02: 222,
		03: 333,
	}
	value, ok := a[03]
	if ok == false {
		fmt.Println("key 03 is not exist\n")
	} else {
		fmt.Printf("key 03 is %d\n", value)
	}
}
