package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var str string
	reader := bufio.NewReader(os.Stdin) //通过Bufio读取一行字符
	str, _ = reader.ReadString('\n')    //指定分割符
	fmt.Println(str)
}
