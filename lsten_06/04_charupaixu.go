package main

import "fmt"

func inset_pauxi(inset [8]int) [8]int {
	for i := 1; i < len(inset); i++ {
		for j := i; j > 0; j-- { //
			//判断inset[1]与前一位的大小,前一位等于inset[1-1]
			if inset[j] < inset[j-1] {
				inset[j], inset[j-1] = inset[j-1], inset[j]
			} else {
				break
			}
		}
	}
	return inset
}

func main() {
	var a [8]int = [8]int{5, 6, 0, 4, 2, 3, 5, 8}
	b := inset_pauxi(a)
	fmt.Println(a)
	fmt.Println(b)

}
