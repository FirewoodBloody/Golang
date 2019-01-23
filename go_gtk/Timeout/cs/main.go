package main

import (
	"os"
	"strconv"

	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()       //新建builder
	builder.AddFromFile("test.glade") //读取glade文件

	// 获取相应控件
	window := gtk.WindowFromObject(builder.GetObject("window1"))
	label := gtk.LabelFromObject(builder.GetObject("label1"))
	buttonStart := gtk.ButtonFromObject(builder.GetObject("button1"))
	buttonStop := gtk.ButtonFromObject(builder.GetObject("button2"))

	label.ModifyFontSize(50)       //设置label字体大小
	buttonStop.SetSensitive(false) //停止按钮不能按

	var id int      //定时器id
	var num int = 1 //累加标记

	//信号处理
	//启动按钮
	buttonStart.Connect("clicked", func() {
		//启动定时器, 500毫秒为时间间隔，回调函数为匿名函数
		id = glib.TimeoutAdd(500, func() bool {
			num++
			label.SetText(strconv.Itoa(num)) //给标签设置内容
			return true                      //只要定时器没有停止，时间到自动调用回调函数
		})

		buttonStart.SetSensitive(false) //启动按钮变灰，不能按
		buttonStop.SetSensitive(true)   //定时器启动后，停止按钮可以按
	})

	//停止按钮
	buttonStop.Connect("clicked", func() {
		//停止定时器
		glib.TimeoutRemove(id)

		buttonStart.SetSensitive(true)
		buttonStop.SetSensitive(false)
	})

	window.Connect("destroy", gtk.MainQuit) //关闭窗口

	window.ShowAll()

	gtk.Main()
}
