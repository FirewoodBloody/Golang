package main

//异或加密算法
import (
	"fmt"
	"strconv"
)

var XorKey []byte = []byte{0xA1, 0xB7, 0xAC, 0x57, 0x1C, 0x63, 0x3B, 0x81}

type Xor struct {
}

//加密
func (a *Xor) enc(src string) string {
	var result string
	j := 0
	s := ""
	bt := []rune(src)
	for i := 0; i < len(bt); i++ {
		if i == 0 {
			fmt.Println(bt[i], "###", byte(bt[i]))
		}
		s = strconv.FormatInt(int64(byte(bt[i])^XorKey[j]), 16)
		if len(s) == 1 {
			s = "0" + s
		}
		result = result + (s)
		j = (j + 1) % 8
	}
	return result
}

//解密
func (a *Xor) dec(src string) string {
	var result string
	var s int64
	j := 0
	bt := []rune(src)
	//fmt.Println(bt)
	for i := 0; i < len(src)/2; i++ {
		s, _ = strconv.ParseInt(string(bt[i*2:i*2+2]), 16, 0)
		result = result + string(byte(s)^XorKey[j])
		j = (j + 1) % 8
	}
	return result
}
func main() {
	xor := Xor{}
	fmt.Println(xor.enc("13991252603,20190724,094557"))
	fmt.Println(xor.dec("908f9a6e2d5609b6998f947b2e5309b191819c63305202b1978594"))
}
