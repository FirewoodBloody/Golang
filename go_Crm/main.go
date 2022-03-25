package main

import (
	"fmt"
	"github.com/FirewoodBloody/Golang/go_Crm/crmapi/KuaiDi/jd"
)

type T struct {
	Codecustomer string `json:"customerCode"`
	WaybillCode  string `json:"waybillCode"`
	JosPin       string `json:"josPin"`
}

func main() {

	s, err := jd.SendRequestS("https://api.jdl.cn", "/query/dynamictraceinfo", "[\n\t\"029K708772\",\n\t\"JDVG01178821629\",\n\t\"jd_12\"\n\t]", "express")
	if err != nil {
		return
	}
	fmt.Println(s)
}
