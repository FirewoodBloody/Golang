package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Luxurioust/excelize"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"

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
	w.win.SetCaption("录音下载")             //程序名
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
			return
		}
		nomao, _ := w.Engine.Engine.Query(fmt.Sprintf("SELECT\n\td1.NAME AS 部门,\n\td.NAME AS 二级部门,\n\tu.nickname AS 员工姓名,\n\tu.login_name AS 员工工号,\n\tCOUNT(\n\tIF\n\t( cqv.call_duration = 0, cqv.id, NULL )) AS 未接通话总数,\n\tSUM(\n\tIF\n\t( cqv.call_duration = 0, cqv.call_duration, 0 )) AS 未接通时长,\n\tCOUNT(\n\tIF\n\t( cqv.call_duration > 0 && cqv.call_duration <= 60, cqv.id, NULL )) AS 无效通话总数,\n\tSUM(\n\tIF\n\t( cqv.call_duration > 0 && cqv.call_duration <= 60, cqv.call_duration, 0 )) AS 无效通话时长,\n\tCOUNT(\n\tIF\n\t( cqv.call_duration > 60 && cqv.call_duration <= 180, cqv.id, NULL )) AS 有效通话总数,\n\tSUM(\n\tIF\n\t( cqv.call_duration > 60 && cqv.call_duration <= 180, cqv.call_duration, 0 )) AS 有效通话时长,\n\tCOUNT(\n\tIF\n\t( cqv.call_duration > 180, cqv.id, NULL )) AS 优质通话总数,\n\tSUM(\n\tIF\n\t( cqv.call_duration > 180, cqv.call_duration, 0 )) AS 优质通话时长,\n\tCOUNT( cqv.id ) AS 总通话数,\n\tSUM( cqv.call_duration ) AS 通话总时长 \nFROM\n\tbl_crm_quality_voice cqv\n\tINNER JOIN bl_users u ON cqv.user_id = u.id\n\tINNER JOIN bl_depart d ON u.depart_id = d.id\n\tINNER JOIN bl_depart d1 ON d.parent_id = d1.id \nWHERE\n\tstart_at > '%v 00:00:00' \n\tAND start_at < '%v 23:59:59' \n\tAND call_type = 'extension_outbound'\t\nGROUP BY\n\td1.NAME,\n\td.NAME,\n\tu.nickname,\n\tu.login_name \nORDER BY\n\td1.NAME,\n\td.NAME,\n\tu.nickname,\n\tu.login_name", time_strat, time_stop))
		w.Engine.Engine.Clone()
		fmt.Println(len(nomao))
		f := excelize.NewFile()
		go func() {

			for k, v := range nomao {

				if k == 0 {
					f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+1), "部门")
					f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+1), "二级部门")
					f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), "员工姓名")
					f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+1), "员工工号")
					f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+1), "未接通话总数")
					f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+1), "未接通时长")
					f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+1), "无效通话总数")
					f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+1), "无效通话时长")
					f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+1), "有效通话总数")
					f.SetCellValue("Sheet1", fmt.Sprintf("J%v", k+1), "有效通话时长,")
					f.SetCellValue("Sheet1", fmt.Sprintf("K%v", k+1), "优质通话总数")
					f.SetCellValue("Sheet1", fmt.Sprintf("L%v", k+1), "优质通话时长")
					f.SetCellValue("Sheet1", fmt.Sprintf("M%v", k+1), "总通话数")
					f.SetCellValue("Sheet1", fmt.Sprintf("N%v", k+1), "通话总时长")
				}
				f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+2), v["部门"])
				f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+2), v["二级部门"])
				f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+2), v["员工姓名"])
				f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+2), v["员工工号"])
				f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+2), v["未接通话总数"])
				f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+2), v["未接通时长"])
				f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+2), v["无效通话总数"])
				f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+2), v["无效通话时长"])
				f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+2), v["有效通话总数"])
				f.SetCellValue("Sheet1", fmt.Sprintf("J%v", k+2), v["有效通话时长"])
				f.SetCellValue("Sheet1", fmt.Sprintf("K%v", k+2), v["优质通话总数"])
				f.SetCellValue("Sheet1", fmt.Sprintf("L%v", k+2), v["优质通话时长"])
				f.SetCellValue("Sheet1", fmt.Sprintf("M%v", k+2), v["总通话数"])
				f.SetCellValue("Sheet1", fmt.Sprintf("N%v", k+2), v["通话总时长"])

				vcl.ThreadSync(func() {
					w.progress_bar.SetPosition(int32(float64(k) / float64(len(nomao)) * 100))
				})

			}
			f.SaveAs("电话量统计报表.xlsx")
			time.Sleep(time.Second)
			vcl.ThreadSync(func() {
				vcl.ShowMessageFmt("下载完成！")
			})

			//w.button.SetEnabled(false)
			//w.win.SetEnabled(true)
			w.date_stop_label.SetEnabled(true)
			w.date_strat_label.SetEnabled(true)
			w.button.SetEnabled(true)
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
