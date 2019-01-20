package main

import (
	"fmt"
	"time"
)

func num() {
	for i := 0; i <= 5; i++ {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("%d\n", i)
	}
}

func alp() {
	for i := 'a'; i <= 'f'; i++ {
		time.Sleep(time.Millisecond * 400)
		fmt.Printf("%c\n", i)
	}
}

func main() {
	go num()
	go alp()
	time.Sleep(time.Millisecond * 3000)
	fmt.Println("main")
}
