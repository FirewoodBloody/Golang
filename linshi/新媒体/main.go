package main

import (
	"github.com/FirewoodBloody/Golang/linshi/新媒体/module"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ying32/govcl/vcl"
)

func main() {
	TForm := new(module.Windows)
	TForm.Init()
	TForm.Onclick()

	TForm.Win.Show()
	vcl.Application.Run()
}
