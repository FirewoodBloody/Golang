package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
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

	time1 := fmt.Sprintf("2021-03-01") + " 08:00:00"

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

			fmt.Printf("%#v", v)
			return
		}
	}

}
