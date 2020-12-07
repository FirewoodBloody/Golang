package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

//员工信息结构体
type GetAllUser struct {
	Error_code int    `json:"error_code"`
	Message    string `json:"message"`
	Data       Datas
}

//员工信息结构体
type Datas struct {
	CurrentPage int `json:"currentpage"` //当前页数
	TotalCount  int `json:"totalcount"`  //总条数
	TotalPage   int `json:"totalpage"`   //总页数
	Result      []User
}

//员工信息结构体
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

//通话记录
type Data struct {
	Id            int    `json:"id"`            //记录id
	User_id       string `json:"user_id"`       //员工id
	Type          int    `json:"type"`          //通话类型 1 呼出 2呼入 3呼出未接 4呼入未接
	Duration      int    `json:"duration"`      //通话时长(单位秒)
	File          string `json:"file"`          //通话文件
	Contact_name  string `json:"contact_name"`  //通话人名称
	Contact_phone string `json:"contact_phone"` //通话人号码
	Start_time    string `json:"start_time"`    //通话开始时间
	End_time      string `json:"end_time"`      //通话结束时间
	Update_time   string `json:"update_time"`   //更新时间 (该时间为查询时间）,获取了相同的id进行更新操作 用于下次请求通话记录
	Data_type     int    `json:"data_type"`     //通话类别 0 系统电话 1.微信语音 2.微信通话
	//Riend_id          int    `json:"riend_id"`          //好友id
	Friend_wx_id      string `json:"friend_wx_id"`      //String	好友微信Id
	Friend_alias      string `json:"friend_alias"`      //String	好友微信号
	Friend_chat_title string `json:"friend_chat_title"` //String	好友备注
}

const (
	key  = "si_bolong0911"
	url1 = "http://siai.aihujing.com:9989/phone/list" //通话记录，根据更新时间正序排序
	url2 = "http://siai.aihujing.com:9989/file/url"   //获取文件
	url3 = "http://siai.aihujing.com:9989/api/user/getAllUser"
)

func init() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码
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

	time1 := fmt.Sprintf("2020-08-01") + " 08:00:00"

	for {

		//获取通话记录
		CallData := &CallJl{}
		CallData, err := CallPost(time1, url1)
		if err != nil {
			fmt.Println("2:", err)
		}

		if CallData.Code != 0 || CallData.Message != "" {
			fmt.Printf("3:%#v\n", CallData)

		}

		for _, v := range CallData.Data {

			if v.File == "" {
				continue
			}

			fmt.Println(v)
			//拼接文件请求地址，获
			resp, _ := http.Get(url2 + "?file_path=" + v.File)
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)

			resp, _ = http.Get(string(body))
			defer resp.Body.Close()
			data, _ := ioutil.ReadAll(resp.Body)

			CreateFile(data, "./file.mp3")
			return
		}

		//if len(CallData.Data) == 1000 {
		//	//Update_time  这个字段为下一次请求时间开始的字段，之前我的程序是1天定时计划任务，我们可以改成1小时，数量应该不会超过1000  获取数量等于1000，那就需要使用Update_time 这个的值再次进行获取Update_time 截止当前时刻的通话记录
		//	time1 = CallData.Data[len(CallData.Data)-1].Update_time
		//} else {
		//	return
		//}

	}

}
