package main

import (
	"fmt"
	"github.com/fatih/color"
)

//func ListDir(dirName string, deep int) (err error) {
//	dir, err := ioutil.ReadDir(dirName)
//	if err != nil {
//		return err
//	} else {
//		if deep == 1 {
//			fmt.Printf("!---%s")
//			color.Set(color.FgBlue, color.Bold)
//			fmt.Println(filepath.Base(dirName))
//			defer color.Unset()
//			fmt.Printf("!---%s")
//		} else {
//
//		}
//	}
//	return nil
//}

func main() {
	fmt.Printf("!---")
	color.Set(color.FgBlue, color.Bold)

	fmt.Println("00000")
	color.Set(color.FgWhite)
	fmt.Printf("!---")
	color.Unset()
}
