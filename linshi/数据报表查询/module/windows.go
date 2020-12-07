package module

import (
	"encoding/json"
	"fmt"
	"github.com/Luxurioust/excelize"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/types"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"xorm.io/xorm"
)

type Engine struct {
	Engine *xorm.Engine
	Err    error
}

type Users struct {
	Win           []Win  `json:"win"`
	Err           string `json:"err"`
	Client_Models string `json:"client_models"`
}

type Win struct {
	Rd string `json:"rd"`
}

type Windows struct {
	Win        *vcl.TForm //窗口
	LoginTForm LoginTForm

	Progress_bar *vcl.TProgressBar //进度条

	Label_strat *vcl.TLabel //开始标签
	Label_stop  *vcl.TLabel //结束标签

	Date_strat_label *vcl.TDateTimePicker //开始日期菜单
	Date_stop_label  *vcl.TDateTimePicker //结束日期

	Button      *vcl.TButton     //下载
	TsaveDialog *vcl.TSaveDialog //b保存文件

	Taskdlg *vcl.TTaskDialog

	Engine        Engine
	Client_Models string
	LoginName     string
}

const (
	TimeFormat = "2006-01-02"
	driverName = "mysql"
	dBconnect  = "root:dcf4aa3f7b982ce4@tcp(192.168.0.11:3306)/bl_crm?charset=utf8"
)

var (
	isLogin   bool
	tbMappers core.PrefixMapper
)

type LoginTForm struct {
	WinLogin *vcl.TForm
	TEdit1   *vcl.TEdit
	TEdit2   *vcl.TEdit
	TButton1 *vcl.TButton
	TButton2 *vcl.TButton
	TLabel1  *vcl.TLabel
	TLabel2  *vcl.TLabel
}

type Statu struct {
	Err    string `json:"err"`
	Status bool   `json:"status"`
}

func init() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码
}

//登陆窗口控件初始化
func (l *LoginTForm) Init() {
	l.WinLogin = vcl.Application.CreateForm() //新建窗口
	l.WinLogin.SetCaption("数据报表查询")           //程序名
	//w.Win.SetFormStyle(2)
	l.WinLogin.SetHeight(300)     //高
	l.WinLogin.SetWidth(400)      //宽
	l.WinLogin.ScreenCenter()     //居于当前屏幕中心
	l.WinLogin.SetBorderIcons(3)  //设置窗口最大化 最小化 关闭按钮  3代表是 最大化按钮和功能无效
	l.WinLogin.Font().SetSize(11) //整体字体大小
	l.WinLogin.Font().SetColor(255)
	l.WinLogin.Font().SetStyle(16) //字体样式
	l.WinLogin.SetColor(16775388)

	l.TButton1 = vcl.NewButton(l.WinLogin)
	l.TButton1.SetParent(l.WinLogin)
	l.TButton1.SetCaption("登陆")
	l.TButton1.SetLeft(80)
	l.TButton1.SetTop(200)

	l.TButton2 = vcl.NewButton(l.WinLogin)
	l.TButton2.SetParent(l.WinLogin)
	l.TButton2.SetCaption("取消")
	l.TButton2.SetLeft(240)
	l.TButton2.SetTop(200)

	l.TEdit1 = vcl.NewEdit(l.WinLogin)
	l.TEdit1.SetParent(l.WinLogin)
	l.TEdit1.SetTop(80)
	l.TEdit1.SetLeft(165)
	l.TEdit1.SetWidth(100)

	l.TEdit2 = vcl.NewEdit(l.WinLogin)
	l.TEdit2.SetParent(l.WinLogin)
	l.TEdit2.SetTop(125)
	l.TEdit2.SetLeft(165)
	l.TEdit2.SetPasswordChar(42)
	l.TEdit2.SetWidth(100)

	l.TLabel1 = vcl.NewLabel(l.WinLogin)
	l.TLabel1.SetParent(l.WinLogin)
	l.TLabel1.SetCaption("工号")
	l.TLabel1.SetLeft(110)
	l.TLabel1.SetTop(83)

	l.TLabel2 = vcl.NewLabel(l.WinLogin)
	l.TLabel2.SetParent(l.WinLogin)
	l.TLabel2.SetCaption("密码")
	l.TLabel2.SetLeft(110)
	l.TLabel2.SetTop(128)

}

//工作窗口控件初始化
func (w *Windows) Init() {
	w.Win = vcl.Application.CreateForm() //新建窗口
	w.Win.SetCaption("数据报表查询")           //程序名
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

//工作窗口事件
func (w *Windows) Onclick() {

	//	查询下载按钮
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
					} else if w.Taskdlg.RadioButton().Caption() == "礼品订单明细" {
						vcl.ThreadSync(func() {
							w.Progress_bar.SetPosition(0)
						})
						go func() { Order_gift_select(w, file) }()
					} else if w.Taskdlg.RadioButton().Caption() == "订单出库明细" {
						vcl.ThreadSync(func() {
							w.Progress_bar.SetPosition(0)
						})
						go func() { Order_warehouse_select(w, file) }()
					} else if w.Taskdlg.RadioButton().Caption() == "通话记录统计报表" {
						vcl.ThreadSync(func() {
							w.Progress_bar.SetPosition(0)
						})
						go func() { Call_log_statistics(w, file) }()
					} else if w.Taskdlg.RadioButton().Caption() == "客户积分报表" {
						vcl.ThreadSync(func() {
							w.Progress_bar.SetPosition(0)
						})
						go func() { Client_coin(w, file) }()
					} else if w.Taskdlg.RadioButton().Caption() == "客户生日报表" {
						vcl.ThreadSync(func() {
							w.Progress_bar.SetPosition(0)
						})
						go func() { Client_birthday(w, file) }()
					} else if w.Taskdlg.RadioButton().Caption() == "客户名单构成报表" {
						vcl.ThreadSync(func() {
							w.Progress_bar.SetPosition(0)
						})
						go func() { Client_prestigious_university(w, file) }()
					}
				}
			}
			return
		}
	})

	w.Win.SetOnCloseQuery(func(sender vcl.IObject, canClose *bool) {
		if *canClose {
			*canClose = vcl.MessageDlg("是否退出?", types.MtConfirmation, types.MbYes, types.MbNo) == types.MrYes
		}

		if *canClose {
			os.Exit(0)
		}
	})

}

//登录窗口事件
func (w *Windows) LoginOnclick() {
	//登录按钮
	w.LoginTForm.TButton1.SetOnClick(func(sender vcl.IObject) {
		w.LoginTForm.WinLogin.SetEnabled(false)
		w.LoginTForm.TButton1.SetEnabled(false)
		w.LoginTForm.TButton2.SetEnabled(false)

		if w.LoginTForm.TEdit1.Text() == "" || w.LoginTForm.TEdit2.Text() == "" {
			vcl.ShowMessage("账号或密码不能为空！")
			w.LoginTForm.WinLogin.SetEnabled(true)
			w.LoginTForm.TButton1.SetEnabled(true)
			w.LoginTForm.TButton2.SetEnabled(true)
			return
		}

		bady := url.Values{}
		bady.Add("Login_Name", w.LoginTForm.TEdit1.Text())
		bady.Add("Password", w.LoginTForm.TEdit2.Text())

		resp, err := http.PostForm(Url, bady)
		if err != nil {
			vcl.ShowMessage(fmt.Sprintf("%v：请联系管理员！", err))
			w.LoginTForm.WinLogin.SetEnabled(true)
			w.LoginTForm.TButton1.SetEnabled(true)
			w.LoginTForm.TButton2.SetEnabled(true)
			return
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)

		Stat := new(Statu)
		err = json.Unmarshal(data, Stat)
		if err != nil {
			vcl.ShowMessage(fmt.Sprintf("%v：请联系管理员！", err))
			w.LoginTForm.WinLogin.SetEnabled(true)
			w.LoginTForm.TButton1.SetEnabled(true)
			w.LoginTForm.TButton2.SetEnabled(true)
			isLogin = false
			return
		}

		if Stat.Status {
			isLogin = true
			w.LoginName = w.LoginTForm.TEdit1.Text()

			//w.LoginTForm.WinLogin.Close()

			w.Taskdlg = vcl.NewTaskDialog(w.Win)
			//defer w.Taskdlg.Free()
			w.Taskdlg.SetCaption("询问")
			w.Taskdlg.SetTitle("报表选择")
			w.Taskdlg.SetText("请选择查询导出的报表？")
			//w.Taskdlg.SetExpandButtonCaption("展开")
			//w.Taskdlg.SetExpandedText("展开的文本")
			//w.Taskdlg.SetFooterText("")//部门

			bady := url.Values{}
			bady.Add("Client_Models", "Client_Models")
			bady.Add("Login_Name", w.LoginTForm.TEdit1.Text())
			//bady.Add("Mysql_Select", "销售订单明细")
			//bady.Add("Start_Time", "2020-09-22")
			//bady.Add("Stop_Time", "2020-09-22")
			resp, err := http.PostForm(Url1, bady)
			if err != nil {
				vcl.ShowMessage(fmt.Sprintf("%v", err))
			}
			defer resp.Body.Close()
			data, err := ioutil.ReadAll(resp.Body)
			//fmt.Println(string(data))

			rds := Users{}

			err = json.Unmarshal(data, &rds)
			if err != nil {
				vcl.ShowMessage(fmt.Sprintf("%v", err))
			}

			w.Client_Models = rds.Client_Models

			for _, k := range rds.Win {
				rd := vcl.AsTaskDialogRadioButtonItem(w.Taskdlg.RadioButtons().Add())
				rd.SetCaption(k.Rd)
			}

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

			w.Win.Show()
			w.LoginTForm.WinLogin.Hide()
			return
		} else {
			vcl.ShowMessage(fmt.Sprintf("%v", Stat.Err))
			w.LoginTForm.WinLogin.SetEnabled(true)
			w.LoginTForm.TButton1.SetEnabled(true)
			w.LoginTForm.TButton2.SetEnabled(true)
			return
		}

	})

	w.LoginTForm.TEdit2.SetOnKeyDown(func(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
		if *key == 13 {
			w.LoginTForm.TButton1.Click()
		}
	})

	w.LoginTForm.TEdit1.SetOnKeyDown(func(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
		if *key == 13 {
			w.LoginTForm.TButton1.Click()
		}
	})

	w.LoginTForm.WinLogin.SetOnKeyDown(func(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
		if *key == 13 {
			w.LoginTForm.TButton1.Click()
		}
	})

	//取消按钮退出
	w.LoginTForm.TButton2.SetOnClick(func(sender vcl.IObject) {
		os.Exit(0)
	})

	w.LoginTForm.WinLogin.SetOnClose(func(sender vcl.IObject, action *types.TCloseAction) {
		os.Exit(0)
	})
}
