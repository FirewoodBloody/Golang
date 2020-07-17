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

//异常单请求参数
type FormatAn struct {
	CustomerCode    string `json:"customerCode"`    //是	010k12323	商家编码（长度：不超过20个字符
	DeliveryId      string `json:"deliveryId"`      //是	VC012323123	运单号（长度：不超过50个字符）
	ResponseComment string `json:"responseComment"` //是	审批通过	返回描述（长度：不超过100个字符）
	Type            int    `json:"type"`            //是	1	异常处理结果（1：再投，2：退回，3：站点报废，5：再投后退回，6：再投后报废）
}
type Jlaar struct {
	Jingdong_ldop_abnormal_approval_responce Jingdong_ldop_abnormal_approval_responce `json:"jingdong_ldop_abnormal_approval_responce"`
}

type Jingdong_ldop_abnormal_approval_responce struct {
	Approval_result struct { //返回结果
		StatusCode    int    `json:"statusCode"`    //状态码
		StatusMessage string `json:"statusMessage"` //	状态描述
	} `json:"approval_result"`
	Code string `json:"code"`
}

type Fta struct {
	CustomerCode string `json:"customerCode"`
}

type Datess struct {
	Jingdong_ldop_abnormal_get_responce struct {
		Code                    string `json:"code"`
		StatusMessage           string `json:"statusMessage"`
		Querybycondition_result struct {
			StatusCode    int    `json:"statusCode"`
			StatusMessage string `json:"statusMessage"`
			Data          []struct {
				OrderId             string `json:"orderId"`             //	102022123	订单号（长度：不超过50个字符）
				DeliveryId          string `json:"deliveryId"`          //	String[]	VC102020102	运单号（长度：不超过50个字符）
				OperateTime         int64  `json:"operateTime"`         //	Date[]		操作时间
				MainTypeName        string `json:"mainTypeName"`        //	String[]	整单协商再投	类型名称
				ReqestComment       string `json:"reqestComment"`       //	String[]	通过	请求描述（长度：不超过100个字符）
				CurrentAuditCounter int    `json:"currentAuditCounter"` //		Number	1	当前审核次
				TotalAuditCounter   int    `json:"totalAuditCounter"`   //	Number	2	审批总次数
				//TotalAuditCounter1  int    `json:"totalAuditCounter1"`  //	Number	2	审批总次数
			} `json:"data"`
		} `json:"querybycondition_result"`
	} `json:"jingdong_ldop_abnormal_get_responce"`
}

const (
	Anomaly_modules      = "jingdong.ldop.abnormal.approval"
	Anomaly_modules_list = "jingdong.ldop.abnormal.get"
)

//处理异常订单
func Anomaly_Post(code string) *Jlaar {
	//请求参数
	f := new(FormatAn)
	f.CustomerCode = JdLoginId
	f.DeliveryId = code
	f.ResponseComment = "审批通过"
	f.Type = 1
	data, _ := json.Marshal(f)

	//请求用户校验数据
	p := new(PublicParameters)
	p.App_key = App_key
	p.V = "2.0" //版本
	p.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	p.Access_token = access_token
	p.Method = Anomaly_modules
	p.Buy_param_json = string(data)

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

	resp, err := http.PostForm(Url1, requestParameters)

	if err != nil {
		fmt.Println("1:", err)
	}

	defer resp.Body.Close()
	datas, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("2:", err)
	}
	fmt.Println(string(datas))
	jla := new(Jlaar)
	err = json.Unmarshal(datas, jla)
	fmt.Println(err)
	return jla
}

func Anomaly_Post_list(code string) []byte {
	//请求参数
	f := new(Fta)
	f.CustomerCode = code
	data, _ := json.Marshal(f)

	//请求用户校验数据
	p := new(PublicParameters)
	p.App_key = App_key
	p.V = "2.0" //版本
	p.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	p.Access_token = access_token
	p.Method = Anomaly_modules_list
	p.Buy_param_json = string(data)

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

	resp, err := http.PostForm(Url1, requestParameters)

	fmt.Println(requestParameters.Encode())

	if err != nil {
		fmt.Println("1:", err)
	}

	defer resp.Body.Close()
	datas, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("2:", err)
	}
	//fmt.Println(string(datas))
	return datas
}
