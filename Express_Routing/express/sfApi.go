package express

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	UserId  = "BLWHYSP_tipea"                                            //顾客编码
	UserKey = "eJXycwXzucwcxit4a8WK7b3qGl4UfkB1"                         //秘钥
	ApiUrl  = "http://bsp-oisp.sf-express.com/bsp-oisp/sfexpressService" //请求地址
)

//请求接口基本信息
type Request struct {
	Service string `xml:"service,attr"` //请求服务名称
	Lang    string `xml:"lang,attr"`    //路由支持语言类型zh-CN
	Head    string `xml:"Head"`         //顾客编码
	Body    Body   `xml:"Body"`         //请求数据XML
}

//请求数据XML
type RouteRequest struct {
	//用于定义请求信息的参数
	Tracking_type   int    `xml:"tracking_type,attr"`   //查询号类别：
	Method_type     int    `xml:"method_type,attr"`     //路由查询类别：
	Tracking_number string `xml:"tracking_number,attr"` //查询单号
}

//Body数据XML
type Body struct {
	RouteRequest RouteRequest `xml:"RouteRequest"` //请求数据XML

}

//接收数据接口
type Response struct {
	Service string `xml:"service,attr"`
	Head    string `xml:"Head"` //顾客编码
	Body    struct {
		RouteResponse RouteResponse `xml:"RouteResponse"` // 接收数据的参数
	} `xml:"Body"`
}

//接收数据
type RouteResponse struct {
	Mailno string `xml:"mailno,attr"` //接收数据的单号
	//用于定义接收信息的参数
	Route []Route //路由更新状态参数
}

//路由更新状态
type Route struct {
	Remark         string `xml:"remark,attr"`         //接收路由变更信息
	Accept_Time    string `xml:"accept_time,attr"`    //接收路由变更对应的时间
	Accept_Address string `xml:"accept_address,attr"` //接收路由变更时所在城市
	Opcode         string `xml:"opcode,attr"`         //接收路由变更时状态代码,路由状态代码
}

//Post 查询请求路由请求（顺丰）
//resp 返回快递路由信息
func SfPost(requestParameters url.Values) (SfDataStruct *Response, err error) {
	resp, err := http.PostForm(ApiUrl, requestParameters)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}
	SfDataStruct = &Response{}

	err = xml.Unmarshal(data, SfDataStruct) //反序列化
	if err != nil {
		return nil, err
	}

	return SfDataStruct, nil
}

//初始化 请求信息，进行加密处理
func SfCreateData(CourierNumber string) (SfDataStruct *Response, err error) {
	//定义快递请求信息
	body := &Request{
		Service: "RouteService",
		Lang:    "zh-CN",
		Head:    UserId,
		Body: Body{
			RouteRequest{
				Tracking_type:   1,
				Method_type:     1,
				Tracking_number: CourierNumber,
			},
		},
	}

	// xml 序列化
	dataXml, err := xml.MarshalIndent(body, "", "")
	if err != nil {

		return nil, err
	}

	//对xml报头和Key进行加密
	md5Key := md5.New()
	md5Key.Write([]byte(fmt.Sprintf("%s%s", dataXml, UserKey)))
	xmlKey := base64.StdEncoding.EncodeToString(md5Key.Sum(nil))

	//配置请求参数,
	requestParameters := url.Values{}
	requestParameters.Set("xml", string(dataXml))
	requestParameters.Set("verifyCode", xmlKey)

	return SfPost(requestParameters)
}
