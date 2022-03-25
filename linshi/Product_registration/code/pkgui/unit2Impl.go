package pkgui

import (
	"fmt"
	"github.com/FirewoodBloody/Golang/linshi/Product_registration/code/pkgui/sql"
	"github.com/Luxurioust/excelize"
	"github.com/ying32/govcl/vcl"
	"time"
)

//::private::
type TForm2Fields struct {
}

func (f *TForm2) OnButton1Click(sender vcl.IObject) {
	fmt.Println(f.DateTimePicker1.Date().Format("2006-01-02"))
	fmt.Println(f.DateTimePicker2.Date().Format("2006-01-02"))

	e := new(sql.Engine)
	fmt.Println(e.SelectOrder_x(f.DateTimePicker1.Date().Format("2006-01-02"), f.DateTimePicker2.Date().Format("2006-01-02")))
	return
	if f.SaveDialog1.Execute() {
		file := excelize.NewFile()
		//进行数据查询
		err := file.SaveAs(f.SaveDialog1.FileName())
		if err != nil {
			return
		}
	}
}

func (f *TForm2) OnButton2Click(sender vcl.IObject) {
	f.Close()
}

func (f *TForm2) OnFormShow(sender vcl.IObject) {
	f.DateTimePicker1.SetDate(time.Now())
	f.DateTimePicker2.SetDate(time.Now())
}
