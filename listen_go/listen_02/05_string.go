package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Println(a)
	a = "a"
	b := a
	fmt.Println(b)

	fmt.Printf("%s\n", b)

	b = "Hello World"

	fmt.Println("len b is :", len(b))

	ip := "192.12.12.3;123.12.123.3"
	ipsrt := strings.Split(ip, ";") //字符串的分割,切片
	fmt.Println(ipsrt[0], ipsrt[1])

	//c := "hello\n"
	c := "hello"
	fmt.Println(c)

	d := len(ip)
	fmt.Println(d)

	a = "我是谁"
	fmt.Println(len(a))
	//c = c + a
	c = fmt.Sprintf("%s,%s", c, a)
	fmt.Println(c)
}
