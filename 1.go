package main

import (
	"Golang/Express_Routing/express"
	"fmt"
)

func main() {

	data, _ := express.SfCreateData("016229745333")
	a := data.Body.RouteResponse.Route
	fmt.Println(data)
	for _, v := range a {
		fmt.Println(v)
	}

}
