package main

import (
	"github.com/ying32/govcl/vcl"
)

func main() {
	vcl.Application.Initialize()
	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.EnabledMaximize(false)
	mainForm.ScreenCenter()
	btn := vcl.NewButton(mainForm)
	btn.SetParent(mainForm)
	btn.SetCaption("Hello")
	btn.SetOnClick(func(sender vcl.IObject) {
		vcl.ShowMessage("Hello!")
	})
	vcl.Application.Run()
}
