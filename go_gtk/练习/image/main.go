package main

import (
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args)

	//新建Builder
	builder := gtk.NewBuilder()
	builder.AddFromFile("image.glade")

	var w, h int = 500, 300
	//获取窗口控件指针
	window := gtk.WindowFromObject(builder.GetObject("window1"))
	//window.GetSizeRequest(&w, &h)
	window.SetSizeRequest(w, h)

	//获取image控件
	image := gtk.ImageFromObject(builder.GetObject("image1"))
	//image.GetSizeRequest(&w, &h)

	//创建pixbuf，指定大小（宽度和高度），image有多大就设置多大
	//最后一个参数false代表不保存图片原来的尺寸
	pixbuf, _ := gdkpixbuf.NewPixbufFromFileAtScale("bx.png", w, h, false)

	//image设置pixbuf
	image.SetFromPixbuf(pixbuf)

	//pixbuf使用完毕，需要释放资源
	pixbuf.Unref()

	//按窗口关闭按钮，自动触发"destroy"信号
	window.Connect("destroy", gtk.MainQuit)

	window.Show()

	gtk.Main()
}
