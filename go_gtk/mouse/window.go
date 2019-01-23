package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"unsafe"
)

func main() {
	gtk.Init(&os.Args)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetDefaultSize(500, 300)
	window.Connect("destroy", gtk.MainQuit)

	var x, y int
	//鼠标按下事件处理
	window.Connect("button-press-event", func(ctx *glib.CallbackContext) {
		//获取鼠键按下属性结构体变量，系统内部的变量，不是用户传参变量
		arg := ctx.Args(0)
		event := *(**gdk.EventButton)(unsafe.Pointer(&arg))

		if event.Button == 1 { //左键
			x, y = int(event.X), int(event.Y)
		} else if event.Button == 3 { //右键
			gtk.MainQuit()
		}
	})

	//鼠标移动事件处理
	window.Connect("motion-notify-event", func(ctx *glib.CallbackContext) {
		arg := ctx.Args(0)
		event := *(**gdk.EventButton)(unsafe.Pointer(&arg))

		//移动窗口
		window.Move(int(event.XRoot)-x, int(event.YRoot)-y)
		fmt.Printf("%v %v \n %v %v\n", event.XRoot, event.YRoot, x, y)

		// event.XRoot  鼠标在当前屏幕显示的位置坐标
		// event.X      鼠标在当前窗口显示的位置坐标
	})
	//添加鼠标按下事件
	window.SetEvents(int(gdk.BUTTON_PRESS_MASK | gdk.BUTTON1_MOTION_MASK))

	window.ShowAll()

	gtk.Main()

}
