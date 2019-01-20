package main

import (
	"fmt"       //格式化输出
	"math/rand" //随机数
	"sort"      //排序
	"time"      //时间包
)

func main() {
	rand.Seed(time.Now().Unix())
	a := make(map[int]int)

	for i := 0; i < 128; i++ {
		b := rand.Intn(int(time.Now().Unix()))
		a[i] = b
	}
	//for key, value := range a {
	//	fmt.Printf("a[%d] = %d \n", key, value)
	//}
	var sorts []int
	for key, _ := range a {
		sorts = append(sorts, key)
	}
	sort.Ints(sorts)
	for _, key := range sorts {
		fmt.Printf("a[%d] = %d \n", key, a[key])
	}
}
