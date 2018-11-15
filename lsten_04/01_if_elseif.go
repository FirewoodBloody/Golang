package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a/2 == 0 {
		fmt.Println("then is 偶数")
	} else {
		fmt.Println("then is 奇数")
	}
}
