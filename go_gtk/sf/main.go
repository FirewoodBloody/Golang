package main

import (
	"github.com/mattn/go-gtk/gtk"
	"os"
)

type Control struct {
	window            *gtk.Window
	filechooserbutton *gtk.FileChooserButton
}

func main() {

	gtk.Init(&os.Args)
	builder := gtk.NewBuilder()
	builder.AddFromFile("./sf.glade")

	var sf Control
	sf.window = gtk.WindowFromObject(builder.GetObject("window1"))

	sf.window.ShowAll()

	gtk.Main()

}
