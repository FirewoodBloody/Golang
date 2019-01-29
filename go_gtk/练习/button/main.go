package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args) //环境初始化

	builder := gtk.NewBuilder()          //新建builder
	builder.AddFromFile(" button.glade") //读取glade文件

	window := gtk.WindowFromObject(builder.GetObject("window1")) //获取窗口控件之战
	window.SetSizeRequest(600, 400)                              //设置窗口大小
	window.SetPosition(gtk.WIN_POS_CENTER)                       //设置窗口居中
	window.Connect("destroy", gtk.MainQuit)                      //设置按窗口关闭按钮，自动触发"destroy"信号

	button1 := gtk.ButtonFromObject(builder.GetObject("button1")) //获取按钮控件 button
	button2 := gtk.ButtonFromObject(builder.GetObject("button2")) //获取按钮控件 button

	button1.SetLabel("Hello")                     //设置按钮显示内容
	fmt.Println("button1 = ", button1.GetLabel()) //获取按钮内容
	button1.SetSensitive(false)                   //设置按钮变灰色，不能按

	//获取按钮大小
	var w, h int
	button2.GetSizeRequest(&w, &h)

	//创建pixbuf指定大小宽度和高度,false 不保存图片原来的尺寸
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("bx.png", w-10, h-10, false)

	//通过pixbuf新建image
	image := gtk.NewImageFromPixbuf(pixbuf)

	pixbuf.Unref() //释放pinbuf资源

	button2.SetImage(image) //按钮设置image

	//按钮信号处理
	button2.Connect("clicked", func() {
		fmt.Println("按钮2按下!")
	})

	window.ShowAll()

	gtk.Main()
}
