package main

import (
	"Golang/Import_Excel/event_processing"
	"Golang/Import_Excel/module"
	"fmt"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/wendal/go-oci8"
	"github.com/ying32/govcl/vcl"
	"os"
)

const (
	dBconnect  = "BLCRM/BLCRM2012@192.168.0.9:1521/BLDB"
	driverName = "oci8"
	tbMapper   = "BLCRM."
)

var (
	tbMappers core.PrefixMapper
	Number    int
)

type Engine struct {
	Engine *xorm.Engine
	Err    error
}

type LoginTForm struct {
	WinLogin *vcl.TForm
	TEdit1   *vcl.TEdit
	TEdit2   *vcl.TEdit
	TButton1 *vcl.TButton
	TButton2 *vcl.TButton
	TLabel1  *vcl.TLabel
	TLabel2  *vcl.TLabel
}

func init() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码
	tbMappers = core.NewPrefixMapper(core.SnakeMapper{}, tbMapper)
}

//新建连接
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

//登陆窗口控件初始化
func (l *LoginTForm) LoginClient() {
	vcl.Application.Initialize()
	l.WinLogin = vcl.Application.CreateForm()
	l.WinLogin.SetAlphaBlend(true)
	l.WinLogin.SetName("登陆")
	l.WinLogin.SetHeight(300)
	l.WinLogin.SetWidth(400)
	l.WinLogin.ScreenCenter()
	l.WinLogin.SetBorderIcons(3)  //设置窗口最大化 最小化 关闭按钮  3代表是 最大化按钮和功能无效
	l.WinLogin.Font().SetSize(11) //整体字体大小
	l.WinLogin.Font().SetStyle(16)
	//l.WinLogin.SetColor(100)
	l.WinLogin.SetFormStyle(3)

	l.TButton1 = vcl.NewButton(l.WinLogin)
	l.TButton1.SetParent(l.WinLogin)
	l.TButton1.SetName("登陆")
	l.TButton1.SetLeft(80)
	l.TButton1.SetTop(200)

	l.TButton2 = vcl.NewButton(l.WinLogin)
	l.TButton2.SetParent(l.WinLogin)
	l.TButton2.SetName("取消")
	l.TButton2.SetLeft(240)
	l.TButton2.SetTop(200)

	l.TEdit1 = vcl.NewEdit(l.WinLogin)
	l.TEdit1.SetParent(l.WinLogin)
	l.TEdit1.SetTop(90)
	l.TEdit1.SetLeft(150)

	l.TLabel1 = vcl.NewLabel(l.WinLogin)
	l.TLabel1.SetParent(l.WinLogin)
	l.TLabel1.SetName("工号")
	l.TLabel1.SetLeft(100)
	l.TLabel1.SetTop(90)

}

func (l *LoginTForm) LoginTFormSet(t *module.TForm) {
	l.TButton1.SetOnClick(func(sender vcl.IObject) {
		if l.TEdit1.Text() == "" {
			vcl.ShowMessageFmt("请输入工号！")
			return
		} else if l.TEdit1.Text() != "1973" {
			t.RadioButton1.Hide()
		}
		e := new(Engine)
		e.NewEngine()
		str := e.Select_Name(l.TEdit1.Text())
		e.Engine.Close()

		if str == "" {
			t.RadioButton1.SetEnabled(true)
			return
		}

		t.TLabel.SetName(str)
		t.Operator = l.TEdit1.Text()
		l.WinLogin.Free()
		t.Windows.Show()

	})

	l.TButton2.SetOnClick(func(sender vcl.IObject) {
		os.Exit(0)
	})

}

func (e *Engine) Select_Name(str string) string {
	rows, _ := e.Engine.Query(fmt.Sprintf("SELECT NAME FROM BLCRM.CRM_SYS02 WHERE NO = '%v'", str))
	for _, v := range rows {
		for _, k := range v {
			if string(k) != "" {
				return string(k)
			}
		}
	}
	return ""
}

func main() {
	vcl.Application.Initialize() //初始化VCL环境

	l := new(LoginTForm)
	l.LoginClient()

	TFrom := new(module.TForm)
	TFrom.Init()
	event_processing.TheEventInit(TFrom)

	l.LoginTFormSet(TFrom)

	TFrom.Windows.Hide()
	l.WinLogin.Show()

	vcl.Application.Run()
}
