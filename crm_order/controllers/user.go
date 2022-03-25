package controllers

import (
	"github.com/FirewoodBloody/Golang/crm_order/models"

	"github.com/astaxie/beego"
)

// UserController Operations about Users
type UserController struct {
	beego.Controller
}

// Post @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	//var user models.User
	//json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//uid := models.AddUser(user)
	//fmt.Println(u.GetString("orderid"))

	login_name := u.GetString("Login_Name")
	pass := u.GetString("Password")
	str2 := u.GetString("Version")

	if str2 == "Version" {
		type v struct {
			V string
		}
		a := v{V: beego.AppConfig.String(str2)}

		u.Data["json"] = a
	} else if login_name != "" && pass != "" {
		u.Data["json"] = models.Login(login_name, pass)
	} else if login_name == "" {
		u.Data["json"] = models.Statu{Err: "账号不能为空！"}
	} else if pass == "" {
		u.Data["json"] = models.Statu{Err: "密码不能为空！"}
	} else {
		u.Data["json"] = ""
	}

	u.ServeJSON()
}

// Version @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /Version [post]
func (u *UserController) Version() {
	//var user models.User
	//json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//uid := models.AddUser(user)
	//fmt.Println(u.GetString("orderid"))
	version := u.GetString("Version")

	if version == "Version" {
		type v struct {
			V      string
			Remark string
		}
		a := v{V: beego.AppConfig.String(version), Remark: beego.AppConfig.String("Remark")}

		u.Data["json"] = a
	} else {
		u.Data["json"] = ""
	}

	u.ServeJSON()
}

// Client_Models @Title Client_ModelsUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /Client_Models [post]
func (u *UserController) Client_Models() {
	//var user models.User
	//json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//uid := models.AddUser(user)
	//fmt.Println(u.GetString("orderid"))
	client_Models := u.GetString("Client_Models")

	login_name := u.GetString("Login_Name")

	u.Data["json"] = models.GetClientMessage(client_Models, login_name)

	u.ServeJSON()
}

// GetAll @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	//users := models.GetAllUsers()
	//u.Data["json"] = users
	//u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	//uid := u.GetString(":uid")
	//if uid != "" {
	//	user, err := models.GetUser(uid)
	//	if err != nil {
	//		u.Data["json"] = err.Error()
	//	} else {
	//		u.Data["json"] = user
	//	}
	//}
	//u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	//uid := u.GetString(":uid")
	//if uid != "" {
	//	var user models.User
	//	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//	uu, err := models.UpdateUser(uid, &user)
	//	if err != nil {
	//		u.Data["json"] = err.Error()
	//	} else {
	//		u.Data["json"] = uu
	//	}
	//}
	//u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	//uid := u.GetString(":uid")
	//models.DeleteUser(uid)
	//u.Data["json"] = "delete success!"
	//u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	//username := u.GetString("username")
	//password := u.GetString("password")
	//if models.Login(username, password) {
	//	u.Data["json"] = "login success"
	//} else {
	//	u.Data["json"] = "user not exist"
	//}
	//u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	//u.Data["json"] = "logout success"
	//u.ServeJSON()
}
