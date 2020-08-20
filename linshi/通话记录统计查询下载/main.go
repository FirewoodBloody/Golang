package main

import (
	"fmt"
	"github.com/Luxurioust/excelize"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"os"
	"time"

	"github.com/ying32/govcl/vcl"
)

type Engine struct {
	Engine *xorm.Engine
	Err    error
}

type Windows struct {
	win *vcl.TForm //窗口

	label_strat *vcl.TLabel //开始标签
	label_stop  *vcl.TLabel //结束标签
	//label_time  *vcl.TLabel

	//time_tedit *vcl.TEdit

	date_strat_label *vcl.TDateTimePicker //开始日期菜单
	date_stop_label  *vcl.TDateTimePicker //结束日期

	progress_bar *vcl.TProgressBar //进度条

	button *vcl.TButton //下载

	Engine Engine
}

const (
	TimeFormat = "2006-01-02"
	driverName = "mysql"
	dBconnect  = "root:dcf4aa3f7b982ce4@tcp(192.168.0.11:3306)/bl_crm?charset=utf8"
)

var tbMappers core.PrefixMapper

func init() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码

}

func (w *Windows) init() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	w.win = vcl.Application.CreateForm() //新建窗口
	w.win.SetCaption("出库记录下载")           //程序名
	//w.win.SetFormStyle(2)
	w.win.SetHeight(300)     //高
	w.win.SetWidth(400)      //宽
	w.win.ScreenCenter()     //居于当前屏幕中心
	w.win.SetBorderIcons(3)  //设置窗口最大化 最小化 关闭按钮  3代表是 最大化按钮和功能无效
	w.win.Font().SetSize(11) //整体字体大小
	w.win.Font().SetColor(255)
	w.win.Font().SetStyle(16) //字体样式
	w.win.SetColor(16775388)
	//w.win.SetTransparentColor(true)
	//w.win.SetTransparentColorValue(1)

	w.label_strat = vcl.NewLabel(w.win)
	w.label_strat.SetParent(w.win)
	w.label_strat.SetCaption("开始时间")
	w.label_strat.SetLeft(50) //设置按钮位置  横向
	w.label_strat.SetTop(50)  //设置按钮位置 竖向

	w.label_stop = vcl.NewLabel(w.win)
	w.label_stop.SetParent(w.win)
	w.label_stop.SetCaption("结束时间")
	w.label_stop.SetLeft(50)
	w.label_stop.SetTop(100)

	w.date_strat_label = vcl.NewDateTimePicker(w.win)
	w.date_strat_label.SetParent(w.win)
	w.date_strat_label.SetLeft(130)
	w.date_strat_label.SetTop(50)

	w.date_stop_label = vcl.NewDateTimePicker(w.win)
	w.date_stop_label.SetParent(w.win)
	w.date_stop_label.SetLeft(130)
	w.date_stop_label.SetTop(100)

	w.button = vcl.NewButton(w.win)
	w.button.SetParent(w.win)
	w.button.SetHeight(50)
	w.button.SetWidth(100)
	w.button.SetTop(150)
	w.button.SetLeft(150)
	w.button.SetCaption("开始下载")

	w.progress_bar = vcl.NewProgressBar(w.win)
	w.progress_bar.SetParent(w.win)
	w.progress_bar.SetPosition(0)
	w.progress_bar.SetWidth(400)
	w.progress_bar.SetHeight(20)
	w.progress_bar.SetLeft(0)
	w.progress_bar.SetTop(230)

}

func (e *Engine) NewEngine() error {

	e.Engine, e.Err = xorm.NewEngine(driverName, dBconnect)
	if e.Err != nil {
		return e.Err
	}
	//tbMappe := core.NewPrefixMapper(core.SnakeMapper{}, tbMapper)
	e.Engine.ShowSQL(true)
	e.Engine.SetTableMapper(tbMappers)
	return nil
}

func (w *Windows) Onclick() {

	w.button.SetOnClick(func(sender vcl.IObject) {

		//w.win.SetEnabled(false)
		w.button.SetEnabled(false)
		w.date_stop_label.SetEnabled(false)
		w.date_strat_label.SetEnabled(false)
		w.win.SetEnabled(false)

		time_strat := w.date_strat_label.DateTime().Format("2006-01-02")
		time_stop := w.date_stop_label.DateTime().Format("2006-01-02")

		err := w.Engine.NewEngine()
		if err != nil {
			vcl.ThreadSync(func() {
				vcl.ShowMessageFmt("数据库连接失败！")
			})
			w.button.SetEnabled(true)
			w.date_stop_label.SetEnabled(true)
			w.date_strat_label.SetEnabled(true)
			w.win.SetEnabled(true)
			return
		}
		nomao, _ := w.Engine.Engine.Query(fmt.Sprintf("SELECT\n\tmo.created_at AS 订单创建时间,\n\tumc.nickname AS 订单创建人,\n\tumc.login_name AS 订单创建人工号,\n\tdmc.`name` AS 订单创建人部门,\n\tcc.`name` AS 客户姓名,\n\tcc.customer_code AS 客户编码,\n\tmo.order_no AS 系统订单号,\n\tdi2.`name` AS 订单类型,\n\tdi.`name` AS 订单状态,\n\tmo.performance_at AS 回款日期,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"来源渠道\"' ) AS 线上下单渠道,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) AS 线上下单时间,\n\tumm.nickname AS 线上下单核单人,\n\tumm.login_name AS 线上下单核单人工号,\n\tdmm.`name` AS 线上下单核单人部门,\n\tmo.`name` AS 订单备注,\n\tump.nickname AS 业绩归属人,\n\tump.login_name AS 业绩归属人工号,\n\tdmp.`name` AS 业绩归属人部门,\n\tmo.total_amount / 100 AS 订单总金额,\n\tmc.`name` AS 订单商品,\n\tmc.goods_count AS 商品数量,\n\tmc.goods_price / 100 AS 商品原价,\n\tmc.price_sale / 100 AS 商品售价,\n\tmc.total_price / 100 AS 商品总价,\n\tei.ship_channel_name AS 配送方式,\n\tei.ship_channel_no AS 快递单号,\n\tei.last_trace AS 快递状态 \nFROM\n\tbl_mall_order mo\n\tLEFT JOIN bl_mall_cart mc ON mo.id = mc.order_id\n\tLEFT JOIN bl_express_invoice ei ON mo.id = ei.order_id\n\tLEFT JOIN bl_mall_order_media mom ON mo.id = mom.mall_order_id\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 43 ) di ON mo.`status` = di.item_value\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 53 ) di2 ON mo.type = di2.item_value\n\tLEFT JOIN bl_crm_customer cc ON mo.customer_id = cc.id\n\tLEFT JOIN bl_users umc ON mo.create_user_id = umc.id\n\tLEFT JOIN bl_depart dmc ON mo.create_user_depart_id = dmc.id\n\tLEFT JOIN bl_users ump ON mo.performance_user_id = ump.id\n\tLEFT JOIN bl_depart dmp ON mo.performance_user_depart_id = dmp.id\n\tLEFT JOIN bl_users umm ON mom.assign_user_id = umm.id\n\tLEFT JOIN bl_depart dmm ON umm.depart_id = dmm.id \nWHERE\n\tmo.deleted_at IS NULL \n\tAND mo.delete_user_id IS NULL \n\tAND mc.deleted_at IS NULL \n\tAND ei.created_at >= '%v 00:00:00' \n\tAND ei.created_at < '%v 23:59:59' \nORDER BY\n\tmo.created_at DESC", time_strat, time_stop))
		w.Engine.Engine.Clone()
		fmt.Println(len(nomao))
		f := excelize.NewFile()
		go func() {

			for k, v := range nomao {

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
					f.SetCellValue("Sheet1", fmt.Sprintf("J%v", k+1), "回款日期,")
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
				}
				f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+2), v["订单创建时间"])
				f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+2), v["订单创建人"])
				f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+2), v["订单创建人工号"])
				f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+2), v["订单创建人部门"])
				f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+2), v["客户姓名"])
				f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+2), v["客户编码"])
				f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+2), v["系统订单号"])
				f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+2), v["订单类型"])
				f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+2), v["订单状态"])
				f.SetCellValue("Sheet1", fmt.Sprintf("J%v", k+2), v["回款日期,"])
				f.SetCellValue("Sheet1", fmt.Sprintf("K%v", k+2), v["线上下单渠道"])
				f.SetCellValue("Sheet1", fmt.Sprintf("L%v", k+2), v["线上下单时间"])
				f.SetCellValue("Sheet1", fmt.Sprintf("M%v", k+2), v["线上下单核单人"])
				f.SetCellValue("Sheet1", fmt.Sprintf("N%v", k+2), v["线上下单核单人工号"])
				f.SetCellValue("Sheet1", fmt.Sprintf("O%v", k+2), v["线上下单核单人部门"])
				f.SetCellValue("Sheet1", fmt.Sprintf("P%v", k+2), v["订单备注"])
				f.SetCellValue("Sheet1", fmt.Sprintf("Q%v", k+2), v["业绩归属人"])
				f.SetCellValue("Sheet1", fmt.Sprintf("R%v", k+2), v["业绩归属人工号"])
				f.SetCellValue("Sheet1", fmt.Sprintf("S%v", k+2), v["业绩归属人部门"])
				f.SetCellValue("Sheet1", fmt.Sprintf("T%v", k+2), v["订单总金额"])
				f.SetCellValue("Sheet1", fmt.Sprintf("U%v", k+2), v["订单商品"])
				f.SetCellValue("Sheet1", fmt.Sprintf("V%v", k+2), v["商品数量"])
				f.SetCellValue("Sheet1", fmt.Sprintf("W%v", k+2), v["商品原价"])
				f.SetCellValue("Sheet1", fmt.Sprintf("X%v", k+2), v["商品售价"])
				f.SetCellValue("Sheet1", fmt.Sprintf("T%v", k+2), v["商品总价"])
				f.SetCellValue("Sheet1", fmt.Sprintf("Z%v", k+2), v["配送方式"])
				f.SetCellValue("Sheet1", fmt.Sprintf("AA%v", k+2), v["快递单号"])
				f.SetCellValue("Sheet1", fmt.Sprintf("AB%v", k+2), v["快递状态"])

				vcl.ThreadSync(func() {
					w.progress_bar.SetPosition(int32(float64(k) / float64(len(nomao)) * 100))
				})

			}
			if w.date_strat_label.DateTime().Format("2006-01-02") == w.date_stop_label.DateTime().Format("2006-01-02") {
				f.SaveAs(fmt.Sprintf("%v-出库记录明细.xlsx", w.date_strat_label.DateTime().Format("2006-01-02")))
				time.Sleep(time.Second)
				vcl.ThreadSync(func() {
					vcl.ShowMessageFmt("下载完成！")
				})
			} else {
				f.SaveAs(fmt.Sprintf("%v-%v-出库记录明细.xlsx", w.date_strat_label.DateTime().Format("2006-01-02"), w.date_stop_label.DateTime().Format("2006-01-02")))
				time.Sleep(time.Second)
				vcl.ThreadSync(func() {
					vcl.ShowMessageFmt("下载完成！")
				})
			}

			//w.button.SetEnabled(false)
			//w.win.SetEnabled(true)
			w.date_stop_label.SetEnabled(true)
			w.date_strat_label.SetEnabled(true)
			w.button.SetEnabled(true)
			w.win.SetEnabled(true)
			os.Exit(0)
		}()

	})

}

func main() {
	TForm := new(Windows)
	TForm.init()
	TForm.Onclick()

	TForm.win.Show()
	vcl.Application.Run()
}
