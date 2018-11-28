package main

import "fmt"

func main() {
	a := map[string]int{
		"s": 111,
		"a": 222,
		"p": 333,
	}
	for key, value := range a {
		fmt.Printf("a[%s] = %d\n", key, value)
	}
}
