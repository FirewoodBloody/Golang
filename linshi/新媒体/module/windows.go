package module

import (
	"github.com/Luxurioust/excelize"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/types"
	"os"
	"xorm.io/xorm"
)

type Engine struct {
	Engine *xorm.Engine
	Err    error
}

type Windows struct {
	Win *vcl.TForm //窗口

	Progress_bar *vcl.TProgressBar //进度条

	Label_strat *vcl.TLabel //开始标签
	Label_stop  *vcl.TLabel //结束标签

	Date_strat_label *vcl.TDateTimePicker //开始日期菜单
	Date_stop_label  *vcl.TDateTimePicker //结束日期

	Button      *vcl.TButton     //下载
	TsaveDialog *vcl.TSaveDialog //b保存文件

	Taskdlg *vcl.TTaskDialog

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

func (w *Windows) Init() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	w.Win = vcl.Application.CreateForm() //新建窗口
	w.Win.SetCaption("新媒体")              //程序名
	//w.Win.SetFormStyle(2)
	w.Win.SetHeight(300)     //高
	w.Win.SetWidth(400)      //宽
	w.Win.ScreenCenter()     //居于当前屏幕中心
	w.Win.SetBorderIcons(3)  //设置窗口最大化 最小化 关闭按钮  3代表是 最大化按钮和功能无效
	w.Win.Font().SetSize(11) //整体字体大小
	w.Win.Font().SetColor(255)
	w.Win.Font().SetStyle(16) //字体样式
	w.Win.SetColor(16775388)
	//w.Win.SetTransparentColor(true)
	//w.Win.SetTransparentColorValue(1)

	w.Button = vcl.NewButton(w.Win)
	w.Button.SetParent(w.Win)
	w.Button.SetHeight(50)
	w.Button.SetWidth(100)
	w.Button.SetTop(150)
	w.Button.SetLeft(150)
	w.Button.SetCaption("查询下载")

	w.Label_strat = vcl.NewLabel(w.Win)
	w.Label_strat.SetParent(w.Win)
	w.Label_strat.SetCaption("开始时间:")
	w.Label_strat.SetLeft(100) //设置按钮位置  横向
	w.Label_strat.SetTop(50)   //设置按钮位置 竖向

	w.Label_stop = vcl.NewLabel(w.Win)
	w.Label_stop.SetParent(w.Win)
	w.Label_stop.SetCaption("结束时间:")
	w.Label_stop.SetLeft(100)
	w.Label_stop.SetTop(100)

	w.Date_strat_label = vcl.NewDateTimePicker(w.Win)
	w.Date_strat_label.SetParent(w.Win)
	w.Date_strat_label.SetLeft(180)
	w.Date_strat_label.SetTop(50)

	w.Date_stop_label = vcl.NewDateTimePicker(w.Win)
	w.Date_stop_label.SetParent(w.Win)
	w.Date_stop_label.SetLeft(180)
	w.Date_stop_label.SetTop(100)

	w.Progress_bar = vcl.NewProgressBar(w.Win)
	w.Progress_bar.SetParent(w.Win)
	w.Progress_bar.SetPosition(0)
	w.Progress_bar.SetWidth(400)
	w.Progress_bar.SetHeight(20)
	w.Progress_bar.SetLeft(0)
	w.Progress_bar.SetTop(230)

	w.TsaveDialog = vcl.NewSaveDialog(w.Win)
	w.TsaveDialog.SetFilter("office Excel (*.xlsx)|*.xlsx")
	//    dlgOpen.SetInitialDir()
	//	dlgOpen.SetFilterIndex()

	w.TsaveDialog.SetOptions(w.TsaveDialog.Options().Include(types.OfShowHelp, types.OfAllowMultiSelect)) //rtl.Include(, types.OfShowHelp))
	w.TsaveDialog.SetTitle("保存")

	w.Taskdlg = vcl.NewTaskDialog(w.Win)
	//defer w.Taskdlg.Free()
	w.Taskdlg.SetCaption("询问")
	w.Taskdlg.SetTitle("报表选择")
	w.Taskdlg.SetText("请选择查询导入的报表？")
	//w.Taskdlg.SetExpandButtonCaption("展开")
	//w.Taskdlg.SetExpandedText("展开的文本")
	w.Taskdlg.SetFooterText("新媒体")

	rd := vcl.AsTaskDialogRadioButtonItem(w.Taskdlg.RadioButtons().Add())
	rd.SetCaption("销售订单明细")
	rd = vcl.AsTaskDialogRadioButtonItem(w.Taskdlg.RadioButtons().Add())
	rd.SetCaption("新媒体线上明细")
	rd = vcl.AsTaskDialogRadioButtonItem(w.Taskdlg.RadioButtons().Add())
	rd.SetCaption("新媒体线上统计")
	rd = vcl.AsTaskDialogRadioButtonItem(w.Taskdlg.RadioButtons().Add())
	rd.SetCaption("新媒体线下二开统计")

	w.Taskdlg.SetCommonButtons(0) //rtl.Include(0, 0))
	btn := vcl.AsTaskDialogButtonItem(w.Taskdlg.Buttons().Add())
	btn.SetCaption("确定")
	btn.SetModalResult(types.MrYes)

	btn = vcl.AsTaskDialogButtonItem(w.Taskdlg.Buttons().Add())
	btn.SetCaption("取消")
	btn.SetModalResult(types.MrNo)

	if rtl.LcLLoaded() {
		w.Taskdlg.SetMainIcon(types.TdiQuestion)
	} else {
		w.Taskdlg.SetMainIcon(types.TdiInformation)
	}

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

	w.Button.SetOnClick(func(sender vcl.IObject) {
		if w.Taskdlg.Execute() {
			if w.Taskdlg.ModalResult() == types.MrYes {
				//fmt.Println(w.Taskdlg.RadioButton().Caption())
				if w.TsaveDialog.Execute() { //选择打开文件
					//fmt.Println("filename: ", w.TsaveDialog.FileName())
					file := excelize.NewFile()

					w.Button.SetEnabled(false)
					w.Win.SetEnabled(false)

					if w.Taskdlg.RadioButton().Caption() == "销售订单明细" {
						vcl.ThreadSync(func() {
							w.Progress_bar.SetPosition(0)
						})
						go func() { Order_select(w, file) }()
					} else if w.Taskdlg.RadioButton().Caption() == "新媒体线上明细" {
						vcl.ThreadSync(func() {
							w.Progress_bar.SetPosition(0)
						})
						go func() { Order_xs_select(w, file) }()
					} else if w.Taskdlg.RadioButton().Caption() == "新媒体线上统计" {
						vcl.ThreadSync(func() {
							w.Progress_bar.SetPosition(0)
						})
						go func() { Order_xsTJ_select(w, file) }()
					} else if w.Taskdlg.RadioButton().Caption() == "新媒体线下二开统计" {
						vcl.ThreadSync(func() {
							w.Progress_bar.SetPosition(0)
						})
						go func() { Order_xxTJ_select(w, file) }()
					}
				}
			}
			return
		}
	})

}
