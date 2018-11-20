package main

import "fmt"

//水仙花数:水仙花数是指一个 3 位数，它的每个位上的数字的 3次幂之和等于它本身
// （例如：1^3 + 5^3+ 3^3 = 153）。

func example1() {
	for i := 100; i < 1000; i++ {
		if is_shuixianhua(i) == true {
			fmt.Printf("%d is 水仙花数 \n", i)
		}

	}
}

func is_shuixianhua(n int) bool {
	fist := n % 10          //个位
	second := (n / 10) % 10 //十位
	third := (n / 100) % 10 //百位
	sum := fist*fist*fist + second*second*second + third*third*third
	if sum == n {
		return true
	}
	return false
}

func main() {
	example1()
}
