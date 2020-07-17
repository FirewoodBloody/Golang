package main

import (
	"Golang/express_api/modules"
	"fmt"
	"github.com/Luxurioust/excelize"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"xorm.io/xorm"
)

const (
	TimeFormat = "2006-01-02"
	driverName = "mysql"
	dBconnect  = "root:dcf4aa3f7b982ce4@tcp(192.168.0.11:3306)/bl_crm?charset=utf8"
)

type Delivery_routing struct {
	order_no         string
	order_no_routing string
	ship_channel_no  string
	ship_channel_id  string
	Strat            Strat
}

type Strat struct {
	Strat_T string
	Strat_Q string
}

func init() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码
}

func main() {
	file, _ := excelize.OpenFile("./Delivery_routing.xlsx")

	var err error

	engine, err := xorm.NewEngine(driverName, dBconnect)
	if err != nil {
		fmt.Println(err)
	}

	Delivery_routing := new(Delivery_routing)

	engine.ShowSQL(true)
	KD, _ := engine.Query("SELECT mo.order_no,di.`name`,ei.ship_channel_no,ei.ship_channel_id FROM bl_mall_order mo INNER JOIN bl_express_invoice ei ON mo.id = ei.order_id INNER JOIN bl_dict_item di ON mo.`status` = di.item_value WHERE ei.ship_channel_no IS NOT NULL AND di.dict_id = 43 and mo.deleted_at IS  NULL")

	engine.Close()
	s := 0
	for _, v := range KD {

		Delivery_routing.order_no = string(v["order_no"])
		Delivery_routing.order_no_routing = string(v["name"])
		Delivery_routing.ship_channel_no = string(v["ship_channel_no"])
		Delivery_routing.ship_channel_id = string(v["ship_channel_id"])
		Delivery_routing.Strat.Strat_Q = ""
		Delivery_routing.Strat.Strat_T = ""
		Delivery_routing_list(Delivery_routing)

		file.SetCellValue("Sheet1", fmt.Sprintf("A%v", s+1), Delivery_routing.order_no)
		file.SetCellValue("Sheet1", fmt.Sprintf("B%v", s+1), Delivery_routing.order_no_routing)
		file.SetCellValue("Sheet1", fmt.Sprintf("C%v", s+1), Delivery_routing.ship_channel_no)
		file.SetCellValue("Sheet1", fmt.Sprintf("D%v", s+1), Delivery_routing.Strat.Strat_T)
		file.SetCellValue("Sheet1", fmt.Sprintf("E%v", s+1), Delivery_routing.Strat.Strat_Q)
		s++
	}
	file.Save()
}

func Delivery_routing_list(Delivery_routing *Delivery_routing) {

	if Delivery_routing.ship_channel_id == "1" {
		sf, _ := modules.SfCreateData(Delivery_routing.ship_channel_no)
		for _, i := range sf.Body.RouteResponse.Route {
			if i.Opcode == "648" || i.Opcode == "99" {
				Delivery_routing.Strat.Strat_T = "快递退回"
				continue
			} else if i.Opcode == "80" {
				Delivery_routing.Strat.Strat_Q = "已签收"
				continue
			} else if i.Opcode == "8080" {
				Delivery_routing.Strat.Strat_Q = "已签收"
				continue
			} else {
				if Delivery_routing.Strat.Strat_Q != "" {
					continue
				}
				Delivery_routing.Strat.Strat_Q = "转运中"
				continue
			}

		}

	} else if Delivery_routing.ship_channel_id == "2" {
		jd := modules.SelectData(Delivery_routing.ship_channel_no)
		for _, i := range jd.Data {
			if i.State == "160" {
				Delivery_routing.Strat.Strat_T = "快递退回"
				continue
			} else if i.State == "150" {
				Delivery_routing.Strat.Strat_Q = "已签收"
				continue
			} else {
				if Delivery_routing.Strat.Strat_Q != "" {
					continue
				}
				Delivery_routing.Strat.Strat_Q = "转运中"
				continue
			}
		}

	}
}
