package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
)

func main() {
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()
	builder.AddFromFile("image.glade")

	window := gtk.WindowFromObject(builder.GetObject("window1"))
	window.SetSizeRequest(600, 400)
	window.SetPosition(gtk.WIN_POS_CENTER)

	image := gtk.ImageFromObject(builder.GetObject("image1"))

	image.SetFromFile("bx.png")
	//image.SetSizeRequest(300, 200)
	image.SetWindow(window.GetWindow())

	window.Connect("destroy", gtk.MainQuit)

	window.ShowAll()

	gtk.Main()

}
