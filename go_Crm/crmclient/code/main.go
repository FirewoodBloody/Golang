// 由res2go IDE插件自动生成。
package main

import (
    "github.com/ying32/govcl/vcl"
    "github.com/FirewoodBloody/Golang/go_Crm/crmclient/code/pkgui"
)

func main() {
    vcl.Application.SetScaled(true)
    vcl.Application.SetTitle("project1")
    vcl.Application.Initialize()
    vcl.Application.SetMainFormOnTaskBar(true)
    vcl.Application.CreateForm(&pkgui.Form1)
    vcl.Application.Run()
}
