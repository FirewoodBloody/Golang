package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello world goroutine!")
}

func main() {
	go hello()
	fmt.Println("hello main")
	time.Sleep(time.Second)
}
