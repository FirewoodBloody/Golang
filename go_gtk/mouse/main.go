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
	window.SetSizeRequest(500, 320)

	//鼠标按下事件处理
	window.Connect("button-press-event", func(cty *glib.CallbackContext) {
		//获取鼠键按下属性结构体变量，系统内部的变量，不是用户传参变量.
		arg := cty.Args(0)
		event := *(**gdk.EventButton)(unsafe.Pointer(&arg))

		if event.Type == int(gdk.BUTTON_PRESS) {
			fmt.Println("单击")
		} else if event.Type == int(gdk.BUTTON2_PRESS) {
			fmt.Println("双击")
		}

		if event.Button == 1 {
			fmt.Println("左键")
		} else if event.Button == 2 {
			fmt.Println("中键")
		} else if event.Button == 3 {
			fmt.Println("右键")
		}

		fmt.Println("坐标", int(event.X), event.Y)
	})

	//鼠标移动事件处理
	window.Connect("motion-notify-event", func(cty *glib.CallbackContext) {
		//获取鼠标移动属性结构体变量，系统内部的变量，不是用户传参变量
		arg := cty.Args(0)
		event := *(**gdk.EventButton)(unsafe.Pointer(&arg))

		fmt.Println("坐标:", event.X, event.Y)
	})

	//添加鼠标按下事件
	//BUTTON_PRESS_MASK: 鼠标按下，触发信号"button-press-event"
	//BUTTON_RELEASE_MASK：鼠标抬起，触发"button-release-event"
	//鼠标移动都是触发"motion-notify-event"
	//BUTTON_MOTION_MASK: 鼠标移动，按下任何键移动都可以
	//BUTTON1_MOTION_MASK：鼠标移动，按住左键移动才触发
	//BUTTON2_MOTION_MASK：鼠标移动，按住中间键移动才触发
	//BUTTON3_MOTION_MASK：鼠标移动，按住右键移动才触发
	window.SetEvents(int(gdk.BUTTON_PRESS_MASK | gdk.BUTTON1_MOTION_MASK))
	window.Connect("destroy", gtk.MainQuit)
	window.ShowAll() //显示控件

	gtk.Main()

}
