package models

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TForm struct {
	Win         *vcl.TForm
	Combobox    *vcl.TComboBox //下拉选项
	Button      *vcl.TButton   //按钮
	Tedit       *vcl.TEdit     //单行文本框
	ListView    *vcl.TListView //高级列表框
	TPanel1     *vcl.TPanel
	TPanel2     *vcl.TPanel
	TRadioGroup *vcl.TRadioGroup
}

func (T *TForm) InItWin() {
	vcl.Application.Initialize() //初始化环境
	T.Win = vcl.Application.CreateForm()
	T.Win.SetName("客户信息查询")
	T.Win.SetFormStyle(0)
	T.Win.SetHeight(668)
	T.Win.SetWidth(1160)
	T.Win.ScreenCenter()      //居中
	T.Win.SetBorderIcons(3)   //设置窗口最大化 最小化 关闭按钮  3代表是 最大化按钮和功能无效
	T.Win.Font().SetSize(11)  //整体字体大小
	T.Win.Font().SetStyle(16) //字体样式

	T.TPanel1 = vcl.NewPanel(T.Win)
	T.TPanel1.SetAlign(types.AlRight)
	//T.TPanel1.SetCaption("左")
	T.TPanel1.SetParent(T.Win)
	T.TPanel1.SetAlignWithMargins(true) //设置边距

	T.TPanel2 = vcl.NewPanel(T.Win)
	T.TPanel2.SetAlign(types.AlClient)
	T.TPanel2.SetCaption("客户")
	T.TPanel2.SetParent(T.Win)
	T.TPanel2.SetAlignWithMargins(true)

	//T.Combobox = vcl.NewComboBox(T.Win)
	//T.Combobox.SetParent(T.TPanel1)
	//T.Combobox.SetLeft((T.TPanel1.Width() - T.Combobox.Width()) / 2)
	//T.Combobox.SetTop(T.TPanel1.Height() / 3)
	//T.Combobox.SetText("客户姓名")
	//T.Combobox.AddItem("客户姓名", nil)
	//T.Combobox.AddItem("客户电话", nil)
	//T.Combobox.AddItem("客户编码", nil)

	T.TRadioGroup = vcl.NewRadioGroup(T.Win)
	T.TRadioGroup.SetParent(T.TPanel1)
	T.TRadioGroup.Items().Add("客户姓名")
	T.TRadioGroup.Items().Add("客户电话")
	T.TRadioGroup.Items().Add("客户编码")
	T.TRadioGroup.SetWidth(T.TPanel1.Width() - 5)
	T.TRadioGroup.SetAlignWithMargins(true)
	T.TRadioGroup.SetStyleElements(8)
	T.TRadioGroup.SetControlStyle(8)
	T.TRadioGroup.SetTop(T.TPanel1.Height() / 4 * 2)

}
