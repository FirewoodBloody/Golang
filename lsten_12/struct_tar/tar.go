package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username string `json:"username",bd:"user_name"`
	Sex      string `json:"sex"`
	Age      int    `json:"age""`
	Score    int
}

func main() {
	user := User{
		Username: "张三",
		Sex:      "男",
		Age:      18,
		Score:    89,
	}
	fmt.Printf("%#v\n", user)
	date, err := json.Marshal(user)
	if err == nil {
		fmt.Printf("%#s\n", string(date))
	}
}
