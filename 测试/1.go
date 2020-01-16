package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	Url1  = "http://opentestapi.yto.net.cn/service/waybill_query/v1/45qTKO"
	url23 = "http://opentestapi.yto.net.cn/service/waybill_query/v1/D2nvLq"
	url22 = "http://opentestapi.yto.net.cn/service/charge_query/v1/D2nvLq"
)

func main() {
	xmlValue := "<ufinterface><Result><WaybillCode><Number>111111111111</Number></WaybillCode></Result></ufinterface>"
	xml := "<ufinterface><Result><TransportInfo><StartProvince>陕西省</StartProvince><StartCity>西安市</StartCity><EndProvince>陕西省</EndProvince><EndCity>西安市</EndCity><GoodsWeight>1</GoodsWeight><GoodsWidth>1</GoodsWidth><GoodsHeight /><GoodsLength>1</GoodsLength></TransportInfo></Result></ufinterface>"

	_ = url.QueryEscape(xmlValue)

	dataStr := string(xmlValue) + "1QLlIZ"

	md5Key := md5.New()
	md5Key.Write([]byte(dataStr))

	// xmlKey := url.QueryEscape(base64.StdEncoding.EncodeToString(md5Key.Sum(nil)[:]))
	// xmlKey := base64.StdEncoding.EncodeToString(md5Key.Sum(nil)[:])
	a := "1QLlIZapp_keysF1JznformatXMLmethodyto.Marketing.TransportPricetimestamp2019-12-31 16:10:14user_idYTOTESTv1.01"

	requestParameters := url.Values{}
	requestParameters.Set("app_key", "sF1Jzn")
	requestParameters.Set("format", "XML")
	requestParameters.Add("method", "yto.Marketing.TransportPrice")
	requestParameters.Add("timestamp", "2019-12-31 16:10:14")
	requestParameters.Add("user_id", "YTOTEST")
	requestParameters.Add("v", "1.01")

	// aa, _ := url.QueryUnescape("%3C%3Fxml+version%3D%221.0%22+encoding%3D%22UTF-8%22+standalone%3D%22yes%22%3F%3E%3Cufinterface%3E%3CResult%3E%3CWaybillCode%3E%3CNumber%3E1111111111%3C%2FNumber%3E%3C%2FWaybillCode%3E%3C%2FResult%3E%3C%2Fufinterface%3E")
	requestParameters.Add("param", xml)

	// requestParameters.Add("Secret_Key", "1QLlIZ")
	requestParameters.Add("sign", fmt.Sprintf("%X", md5.Sum([]byte(a))))

	fmt.Printf("%X\n", md5.Sum([]byte(a)))

	// req, err := http.Get(string(Url1 + "?" + requestParameters.Encode()))
	// fmt.Println(Url1 + "?" + requestParameters.Encode())

	req, err := http.NewRequest("POST", url22, strings.NewReader(requestParameters.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	// 读取返回值
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(result))
}
