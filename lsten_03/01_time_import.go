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
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d %s\n", year, month, day, hour, minute, second, week)
	timesecon := now.Unix() //获取当前时间的时间戳,时间戳单位为秒
	fmt.Printf("%d\n", timesecon)
} //%02d 打印时不足两位默认用0补齐

func testtimestamp(timestamp int64) { //定义方法,传入一个参数,参数为时间戳,类型为int64
	timeobj := time.Unix(timestamp, 0) //将时间戳(秒)转换为time(y-M-d h:m:s)时间格式
	//在转换的时间内获取每个分段的时间
	year := timeobj.Year()     //年
	month := timeobj.Month()   //月
	day := timeobj.Day()       //日
	hour := timeobj.Hour()     //时
	minute := timeobj.Minute() //分
	second := timeobj.Second() //秒
	week := timeobj.Weekday()  //星期
	fmt.Printf("%02d-%02d-%02d %02d:%02d:%02d %s\n", year, month, day, hour, minute, second, week)
}

func testticke() {
	ticker := time.Tick(time.Second * 2)
	//每10秒将当前时间传入变量
	fmt.Printf("%v", ticker)
	for i := range ticker {
		fmt.Printf("%v\n", i)
		fmt.Printf("%v\n", ticker)
	}
}

func testFormat() {
	now := time.Now().Format("2006-01-02 15:04:05")
	fmt.Printf("%v\n", now)
}

func testConst() { //time 内置的常量
	fmt.Printf("nano second:%d\n", time.Nanosecond)
	fmt.Printf("micro second:%d\n", time.Microsecond)
	fmt.Printf("mili second:%d\n", time.Millisecond)
	fmt.Printf("second :%d\n", time.Second)
}

func main() {
	start := time.Now().Nanosecond()

	testTime()
	timestamp := time.Now().Unix() //获取当前时间的时间戳
	testtimestamp(timestamp)       //调用函数,并传入参数,当前时间的时间戳
	//testticke()
	testConst()
	testFormat()

	end := time.Now().Nanosecond()

	cost := end - start
	fmt.Printf("%d\n", cost)
}
