package routers

import (
	"github.com/FirewoodBloody/Golang/data_analyst/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
