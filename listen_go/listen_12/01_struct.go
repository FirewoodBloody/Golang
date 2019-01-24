package main

import "fmt"

type Users struct {
	Name string
	Sex  string
	Age  int
	Url  string
}

func main() {
	var user Users
	user.Name = "张三"
	user.Sex = "女"
	user.Age = 18
	user.Url = "www.baidu.com"
	fmt.Printf("%s , %s , %d , %s \n", user.Name, user.Sex, user.Age, user.Url)
	fmt.Printf("%#v\n", user)

	var user1 User = User{
		Name: string("王五"),
		//Sex:  "男",
		Age: 18,
		Url: "www.xxx.com",
	}
	fmt.Printf("%#v\n", user1)
}
