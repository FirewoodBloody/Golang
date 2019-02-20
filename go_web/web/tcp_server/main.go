package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:50000")
	if err != nil {
		fmt.Printf("listen failed ,err:%v\n", err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed,err:%s\n", err)
			continue
		}
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Printf("read form conn failed ,err:%s\n", err)
		}
		str := string(buf[:n])
		fmt.Printf("recv for client,data:%s\n", str)
	}
}
