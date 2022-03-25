package pkgui

import (
	"github.com/ying32/govcl/vcl"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	if f.Panel3.Enabled() {
		f.Panel3.Hide()
		f.Panel3.SetEnabled(false)
	} else {
		f.Panel3.Show()
		f.Panel3.SetEnabled(true)
	}
}

func (f *TForm1) OnButton2Click(sender vcl.IObject) {
	f.TabSheet1.Show()
	f.Panel3.Hide()
	f.Panel3.SetEnabled(false)
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	f.Panel3.SetEnabled(false)
	f.Panel3.Hide()
}

func (f *TForm1) OnButton4Click(sender vcl.IObject) {
	f.TabSheet2.Show()
}
