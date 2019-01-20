package main

import (
	"fmt"
	"math/rand"
	"time"
)

//求数组元素之和
func sumArray(a [10]int) int {
	var sum int
	//for _, value := range a {
	//	sum += value
	//}
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	return sum
}

func vararray() {
	var input [10]int
	for i := 0; i < len(input); i++ {
		rand.Seed(time.Now().Unix()) //初始化随机数种子
		input[i] = rand.Intn(10000)  //产生一个0-999的伪随机数
	}
	sum := sumArray(input)
	fmt.Printf("input[:] sum = %d \n", sum)
}

func main() {
	vararray()

}
