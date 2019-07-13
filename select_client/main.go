package main

import (
	"Golang/select_client/models"
	"fmt"
	"github.com/ying32/govcl/vcl"
)

func main() {
	TForm := models.TForm{}
	TForm.InItWin()
	TForm.Win.Show()
	TForm.TRadioGroup.SetOnClick(func(sender vcl.IObject) {
		fmt.Println(TForm.TRadioGroup.Items().Strings(TForm.TRadioGroup.ItemIndex()))
		fmt.Println(TForm.TRadioGroup.StyleElements())
	})

	vcl.Application.Run()
}
