package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args) //初始化环境

	Builder := gtk.NewBuilder()          //新建builder
	Builder.AddFromFile("windows.glade") //读取glade文件

	// 获取窗口控件指针，注意"window1"要和glade里的标志名称匹配
	window := gtk.WindowFromObject(Builder.GetObject("window1"))

	window.SetSizeRequest(300, 200)        //设置窗口宽度和高度
	window.SetIconFromFile("windows.png")  //设置icon图标
	window.SetTitle("hello world")         //设置窗口标题
	window.SetPosition(gtk.WIN_POS_CENTER) //设置居中显示
	window.SetResizable(false)             //设置不可伸缩

	window.Connect("destroy", gtk.MainQuit) //按窗口关闭按钮，自动触发"destroy"信号

	window.ShowAll() //显示所有控件和窗口
	gtk.Main()       //主事件循环
}
