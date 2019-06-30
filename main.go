package main

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"time"
)

func main() {
	c, err := ftp.Dial("2shoucang.f3322.net:21", ftp.DialWithTimeout(5*time.Second)) //创建连接
	if err != nil {
		fmt.Println(err)
	}
	err = c.Login("BOLONG", "131420") //登陆
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("OK")
	time.Sleep(time.Second * 5)
}
