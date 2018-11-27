package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	lenght  int
	charset string
)

const (
	NumSet    = "1234567890"
	CharSeeet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SpecSet   = "!@#%*&$()_+.?"
)

func example() {
	flag.IntVar(&lenght, "1", 16, "-1 生成密码的长度")
	flag.StringVar(&charset, "t", "num",
		"-t 制定密码生成的字符集, num只是用数字[0-9], char只是用英文字母[a-zA-Z],"+
			"mix 使用数字和字母, advance:使用数字字母以及特殊字符")
	flag.Parse()
}

func generatePassword() string {
	var password []byte = make([]byte, lenght, lenght)
	var sourceSet string
	if charset == "num" {
		sourceSet = NumSet
	} else if charset == "char" {
		sourceSet = CharSeeet
	} else if charset == "mix" {
		sourceSet = fmt.Sprintf("%s%s", NumSet, CharSeeet)
	} else if charset == "advance" {
		sourceSet = fmt.Sprintf("%s%s%s", NumSet, CharSeeet, SpecSet)
	} else {
		sourceSet = NumSet
	}
	for i := 0; i < lenght; i++ {
		index := rand.Intn(len(sourceSet))
		password[i] = sourceSet[index]
	}
	return string(password)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	example()
	//fmt.Printf("lenght:%d  charset:%s", lenght, charset)
	password := generatePassword()
	fmt.Printf("%s", password)
}
