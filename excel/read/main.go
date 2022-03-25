package main

import (
	"fmt"
	"github.com/FirewoodBloody/Golang/express_api/modules"
	"github.com/Luxurioust/excelize"
	"os"
)

func main() {
	excel, err := excelize.OpenFile("1.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//cell := excel.GetCellValue("Sheet1", "B1")
	//cell := excel.GetCellValue("Sheet1", "B1")
	//fmt.Println(cell)

	//index := excel.GetSheetIndex("Sheet2")

	rows := excel.GetRows("Sheet1")

	for k, _ := range rows {

		fmt.Println(k)
		if excel.GetCellValue("Sheet1", fmt.Sprintf("A%v", k+1)) == "顺丰快递" {
			a, err := modules.SfCreateData(excel.GetCellValue("Sheet1", fmt.Sprintf("B%v", k+1)))
			fmt.Println(excel.GetCellValue("Sheet1", fmt.Sprintf("B%v", k+1)))
			if err != nil {
				continue
			}

			if len(a.Body.RouteResponse.Route) == 0 {
				continue
			}
			fmt.Println(a.Body.RouteResponse.Route[len(a.Body.RouteResponse.Route)-1].Remark)
			excel.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), a.Body.RouteResponse.Route[len(a.Body.RouteResponse.Route)-1].Remark)
		} else if excel.GetCellValue("Sheet1", fmt.Sprintf("A%v", k+1)) == "京东快递" {
			a := modules.SelectData(excel.GetCellValue("Sheet1", fmt.Sprintf("B%v", k+1)))
			fmt.Println(excel.GetCellValue("Sheet1", fmt.Sprintf("B%v", k+1)))

			if len(a.Data) == 0 {
				continue
			}
			fmt.Println(a.Data[len(a.Data)-1].OpeRemark)
			excel.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), a.Data[len(a.Data)-1].OpeRemark)
		}
	}

	excel.Save()

}
