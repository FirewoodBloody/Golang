package main

import (
	"Golang/express_api/modules"
	"encoding/json"
	"fmt"
)

func main() {
	data, _ := modules.SfCreateData("SF1026386232204")
	a, _ := json.Marshal(data)
	fmt.Println(string(a))

}
