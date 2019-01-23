package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args) //初始化环境

	builder := gtk.NewBuilder()
	builder.AddFromFile("hbox.glade")

	//获取窗口空间指针，注意：window1 要与 glade 中的名称相同
	window := gtk.WindowFromObject(builder.GetObject("window1"))
	window.SetSizeRequest(480, 320)
	window.SetTitle("水平布局")

	//获取 hbox 水平框控件
	hbox := gtk.HBoxFromObject(builder.GetObject("hbox1"))

	//新建按钮 burron
	button := gtk.NewButtonWithLabel("按钮A")

	//将按钮添加至水平框
	hbox.Add(button)

	window.Connect("destroy", gtk.MainQuit)

	window.ShowAll()

	gtk.Main()
}
