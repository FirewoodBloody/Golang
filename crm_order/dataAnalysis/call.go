package dataAnalysis

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

//通话记录
type CallJl struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []Data `json:"data"`
}

//通话记录
type Data struct {
	Id                int    `json:"id"`                //记录id
	Type              int    `json:"type"`              //通话类型 1 呼出 2呼入 3呼出未接 4呼入未接
	Duration          int    `json:"duration"`          //通话时长(单位秒)
	File              string `json:"file"`              //通话文件
	User_id           string `json:"user_id"`           //员工id
	Contact_name      string `json:"contact_name"`      //通话人名称
	Contact_phone     string `json:"contact_phone"`     //通话人号码
	Start_time        string `json:"start_time"`        //通话开始时间
	End_time          string `json:"end_time"`          //通话结束时间
	Update_time       string `json:"update_time"`       //更新时间 (该时间为查询时间）,获取了相同的id进行更新操作 用于下次请求通话记录
	Data_type         int    `json:"data_type"`         //通话类别 0 系统电话 1.微信语音 2.微信通话
	Riend_id          int    `json:"riend_id"`          //好友id
	Friend_wx_id      string `json:"friend_wx_id"`      //String	好友微信Id
	Friend_alias      string `json:"friend_alias"`      //String	好友微信号
	Friend_chat_title string `json:"friend_chat_title"` //String	好友备注
}

const (
	key           = "si_bolong0911"
	url1          = "http://siai.aihujing.com:9989/phone/list" //通话记录，根据更新时间正序排序
	TimeFormat    = "2006-01-02"
	driverName    = "sql"
	dBconnect     = "root:dcf4aa3f7b982ce4@tcp(192.168.0.11:3306)/bl_crm?charset=utf8"
	callconnect   = "root:123456@tcp(192.168.0.19:3306)/data_analysis_library?charset=utf8"
	start_at      = "2021-01-01 00:00:00"
	start_at_type = "%Y-%m-%d %H:%i:%s"
	end_at        = "2021-03-19 23:59:59"
)

type BlCrm struct {
	Engine *xorm.Engine
	Err    error
}

type Call struct {
	Engine *xorm.Engine
	Err    error
}

var tbMappers core.PrefixMapper

func init() {
	_ = os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码
}

//获取通话记录
func CallPost(time, Url string) (*CallJl, error) {
	//增加header选项
	query := url.Values{}
	query.Add("start_time", time)
	query.Add("limit", "1000")
	query.Add("appid", key)

	response, err := http.PostForm(Url, query)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	CallAll := &CallJl{}
	err = json.Unmarshal(body, CallAll)

	return CallAll, nil
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

func (e *Call) NewEngine_Call() error {

	e.Engine, e.Err = xorm.NewEngine(driverName, callconnect)
	if e.Err != nil {
		return e.Err
	}
	//tbMappe := core.NewPrefixMapper(core.SnakeMapper{}, tbMapper)
	e.Engine.ShowSQL(true)
	e.Engine.SetTableMapper(tbMappers)
	return nil
}

//数据库连接关闭
func (e *BlCrm) Close() {
	_ = e.Engine.Close()
}

//数据库连接关闭
func (e *Call) Close() {
	_ = e.Engine.Close()
}

//虎鲸通话记录同步
func CallList() {
	//定时器    首次启动时间设定
	now := time.Now()                                                                  //获取当前时间，放到now里面，要给next用
	next := time.Date(now.Year(), now.Month(), now.Day(), 23, 0, 0, 0, now.Location()) //获取晚上23点的日期
	t := time.NewTimer(next.Sub(now))                                                  //计算当前时间到凌晨的时间间隔，设置一个定时器
	<-t.C

	time1 := time.Now().Format(TimeFormat) + " 00:00:00"
	//time1 := "2021-03-01 00:00:00"
	b := new(BlCrm)
	c := new(Call)

	for {
		_ = b.NewEngine()
		_ = c.NewEngine_Call()
		//获取通话记录
		//CallData := &CallJl{}
		CallData, err := CallPost(time1, url1)
		if err != nil {
			fmt.Println("2:", err)
		}

		if CallData.Code != 0 || CallData.Message != "" || len(CallData.Data) == 0 {
			fmt.Printf("3:%#v\n", CallData)
			goto breakHere
		}

		for k, v := range CallData.Data {

			if v.Data_type != 0 { //跳过非手机的通话
				continue
			}
			if v.Type == 2 { //跳过呼入通话
				continue
			}
			if v.Type == 4 { //跳过呼入未接通话
				continue
			}
			if k%500 == 0 { //每循环500次进行一次数据链接重置
				b.Close()
				c.Close()
				time.Sleep(time.Second * 2)
				_ = b.NewEngine()
				_ = c.NewEngine_Call()
			}

			//通过员工呼叫的电话 查询客户信息
			customer_resulist, _ := b.Engine.Query(fmt.Sprintf("SELECT\n\tcc.id as id,\n\tcc.customer_code as customer_code,\n\tcc.name as name\nFROM\n\tbl_crm_customer cc\n\tLEFT JOIN bl_crm_customer_phone ccp ON ccp.customer_id = cc.id \nWHERE\n\tccp.deleted_at IS NULL \n\tAND cc.deleted_at IS NULL \n\tAND ccp.phone = '%v'", v.Contact_phone))
			if len(customer_resulist) == 0 { //如果系统不存在次客户信息，则跳过此条数据
				continue
			}

			//通过员工ID查询员工姓名
			user_resulist, _ := c.Engine.Query(fmt.Sprintf("SELECT\n\tdepart_v1_name,\n\tdepart_v2_name,\n\tnickname \nFROM\n\t`users` \nWHERE\n\tlogin_name = '%v'", v.User_id))

			if len(user_resulist) == 0 {
				user_resulist, _ = b.Engine.Query(fmt.Sprintf("SELECT nickname,status FROM bl_users WHERE login_name = '%v'", v.User_id))
				if len(user_resulist) == 0 {
					continue
				}
				_, _ = c.Engine.Query(fmt.Sprintf("INSERT INTO users (nickname,login_name,status) VALUES('%v','%v',%v)", string(user_resulist[0]["nickname"]), v.User_id, string(user_resulist[0]["status"])))
			}

			is_effective_call := 0
			is_standard_call := 0
			if v.Duration != 0 { //判断通话是否有效
				is_effective_call = 1
			}

			if v.Duration >= 180 { //判断通话是否达标
				is_standard_call = 1
			}

			_, _ = c.Engine.Query(fmt.Sprintf("INSERT INTO `quality_voice`( `stat_at`, `end_at`, `customer_id`, `customer_code`, `customer_name`, `depart_v1_name`, `depart_v2_name`,`nickname`, `login_name`, `call_duration`, `type`, `is_effective_call`, `is_standard_call`) VALUES ( '%v', '%v', %v, '%v', '%v', '%v','%v','%v', '%v', %v, %v, %v, %v);", v.Start_time, v.End_time, string(customer_resulist[0]["id"]), string(customer_resulist[0]["customer_code"]), string(customer_resulist[0]["name"]), string(user_resulist[0]["depart_v1_name"]), string(user_resulist[0]["depart_v2_name"]), string(user_resulist[0]["nickname"]), v.User_id, v.Duration, v.Type, is_effective_call, is_standard_call))
		}
	breakHere:
		if len(CallData.Data) < 1000 {
			b.Close()
			c.Close()
			//定时器
			now := time.Now()                                                                     //获取当前时间，放到now里面，要给next用
			next := now.Add(time.Hour * 24)                                                       //通过now偏移24小时
			next = time.Date(next.Year(), next.Month(), next.Day(), 23, 0, 0, 0, next.Location()) //获取下一个晚上23点的日期
			t := time.NewTimer(next.Sub(now))                                                     //计算当前时间到凌晨的时间间隔，设置一个定时器
			<-t.C
		}

		if CallData.Code == 0 && CallData.Message == "" && len(CallData.Data) != 0 {
			time1 = CallData.Data[len(CallData.Data)-1].Update_time
		}

	}

}
