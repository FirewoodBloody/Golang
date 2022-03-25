package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	urlAppAccess_token = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?"
	wxId               = "ww368e805e630821c5"
	blcrmSecret        = "zKSe7jzng4gLwZR7rh63L_sNdzvWrDSa6Mt9RWvQWSA"
	urlMessage         = "https://qyapi.weixin.qq.com/cgi-bin/message/send?"
	TimeFormat         = "2006-01-02 15:04:05"
)

// AccessToken 存储AccessToken
type AccessToken struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// UserInformation 员工信息存储
type UserInformation struct {
	Errcode  int    `json:"errcode"`
	Errmsg   string `json:"errmsg"`
	Userlist []struct {
		Userid     string `json:"userid"`
		Name       string `json:"name"`
		Department []int  `json:"department"`
	} `json:"userlist"`
}

type Message struct {
	Touser  string `json:"touser"`
	Toparty string `json:"toparty"`
	Totag   string `json:"totag"`
	Msgtype string `json:"msgtype"`
	Agentid int    `json:"agentid"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
	Safe                   int `json:"safe"`
	EnableIdTrans          int `json:"enable_id_trans"`
	EnableDuplicateCheck   int `json:"enable_duplicate_check"`
	DuplicateCheckInterval int `json:"duplicate_check_interval"`
}

// GetAccessToken 获取APP   access_token
func GetAccessToken() (string, error) {
	//请求HTTPS 获取APP   access_token
	response, err := http.Get(urlAppAccess_token + "corpid=" + wxId + "&corpsecret=" + blcrmSecret)
	if err != nil {
		return "", err
	}

	a := new(AccessToken)

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	//json 格式化
	_ = json.Unmarshal(body, a)
	return a.AccessToken, nil
}

//func main() {
//	a := new(Message)
//	a.Touser = "MaXiaoChen|BoLongDianCangZongBuKeFu"
//	a.Msgtype = "text"
//	a.Agentid = 1000010
//	a.Text.Content = "CRM客户管理系统升级提示：\n1. 新增新媒体订单部门权限功能\n2. 优化新媒体批量分配问题\n3. 修复公海查询客户问题\n4. 修复订单商品名称显示问题"
//
//	marshal, err := json.Marshal(a)
//	if err != nil {
//		fmt.Println(err)
//	}
//	ody := bytes.NewBuffer(marshal)
//
//	token, err := GetAccessToken()
//	if err != nil {
//		return
//	}
//
//	//fmt.Println(token)
//
//	post, err := http.Post(urlMessage+"access_token="+token, "", ody)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	defer post.Body.Close()
//	all, err := ioutil.ReadAll(post.Body)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println(string(all))
//}
type T struct {
	Attr      string `json:"attr"`
	Code      string `json:"code"`
	CostPrice string `json:"costPrice"`
	CreatedAt int64  `json:"createdAt"`
	Deduct    int    `json:"deduct"`
	Id        string `json:"id"`
	Inventory int    `json:"inventory"`
	KingdeeId string `json:"kingdeeId"`
	Name      string `json:"name"`
	Price     string `json:"price"`
	SalePrice string `json:"salePrice"`
	SpuId     string `json:"spuId"`
	Status    int    `json:"status"`
	UpdatedAt int64  `json:"updatedAt"`
}

const (
	//urlAppAccess_token = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?" //
	//urlAppAccess_token = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=ww368e805e630821c5&corpsecret=zKSe7jzng4gLwZR7rh63L_sNdzvWrDSa6Mt9RWvQWSA"
	//urlMessage   = "https://qyapi.weixin.qq.com/cgi-bin/message/send?"
	//urlUsers     = "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?"
	//wxId         = "ww368e805e630821c5"
	//blcrmSecret  = "zKSe7jzng4gLwZR7rh63L_sNdzvWrDSa6Mt9RWvQWSA"
	dBconnect    = "root:dcf4aa3f7b982ce4@tcp(192.168.0.11:3306)/bl_crm?charset=utf8"
	NewdBconnect = "bolong:bolong2021!@#@tcp(192.168.0.17:3306)/crm?charset=utf8"
	driverName   = "sql"
)

type BlCrm struct {
	Engine *xorm.Engine
	Err    error
}

var tbMappers core.PrefixMapper

func init() {
	_ = os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码
}

func (e *BlCrm) NewEngine() error {

	e.Engine, e.Err = xorm.NewEngine(driverName, dBconnect)
	if e.Err != nil {
		return e.Err
	}
	//tbMappe := core.NewPrefixMapper(core.SnakeMapper{}, tbMapper)
	e.Engine.ShowSQL(true)
	e.Engine.SetTableMapper(tbMappers)
	return nil
}

// NEWCrmDbCommit 建立mysql连接池
func (e *BlCrm) NEWCrmDbCommit() error {

	e.Engine, e.Err = xorm.NewEngine(driverName, NewdBconnect)
	if e.Err != nil {
		return e.Err
	}
	//tbMappe := core.NewPrefixMapper(core.SnakeMapper{}, tbMapper)
	e.Engine.ShowSQL(false)
	e.Engine.SetTableMapper(tbMappers)
	return nil
}

func main() {
	e := new(BlCrm)
	e.NEWCrmDbCommit()

	for {
		//time.Sleep(time.Second)
		query, err := e.Engine.Query(fmt.Sprintf("show processlist;"))
		//query, err := e.Engine.Query(fmt.Sprintf("SELECT COUNT(1) FROM `mall_order`;"))

		if err != nil {
			return
		}
		if len(query) == 0 {
			continue
		}
		for _, v := range query {
			if string(v["Command"]) == "Sleep" {
				continue
			}
			if string(v["Info"]) == "show processlist" {
				continue
			}

			fmt.Println(string(v["Info"]))

		}
	}
}
