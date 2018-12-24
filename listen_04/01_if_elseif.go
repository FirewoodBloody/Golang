package main

import (
	"fmt"
)

func main() {
	var a int
	for {
		fmt.Print("请输入一个非零的整数:")
		fmt.Scan(&a)

		//typea := fmt.Sprintf("%T", a)
		//fmt.Println(a)
		//fmt.Println(typea)
		if a != 0 {
			if a%2 == 0 {
				fmt.Println("then is 偶数")
				break
			} else {
				fmt.Println("then is 奇数")
				break
			}
		} else {
			fmt.Println("您的输入有误,请重新输入!")
		}
	}
}
