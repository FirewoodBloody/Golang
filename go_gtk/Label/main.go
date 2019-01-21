package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()        //新建builder
	builder.AddFromFile("label.glade") //读取glade文件

	window := gtk.WindowFromObject(builder.GetObject("window1")) // 获取窗口控件指针，注意"window1"要和glade里的标志名称匹配
	window.SetTitle("Go GTK")                                    //设置窗口标签

	labelone := gtk.LabelFromObject(builder.GetObject("label1")) //获取label控件   清楚的区分 label 链接按钮和标签
	labeltwo := gtk.LabelFromObject(builder.GetObject("label2")) //获取label控件

	fmt.Println("labelone=", labelone.GetText()) //获取label内容
	labelone.SetText("你大爷")                      //设置内容
	labeltwo.SetText("you")                      //设置内容

	window.Connect("destroy", gtk.MainQuit) //点击窗口关闭按钮，出发 destroy 信号

	window.ShowAll() //显示所有窗口和控件

	gtk.Main() //主事件循环，等待用户输入

}
