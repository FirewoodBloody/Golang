package main

import (
	"fmt"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"time"
)

//按钮b1信号处理的回调函数
func HandleButton(ctx *glib.CallbackContext) {
	arg := ctx.Data()   //获取用户传递的参数，是空接口类型
	p, ok := arg.(*int) //类型断言
	if ok {             //如果ok为true，说明类型断言正确
		fmt.Println("*p = ", *p) //用户传递传递的参数为&tmp，是一个变量的地址
		*p = 250                 //操作指针所指向的内存
	}

	fmt.Println("按钮A被按下")

	//gtk.MainQuit() //关闭gtk程序
}

func Hello() {
	fmt.Println("Hello,窗口即将关闭")
	time.Sleep(time.Second * 2)
	gtk.MainQuit()
}

func main() {
	gtk.Init(&os.Args) //初始化环境

	//主窗口
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL) //创建窗口
	window.SetPosition(gtk.WIN_POS_CENTER)       //设置窗口居中显示
	window.SetTitle("信号控制")                      //设置窗口标签
	window.SetSizeRequest(300, 200)              //设置窗口宽度和高度

	layout := gtk.NewFixed() //创建固定布局

	//--------------------------------------------------------
	// GtkButton
	//--------------------------------------------------------
	b1 := gtk.NewButton()      //新建按钮
	b1.SetLabel("按钮A")         //设置按钮内容
	b1.SetSizeRequest(100, 50) //设置按钮高度与宽度（大小）

	b2 := gtk.NewButtonWithLabel("按钮B") //新建按钮同时设置内容
	b2.SetSizeRequest(100, 50)          //设置按钮高度与宽度（大小）

	//--------------------------------------------------------
	// 添加布局、添加容器
	//--------------------------------------------------------
	window.Add(layout)      //把布局添加到主窗口中
	layout.Put(b1, 50, 50)  //设置按钮在容器的位置
	layout.Put(b2, 50, 100) //设置按钮在容器的位置

	//--------------------------------------------------------
	// 信号处理
	//--------------------------------------------------------
	tmp := 10
	//按钮按下自动触发"pressed"，自动调用 HandleButton, 同时将 &tmp 传递给(HandleButton)
	//	b1.Connect("pressed", HandleButton, &tmp)

	//回调函数为匿名函数，推荐写法
	//按钮按下自动触发"pressed"，自动调用匿名函数
	b2.Connect("pressed", func() {

		fmt.Println("B被按下")
		fmt.Println("tmp = ", tmp)

	}) //注意：}和)在同一行

	b1.Connect("pressed", Hello)

	window.ShowAll() //显示所有的控件

	gtk.Main() //主事件循环，等待用户操作
}
