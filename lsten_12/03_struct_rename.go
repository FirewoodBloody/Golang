package main

import "fmt"

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
}
