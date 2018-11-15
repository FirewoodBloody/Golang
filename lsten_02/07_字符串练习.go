package main

import "fmt"

func main() {
	var str = "hello"
	var bytes = []byte(str)

	for i := 0; i < len(str)/2; i++ { //i < len(str) 循环次数,只循环一般
		tmp := bytes[len(str)-i-1]
		bytes[len(str)-i-1] = bytes[i]
		bytes[i] = tmp
	}

	str = string(bytes)

	fmt.Println(str)
}
