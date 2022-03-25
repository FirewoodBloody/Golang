package models

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	SelectUrl = "https://poll.kuaidi100.com/poll/query.do" //请求链接
)

type Param struct {
	Com      string `json:"com"`      // 是 查询的快递公司的编码， 一律用小写字母 下载编码表格
	Num      string `json:"num"`      // 是 查询的快递单号， 单号的最大长度是32个字符
	Phone    string `json:"phone"`    // 否 收、寄件人的电话号码（手机和固定电话均可，只能填写一个，顺丰速运和丰网速运必填，其他快递公司选填。如座机号码有分机号，分机号无需上传。）
	From     string `json:"from"`     // 否 出发地城市
	To       string `json:"to"`       // 否 目的地城市，到达目的地后会加大监控频率
	Resultv2 string `json:"resultv2"` // 否 添加此字段表示开通行政区域解析功能。0：关闭（默认），1：开通行政区域解析功能以及物流轨迹增加物流状态值，2：开通行政解析功能以及物流轨迹增加物流状态值并且返回出发、目的及当前城市信息 4：开通行政解析功能以及物流轨迹增加物流高级状态名称并且返回出发、目的及当前城市信息 6:开通行政解析功能以及物流轨迹增加物流高级状态名称、状态值并且返回出发、目的及当前城市信息
	Show     string `json:"show"`     // 否 返回格式：0：json格式（默认），1：xml，2：html，3：text
	Order    string `json:"order"`    // 否 返回结果排序:desc降序（默认）,asc 升序
}

type T struct {
	Message   string     `json:"message"`   //消息体，请忽略
	Nu        string     `json:"nu"`        //单号
	Ischeck   string     `json:"ischeck"`   //是否签收标记，请忽略，明细状态请参考state字段
	Condition string     `json:"condition"` //快递单明细状态标记，暂未实现，请忽略
	Com       string     `json:"com"`       //快递公司编码,一律用小写字母
	Status    string     `json:"status"`    //通讯状态，请忽略
	State     string     `json:"state"`     //快递单当前状态，包括0在途，1揽收，2疑难，3签收，4退签，5派件，6退回，7转单，10待清关，11清关中，12已清关，13清关异常，14拒签等36个状态
	Data      []struct { //最新查询结果，数组，包含多项，全量，倒序（即时间最新的在最前），每项都是对象，对象包含字段请展开
		Time    string `json:"time"`    //
		Ftime   string `json:"ftime"`   //
		Context string `json:"context"` //
	} `json:"data"`
}

// Post 请求进行快递快递路由查询
func Post() {
	p := new(Param)
	p.Com = "shunfeng"
	p.Num = "SF1327875798237"
	data, _ := json.Marshal(p)
	str := "?customer=" + Customer + "&sign=" + Md5(string(data), Customer, AppKey) + "&param=" + string(data)

	client := http.Client{}

	requst, _ := http.NewRequest("POST", SelectUrl+str, nil)
	requst.Header.Add("Content-Type", "application/json;charset=UTF-8")

	response, err := client.Do(requst)
	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))

}

// Md5 sign 加密
func Md5(param, customer, appKey string) string {
	str := param + appKey + customer
	m := md5.New()
	m.Write([]byte(str))
	return fmt.Sprintf("%X", m.Sum(nil))
}
