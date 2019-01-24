package main

import (
	"Golang/listen_go/listen_12/example1"
	"fmt"
)

type User struct {
	Username string
	Age      int
	Sex      int
	int
	string
}

func main() {
	a := User{}
	a.Username = "张三"
	a.Age = 18
	a.Sex = 89
	a.int = 100
	a.string = "a"
	fmt.Printf("%#v\n", a)
	b := example1.User{}
	b.City = "sss"
	fmt.Printf("%#v\n", b)

}
