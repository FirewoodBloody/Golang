package iniconfig

import (
	"fmt"
	"strings"
)

func main() {

}
func Marshal(data interface{}) (result []byte, err error) {
	return
}

func UnMarshal(data []byte, result interface{}) (err error) {
	lineArr := strings.Split(string(data), "\n")
	for _, v := range lineArr {
		fmt.Println(v)
	}
	return
}
