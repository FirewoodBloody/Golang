package main

import "fmt"

func main() {
	var stuMap map[int]map[string]interface{} //空数据类型，interface 可以插入任何类型的数据
	stuMap = make(map[int]map[string]interface{}, 16)
	var id int = 1
	var name string = "张三"
	var age int = 18
	var scoue int = 89
	value, ok := stuMap[id]
	if !ok {
		value = make(map[string]interface{}, 8)
		value["name"] = name
		value["age"] = age
		value["scoue"] = scoue
	} else {
		value["name"] = name
		value["age"] = age
		value["scoue"] = scoue
	}
	stuMap[id] = value
	fmt.Printf("%v", stuMap)
}
