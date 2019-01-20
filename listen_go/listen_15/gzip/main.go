package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	flie, err := os.Open("./main_server.go.gz")
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer flie.Close()
	reads, err := gzip.NewReader(flie)
	if err != nil {
		fmt.Println("gzip open file error: ", err)
		return
	}

	var buf [128]byte
	var conts []byte
	for {
		n, err := reads.Read(buf[:])
		if err == io.EOF {
			break

		}
		if err != nil {
			fmt.Println("read file error :", err)
			return
		}
		conts = append(conts, buf[:n]...)
	}
	fmt.Println(string(conts))
}
