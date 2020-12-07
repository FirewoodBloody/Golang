package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type FileUpLoadControllers struct {
	beego.Controller
}

func (this *FileUpLoadControllers) Download_64() {
	this.Ctx.Output.Download(fmt.Sprintf("./update/%v_64.exe", beego.AppConfig.String("filename")))
}

func (this *FileUpLoadControllers) Download_32() {
	this.Ctx.Output.Download(fmt.Sprintf("./update/%v_32.exe", beego.AppConfig.String("filename")))
}
