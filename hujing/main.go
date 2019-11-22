package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/wendal/go-oci8"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type GetAllUser struct {
	Error_code int    `json:"error_code"`
	Message    string `json:"message"`
	Data       Datas
}

type Datas struct {
	CurrentPage int `json:"currentpage"` //当前页数
	TotalCount  int `json:"totalcount"`  //总条数
	TotalPage   int `json:"totalpage"`   //总页数
	Result      []User
}

type User struct {
	CreateTime      string `json:"create_time"`     //销售创建时间（yyyy-MM-dd HH:mm:ss）
	User_id         string `json:"user_id"`         //销售账号
	Username        string `json:"username"`        //销售名称
	Department_id   int    `json:"department_id"`   //部门id
	Department_name string `json:"department_name"` //部门名称
}

//通话记录
type CallJl struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []Data
}

type Data struct {
	Id            int    `json:"id"` //记录id
	User_id       string `json:"user_id"`
	Type          int    `json:"type"`          //通话类型 1 呼出 2呼入 3呼出未接 4呼入未接
	Duration      int    `json:"duration"`      //通话时长(单位秒)
	File          string `json:"file"`          //通话文件
	Contact_name  string `json:"contact_name"`  //通话人名称
	Contact_phone string `json:"contact_phone"` //通话人号码
	Start_time    string `json:"start_time"`    //通话开始时间
	End_time      string `json:"end_time"`      //通话结束时间
	Update_time   string `json:"update_time"`   //更新时间 (该时间为查询时间）,获取了相同的id进行更新操作
}

type Engine struct {
	Engine *xorm.Engine
	Err    error
}

const (
	key  = "si_bolong0911"
	url1 = "http://siai.aihujing.com:9989/phone/list" //通话记录，根据更新时间正序排序
	url2 = "http://siai.aihujing.com:9989/file/down"  //获取文件
	url3 = "http://siai.aihujing.com:9989/api/user/getAllUser"

	TimeFormat = "2006-01-02"
	driverName = "oci8"
	dBconnect  = "BLCRM/BLCRM2012@192.168.0.9:1521/BLDB"
	tbMapper   = "BLCRM."
)

var tbMappers core.PrefixMapper

func init() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码
	tbMappers = core.NewPrefixMapper(core.SnakeMapper{}, tbMapper)
}

//初始化化
func (e *Engine) NewEngine() error {

	e.Engine, e.Err = xorm.NewEngine(driverName, dBconnect)
	if e.Err != nil {
		return e.Err
	}
	//tbMappe := core.NewPrefixMapper(core.SnakeMapper{}, tbMapper)
	e.Engine.ShowSQL(false)
	e.Engine.SetTableMapper(tbMappers)
	return nil
}

//查询员工id
func (e *Engine) SelectId(mobil string) (no, user string, err error) {
	nomao, err := e.Engine.Query(fmt.Sprintf("SELECT NO FROM BLCRM.CRM_SYS02 WHERE EMERGENCY_PHONE = '%s'", mobil))
	if err != nil {
		return no, user, err
	}
	for _, v := range nomao {
		for _, i := range v {
			if i != nil {
				no = string(i)
			}
		}
	}
	maps, err := e.Engine.Query(fmt.Sprintf("SELECT LOGIN_NAME FROM BLCRM.CRM_SYS04_N WHERE OPER_NO = '%s'", no))
	if err != nil {
		return no, user, err
	}

	for _, v := range maps {
		for _, i := range v {
			if string(i) != "" {
				user = string(i)
			}
		}
	}
	return no, user, err
}

//查询客户ID
func (e *Engine) SelectClientId(mobil string) (khid, gonghao string, err error) {
	maps, err := e.Engine.Query(fmt.Sprintf("SELECT KHID FROM BLCRM.CRM_DAT001 WHERE MOBIL = '%s'", mobil))
	if err != nil {
		return khid, gonghao, err
	}

	for _, v := range maps {
		for _, i := range v {
			if string(i) != "" {
				khid = string(i)
			}
		}
	}
	maps, err = e.Engine.Query(fmt.Sprintf("SELECT GONGHAO FROM BLCRM.CRM_DAT001 WHERE MOBIL = '%s'", mobil))
	if err != nil {
		return khid, gonghao, err
	}

	for _, v := range maps {
		for _, i := range v {
			if string(i) != "" {
				gonghao = string(i)
			}
		}
	}

	return khid, gonghao, err
}

//插入通话记录
func (e *Engine) InsertCallList(types, shichang int, time, gonghao, khid, telnomber, filename string) error {
	var sql string
	switch types {
	case 1:
		sql = fmt.Sprintf("INSERT INTO BLCRM.CRM_DAT101 VALUES (TO_DATE('%v','yyyy-MM-dd HH24:mi:ss'),'%s','%v','%s','%v','%s','%s')",
			time, telnomber, shichang, khid, 160004, gonghao, filename)
	case 2:
		sql = fmt.Sprintf("INSERT INTO BLCRM.CRM_DAT101 VALUES (TO_DATE('%v','yyyy-MM-dd HH24:mi:ss'),'%s','%v','%s','%v','%s','%s')",
			time, telnomber, shichang, khid, 160005, gonghao, filename)
	case 3:
		sql = fmt.Sprintf("INSERT INTO BLCRM.CRM_DAT101 VALUES (TO_DATE('%v','yyyy-MM-dd HH24:mi:ss'),'%s','%v','%s','%v','%s','%s')",
			time, telnomber, shichang, khid, 160006, gonghao, filename)
	case 4:
		sql = fmt.Sprintf("INSERT INTO BLCRM.CRM_DAT101 VALUES (TO_DATE('%v','yyyy-MM-dd HH24:mi:ss'),'%s','%v','%s','%v','%s','%s')",
			time, telnomber, shichang, khid, 160007, gonghao, filename)
	}

	_, e.Err = e.Engine.Exec(sql)
	if e.Err != nil {
		return e.Err
	}
	return nil
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

//http请求录音文件下载
func FileGet(Url, fileName string) ([]byte, error) {
	UrlA := Url + "?file_path=" + fileName
	response, err := http.Get(UrlA)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

//获取员工信息
func PostUserAll(Url string) (*GetAllUser, error) {
	query := url.Values{}
	query.Add("app_id", key)
	query.Add("page", "1")
	query.Add("page_size", "200")

	response, err := http.PostForm(Url, query)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	User := &GetAllUser{}
	err = json.Unmarshal(body, User)
	if err != nil {
		return nil, err
	}
	return User, nil
}

//匹配员工
func (e *Engine) SelectUser(mobil string, user *GetAllUser) (no, nomber string, err error) {

	for _, v := range user.Data.Result {
		if v.User_id == mobil {
			no, nomber, err := e.SelectId(v.Username)

			if err != nil {
				return no, nomber, err
			}
			return no, nomber, nil
		}
	}

	return no, nomber, nil
}

//存储录音文件
func CreateFile(data []byte, filename string) error {
	f, err := os.Create(filename)

	if err != nil {
		return err
	} else {
		_, err = f.Write(data)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {

	//获取员工信息
	//User := &GetAllUser{}
	//User, err := PostUserAll(url3)
	//if err != nil {
	//	fmt.Println("1:", err)
	//}

	Engine := &Engine{}

	time1 := fmt.Sprintf(time.Now().Format(TimeFormat)) + " 01:00:00"
	//time1 := fmt.Sprintf("2019-11-20") + " 01:00:00"

	for {

		//获取通话记录
		CallData := &CallJl{}
		CallData, err := CallPost(time1, url1)
		if err != nil {
			fmt.Println("2:", err)
		}

		if CallData.Code != 0 || CallData.Message != "" {
			fmt.Printf("3:%+v,%+v\n", CallData.Code, CallData.Message)
			fmt.Printf("3:%#v\n", CallData)

		}

		err = Engine.NewEngine()
		if err != nil {
			fmt.Println("4:", err)
		}
		defer Engine.Engine.Close()

		for _, v := range CallData.Data {
			no, number, err := Engine.SelectId(v.User_id)
			if err != nil {
				fmt.Println("5:", err)
			}

			client, gonghao, err := Engine.SelectClientId(v.Contact_phone)
			if err != nil {
				fmt.Println("6:", err)
			}

			if no == "" {
				if gonghao != "" {
					no = gonghao

					maps, _ := Engine.Engine.Query(fmt.Sprintf("SELECT LOGIN_NAME FROM BLCRM.CRM_SYS04_N WHERE OPER_NO = '%s'", no))

					for _, v := range maps {
						for _, i := range v {
							if string(i) != "" {
								number = string(i)
							}
						}
					}
				} else if gonghao == "" {
					continue
				}
			}

			if v.File != "" {
				//存储录音文件
				data, err := FileGet(url2, v.File)
				if err != nil {
					fmt.Println("7:", err)
				}

				str := strings.Split(v.Start_time, " ")
				dateS := strings.Split(str[0], "-")
				timeS := strings.Split(str[1], ":")

				err = CreateFile(data, fmt.Sprintf("D:/FTPROOT/%v%v%v/%v_%v%v%v.mp3", dateS[0], dateS[1], dateS[2], number, timeS[0], timeS[1], timeS[2]))
				if err != nil {
					fmt.Println("8:", err)
				}
				v.File = fmt.Sprintf("%v_%v%v%v.mp3", number, timeS[0], timeS[1], timeS[2])

			}

			err = Engine.InsertCallList(v.Type, v.Duration, v.Start_time, no, client, v.Contact_phone, v.File)
			if err != nil {
				fmt.Println("9:", err)
			}

			//time.Sleep(time.Second * 5)
		}
		if len(CallData.Data) == 1000 {
			time1 = CallData.Data[len(CallData.Data)-1].Update_time
		} else {
			return
		}

	}

}
