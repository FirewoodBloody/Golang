package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args) //环境初始化

	windows := gtk.NewWindow(gtk.WINDOW_TOPLEVEL) //新建空白窗口(有边框)
	windows.SetPosition(gtk.WIN_POS_CENTER)       //设置窗口居中
	windows.SetTitle("GO Gtk")                    //设置窗口标题
	windows.SetSizeRequest(300, 200)              //设置窗口高度和宽度

	windows.Show() //显示窗口

	gtk.Main() //主事件循环，等待用户输入
}
