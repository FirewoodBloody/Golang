package main

import "fmt"

func array_bl() {
	a := [...]int{2, 3, 7, 5, 3, 7, 9}
	for i := 0; i < len(a); i++ {
		fmt.Printf("a[%d] = %d \n", i, a[i])
	}
}

func array_Range() {
	a := [...]int{2, 3, 7, 5, 3, 7, 9}
	for index, vlaue := range a {
		fmt.Printf("a[%d] = %d \n", index, vlaue)
	}
}

func main() {
	array_bl()
	array_Range()

}
