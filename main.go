package main

import (
	"Golang/Express_Routing/models"
	"fmt"
)

func main() {
	e := models.Engine{}

	e.NewEngine()

	var no string
	nomao, err := e.Engine.Query(fmt.Sprintf("SELECT NO FROM BLCRM.CRM_SYS02 WHERE NAME = '%s'", "柴雪新"))
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range nomao {
		for _, i := range v {
			if i != nil {
				no = string(i)
			}
		}
	}
	maps, err := e.Engine.Query(fmt.Sprintf("SELECT LOGIN_NAME FROM BLCRM.CRM_SYS04 WHERE OPER_NO = '%s'", no))
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range maps {
		for _, i := range v {
			if string(i) != "" {
				fmt.Println(string(i))
			}
		}
	}

}
