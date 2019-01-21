package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gtk"
	"os"
	"time"
)

func main() {
	gtk.Init(&os.Args)

	builder := gtk.NewBuilder()
	builder.AddFromFile("progressbar.glade")

	window := gtk.WindowFromObject(builder.GetObject("window1"))
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetSizeRequest(320, 200)

	progressbar := gtk.ProgressBarFromObject(builder.GetObject("progressbar1"))
	window.ShowAll()
	for i := 0.0; i <= 1.0; i += 0.1 {

		progressbar.SetFraction(i)
		time.Sleep(time.Second)
		progressbar.SetText(fmt.Sprint(int(i*100), "%"))
	}

	window.Connect("destroy", gtk.MainQuit)

	gtk.Main()
}
