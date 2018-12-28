package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64
	reflect_basic(&x)
	fmt.Println(x)
}

func reflect_basic(a interface{}) {
	t := reflect.ValueOf(a)
	fmt.Println(&t)
	k := t.Kind()
	if k == reflect.Float64 {
		t.Elem().SetFloat(6.4)
	}

}
