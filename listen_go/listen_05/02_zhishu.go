package main

import "fmt"

func justify(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ { //判断数字是否能被其他非1和本身整除
		if n%i == 0 {
			return false
		}
	}
	return true
}

func example() {
	for i := 2; i < 100; i++ { //循环1至100的数字
		if justify(i) == true {
			fmt.Printf("%d is prime\n", i)
		}
	}
}

func main() {
	example() //判断一个数是不是质数 ,:只能被1和本身整除的为质数
}
