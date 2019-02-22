package main

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Request struct {
	Service string `xml:"service,attr"`
	Lang    string `xml:"lang,attr"`
	Head    string `xml:"Head"`
	Body    Body   `xml:"Body"`
}

type Body struct {
	RouteRequest RouteRequest `xml:"RouteRequest"`
}

type RouteRequest struct {
	Tracking_type   int    `xml:"tracking_type,attr"`
	Method_type     int    `xml:"method_type,attr"`
	Tracking_number string `xml:"tracking_number,attr"`
}

var (
	CourierNumber string //快递单号
)

const (
	ClientCode = "BLWHYSP_tipea"
	CheckWord  = "eJXycwXzucwcxit4a8WK7b3qGl4UfkB1"
	SFurl      = "http://bsp-oisp.sf-express.com/bsp-oisp/sfexpressService"
)

//获取快递公司、单号等信息
//获取快递路信息，并将其反序列化，返回一个结构体
func SfExpressInformation() (SfDataStruct *struct{}, err error) {
	RequestParameters, err := SfCreateData()
	if err != nil {
		return
	}

	_, err = SfPost(SFurl, RequestParameters)
	if err != nil {
		return
	}

	return &struct{}{}, nil
}

//Post 查询请求路由请求（快递鸟）
//resp 返回快递路由信息
func SfPost(ApiURL string, value url.Values) (rs []byte, err error) {
	resp, err := http.PostForm(ApiURL, value)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//初始化 请求信息，进行加密处理
func SfCreateData() (requestParameters url.Values, err error) {
	//定义快递请求信息
	body := &Request{
		Service: "RouteService",
		Lang:    "zh-CN",
		Head:    ClientCode,
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
		fmt.Println(err)
		return
	}

	//对xml报头和Key进行加密
	md5Key := md5.New()
	md5Key.Write([]byte(fmt.Sprintf("%s%s", dataXml, CheckWord)))
	xmlKey := base64.StdEncoding.EncodeToString(md5Key.Sum(nil))

	//配置请求参数,
	requestParameters = url.Values{}
	requestParameters.Set("xml", string(dataXml))
	requestParameters.Set("verifyCode", xmlKey)

	return
}

func main() {
	CourierNumber = "547709536321"
	RequestParameters, err := SfCreateData()
	if err != nil {
		return
	}

	data, err := SfPost(SFurl, RequestParameters)
	if err != nil {
		return
	}

	fmt.Printf("%s", data)
}
