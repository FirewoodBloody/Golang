package main

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"strconv"
)

func main() {
	gtk.Init(&os.Args) //初始化环境

	builder := gtk.NewBuilder()       //新建 builder 建筑者
	builder.AddFromFile("test.glade") //读取 glade 文件

	window := gtk.WindowFromObject(builder.GetObject("window1"))      //获取窗口控件指针
	buttonStart := gtk.ButtonFromObject(builder.GetObject("button1")) //获取star 按钮
	buttonStop := gtk.ButtonFromObject(builder.GetObject("button2"))  //获取 stop 按钮
	label := gtk.LabelFromObject(builder.GetObject("label1"))         //获取 label 控件

	buttonStart.SetLabel("Start")
	buttonStop.SetLabel("Stop")

	label.ModifyFontSize(50) //设置 label 字体大小
	label.SetLabel("0")
	buttonStop.SetSensitive(false) //设置 stop 按钮不能按
	window.SetSizeRequest(500, 200)
	window.SetPosition(gtk.WIN_POS_CENTER)

	var id int
	var num int = 1 //定时器id  累加标记
	buttonStart.Connect("clicked", func() {
		//启动定时器, 500毫秒为时间间隔，回调函数为匿名函数
		id = glib.TimeoutAdd(1, func() bool {
			num++
			label.SetLabel(strconv.Itoa(num)) //给标签设置内容
			return true                       //只要定时器没有停止，时间到自动调用回调函数
		})
		buttonStart.SetSensitive(false) //启动按钮变灰，不能按
		buttonStop.SetSensitive(true)   //定时器启动后，停止按钮可以按

	})
	//停止按钮
	buttonStop.Connect("clicked", func() {
		glib.TimeoutRemove(id) //停止定时器

		buttonStop.SetSensitive(false)
		buttonStart.SetSensitive(true)
	})

	window.Connect("destroy", gtk.MainQuit) //关闭窗口

	window.ShowAll()

	gtk.Main()

}
