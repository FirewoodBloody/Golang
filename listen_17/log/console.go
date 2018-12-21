package log

import (
	"fmt"
)

type ConsoleLog struct {
}

func NewConsileLog(file string) LogInterface {
	return &ConsoleLog{}
}

func (f *ConsoleLog) LogDebug(msg string) {
	fmt.Println(msg)
}

func (f *ConsoleLog) LogWarn(msg string) {
	fmt.Println(msg)
}
