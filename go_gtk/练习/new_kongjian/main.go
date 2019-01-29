package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args) //初始化环境

	//--------------------------------------------------------
	// 主窗口
	//--------------------------------------------------------
	windows := gtk.NewWindow(gtk.WINDOW_TOPLEVEL) //创建窗口
	windows.SetPosition(gtk.WIN_POS_CENTER)       //设置窗口居中
	windows.SetTitle("Go GTK")                    //设置窗口标题
	windows.SetSizeRequest(300, 200)              //设置窗口的宽度与高度

	//--------------------------------------------------------
	//GtkFixed
	//--------------------------------------------------------
	layout := gtk.NewFixed() //创建固定布局

	//--------------------------------------------------------
	// GtkButton
	//--------------------------------------------------------
	b1 := gtk.NewButton() //新建按钮
	//b1.SetTitle("Hello")      //设置按钮标签内容
	b1.SetLabel("111")
	b1.SetSizeRequest(100, 50) //设置按钮大小

	b2 := gtk.NewButtonWithLabel("World") //新建按钮并设置内容
	b2.SetSizeRequest(100, 50)            //设置按钮大小

	//--------------------------------------------------------
	// 添加布局、添加容器
	//--------------------------------------------------------
	windows.Add(layout) //把布局添加到主界面中

	layout.Put(b1, 0, 0)    //设置按钮在容器的位置
	layout.Move(b1, 50, 50) //移动按钮的位置，必须先put，再用move

	layout.Put(b2, 50, 100) //设置按钮在容器的位置

	windows.ShowAll() //显示所有的控件

	gtk.Main() //主事件循环，等待用户操作
}
