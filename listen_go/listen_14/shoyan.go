package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err = inputReader.ReadString('S') //func (b *Reader) ReadString(delim byte) (line string, err error) ,‘S’ 这个例子里使用S表示结束符，也可以用其它，如'\n'
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}
