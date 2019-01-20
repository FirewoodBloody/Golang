package main

import "fmt"

func main() {
	a := "hwo do you do"
	var b map[string]int = make(map[string]int, 10000)
	var app string
	for i, key := range a {
		//fmt.Printf("%d %c %T\n", index, key, a)
		if i == len(a)-1 {
			app = fmt.Sprintf("%s%s", app, string(key))
			b[app] += 1
		} else if !(string(key) == " ") {
			app = fmt.Sprintf("%s%s", app, string(key))
		} else if string(key) == " " {
			b[app]++
			app = ""
		}

	}
	for key, value := range b {
		fmt.Printf("%v is %v\n", key, value)
	}
}
