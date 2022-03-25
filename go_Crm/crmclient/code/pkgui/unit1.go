// 由res2go IDE插件自动生成，不要编辑。
package pkgui

import (
    "github.com/ying32/govcl/vcl"
    _ "embed"
)

type TForm1 struct {
    *vcl.TForm
    Panel1       *vcl.TPanel
    Button1      *vcl.TButton
    Button2      *vcl.TButton
    Panel3       *vcl.TPanel
    Button4      *vcl.TButton
    Panel2       *vcl.TPanel
    PageControl1 *vcl.TPageControl `events:"OnButton2Click"`
    TabSheet1    *vcl.TTabSheet
    Button3      *vcl.TButton
    TabSheet2    *vcl.TTabSheet
    Edit1        *vcl.TEdit

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
