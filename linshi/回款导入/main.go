package main

import (
	"fmt"
	"github.com/Luxurioust/excelize"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/ying32/govcl/vcl/types"
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

	progress_bar *vcl.TProgressBar //进度条

	button      *vcl.TButton     //下载
	tOpenDialog *vcl.TOpenDialog //打开文件

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
	w.win.SetCaption("回款导入匹配")           //程序名
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

	w.button = vcl.NewButton(w.win)
	w.button.SetParent(w.win)
	w.button.SetHeight(50)
	w.button.SetWidth(100)
	w.button.SetTop(120)
	w.button.SetLeft(150)
	w.button.SetCaption("打开文件")

	w.progress_bar = vcl.NewProgressBar(w.win)
	w.progress_bar.SetParent(w.win)
	w.progress_bar.SetPosition(0)
	w.progress_bar.SetWidth(400)
	w.progress_bar.SetHeight(20)
	w.progress_bar.SetLeft(0)
	w.progress_bar.SetTop(230)

	w.tOpenDialog = vcl.NewOpenDialog(w.win)
	w.tOpenDialog.SetFilter("所有文件(*.*)")
	//    dlgOpen.SetInitialDir()
	//	dlgOpen.SetFilterIndex()

	w.tOpenDialog.SetOptions(w.tOpenDialog.Options().Include(types.OfShowHelp, types.OfAllowMultiSelect)) //rtl.Include(, types.OfShowHelp))
	w.tOpenDialog.SetTitle("打开")

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

		if w.tOpenDialog.Execute() { //选择打开文件
			//fmt.Println("filename: ", w.tOpenDialog.FileName())
			file, err := excelize.OpenFile(w.tOpenDialog.FileName())
			if err != nil {
				vcl.ThreadSync(func() {
					vcl.ShowMessage(fmt.Sprintln(err))
				})
				return
			}

			w.button.SetEnabled(false)
			w.win.SetEnabled(false)

			err = w.Engine.NewEngine()
			if err != nil {
				vcl.ThreadSync(func() {
					vcl.ShowMessageFmt("数据库连接失败！请联系管理员。")
				})
				w.button.SetEnabled(true)
				w.win.SetEnabled(true)
				return
			}

			for i := 1; i <= len(file.GetRows("Sheet1")); i++ {
				if i == 1 {
					if file.GetCellValue("Sheet1", fmt.Sprintf("A%v", i)) != "订单号" {
						vcl.ThreadSync(func() {
							vcl.ShowMessage(fmt.Sprintf("A%v，表头错误，正确表头信息：订单号", i))
						})
						w.button.SetEnabled(true)
						w.win.SetEnabled(true)
						return
					}
					if file.GetCellValue("Sheet1", fmt.Sprintf("B%v", i)) != "物流面单号" {
						vcl.ThreadSync(func() {
							vcl.ShowMessage(fmt.Sprintf("A%v，表头错误，正确表头信息：物流面单号", i))
						})
						w.button.SetEnabled(true)
						w.win.SetEnabled(true)
						return
					}
					if file.GetCellValue("Sheet1", fmt.Sprintf("C%v", i)) != "待收货款金额" {
						vcl.ThreadSync(func() {
							vcl.ShowMessage(fmt.Sprintf("A%v，表头错误，正确表头信息：待收货款金额", i))
						})
						w.button.SetEnabled(true)
						w.win.SetEnabled(true)
						return
					}
					if file.GetCellValue("Sheet1", fmt.Sprintf("D%v", i)) != "回款时间" {
						vcl.ThreadSync(func() {
							vcl.ShowMessage(fmt.Sprintf("A%v，表头错误，正确表头信息：回款时间", i))
						})
						w.button.SetEnabled(true)
						w.win.SetEnabled(true)
						return
					}
					if file.GetCellValue("Sheet1", fmt.Sprintf("E%v", i)) != "快递公司" {
						vcl.ThreadSync(func() {
							vcl.ShowMessage(fmt.Sprintf("A%v，表头错误，正确表头信息：快递公司", i))
						})
						w.button.SetEnabled(true)
						w.win.SetEnabled(true)
						return
					}
					vcl.ThreadSync(func() {
						w.progress_bar.SetPosition(int32(float64(i) / float64(len(file.GetRows("Sheet1"))) * 100))
					})
					continue
				}

				if file.GetCellValue("Sheet1", fmt.Sprintf("E%v", i)) == "顺丰快递" {
					if file.GetCellValue("Sheet1", fmt.Sprintf("B%v", i)) == "" {
						vcl.ThreadSync(func() {
							vcl.ShowMessage(fmt.Sprintf("D%v :快递面单信息为空，跳过此列！", i))
							w.progress_bar.SetPosition(int32(float64(i) / float64(len(file.GetRows("Sheet1"))) * 100))
						})
						continue
					}
					number, _ := w.Engine.Engine.Query(fmt.Sprintf("SELECT\n\tmo.order_no AS order_no \nFROM\n\tbl_mall_order mo\n\tLEFT JOIN bl_express_invoice ei ON mo.id = ei.order_id \nWHERE\n\tei.ship_channel_no = '%v'", file.GetCellValue("Sheet1", fmt.Sprintf("B%v", i))))
					for _, v := range number {
						fmt.Println(string(v["order_no"]))
						file.SetCellValue("Sheet1", fmt.Sprintf("A%v", i), string(v["order_no"]))
					}
					vcl.ThreadSync(func() {
						w.progress_bar.SetPosition(int32(float64(i) / float64(len(file.GetRows("Sheet1"))) * 100))
					})
				} else {

					vcl.ThreadSync(func() {
						w.progress_bar.SetPosition(int32(float64(i) / float64(len(file.GetRows("Sheet1"))) * 100))
					})
					continue
				}

			}
			vcl.ShowMessageFmt("匹配完成，正在保存！")
			file.Save()
			time.Sleep(time.Second * 1)
			vcl.ShowMessageFmt("保存完成！")

			w.button.SetEnabled(true)
			w.win.SetEnabled(true)
			os.Exit(0)
		}

		return
	})

}

func main() {
	TForm := new(Windows)
	TForm.init()
	TForm.Onclick()

	TForm.win.Show()
	vcl.Application.Run()
}
