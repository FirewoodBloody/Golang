package main

import (
	"Golang/ExpressAIP/sfApi"
	"fmt"
	"github.com/Luxurioust/excelize"
	"io/ioutil"
	"strings"
)

const (
	UserId  = "BLWHYSP_tipea"                                            //顾客编码
	UserKey = "eJXycwXzucwcxit4a8WK7b3qGl4UfkB1"                         //秘钥
	ApiUrl  = "http://bsp-oisp.sf-express.com/bsp-oisp/sfexpressService" //请求地址
)

func main() {
	data, err := ioutil.ReadFile("./1.txt")
	if err != nil {
		fmt.Println(err)
	}

	dataStr := strings.Split(fmt.Sprintf("%s", data), ",")

	Xlsx := excelize.NewFile()
	// 创建一个工作表
	index := Xlsx.NewSheet("Sheet1")
	Xlsx.SetActiveSheet(index)
	//num := 2
	//Sheet := fmt.Sprintf("Sheet1")
	Xlsx.SetActiveSheet(index)
	for _, v := range dataStr {

		RequestParameters, err := sfApi.SfCreateData(UserId, UserKey, v) //初始化
		if err != nil {
			fmt.Println("create:", err)
		}
		fmt.Printf("%s%d\n", v, len(v))
		SfDataStruct, err := sfApi.SfPost(ApiUrl, RequestParameters) //post请求
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(SfDataStruct)

		//for i := 0; i < len(SfDataStruct.Body.RouteResponse.Route); i++ {
		//	Remark := SfDataStruct.Body.RouteResponse.Route[i].Remark
		//	Opcode := SfDataStruct.Body.RouteResponse.Route[i].Opcode
		//
		//	A := fmt.Sprintf("A%v", num)
		//	B := fmt.Sprintf("B%v", num)
		//
		//	// 设置单元格的值
		//	Xlsx.SetCellValue(Sheet, A, Opcode)
		//	Xlsx.SetCellValue(Sheet, B, Remark)
		//	num = num + 1
		//
		//}
	}
	//err = Xlsx.SaveAs("./1.xlsx")
	//if err != nil {
	//	fmt.Println("save", err)
	//}
}
