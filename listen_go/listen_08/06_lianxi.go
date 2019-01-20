package main

import "fmt"

func main() {
	var sa = make([]string, 5, 10)
	fmt.Println(sa)
	for i := 0; i < 10; i++ {
		sa = append(sa, fmt.Sprintf("%v", i))
	}
	fmt.Println(sa)
}
