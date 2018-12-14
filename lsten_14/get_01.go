package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args) //初始化环境变量

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL) //创建窗口
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("俄罗斯方块")
	window.SetSizeRequest(300, 200) //设置窗口的宽度和高度

	window.Show() //显示窗口
	gtk.Main()    //主事件循环，等待用户操作
}
