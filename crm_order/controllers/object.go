package controllers

import (
	"github.com/FirewoodBloody/Golang/crm_order/models"

	"github.com/astaxie/beego"
)

// ObjectController Operations about object
type ObjectController struct {
	beego.Controller
}

// Post @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
//查询数据需求
func (o *ObjectController) Post() {
	//var ob models.Object
	//json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	//objectid := models.AddOne(ob)
	//o.Data["json"] = map[string]string{"ObjectId": objectid}
	id := o.GetString("Mysql_Select")
	start_time := o.GetString("Start_Time")
	stop_time := o.GetString("Stop_Time")
	login_name := o.GetString("Login_Name")
	cient_Models := o.GetString("Client_Models")

	resulist, err := models.Select(id, start_time, stop_time, login_name, cient_Models)
	if len(resulist) == 0 {
		o.Data["json"] = err
	} else {
		o.Data["json"] = resulist
	}

	o.ServeJSON()
}

// NewSelect @Title NewSelect
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router /NewSelect [post]
func (o *ObjectController) NewSelect() {
	//id := o.GetString("Mysql_Select")
	start_time := o.GetString("Start_Time")
	stop_time := o.GetString("Stop_Time")
	login_name := o.GetString("Login_Name")
	//cient_Models := o.GetString("Client_Models")

	resulist, err := models.Select("新媒体线上明细", start_time, stop_time, login_name, "")
	if len(resulist) == 0 {
		o.Data["json"] = err
	} else {
		o.Data["json"] = resulist
	}

	o.ServeJSON()

}

// Get @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *ObjectController) Get() {
	//objectId := o.Ctx.Input.Param(":objectId")
	//if objectId != "" {
	//	ob, err := models.GetOne(objectId)
	//	if err != nil {
	//		o.Data["json"] = err.Error()
	//	} else {
	//		o.Data["json"] = ob
	//	}
	//}
	//o.ServeJSON()
}

// GetAll @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *ObjectController) GetAll() {
	//obs := models.GetAll()
	//o.Data["json"] = obs
	//o.ServeJSON()
}

// Put @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (o *ObjectController) Put() {
	//objectId := o.Ctx.Input.Param(":objectId")
	//var ob models.Object
	//json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	//
	//err := models.Update(objectId, ob.Score)
	//if err != nil {
	//	o.Data["json"] = err.Error()
	//} else {
	//	o.Data["json"] = "update success!"
	//}
	//o.ServeJSON()
}

// Delete @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (o *ObjectController) Delete() {
	//objectId := o.Ctx.Input.Param(":objectId")
	//models.Delete(objectId)
	//o.Data["json"] = "delete success!"
	//o.ServeJSON()
}
