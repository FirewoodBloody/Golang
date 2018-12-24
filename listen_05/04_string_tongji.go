package main

import "fmt"

func calc(str string) (charcount int, numcount int, spacecount int, othrecount int) {
	uftchar := []rune(str)
	for i := 0; i < len(uftchar); i++ {
		switch uftchar[i] {
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			charcount++
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			charcount++
		case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
			numcount++
		case ' ':
			spacecount++
		default:
			othrecount++
		}
	}
	return
}

func example() {
	var str = "asdasd    AAAAA  中国是打过  12345678"
	charcount, numcount, spaleconunt, othercount := calc(str)
	fmt.Printf("%v , %v , %v , %v ", charcount, numcount, spaleconunt, othercount)
}

func main() {
	example()
}
