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
	w.label_strat.SetLeft(100) //设置按钮位置  横向
	w.label_strat.SetTop(50)   //设置按钮位置 竖向

	w.label_stop = vcl.NewLabel(w.win)
	w.label_stop.SetParent(w.win)
	w.label_stop.SetCaption("结束时间")
	w.label_stop.SetLeft(100)
	w.label_stop.SetTop(100)

	w.date_strat_label = vcl.NewDateTimePicker(w.win)
	w.date_strat_label.SetParent(w.win)
	w.date_strat_label.SetLeft(180)
	w.date_strat_label.SetTop(50)

	w.date_stop_label = vcl.NewDateTimePicker(w.win)
	w.date_stop_label.SetParent(w.win)
	w.date_stop_label.SetLeft(180)
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
		nomao, _ := w.Engine.Engine.Query(fmt.Sprintf("SELECT\n\teic.created_at AS 出库日期,\n\teic.order_no AS 订单编号,\n\tIF( eic.`status` = 1 , \"订单\" , IF( eic.`status` = 2, \"工单\", \"\" )) AS 类型,\n\tdi.`name` as 订单类型,\n\tu.nickname AS 员工姓名,\n\tu.login_name AS 员工工号,\n\td.`name` AS 员工部门,\n\teic.goods_name AS 商品名称,\n\teic.goods_count AS 商品数量,\n\teic.amount AS 商品售价,\n\teic.warehouse_name AS 出货仓库,\n\teic.ship_channel_name AS 配送方式,\n\teic.ship_channel_no AS 快递单号 \nFROM\n\tbl_express_invoice_cart eic\n\tLEFT JOIN bl_mall_order mo ON eic.order_id = mo.id\n\tLEFT JOIN bl_users u ON mo.performance_user_id = u.id\n\tLEFT JOIN bl_depart d ON mo.performance_user_depart_id = d.id\n\tLEFT JOIN (SELECT id,name,item_value FROM bl_dict_item where dict_id = 53) di on mo.type = di.item_value\t\nWHERE\n\teic.created_at >= '%v 00:00:00' \n\tAND eic.created_at < '%v 23:59:59' \n\tAND mo.deleted_at IS NULL \nORDER BY\n\teic.created_at,\n\teic.order_no,\n\teic.goods_name,\n\tu.nickname,\n\td.`name`", time_strat, time_stop))
		w.Engine.Engine.Clone()
		fmt.Println(len(nomao))
		f := excelize.NewFile()
		go func() {

			for k, v := range nomao {

				if k == 0 {
					f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+1), "出库日期")
					f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+1), "订单编号")
					f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), "类型")
					f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+1), "订单类型")
					f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+1), "员工姓名")
					f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+1), "员工工号")
					f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+1), "员工部门")
					f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+1), "商品名称")
					f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+1), "商品数量")
					f.SetCellValue("Sheet1", fmt.Sprintf("J%v", k+1), "商品售价")
					f.SetCellValue("Sheet1", fmt.Sprintf("K%v", k+1), "出货仓库")
					f.SetCellValue("Sheet1", fmt.Sprintf("L%v", k+1), "配送方式")
					f.SetCellValue("Sheet1", fmt.Sprintf("M%v", k+1), "快递单号")
				}
				f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+2), v["出库日期"])
				f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+2), v["订单编号"])
				f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+2), v["类型"])
				f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+2), v["订单类型"])
				f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+2), v["员工姓名"])
				f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+2), v["员工工号"])
				f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+2), v["员工部门"])
				f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+2), v["商品名称"])
				f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+2), v["商品数量"])
				f.SetCellValue("Sheet1", fmt.Sprintf("J%v", k+2), v["商品售价"])
				f.SetCellValue("Sheet1", fmt.Sprintf("K%v", k+2), v["出货仓库"])
				f.SetCellValue("Sheet1", fmt.Sprintf("L%v", k+2), v["配送方式"])
				f.SetCellValue("Sheet1", fmt.Sprintf("M%v", k+2), v["快递单号"])

				vcl.ThreadSync(func() {
					w.progress_bar.SetPosition(int32(float64(k) / float64(len(nomao)) * 100))
				})

			}
			if w.date_strat_label.DateTime().Format("2006-01-02") == w.date_stop_label.DateTime().Format("2006-01-02") {
				f.SaveAs(fmt.Sprintf("%v_出库记录明细.xlsx", w.date_strat_label.DateTime().Format("2006-01-02")))
				time.Sleep(time.Second)
				vcl.ThreadSync(func() {
					w.progress_bar.SetPosition(1)
					vcl.ShowMessageFmt("下载完成！")
				})
			} else {
				f.SaveAs(fmt.Sprintf("%v_%v_出库记录明细.xlsx", w.date_strat_label.DateTime().Format("2006-01-02"), w.date_stop_label.DateTime().Format("2006-01-02")))
				time.Sleep(time.Second)
				vcl.ThreadSync(func() {
					w.progress_bar.SetPosition(1)
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
