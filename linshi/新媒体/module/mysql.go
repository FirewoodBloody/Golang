package module

import (
	"fmt"
	"github.com/FirewoodBloody/Golang/Express_Routing/express"
	"github.com/FirewoodBloody/Golang/express_api/modules"
	"github.com/Luxurioust/excelize"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ying32/govcl/vcl"
	"time"
)

func Order_select(w *Windows, f *excelize.File) {
	err := w.Engine.NewEngine()
	if err != nil {
		vcl.ThreadSync(func() {
			vcl.ShowMessageFmt("数据库连接失败！请联系管理员。")
		})

		w.Button.SetEnabled(true)
		w.Win.SetEnabled(true)
		return
	}

	resulist, _ := w.Engine.Engine.Query(fmt.Sprintf("SELECT\n\tmo.created_at AS 订单创建时间,\n\tumc.nickname AS 订单创建人,\n\tumc.login_name AS 订单创建人工号,\n\tdmc.`name` AS 订单创建人部门,\n\tcc.`name` AS 客户姓名,\n\tcc.customer_code AS 客户编码,\n\tmo.order_no AS 系统订单号,\n\tdi2.`name` AS 订单类型,\n\tdi.`name` AS 订单状态,\n\tmo.performance_at AS 回款日期,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"来源渠道\"' ) AS 线上下单渠道,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) AS 线上下单时间,\n\tumm.nickname AS 线上下单核单人,\n\tumm.login_name AS 线上下单核单人工号,\n\tdmm.`name` AS 线上下单核单人部门,\n\tmo.`name` AS 订单备注,\n\tump.nickname AS 业绩归属人,\n\tump.login_name AS 业绩归属人工号,\n\tdmp.`name` AS 业绩归属人部门,\n\tmo.total_amount / 100 AS 订单总金额,\n\tmc.`name` AS 订单商品,\n\tmc.goods_count AS 商品数量,\n\tmc.goods_price / 100 AS 商品原价,\n\tmc.price_sale / 100 AS 商品售价,\n\tmc.total_price / 100 AS 商品总价,\n\tei.ship_channel_name AS 配送方式,\n\tei.ship_channel_no AS 快递单号,\n\tei.last_trace AS 快递状态 \nFROM\n\tbl_mall_order mo\n\tLEFT JOIN bl_mall_cart mc ON mo.id = mc.order_id\n\tLEFT JOIN bl_express_invoice ei ON mo.id = ei.order_id\n\tLEFT JOIN bl_mall_order_media mom ON mo.id = mom.mall_order_id\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 43 ) di ON mo.`status` = di.item_value\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 53 ) di2 ON mo.type = di2.item_value\n\tLEFT JOIN bl_crm_customer cc ON mo.customer_id = cc.id\n\tLEFT JOIN bl_users umc ON mo.create_user_id = umc.id\n\tLEFT JOIN bl_depart dmc ON mo.create_user_depart_id = dmc.id\n\tLEFT JOIN bl_users ump ON mo.performance_user_id = ump.id\n\tLEFT JOIN bl_depart dmp ON mo.performance_user_depart_id = dmp.id\n\tLEFT JOIN bl_users umm ON mom.assign_user_id = umm.id\n\tLEFT JOIN bl_depart dmm ON umm.depart_id = dmm.id \nWHERE\n\tmo.deleted_at IS NULL \n\tAND mo.delete_user_id IS NULL \n\tAND mc.deleted_at IS NULL \n\tAND dmp.id IN ( 33, 91, 92, 107) \n\tAND ((mo.created_at > '%v 00:00:00' AND mo.created_at <= '%v 23:59:59') OR (mo.created_at < '%v 00:00:00' AND mo.performance_at IS NULL AND mo.`status` <= 90) OR (mo.created_at < '%v 00:00:00' AND mo.performance_at >'%v 00:00:00' and mo.performance_at <= '%v 23:59:59' )) \nORDER BY\n\tmo.created_at DESC", w.Date_strat_label.DateTime().Format("2006-01-02"), w.Date_stop_label.DateTime().Format("2006-01-02"), w.Date_strat_label.DateTime().Format("2006-01-02"), w.Date_strat_label.DateTime().Format("2006-01-02"), w.Date_strat_label.DateTime().Format("2006-01-02"), w.Date_stop_label.DateTime().Format("2006-01-02")))

	w.Engine.Engine.Close()
	fmt.Println(len(resulist))
	time.Sleep(time.Second * 5)

	for k, v := range resulist {
		if k == 0 {
			f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+1), "订单创建时间")
			f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+1), "订单创建人")
			f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), "订单创建人工号")
			f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+1), "订单创建人部门")
			f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+1), "客户姓名")
			f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+1), "客户编码")
			f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+1), "系统订单号")
			f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+1), "订单类型")
			f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+1), "订单状态")
			f.SetCellValue("Sheet1", fmt.Sprintf("J%v", k+1), "回款日期")
			f.SetCellValue("Sheet1", fmt.Sprintf("K%v", k+1), "线上下单渠道")
			f.SetCellValue("Sheet1", fmt.Sprintf("L%v", k+1), "线上下单时间")
			f.SetCellValue("Sheet1", fmt.Sprintf("M%v", k+1), "线上下单核单人")
			f.SetCellValue("Sheet1", fmt.Sprintf("N%v", k+1), "线上下单核单人工号")
			f.SetCellValue("Sheet1", fmt.Sprintf("O%v", k+1), "线上下单核单人部门")
			f.SetCellValue("Sheet1", fmt.Sprintf("P%v", k+1), "订单备注")
			f.SetCellValue("Sheet1", fmt.Sprintf("Q%v", k+1), "业绩归属人")
			f.SetCellValue("Sheet1", fmt.Sprintf("R%v", k+1), "业绩归属人工号")
			f.SetCellValue("Sheet1", fmt.Sprintf("S%v", k+1), "业绩归属人部门")
			f.SetCellValue("Sheet1", fmt.Sprintf("T%v", k+1), "订单总金额")
			f.SetCellValue("Sheet1", fmt.Sprintf("U%v", k+1), "订单商品")
			f.SetCellValue("Sheet1", fmt.Sprintf("V%v", k+1), "商品数量")
			f.SetCellValue("Sheet1", fmt.Sprintf("W%v", k+1), "商品原价")
			f.SetCellValue("Sheet1", fmt.Sprintf("X%v", k+1), "商品售价")
			f.SetCellValue("Sheet1", fmt.Sprintf("Y%v", k+1), "商品总价")
			f.SetCellValue("Sheet1", fmt.Sprintf("Z%v", k+1), "配送方式")
			f.SetCellValue("Sheet1", fmt.Sprintf("AA%v", k+1), "快递单号")
			f.SetCellValue("Sheet1", fmt.Sprintf("AB%v", k+1), "快递状态")
			vcl.ThreadSync(func() {
				w.Progress_bar.SetPosition(int32(float64(k+1) / float64(len(resulist)) * 100))
			})
			continue
		}
		f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+1), v["订单创建时间"])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+1), v["订单创建人"])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), v["订单创建人工号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+1), v["订单创建人部门"])
		f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+1), v["客户姓名"])
		f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+1), v["客户编码"])
		f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+1), v["系统订单号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+1), v["订单类型"])
		f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+1), v["订单状态"])
		f.SetCellValue("Sheet1", fmt.Sprintf("J%v", k+1), v["回款日期"])
		f.SetCellValue("Sheet1", fmt.Sprintf("K%v", k+1), v["线上下单渠道"])
		f.SetCellValue("Sheet1", fmt.Sprintf("L%v", k+1), v["线上下单时间"])
		f.SetCellValue("Sheet1", fmt.Sprintf("M%v", k+1), v["线上下单核单人"])
		f.SetCellValue("Sheet1", fmt.Sprintf("N%v", k+1), v["线上下单核单人工号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("O%v", k+1), v["线上下单核单人部门"])
		f.SetCellValue("Sheet1", fmt.Sprintf("P%v", k+1), v["订单备注"])
		f.SetCellValue("Sheet1", fmt.Sprintf("Q%v", k+1), v["业绩归属人"])
		f.SetCellValue("Sheet1", fmt.Sprintf("R%v", k+1), v["业绩归属人工号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("S%v", k+1), v["业绩归属人部门"])
		f.SetCellValue("Sheet1", fmt.Sprintf("T%v", k+1), v["订单总金额"])
		f.SetCellValue("Sheet1", fmt.Sprintf("U%v", k+1), v["订单商品"])
		f.SetCellValue("Sheet1", fmt.Sprintf("V%v", k+1), v["商品数量"])
		f.SetCellValue("Sheet1", fmt.Sprintf("W%v", k+1), v["商品原价"])
		f.SetCellValue("Sheet1", fmt.Sprintf("X%v", k+1), v["商品售价"])
		f.SetCellValue("Sheet1", fmt.Sprintf("Y%v", k+1), v["商品总价"])
		f.SetCellValue("Sheet1", fmt.Sprintf("Z%v", k+1), v["配送方式"])
		f.SetCellValue("Sheet1", fmt.Sprintf("AA%v", k+1), v["快递单号"])
		if string(v["订单状态"]) == "订单配送中" || string(v["订单状态"]) == "订单未妥投" {
			if string(v["快递单号"]) != "" {
				if string(v["配送方式"]) == "顺丰快递" {
					data, _ := express.SfCreateData(string(v["快递单号"]))
					if len(data.Body.RouteResponse.Route) != 0 {
						fmt.Println(data.Body.RouteResponse.Route[len(data.Body.RouteResponse.Route)-1].Remark)
						f.SetCellValue("Sheet1", fmt.Sprintf("AB%v", k+1), data.Body.RouteResponse.Route[len(data.Body.RouteResponse.Route)-1].Remark)
					}
				} else if string(v["配送方式"]) == "京东快递" {
					data := modules.SelectData(string(v["快递单号"]))
					if len(data.Data) != 0 {
						fmt.Println(data.Data[len(data.Data)-1].OpeRemark)
						f.SetCellValue("Sheet1", fmt.Sprintf("AB%v", k+1), data.Data[len(data.Data)-1].OpeRemark)
					}
				}
			}
		} else {
			f.SetCellValue("Sheet1", fmt.Sprintf("AB%v", k+1), v["快递状态"])
		}
		vcl.ThreadSync(func() {
			w.Progress_bar.SetPosition(int32(float64(k+1) / float64(len(resulist)) * 100))
		})
	}

	vcl.ThreadSync(func() {
		w.Progress_bar.SetPosition(100)
		f.SaveAs(w.TsaveDialog.FileName())
		time.Sleep(time.Second * 1)
		vcl.ShowMessageFmt("保存完成！")
	})

	w.Button.SetEnabled(true)
	w.Win.SetEnabled(true)
	//os.Exit(0)
}

func Order_xs_select(w *Windows, f *excelize.File) {
	err := w.Engine.NewEngine()
	if err != nil {
		vcl.ThreadSync(func() {
			vcl.ShowMessageFmt("数据库连接失败！请联系管理员。")
		})

		w.Button.SetEnabled(true)
		w.Win.SetEnabled(true)
		return
	}

	resulist, _ := w.Engine.Engine.Query(fmt.Sprintf("SELECT\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) AS 线上下单时间,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"来源渠道\"' ) AS 线上下单渠道,\n\tcc.customer_code AS 客户编码,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"客户名称\"' ) AS 下单客户姓名,\n\t#JSON_UNQUOTE( mom.raw_json -> '$[0].\"客户电话\"' ) AS 下单客户电话,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"商品名称\"' ) AS 下单商品,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"渠道订单号\"' ) AS 渠道订单号,\n\tIF(mom.`status`= 0,\"已取消\",IF(mom.`status` = 1,\"已创建\",IF(mom.`status` = 10,\"已分配\",IF(mom.`status` = 30 ,\"核单通过\",IF(mom.`status` = 40 ,\"核单失败\",\"0\"))))) as 核单状态,\n\tdi1.`name` as 核单备注,\n\td.`name` AS 部门,\n\tu.nickname AS 员工姓名,\n\tu.login_name AS 员工工号,\n\tmo.order_no AS 系统订单号,\n\tdi2.`name` AS 订单类型,\n\tdi.`name` as 订单状态,\n\tmo.total_amount / 100 AS 订单总金额,\n\tmc.`name` AS 订单商品,\n\tmc.goods_count AS 商品数量, \n\tmc.goods_price / 100 AS 商品原价,\n\tmc.price_sale / 100 AS 商品售价,\n\tmc.total_price / 100 AS 商品总价,\n\tei.ship_channel_name AS 配送方式,\n\tei.ship_channel_no AS 快递单号,\n\tei.last_trace AS 快递状态 \nFROM\n\t`bl_mall_order_media` mom\n\tLEFT JOIN bl_crm_customer cc on cc.id = mom.customer_id\n\tLEFT JOIN bl_express_invoice ei ON mom.mall_order_id = ei.order_id \n\tLEFT JOIN (SELECT id,name,item_value FROM bl_dict_item where dict_id = 74) di1 ON mom.verify_order_comment = di1.item_value\n\tLEFT JOIN bl_users u ON mom.assign_user_id = u.id\n\tLEFT JOIN bl_mall_order mo ON mom.mall_order_id = mo.id\n\tLEFT JOIN bl_depart d ON u.depart_id = d.id\t\n\tLEFT JOIN (SELECT id,name,item_value FROM bl_dict_item where dict_id = 43) di on mo.`status` = di.item_value\n\tLEFT JOIN (SELECT id,name,item_value FROM bl_dict_item where dict_id = 53) di2 on mo.type = di2.item_value\n\tLEFT JOIN bl_mall_cart mc ON mo.id = mc.order_id\n\t\nWHERE\n\tmom.deleted_at IS NULL \n\tAND mom.delete_user_id IS NULL\n\tAND d.id in (33,91,92)\n\tAND  ((JSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) >= '%v 00:00:00' \tAND JSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) <= '%v 23:59:59' )\n\tOR (JSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) < '%v 00:00:00' AND mom.mall_order_id IS NOT NULL AND mo.`status` <= 90 AND mo.performance_at IS NULL) OR (JSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) < '%v 00:00:00' AND mom.mall_order_id IS  NULL AND mom.`status` > 30 ))\n\tORDER BY\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' )\n\t", w.Date_strat_label.DateTime().Format("2006-01-02"), w.Date_stop_label.DateTime().Format("2006-01-02"), w.Date_strat_label.DateTime().Format("2006-01-02"), w.Date_strat_label.DateTime().Format("2006-01-02")))

	w.Engine.Engine.Close()

	time.Sleep(time.Second * 5)

	for k, v := range resulist {
		if k == 0 {
			f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+1), "线上下单时间")
			f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+1), "线上下单渠道")
			f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), "客户编码")
			f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+1), "下单客户姓名")
			f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+1), "下单商品")
			f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+1), "渠道订单号")
			f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+1), "核单状态")
			f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+1), "核单备注")
			f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+1), "部门")
			f.SetCellValue("Sheet1", fmt.Sprintf("J%v", k+1), "员工姓名")
			f.SetCellValue("Sheet1", fmt.Sprintf("K%v", k+1), "员工工号")
			f.SetCellValue("Sheet1", fmt.Sprintf("L%v", k+1), "系统订单号")
			f.SetCellValue("Sheet1", fmt.Sprintf("M%v", k+1), "订单类型")
			f.SetCellValue("Sheet1", fmt.Sprintf("N%v", k+1), "订单状态")
			f.SetCellValue("Sheet1", fmt.Sprintf("O%v", k+1), "订单总金额")
			f.SetCellValue("Sheet1", fmt.Sprintf("P%v", k+1), "订单商品")
			f.SetCellValue("Sheet1", fmt.Sprintf("Q%v", k+1), "商品数量")
			f.SetCellValue("Sheet1", fmt.Sprintf("R%v", k+1), "商品原价")
			f.SetCellValue("Sheet1", fmt.Sprintf("S%v", k+1), "商品售价")
			f.SetCellValue("Sheet1", fmt.Sprintf("T%v", k+1), "商品总价")
			f.SetCellValue("Sheet1", fmt.Sprintf("U%v", k+1), "配送方式")
			f.SetCellValue("Sheet1", fmt.Sprintf("V%v", k+1), "快递单号")
			f.SetCellValue("Sheet1", fmt.Sprintf("W%v", k+1), "快递状态")
			vcl.ThreadSync(func() {
				w.Progress_bar.SetPosition(int32(float64(k+1) / float64(len(resulist)) * 100))
			})
			continue
		}
		f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+1), v["线上下单时间"])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+1), v["线上下单渠道"])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), v["客户编码"])
		f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+1), v["下单客户姓名"])
		f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+1), v["下单商品"])
		f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+1), v["渠道订单号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+1), v["核单状态"])
		f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+1), v["核单备注"])
		f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+1), v["部门"])
		f.SetCellValue("Sheet1", fmt.Sprintf("J%v", k+1), v["员工姓名"])
		f.SetCellValue("Sheet1", fmt.Sprintf("K%v", k+1), v["员工工号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("L%v", k+1), v["系统订单号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("M%v", k+1), v["订单类型"])
		f.SetCellValue("Sheet1", fmt.Sprintf("N%v", k+1), v["订单状态"])
		f.SetCellValue("Sheet1", fmt.Sprintf("O%v", k+1), v["订单总金额"])
		f.SetCellValue("Sheet1", fmt.Sprintf("P%v", k+1), v["订单商品"])
		f.SetCellValue("Sheet1", fmt.Sprintf("Q%v", k+1), v["商品数量"])
		f.SetCellValue("Sheet1", fmt.Sprintf("R%v", k+1), v["商品原价"])
		f.SetCellValue("Sheet1", fmt.Sprintf("S%v", k+1), v["商品售价"])
		f.SetCellValue("Sheet1", fmt.Sprintf("T%v", k+1), v["商品总价"])
		f.SetCellValue("Sheet1", fmt.Sprintf("U%v", k+1), v["配送方式"])
		f.SetCellValue("Sheet1", fmt.Sprintf("V%v", k+1), v["快递单号"])
		if string(v["订单状态"]) == "订单配送中" || string(v["订单状态"]) == "订单未妥投" {
			if string(v["快递单号"]) != "" {
				if string(v["配送方式"]) == "顺丰快递" {
					data, _ := express.SfCreateData(string(v["快递单号"]))
					if len(data.Body.RouteResponse.Route) != 0 {
						f.SetCellValue("Sheet1", fmt.Sprintf("W%v", k+1), data.Body.RouteResponse.Route[len(data.Body.RouteResponse.Route)-1].Remark)
					}
				} else if string(v["配送方式"]) == "京东快递" {
					data := modules.SelectData(string(v["快递单号"]))
					if len(data.Data) != 0 {
						f.SetCellValue("Sheet1", fmt.Sprintf("W%v", k+1), data.Data[len(data.Data)-1].OpeRemark)
					}
				}
			}
		} else {
			f.SetCellValue("Sheet1", fmt.Sprintf("W%v", k+1), v["快递状态"])
		}
		vcl.ThreadSync(func() {
			w.Progress_bar.SetPosition(int32(float64(k+1) / float64(len(resulist)) * 100))
		})
	}
	vcl.ThreadSync(func() {
		w.Progress_bar.SetPosition(100)
		f.SaveAs(w.TsaveDialog.FileName())
		time.Sleep(time.Second * 1)
		vcl.ShowMessageFmt("保存完成！")
	})

	w.Button.SetEnabled(true)
	w.Win.SetEnabled(true)
	//os.Exit(0)
}

func Order_xsTJ_select(w *Windows, f *excelize.File) {
	err := w.Engine.NewEngine()
	if err != nil {
		vcl.ThreadSync(func() {
			vcl.ShowMessageFmt("数据库连接失败！请联系管理员。")
		})

		w.Button.SetEnabled(true)
		w.Win.SetEnabled(true)
		return
	}

	resulist, _ := w.Engine.Engine.Query(fmt.Sprintf("SELECT\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"来源渠道\"' ) AS 订单渠道,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"商品名称\"' ) AS 投放产品,\n\tcount( DISTINCT mom.id ) AS 后台订单数,\n\tcount( DISTINCT mom.mall_order_no ) AS 发货数量,\n\tsum(JSON_UNQUOTE( mom.raw_json -> '$[0].\"总金额\"' )) AS 原订单总金额,\n\tsum(  mc.total_price/ 100) AS 成交金额\nFROM\n\t`bl_mall_order_media` mom\n\tLEFT JOIN bl_mall_order mo ON mom.mall_order_id = mo.id\n\tLEFT JOIN bl_mall_cart mc ON mo.id = mc.order_id\n\tLEFT JOIN bl_mall_goods mg ON mc.goods_id = mg.id \nWHERE\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) >= '%v 00:00:00' \n\tAND JSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) < '%v 23:59:59' \n\tAND mom.deleted_at IS NULL \nGROUP BY\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"来源渠道\"' ),\n\tJSON_UNQUOTE(\n\tmom.raw_json -> '$[0].\"商品名称\"' \n\t)", w.Date_strat_label.DateTime().Format("2006-01-02"), w.Date_stop_label.DateTime().Format("2006-01-02")))

	w.Engine.Engine.Close()

	time.Sleep(time.Second * 5)

	for k, v := range resulist {
		if k == 0 {
			f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+1), "订单渠道")
			f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+1), "投放产品")
			f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), "后台订单数")
			f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+1), "发货数量")
			f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+1), "原订单总金额")
			f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+1), "成交金额")

			vcl.ThreadSync(func() {
				w.Progress_bar.SetPosition(int32(float64(k+1) / float64(len(resulist)) * 100))
			})
			continue
		}
		f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+1), v["订单渠道"])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+1), v["投放产品"])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), v["后台订单数"])
		f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+1), v["发货数量"])
		f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+1), v["原订单总金额"])
		f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+1), v["成交金额"])
		vcl.ThreadSync(func() {
			w.Progress_bar.SetPosition(int32(float64(k+1) / float64(len(resulist)) * 100))
		})
	}
	vcl.ThreadSync(func() {
		w.Progress_bar.SetPosition(100)
		f.SaveAs(w.TsaveDialog.FileName())
		time.Sleep(time.Second * 1)
		vcl.ShowMessageFmt("保存完成！")
	})

	w.Button.SetEnabled(true)
	w.Win.SetEnabled(true)
	//os.Exit(0)
}

func Order_xxTJ_select(w *Windows, f *excelize.File) {
	err := w.Engine.NewEngine()
	if err != nil {
		vcl.ThreadSync(func() {
			vcl.ShowMessageFmt("数据库连接失败！请联系管理员。")
		})

		w.Button.SetEnabled(true)
		w.Win.SetEnabled(true)
		return
	}

	resulist, _ := w.Engine.Engine.Query(fmt.Sprintf("SELECT\n\tmc.`name` AS 二开产品,\n\tCOUNT( DISTINCT mo.id ) AS 下单数, \n\tsum( mc.ratio_price / 100 ) AS 增购订单总金额,\n\tsum( mc.total_price / 100 ) AS 增购成交金额 \nFROM\n\t`bl_mall_order` mo\n\tLEFT JOIN bl_mall_cart mc ON mo.id = mc.order_id \nWHERE\n\tmo.created_at >= '%v 00:00:00' \n\tAND mo.created_at < '%v 23:59:59' \n\tAND mo.deleted_at IS NULL \n\tAND mo.id NOT IN ( SELECT mall_order_id FROM bl_mall_order_media WHERE mall_order_id IS NOT NULL ) \n\tand mo.performance_user_depart_id in (91,92,33,11)\n\tand mo.`status` > 1\nGROUP BY\n\tmc.`name`", w.Date_strat_label.DateTime().Format("2006-01-02"), w.Date_stop_label.DateTime().Format("2006-01-02")))

	w.Engine.Engine.Close()

	time.Sleep(time.Second * 5)

	for k, v := range resulist {
		if k == 0 {
			f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+1), "二开产品")
			f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+1), "下单数")
			f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), "增购订单总金额")
			f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+1), "增购成交金额")
			vcl.ThreadSync(func() {
				w.Progress_bar.SetPosition(int32(float64(k+1) / float64(len(resulist)) * 100))
			})
			continue
		}
		f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+1), v["二开产品"])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+1), v["下单数"])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), v["增购订单总金额"])
		f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+1), v["增购成交金额"])
		vcl.ThreadSync(func() {
			w.Progress_bar.SetPosition(int32(float64(k+1) / float64(len(resulist)) * 100))
		})
	}
	vcl.ThreadSync(func() {
		w.Progress_bar.SetPosition(100)
		f.SaveAs(w.TsaveDialog.FileName())
		time.Sleep(time.Second * 1)
		vcl.ShowMessageFmt("保存完成！")
	})

	w.Button.SetEnabled(true)
	w.Win.SetEnabled(true)
	//os.Exit(0)
}
