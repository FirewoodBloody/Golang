package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("./main", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer file.Close()
	wirte := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		wirte.WriteString("hello world\n")
	}
	wirte.Flush() //刷新缓存区，使得内容写入文件
}
