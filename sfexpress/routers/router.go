package routers

import (
	"Golang/sfexpress/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
