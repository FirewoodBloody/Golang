package main

import (
	"fmt"
	"math/rand"
	"time"
)

func autoTime() (auto int) {
	rand.Seed(time.Now().Unix())
	//使用时间戳,初始化随机种子,如未使用时间戳就需要一个随机数,否则默认使用seed1
	//如未初始化,在取数时,默认调用seed(1)
	auto = rand.Intn(2)
	//auto = auto % 3
	time.Sleep(time.Second * 1)
	return

}

//
//func init63() (a int64) {
//	a = rand.Int63n(100)
//	return a
//}

func main() {
	for {
		a := autoTime()
		fmt.Println(a)
	}
	//for {
	//	//rand.Seed(time.Now().Unix())
	//	a := rand.Intn(100)
	//	//a = a % 3
	//	fmt.Println(a)
	//	time.Sleep(time.Second)
	//}
}
