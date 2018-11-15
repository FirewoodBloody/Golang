package main

import (
	"fmt"
)

func teststring() {
	var str string = "Hello World"
	fmt.Printf("str[0] = %c , len(str) = %d\n", str[0], len(str))

	for index, value := range str {
		fmt.Printf("str index is %d , str value is %c\n", index, value)
	}

	//修改字符串的单个字符
	//修改方法是将字符串转换成切片类型 ,修改至后,再将类型转回字符串(这种理解并不完全正确)
	var byteslect []byte    //切片的定义 格式 与变量的不同之处
	byteslect = []byte(str) //将 str(字符串)的值 以 byte 的类型赋给 byteslect
	byteslect[0] = 'h'      //修改切片的value的值
	str = string(byteslect) //讲切片的值按string的类型赋给str,此时str的类型是string 值已发生改变
	fmt.Printf("str[0] = %c\n", str[0])

	str = "Hello 编程"
	fmt.Println("str is ", str)
	fmt.Printf("len is str : %d\n", len(str))

	str = "中" //一个中文字占3个byte
	fmt.Printf("str len is chan 字符 : %d \n", len(str))

	a := "中国"
	b := 'a'

	fmt.Printf("\"\" type is %T , '' type is %T \n", a, b)

}

func main() {
	teststring()
}
