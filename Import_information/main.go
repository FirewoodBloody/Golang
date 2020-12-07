package main

import (
	"crypto/md5"
	"fmt"
	"github.com/ying32/govcl/vcl"
)

type TForm struct {
	LoginTForm
	ImportTForm
}

type ImportTForm struct {
	WinImport *vcl.TForm
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

func Md5Encryption(data string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(data)))
}

func (t *TForm) Init() {
	t.LoginClient()
}

//登陆窗口控件初始化
func (l *LoginTForm) LoginClient() {
	vcl.Application.Initialize()
	l.WinLogin = vcl.Application.CreateForm()
	l.WinLogin.SetAlphaBlend(true)
	l.WinLogin.SetCaption("数据报表查询")
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
	l.TEdit1.SetTop(90)
	l.TEdit1.SetLeft(150)

	l.TEdit2 = vcl.NewEdit(l.WinLogin)
	l.TEdit2.SetParent(l.WinLogin)
	l.TEdit2.SetTop(120)
	l.TEdit2.SetLeft(170)
	l.TEdit2.SetPasswordChar(42)

	l.TLabel1 = vcl.NewLabel(l.WinLogin)
	l.TLabel1.SetParent(l.WinLogin)
	l.TLabel1.SetCaption("工号")
	l.TLabel1.SetLeft(100)
	l.TLabel1.SetTop(90)

	l.TLabel2 = vcl.NewLabel(l.WinLogin)
	l.TLabel2.SetParent(l.WinLogin)
	l.TLabel2.SetCaption("密码")
	l.TLabel2.SetLeft(90)
	l.TLabel2.SetTop(125)

}

func main() {
	TForm := new(TForm)
	TForm.Init()
	TForm.WinLogin.Show()

	vcl.Application.Run()
}
