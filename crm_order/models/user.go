package models

import (
	"encoding/json"

	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	Win           []Win  `json:"win"`
	Err           error  `json:"err"`
	Client_Models string `json:"client_models"`
}

type Win struct {
	Rd string `json:"rd"`
}

type Statu struct {
	Err    string `json:"err"`
	Status bool   `json:"status"`
}

/////////////////---------------------------------------------------

func GetClientMessage(client_Models, login_name string) *Users {
	rd := new(Users)
	if client_Models != "Client_Models" {
		rd.Err = fmt.Errorf("Client_Models ：参数错误！")
		return rd
	}
	e := new(Engine)
	e.NewEngine()
	defer e.Close()
	depart_id, err := e.Login_depart_id(login_name)

	if err != nil {
		rd.Err = fmt.Errorf("查询出错，请联系管理员：%v", err)
		return rd
	}

	position_level, err := e.Login_position_level(login_name)

	if err != nil {
		rd.Err = fmt.Errorf("查询出错，请联系管理员：%v", err)
		return rd
	}

	//员工
	if position_level <= "1" {
		str := beego.AppConfig.String("Staff_TsaveDialog")
		rd.Err = json.Unmarshal([]byte(str), &rd)
		rd.Client_Models = "Staff_TsaveDialog"

		//管理
	} else if position_level >= "2" {
		//客服  高层管理
		if depart_id == "26" || depart_id == "1" || depart_id == "2" || depart_id == "8" || login_name == "18080001" {
			str := beego.AppConfig.String("TsaveDialog")
			rd.Err = json.Unmarshal([]byte(str), &rd)
			rd.Client_Models = "TsaveDialog"

			//新媒体
		} else if depart_id == "11" || depart_id == "33" {
			str := beego.AppConfig.String("Medium_TsaveDialog")
			rd.Err = json.Unmarshal([]byte(str), &rd)
			rd.Client_Models = "Medium_TsaveDialog"

			//财务部
		} else if depart_id == "10" || depart_id == "30" || depart_id == "31" || depart_id == "32" {
			str := beego.AppConfig.String("Warehouse_TsaveDialog")
			rd.Err = json.Unmarshal([]byte(str), &rd)
			rd.Client_Models = "Warehouse_TsaveDialog"
		} else {
			str := beego.AppConfig.String("Staff_TsaveDialog")
			rd.Err = json.Unmarshal([]byte(str), &rd)
			rd.Client_Models = "Staff_TsaveDialog"
		}
	}

	if len(rd.Win) == 0 {
		rd.Err = fmt.Errorf("Client_Models ：参数错误！")
	}
	return rd
}

func Login(login_name, pass string) *Statu {
	stat := new(Statu)
	e := new(Engine)
	err := e.NewEngine()
	defer e.Close()
	if err != nil {
		stat.Err = fmt.Sprintf("%v", err)
		return stat
	}

	resp, err := e.Engine.Query(fmt.Sprintf("SELECT\n\t`password` \nFROM\n\tbl_users \nWHERE login_name = '%v'\nAND `status` = 1;", login_name))
	if err != nil {
		stat.Err = fmt.Sprintf("%v:请联系管理员！", err)
		return stat
	}
	if len(resp) == 0 {
		stat.Err = "账号或密码不存在!"
		return stat
	}

	err = bcrypt.CompareHashAndPassword(resp[0]["password"], []byte(pass))

	if err != nil {
		stat.Err = "账号或密码错误！"
		return stat
	}

	stat.Status = true

	return stat
}
