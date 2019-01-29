package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"unsafe"
)

//按键时间需要窗口为有边框显示
func main() {
	gtk.Init(&os.Args)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetDefaultSize(500, 300)
	window.SetPosition(gtk.WIN_POS_CENTER)

	window.SetEvents(int(gdk.BUTTON_PRESS_MASK | gdk.BUTTON1_MOTION_MASK))

	window.Connect("button-press-event", func(ctx *glib.CallbackContext) {
		//获取鼠标移动属性结构体变量，系统内部的变量，不是用户传参变量
		arg := ctx.Args(0)
		event := *(**gdk.EventButton)(unsafe.Pointer(&arg))

		if event.Button == 3 {
			gtk.MainQuit()
		}
	})

	window.Connect("key-press-event", func(ctx *glib.CallbackContext) {
		//获取键盘按下时结构体属性，系统内部的变量
		arg := ctx.Args(0)
		event := *(**gdk.EventKey)(unsafe.Pointer(&arg))

		key := event.Keyval //获取按下(释放)键盘键值，每个键值对于一个ASCII码

		if gdk.KEY_Up == key {
			fmt.Println("上")
		} else if gdk.KEY_Down == key {
			fmt.Println("下")
		} else if gdk.KEY_Left == key {
			fmt.Println("左")
		} else if gdk.KEY_Right == key {
			fmt.Println("右")
		}

		fmt.Println("key = ", event.Keyval)
	})

	window.ShowAll()

	gtk.Main()

}
