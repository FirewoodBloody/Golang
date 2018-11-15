package main

import (
	"fmt"
	"time"
)

func testTime() {
	now := time.Now() //获取当前的时间
	//fmt.Println(now)
	//按找不同分段进行不同时间分段的获取
	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //时
	minute := now.Minute() //分
	second := now.Second() //秒
	week := now.Weekday()  //星期
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d %s", year, month, day, hour, minute, second, week)
} //%02d 打印时不足两位默认用0补齐

func main() {
	testTime()
}
