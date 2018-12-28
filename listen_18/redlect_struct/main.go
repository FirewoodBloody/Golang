package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name string
	Sex  int
	Age  int
	sss  string
}

func main() {
	var s student
	v := reflect.ValueOf(s)
	t := v.Type()
	kind := t.Kind()
	switch kind {
	case reflect.Int:
		fmt.Println("type is s int")
	case reflect.Float64:
		fmt.Println("type is s float64")
	case reflect.Struct:
		fmt.Println("type is s struct")
		fmt.Printf("num field s is %d\n", t.NumField())
	}
}
