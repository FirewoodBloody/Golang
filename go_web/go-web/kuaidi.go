package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type DataValue struct {
	OrderCode    string
	ShipperCode  string
	LogisticCode string
}

const (
	EBusinessID = "1402459"
	ApiKye      = "cc5d10c3-4710-4f85-ad26-d49eb5502d2c"
	ApiURL      = "http://api.kdniao.com/Ebusiness/EbusinessOrderHandle.aspx"
)

//获取快递公司、单号等信息
func ExpressInformation() {
	ShipperCode := "YTO"           //快递公司接口编码
	LogisticCode := "818803471597" //快递单号
	RequestParameters, err := CreateData(ShipperCode, LogisticCode)
	if err != nil {
		return
	}

	Data, err := Post(ApiURL, RequestParameters)
	if err != nil {
		return
	}
	err = json.Unmarshal(Data)
	if err != nil {
		fmt.Println(err)
		return
	}

}

//新建请求内容，并进行加密处理
//初始化请求数据，配置请求参数
func CreateData(ShipperCode, LogisticCode string) (requestParameters url.Values, err error) {
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
	//MD5 首先进行MD5加密，返回数据dataStr的MD5校验和，并将其转换为字符串，MD5加密后默认为[16]byte
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
func Post(ApiURL string, value url.Values) (rs []byte, err error) {
	resp, err := http.PostForm(ApiURL, value)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func main() {
	ExpressInformation()
}
