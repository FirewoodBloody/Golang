package main

import "fmt"

var a bool

func main() {
	if !a { //如果bool为真成立，为假不成立
		fmt.Println(a)
	}
}
