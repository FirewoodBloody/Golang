package main

import (
	"fmt"
)

func chanRune() {
	var cha = "hello 世界"
	var aaa = []rune(cha)
	for i := 0; i < len(aaa)/2; i++ { //中文交换注意事项
		tm := aaa[len(aaa)-i-1]
		aaa[len(aaa)-i-1] = aaa[i]
		aaa[i] = tm
	}
	cha = string(aaa)
	fmt.Println(cha)
}

func main() {
	chanRune()
}
