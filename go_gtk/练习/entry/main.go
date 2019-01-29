package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args) //初始化环境

	builder := gtk.NewBuilder()        //新建builder
	builder.AddFromFile("entry.glade") //读取glade文件

	//获取窗口空间指针
	window := gtk.WindowFromObject(builder.GetObject("window1"))
	window.Connect("destroy", gtk.MainQuit)

	//获取entry控件
	entry := gtk.EntryFromObject(builder.GetObject("entry1"))
	entry.SetText("123456")
	//entry.SetVisibility(false)                  //设置不可见字符，即密码模式
	//entry.SetEditable(false)                    //只读，不可编辑
	entry.ModifyFontSize(30) //修改字体大小
	//信号处理，当用户在文本输入控件内部按回车键时引发activate信号
	entry.Connect("activate", func() {
		fmt.Println(entry.GetText())
	})

	window.ShowAll()

	gtk.Main()
}
