package main

import (
	"github.com/FirewoodBloody/Golang/crm_order/controllers"
	_ "github.com/FirewoodBloody/Golang/crm_order/routers"
	"github.com/astaxie/beego"
)

func main() {

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	//go models.CheckOrderData()        //监控订单新增数据
	//go models.RecoveryOfTheCustomer() //新媒体客户回收

	beego.Router("/file/Download", &controllers.FileUpLoadControllers{}, "get:Download_64")
	beego.Router("/file/Download_32", &controllers.FileUpLoadControllers{}, "get:Download_32")
	beego.Run()





}
