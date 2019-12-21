package controllers

import (
	"github.com/astaxie/beego"

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
	//id := u.Ctx.Input.RequestBody
	uid, uid1 := models.GetClientMessage(string(id))
	if uid != nil {
		if uid1 == nil {
			u.Data["json"] = uid
			u.ServeJSON()
		}

	} else if uid1 != nil {
		u.Data["json"] = uid1
		u.ServeJSON()
	}
}
