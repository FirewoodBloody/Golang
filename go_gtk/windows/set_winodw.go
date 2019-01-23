package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetSizeRequest(500, 300)

	window.Connect("configure_event", func() {
		var w, h int
		window.GetSize(&w, &h)
		fmt.Println(w, h)
	})

	window.Connect("destroy", gtk.MainQuit)

	window.ShowAll()

	gtk.Main()
}
