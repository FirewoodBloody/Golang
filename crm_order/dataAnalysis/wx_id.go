// Package dataAnalysis 更新同步  系统用户 和企业微信 用户 的wxId
package dataAnalysis

import (
	"bytes"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	urlAppAccess_token = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?" //
	//urlAppAccess_token = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=ww368e805e630821c5&corpsecret=zKSe7jzng4gLwZR7rh63L_sNdzvWrDSa6Mt9RWvQWSA"
	urlMessage   = "https://qyapi.weixin.qq.com/cgi-bin/message/send?"
	urlUsers     = "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?"
	wxId         = "ww368e805e630821c5"
	blcrmSecret  = "zKSe7jzng4gLwZR7rh63L_sNdzvWrDSa6Mt9RWvQWSA"
	NewdBconnect = "bolong:bolong2021!@#@tcp(192.168.0.17:3306)/crm_prod?charset=utf8"
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

// Message 企业微信消息
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

// NEWCrmDbCommit 建立mysql连接池
func (e *BlCrm) NEWCrmDbCommit() error {

	e.Engine, e.Err = xorm.NewEngine(driverName, NewdBconnect)
	if e.Err != nil {
		return e.Err
	}
	//tbMappe := core.NewPrefixMapper(core.SnakeMapper{}, tbMapper)
	e.Engine.ShowSQL(true)
	e.Engine.SetTableMapper(tbMappers)
	return nil
}

// UserWxIdRefresh 时间定时器 开始执行每日凌晨23点进行员工数据的刷新操作
func UserWxIdRefresh() {
	//定时器    首次启动时间设定
	now := time.Now()                                                                  //获取当前时间，放到now里面，要给next用
	next := time.Date(now.Year(), now.Month(), now.Day(), 23, 0, 0, 0, now.Location()) //获取晚上23点的日期
	t := time.NewTimer(next.Sub(now))                                                  //计算当前时间到凌晨的时间间隔，设置一个定时器
	<-t.C
	// 疑问，这里启动时间的时间差要是 23点之后启动的会不会故障  (#^.^#)   应该是直接运行下面的程序

	//初始化一个对象
	e := new(BlCrm)
	for {
		//先获取当前的时间，用于程序结束后的定时器计算
		now := time.Now()

		//设置因为请求链接错误 重连跳转节点
	userWxIdErr:

		users, err := TheRefreshUsers()
		if err != nil {
			time.Sleep(time.Second * 600)
			//请求错误 休眠  进行节点跳转
			goto userWxIdErr
		}

		//建立 Mysql 数据库连接 （？200个左右的用户一个连接是否存在问题）
		_ = e.NEWCrmDbCommit()

		for _, v := range users.Userlist {
			//查询企业微信ID 是否已存在 (不存在则直接进行更新)
			queryWxId, _ := e.Engine.Query(fmt.Sprintf("SELECT id,real_name,status FROM sys_user WHERE qw_id = '%v';", v.Userid))
			//判断WXid是否存在
			if len(queryWxId) == 0 {
				//查询员工是否存在，存在则更新 不存在 推送企业微信消息
				queryName, _ := e.Engine.Query(fmt.Sprintf("SELECT id,status FROM sys_user WHERE real_name = '%v';", v.Name))
				if len(queryName) == 0 {
					//不存在
					WxMessage("WX:" + v.Name)
					continue
				}

				if string(queryName[0]["status"]) == "1" {
					//更新wxid
					_, _ = e.Engine.Query(fmt.Sprintf("UPDATE sys_user SET qw_id = '%v' WHERE id = '%v';", v.Userid, string(queryName[0]["id"])))
				}

			} else {
				//存在 首先判断是否和wx员工姓名一直
				if string(queryWxId[0]["real_name"]) == v.Name {
					// 用户信息 wxid 一致
					// 查看用户是否启用
					if string(queryWxId[0]["status"]) != "1" {
						//未启用 更新用户 wx id 为空
						_, _ = e.Engine.Query(fmt.Sprintf("UPDATE sys_user SET qw_id = null WHERE id = '%v';", string(queryWxId[0]["id"])))
					}
					//启用则跳过
					continue
				} else {
					//用户信息不一致
					//查询 wx姓名 在系统是否存在
					queryName, _ := e.Engine.Query(fmt.Sprintf("SELECT id,status,qw_id FROM sys_user WHERE real_name = '%v';", v.Name))
					if len(queryName) == 0 {
						// wx 员工在系统不存在 （可能出现两遍的名字不一致需要修改 ）发送企微消息值管理员
						//清空  用户对应wxid
						_, _ = e.Engine.Query(fmt.Sprintf("UPDATE sys_user SET qw_id = null WHERE id = '%v';", string(queryWxId[0]["id"])))
						//发送企微消息 管理员
						WxMessage("WX:" + v.Name + "   CRM:" + string(queryWxId[0]["id"]))
					} else {
						// wxid 对应用户与系统用户姓名不一致
						//更新 wxid 对应系统用户的 wxid为null(有人占用了当前的wxid 由于系统wxid 的唯一性，所以先清空当前占用wxid 的员工wxid）
						_, _ = e.Engine.Query(fmt.Sprintf("UPDATE sys_user SET qw_id = null WHERE id = '%v';", string(queryWxId[0]["id"])))
						if string(queryName[0]["status"]) != "1" {
							//未启用 （可能已经离职，但是没有在wx 删除用户 需要清空这个员工ID）更新用户 wx id 为空
							if string(queryName[0]["qw_id"]) != "" {
								_, _ = e.Engine.Query(fmt.Sprintf("UPDATE sys_user SET qw_id = null WHERE id = '%v';", string(queryName[0]["id"])))
							}
							continue
						} else {
							//启用则直接更新当前wx用户名对应的wxid
							_, _ = e.Engine.Query(fmt.Sprintf("UPDATE sys_user SET qw_id = '%v' WHERE id = '%v';", v.Userid, string(queryName[0]["id"])))
						}
					}
				}

			}
		}

		e.Close()

		//获取当前时间，放到now里面，要给next用
		next := now.Add(time.Hour * 24)                                                       //通过now偏移24小时
		next = time.Date(next.Year(), next.Month(), next.Day(), 23, 0, 0, 0, next.Location()) //获取下一个晚上23点的日期
		t := time.NewTimer(next.Sub(now))                                                     //计算当前时间到凌晨的时间间隔，设置一个定时器
		<-t.C
	}

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

// TheRefreshUsers 获取员工信息
func TheRefreshUsers() (*UserInformation, error) {
	users := new(UserInformation)
	accessToken, err := GetAccessToken()
	if err != nil {
		return users, err
	}

	response, err := http.Get(urlUsers + "access_token=" + accessToken + "&department_id=1&fetch_child=1")
	if err != nil {
		return users, err
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	//json 格式化
	_ = json.Unmarshal(body, users)

	return users, err
}

// WxMessage 企业微信消息推送
func WxMessage(message string) {
	//消息推送无视是否推送成功  O(∩_∩)O哈哈~
	a := new(Message)
	a.Touser = "MaXiaoChen"
	a.Msgtype = "text"
	a.Agentid = 1000010
	a.Text.Content = message
	marshal, _ := json.Marshal(a)
	body := bytes.NewBuffer(marshal)
	token, _ := GetAccessToken()
	_, _ = http.Post(urlMessage+"access_token="+token, "", body)
}
