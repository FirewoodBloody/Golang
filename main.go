package main

import (
	"fmt"
	"strings"
)

func main() {
	//c, err := ftp.Dial("61.136.101.122:21") //创建连接
	//
	//if err != nil {
	//	fmt.Println("1", err)
	//}
	//err = c.Login("enkj", "enkj#EDC") //登陆
	//if err != nil {
	//	fmt.Println("2", err)
	//}
	//
	//fmt.Println("OK")
	//time.Sleep(time.Second * 5)
	//file, err := os.Open("D:/Record/20190703/022005_185000.mp3")
	//if err != nil {
	//	fmt.Println("3", err)
	//}
	//
	//defer file.Close()
	////传输文件，指定传输的文件路径和文件名（针对于接收文件的服务器的根目录之下的），以及需要传输的文件IO
	//err = c.Stor(fmt.Sprintf("%d%02d%02d/%s", time.Now().Year(), int(time.Now().Month()), time.Now().Day(), "010847_122931.mp3"), file)
	//if err != nil {
	//	fmt.Println("4", err)
	//}

	fmt.Println(strings.ToUpper("sf1010392510534"))
	fmt.Println(strings.ToTitle("sf1010392510534"))
}
