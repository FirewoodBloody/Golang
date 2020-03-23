package main

import (
	"Golang/express_api/modules"
	"encoding/json"
	"fmt"
)

func main() {
	P := new(modules.PublicParametersr)

	R := new(modules.RequestData)

	R.CustomerCode = "029K708772"
	R.WaybillCode = "JDVG00272341243"

	data, _ := json.Marshal(R)
	pstr := P.SetUserLogin1(data)
	fmt.Printf("%v", pstr)
}
