package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./main.go", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()
	str := "hello world"
	file.Write([]byte(str))
}
