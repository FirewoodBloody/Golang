package main

import (
	"Golang/Express_Routing/express"
	"fmt"
	"github.com/Luxurioust/excelize"
	"os"
	"strings"
)

func main() {
	excel, err := excelize.OpenFile("快递.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//cell := excel.GetCellValue("Sheet1", "B1")
	//cell := excel.GetCellValue("Sheet1", "B1")
	//fmt.Println(cell)

	//index := excel.GetSheetIndex("Sheet2")
	a := 1
	rows := excel.GetRows("Sheet1")

	for _, row := range rows {
		for _, colcell := range row {

			data, err := express.SfCreateData(colcell)

			if err != nil {

				continue
			}

			for i := 0; i < len(data.Body.RouteResponse.Route); i++ {

				if data.Body.RouteResponse.Route[len(data.Body.RouteResponse.Route)-1].Opcode == "648" {
					str := strings.Split(data.Body.RouteResponse.Route[len(data.Body.RouteResponse.Route)-1].Remark, " ")

					for _, v := range str {
						fmt.Println(v)
						fmt.Println(len(v))
					}
				}
				a++
			}

		}

	}

}
