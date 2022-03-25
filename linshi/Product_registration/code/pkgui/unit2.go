// 由res2go IDE插件自动生成，不要编辑。
package pkgui

import (
    "github.com/ying32/govcl/vcl"
    _ "embed"
)

type TForm2 struct {
    *vcl.TForm
    Panel1          *vcl.TPanel
    DateTimePicker1 *vcl.TDateTimePicker
    DateTimePicker2 *vcl.TDateTimePicker
    Label1          *vcl.TLabel
    Label2          *vcl.TLabel
    Panel2          *vcl.TPanel
    Button1         *vcl.TButton
    Button2         *vcl.TButton
    Panel3          *vcl.TPanel
    ProgressBar1    *vcl.TProgressBar
    SaveDialog1     *vcl.TSaveDialog

    //::private::
    TForm2Fields
}

var Form2 *TForm2




// vcl.Application.CreateForm(&Form2)

func NewForm2(owner vcl.IComponent) (root *TForm2)  {
    vcl.CreateResForm(owner, &root)
    return
}

//go:embed resources/unit2.gfm
var form2Bytes []byte

// 注册Form资源  
var _ = vcl.RegisterFormResource(Form2, &form2Bytes)
