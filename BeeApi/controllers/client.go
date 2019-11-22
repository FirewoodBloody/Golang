package controllers

import (
	"github.com/astaxie/beego"
)
import (
	"Golang/BeeApi/models"
)

type ClientController struct {
	beego.Controller
}

// @Title Accles
// @Description Logs out current logged in user session

// @Success 200 {string} accles success
// @router /Accles [post]
func (u *ClientController) Accles() {
	id := u.GetString("Customer_number")
	//s := u.Ctx.Input.RequestBody
	uid := models.GetClientMessage(id)

	u.Data["json"] = uid
	u.ServeJSON()
}
