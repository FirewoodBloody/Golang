package modules

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	Url2         = "https://api.jd.com/routerjson"
	JdLoginId    = "029K708772"
	App_key      = "5BA6F95488F2BA2655367595505F7057" //应用标识
	App_secret   = "0053e1814a6345a19d7e06009281d5e9" //应用密钥
	access_token = "c5ec850f6c2a4c7288b51d7910df3673yzg1"
	method2      = "jingdong.ldop.receive.trace.get"
)

//用户验证信息
type PublicParametersr struct {
	Method       string `json:"method"`       //API接口名称
	Access_token string `json:"access_token"` //采用OAuth授权方式是必填参数
	App_key      string `json:"app_key"`      //应用的app_key
	//签名算法
	//将所有请求参数按照字母先后顺序排列 access_token,app_key,method,timestamp,v,360buy_param_json ，
	//排序为360buy_param_json,access_token,app_key,method,timestamp,v
	//把所有参数名和参数值进行拼接，例如：360buy_param_jsonxxxaccess_tokenxxxapp_keyxxxmethodxxxxxxtimestampxxxxxxvx
	//把appSecret夹在字符串（上一步拼接串）的两端，例如：appSecret+XXXX+appSecret
	//使用MD5进行加密，再转化成大写
	Sign      string `json:"sign"`      //签名算法：详见“签名算法”描述
	Timestamp string `json:"timestamp"` //时间戳，格式为yyyy-MM-dd HH:mm:ss，例如：2019-05-01 00:00:00。京东API服务端允许客户端请求时间误差为10分钟
	//Format         string `json:"format"`    //暂时只支持json
	V              string `json:"v"` //
	Buy_param_json string `json:"360buy_param_json"`
}

//请求信息
type RequestData struct {
	CustomerCode string `json:"customerCode"` //商家编码
	WaybillCode  string `json:"waybillCode"`  //运单号
}

//返回参数
type ParametersSelect struct {
	jingdong_ldop_receive_trace_get_responce `json:"jingdong_ldop_receive_trace_get_responce"`
}

//返回参数
type jingdong_ldop_receive_trace_get_responce struct {
	Querytrace_result `json:"querytrace_result"` //	返回结果
	Code              string                     `json:"code"` //状态
}

//	返回结果
type Querytrace_result struct {
	Messsage string `json:"messsage"` //	返回信息
	Code     int    `json:"code"`     //返回编码
	Data     []Data `json:"data"`     //返回数据
}

//路由信息
type Data struct {
	OpeTitle    string `json:"opeTitle"`    //操作标题
	OpeRemark   string `json:"opeRemark"`   //操作详情
	OpeName     string `json:"opeName"`     //操作人姓名
	OpeTime     string `json:"opeTime"`     //操作时间
	WaybillCode string `json:"waybillCode"` //	运单号
	Courier     string `json:"courier"`     //配送员
	CourierTel  string `json:"courierTel"`  //配送员电话
}

func (p *PublicParametersr) GetOrder1() *ParametersSelect {
	str := fmt.Sprintf("%v360buy_param_json%vaccess_token%vapp_key%vmethod%vtimestamp%vv%v%v", App_secret, p.Buy_param_json, p.Access_token, p.App_key, p.Method, p.Timestamp, p.V, App_secret)
	//字符串拼接，按照名称App_secret + 参数名称字母顺序和值进行排序 +App_secret
	//将以上字符串进行MD5加密，32位大写
	p.Sign = fmt.Sprintf("%X", md5.Sum([]byte(str)))

	//配置请求参数,按照参数顺序进行排列
	requestParameters := url.Values{}
	requestParameters.Set("360buy_param_json", p.Buy_param_json)
	requestParameters.Add("access_token", p.Access_token)
	requestParameters.Add("app_key", p.App_key)
	requestParameters.Add("method", p.Method)
	requestParameters.Add("sign", p.Sign)
	requestParameters.Add("timestamp", p.Timestamp)
	requestParameters.Add("v", p.V)

	//urlstr := fmt.Sprintf("%v?v=%v&method=%v&app_key=%v&access_token=%v&360buy_param_json=%v&timestamp=%v&sign=%v", Url1, p.V, p.Method, p.App_key, p.Access_token,	p.buy_param_json, p.Timestamp, p.Sign)
	//HTTP POST 请求方式
	resp, err := http.PostForm(Url1, requestParameters)

	if err != nil {
		fmt.Println("1:", err)
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("2:", err)
	}

	dataStruc := &ParametersSelect{}

	err = json.Unmarshal(data, dataStruc) //返回数据进行JSON反序列化
	if err != nil {
		fmt.Println("3:", err)
	}

	return dataStruc
}

//公共参数，用户授权
func (p *PublicParametersr) SetUserLogin1(data []byte) *ParametersSelect {
	p.App_key = App_key
	p.V = "2.0" //版本
	p.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	p.Access_token = access_token
	p.Method = method2
	p.Buy_param_json = string(data)

	return p.GetOrder1()

}

func SelectData(WaybillCode string) *ParametersSelect {
	P := new(PublicParametersr)

	R := new(RequestData)

	R.CustomerCode = "029K708772"
	R.WaybillCode = WaybillCode

	data, _ := json.Marshal(R)
	return P.SetUserLogin1(data)
}
