//go:generate goversioninfo -icon=resource/icon.ico -manifest=resource/goversioninfo.exe.manifest

package main

import (
	"github.com/FirewoodBloody/Golang/linshi/通话记录统计查询/module"
	"github.com/ying32/govcl/vcl"
)

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.SetShowMainForm(false)
	TForm := new(module.Windows)
	TForm.Init()
	TForm.Onclick()
	TForm.LoginTForm.Init()
	TForm.LoginOnclick()
	TForm.LoginTForm.WinLogin.Show()

	vcl.Application.Run()
}
