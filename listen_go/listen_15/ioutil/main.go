package main

import (
	"fmt"
	"io/ioutil" //读取文件
)

func main() {
	file, err := ioutil.ReadFile("./main")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(file))
}
