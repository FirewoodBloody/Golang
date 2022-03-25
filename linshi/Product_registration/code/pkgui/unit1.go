// 由res2go IDE插件自动生成，不要编辑。
package pkgui

import (
    "github.com/ying32/govcl/vcl"
    _ "embed"
)

type TForm1 struct {
    *vcl.TForm
    Panel1       *vcl.TPanel
    Edit1        *vcl.TEdit
    Label4       *vcl.TLabel
    Panel2       *vcl.TPanel
    Button1      *vcl.TButton
    Label1       *vcl.TLabel
    Label2       *vcl.TLabel
    Label3       *vcl.TLabel
    Button2      *vcl.TButton `events:"OnButton1Click"`
    RadioGroup1  *vcl.TRadioGroup
    RadioButton1 *vcl.TRadioButton
    RadioButton2 *vcl.TRadioButton
    MainMenu1    *vcl.TMainMenu
    MenuItem1    *vcl.TMenuItem
    MenuItem2    *vcl.TMenuItem

    //::private::
    TForm1Fields
}

var Form1 *TForm1




// vcl.Application.CreateForm(&Form1)

func NewForm1(owner vcl.IComponent) (root *TForm1)  {
    vcl.CreateResForm(owner, &root)
    return
}

//go:embed resources/unit1.gfm
var form1Bytes []byte

// 注册Form资源  
var _ = vcl.RegisterFormResource(Form1, &form1Bytes)
