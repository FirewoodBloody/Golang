package main

import "fmt"

func main() {
	a := map[string]int{
		"s": 100,
		"a": 200,
	}
	delete(a, "s")
	fmt.Println(a)
}
