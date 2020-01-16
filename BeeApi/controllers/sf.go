package controllers

import (
	"Golang/BeeApi/models"
	"encoding/xml"
	"github.com/astaxie/beego"
)

type SfController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (o *SfController) Post() {
	var ob models.Object
	xml.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := models.AddOne(ob)
	o.Data["xml"] = models.SF(objectid)
	o.ServeXML()
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router /Sf [post]
func (u *SfController) Sf() {
	i := u.GetString("content")
	u.Data["xml"] = models.Sf(i)
	u.ServeXML()
}
