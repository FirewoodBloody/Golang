//快递鸟路由查询接口

package modules

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

//----------------------------------------------

//请求信息结构体
type DataValue struct {
	OrderCode    string
	ShipperCode  string
	LogisticCode string
}

//----------------------------------------------
//定义快递信息接收变量结构体
type RoutingInformation struct {
	LogisticCode string `json:"LogisticCode"` //物流运单号
	ShipperCode  string `json:"ShipperCode"`  //快递公司编码
	//Traces       []Traces //存储快递路由信息的数组
	State       string     `json:"State"`       //物流状态：2-在途中,3-签收,4-问题件
	EBusinessID string     `json:"EBusinessID"` //用户ID
	Success     bool       `json:"Success"`     //成功与否
	OrderCode   string     `json:"OrderCode"`   //订单编号
	Reason      string     `json:"Reason"`      //失败原因
	Traces      []struct { //路由状态接收变量,存储快递路由信息的数组
		AcceptStation string `json:"AcceptStation"` //动态路由信息描述
		AcceptTime    string `json:"AcceptTime"`    //动态路由信息触发时间
		Remark        string `json:"Remark"`        //备注
	}
}

const (
	EBusinessID = "1402459"                                                   //用户ID
	ApiKye      = "cc5d10c3-4710-4f85-ad26-d49eb5502d2c"                      //秘钥
	ApiURL      = "http://api.kdniao.com/Ebusiness/EbusinessOrderHandle.aspx" //请求地址
)

//----------------------------------------------
//获取快递公司、单号等信息
//获取快递路信息，并将其反序列化，返回一个结构体
func KdnExpressInformation(ShipperCode, LogisticCode string) (KdnDataStruct *RoutingInformation, err error) {
	//ShipperCode := "YTO"                                               //快递公司接口编码
	//LogisticCode := "818803471597"                                     //快递单号
	RequestParameters, err := KdnCreateData(ShipperCode, LogisticCode) //初始化请求数据，配置请求参数
	if err != nil {
		return
	}

	//Post 查询请求路由信息（快递鸟）
	Data, err := KdnPost(ApiURL, RequestParameters)
	if err != nil {
		return
	}

	KdnDataStruct = &RoutingInformation{}

	//json反序列化
	err = json.Unmarshal(Data, KdnDataStruct)

	if err != nil {
		os.Exit(1)
	}

	return
}

//新建请求内容，并进行加密处理
//初始化请求数据，配置请求参数
//需求快递的单号，以及快递对应的公司代码
//ShipperCode := "YTO"                                               //快递公司接口编码
//LogisticCode := "818803471597"                                     //快递单号
func KdnCreateData(ShipperCode, LogisticCode string) (requestParameters url.Values, err error) {
	//定义快递请求信息
	data := DataValue{
		OrderCode:    "",
		ShipperCode:  ShipperCode,
		LogisticCode: LogisticCode,
	}

	//将struct序列化为json
	dataA, err := json.Marshal(data)
	if err != nil {
		return
	}
	//字符串转换,DataSign 请求内容(未编码)+AppKey
	dataStr := fmt.Sprintf("%s%s", dataA, ApiKye)

	//Apk 和 请求内容加密  DataSign
	//MD5 首先进行MD5加密，返回数据dataStr的MD5校验和，并将其转换为字符，MD5加密后默认为[16]byte
	//对MD5加密的数据进行base64 编码
	//对base64 加密的数据进行转码使之可以安全的用在URL查询里
	encryption := url.QueryEscape(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%x", md5.Sum([]byte(dataStr))))))

	//初始化请求参数
	requestParameters = url.Values{}

	//配置请求参数,方法内部已处理 urlencode 问题,中文参数可以直接传参
	requestParameters.Set("RequestData", string(dataA)) //设置请求内容需进行URL(utf-8)编码。请求内容JSON格式，须和DataType一致。
	requestParameters.Set("EBusinessID", EBusinessID)   //商户ID
	requestParameters.Set("RequestType", "1002")        //请求指令类型：1002
	requestParameters.Set("DataSign", encryption)       //数据内容签名：把(请求内容(未编码)+AppKey)进行MD5加密，然后Base64编码，最后 进行URL(utf-8)编码。
	return
}

//Post 查询请求路由请求（快递鸟）
//resp 返回快递路由信息
func KdnPost(ApiURL string, value url.Values) (rs []byte, err error) {
	resp, err := http.PostForm(ApiURL, value)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
