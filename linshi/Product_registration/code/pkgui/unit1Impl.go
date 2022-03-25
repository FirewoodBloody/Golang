package pkgui

import (
	"fmt"
	"github.com/FirewoodBloody/Golang/linshi/Product_registration/code/pkgui/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

var (
	Sx            bool
	orderOn       string
	expressNumber string
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnEdit1KeyDown(sender vcl.IObject, key *uint16, shift types.TShiftState) {

	// 键盘输入符 回车判定
	if *key != 13 {
		return
	}

	//出库
	if f.RadioButton1.Checked() {
		//进行文本编辑栏 提示文字修改和 TEXT重置
		if Sx {
			expressNumber = f.Edit1.Text()
			if expressNumber == "" {
				return
			}
			f.Edit1.SetText("")
			f.Edit1.SetTextHint("请输入订单号")
			Sx = false

		} else {
			// 获取用户输入的订单号
			orderOn = f.Edit1.Text()
			if orderOn == "" {
				return
			}

			// 判断用户输入的是否正确
			e := new(sql.Engine)
			err, b := e.SelectOrder(orderOn)

			if err != nil {
				f.Edit1.SetText("")
				return
			}

			if !b {
				f.Edit1.SetText("")
				return
			}

			f.Edit1.SetText("")
			f.Edit1.SetTextHint("请输入快递单号")
			Sx = true
			f.Label1.SetCaption(fmt.Sprintf("订单号：%v", orderOn))
			return
		}
		if !Sx {
			//处理订单信息写入
			e := new(sql.Engine)
			e.InsetrData(orderOn, expressNumber)
			cNumber, tNumber := e.SelectNumber()
			f.Label2.SetCaption(fmt.Sprintf("出库：%v", cNumber))
			f.Label3.SetCaption(fmt.Sprintf("退货：%v", tNumber))
			orderOn = ""
			expressNumber = ""
		}
	}

	//退货
	if f.RadioButton2.Checked() {

		// 获取用户输入的订单号
		// 获取用户输入的订单号
		orderOn = f.Edit1.Text()
		if orderOn == "" {
			return
		}

		// 判断用户输入的是否正确
		e := new(sql.Engine)
		err, b := e.SelectOrder(orderOn)

		if err != nil {
			f.Edit1.SetText("")
			return
		}

		if !b {
			f.Edit1.SetText("")
			return
		}
		f.Label1.SetCaption(fmt.Sprintf("订单号：%v", orderOn))
		f.Edit1.SetText("")
		f.Edit1.SetTextHint("请输入订单号")

		//处理订单信息更新
		e.UpdateData(orderOn)
		cNumber, tNumber := e.SelectNumber()
		f.Label2.SetCaption(fmt.Sprintf("出库：%v", cNumber))
		f.Label3.SetCaption(fmt.Sprintf("退货：%v", tNumber))
		orderOn = ""
		expressNumber = ""

	}

}

func (f *TForm1) OnFormShow(sender vcl.IObject) {
	// 初始化控件数据
	e := new(sql.Engine)
	cNumber, tNumber := e.SelectNumber()
	f.Label2.SetCaption(fmt.Sprintf("出库：%v", cNumber))
	f.Label3.SetCaption(fmt.Sprintf("退货：%v", tNumber))
}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	e := new(sql.Engine)
	err, b := e.DeleteOrder(orderOn)
	if err != nil {
		return
	}
	if b {
		f.Label4.SetCaption(fmt.Sprintf("订单号：%v  删除成功!", orderOn))
		f.Edit1.SetText("")
		f.Label1.SetCaption("订单号：")
		f.Edit1.SetTextHint("请输入订单号")
		orderOn = ""
		expressNumber = ""
		cNumber, tNumber := e.SelectNumber()
		f.Label2.SetCaption(fmt.Sprintf("出库：%v", cNumber))
		f.Label3.SetCaption(fmt.Sprintf("退货：%v", tNumber))
	} else {
		f.Label4.SetCaption(fmt.Sprintf("订单号：%v  订单未出库或不存在!", orderOn))
		f.Edit1.SetText("")
		f.Label1.SetCaption("订单号：")
		f.Edit1.SetTextHint("请输入订单号")
		Sx = false
		orderOn = ""
		expressNumber = ""
		cNumber, tNumber := e.SelectNumber()
		f.Label2.SetCaption(fmt.Sprintf("出库：%v", cNumber))
		f.Label3.SetCaption(fmt.Sprintf("退货：%v", tNumber))
	}
}

func (f *TForm1) OnMenuItem2Click(sender vcl.IObject) {
	Form2.Show()
}

func (f *TForm1) OnRadioButton1Change(sender vcl.IObject) {
	f.Edit1.SetText("")
	f.Label1.SetCaption("订单号：")
	f.Edit1.SetTextHint("请输入订单号")
	f.Label4.SetCaption("欢迎使用！")
	Sx = false
	orderOn = ""
	expressNumber = ""
}

func (f *TForm1) OnRadioButton2Change(sender vcl.IObject) {
	f.Edit1.SetText("")
	f.Label1.SetCaption("订单号：")
	f.Edit1.SetTextHint("请输入订单号")
	f.Label4.SetCaption("欢迎使用！")
	Sx = false
	orderOn = ""
	expressNumber = ""
}
