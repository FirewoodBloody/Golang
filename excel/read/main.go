package main

import (
	"fmt"
	"github.com/Luxurioust/excelize"
	"os"
)

func main() {
	excel, err := excelize.OpenFile("已购客户.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cell := excel.GetCellValue("Sheet1", "B1")
	cell := excel.GetCellValue("Sheet1", "B1")
	fmt.Println(cell)

	//index := excel.GetSheetIndex("Sheet2")

	rows := excel.GetRows("Sheet2")

	for _, row := range rows {
		for _, colcell := range row {
			fmt.Print(colcell, "\t")
		}
		fmt.Println()
	}
}
