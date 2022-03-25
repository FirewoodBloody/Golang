package express

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// T2 查询运单状态
type T2 struct {
	OrderNumber string `json:"orderNumber"` // 订单号
	MailNumber  int    `json:"mailNumber"`  // 顺丰运单号
}

// T3 查询下单结果
type T3 struct {
	OrderNumber     string `json:"orderNumber"`     // 订单单号
	ExInterfaceType int    `json:"exInterfaceType"` // 1 下单结果 2 取消订单结果
}

// T4 查询下单响应字段
type T4 struct {
	Msg    string `json:"msg"`
	Code   string `json:"code"`
	Result struct {
		BspResponse struct {
			Body []struct {
				OrderNumber          string      `json:"orderNumber"`
				MailNumber           string      `json:"mailNumber"`
				ReturnTrackingNumber interface{} `json:"returnTrackingNumber"`
				OriginCode           string      `json:"originCode"`
				DestCode             string      `json:"destCode"`
				FilterResult         int         `json:"filterResult"`
				Remark               interface{} `json:"remark"`
				AgentMailNumber      interface{} `json:"agentMailNumber"`
				MappingMark          interface{} `json:"mappingMark"`
				Url                  interface{} `json:"url"`
				PaymentLink          interface{} `json:"paymentLink"`
				RlsInfo              []struct {
					InvokeResult string      `json:"invokeResult"`
					RlsCode      string      `json:"rlsCode"`
					ErrorDesc    interface{} `json:"errorDesc"`
					RlsDetail    []struct {
						WaybillNumber       string      `json:"waybillNumber"`
						SourceTransferCode  string      `json:"sourceTransferCode"`
						SourceCityCode      string      `json:"sourceCityCode"`
						SourceDeptCode      string      `json:"sourceDeptCode"`
						SourceTeamCode      string      `json:"sourceTeamCode"`
						DestCityCode        string      `json:"destCityCode"`
						DestDeptCode        string      `json:"destDeptCode"`
						DestDeptCodeMapping interface{} `json:"destDeptCodeMapping"`
						DestTeamCode        interface{} `json:"destTeamCode"`
						DestTeamCodeMapping interface{} `json:"destTeamCodeMapping"`
						DestTransferCode    interface{} `json:"destTransferCode"`
						DestRouteLabel      string      `json:"destRouteLabel"`
						ProName             string      `json:"proName"`
						CargoTypeCode       string      `json:"cargoTypeCode"`
						LimitTypeCode       string      `json:"limitTypeCode"`
						ExpressTypeCode     string      `json:"expressTypeCode"`
						CodingMapping       interface{} `json:"codingMapping"`
						CodingMappingOut    interface{} `json:"codingMappingOut"`
						XbFlag              string      `json:"xbFlag"`
						PrintFlag           string      `json:"printFlag"`
						TwoDimensionCode    string      `json:"twoDimensionCode"`
						ProCode             string      `json:"proCode"`
						PrintIcon           string      `json:"printIcon"`
						AbFlag              interface{} `json:"abFlag"`
						ErrMsg              interface{} `json:"errMsg"`
						DestPortCode        interface{} `json:"destPortCode"`
						DestCountry         interface{} `json:"destCountry"`
						DestPostCode        interface{} `json:"destPostCode"`
						GoodsValueTotal     interface{} `json:"goodsValueTotal"`
						CurrencySymbol      interface{} `json:"currencySymbol"`
						GoodsNumber         interface{} `json:"goodsNumber"`
					} `json:"rlsDetail"`
				} `json:"rlsInfo"`
			} `json:"body"`
			Head    string `json:"head"`
			Service string `json:"service"`
		} `json:"bspResponse"`
		MailNumber  string `json:"mailNumber"`
		OrderNumber string `json:"orderNumber"`
	} `json:"result"`
}

// T5 取消下单结果相应字段
type T5 struct {
	Msg    string `json:"msg"`
	Code   string `json:"code"`
	Result struct {
		OrderNumber  string `json:"orderNumber"`
		IsCancelSuss int    `json:"isCancelSuss"`
		MailNumber   string `json:"mailNumber"`
	} `json:"result"`
}

// T7 查询运单状态
type T7 struct {
	Msg    string `json:"msg"`
	Code   string `json:"code"`
	Result struct {
		MailNumber string `json:"mailNumber"`
		Status     int    `json:"status"`
		AcceptTime string `json:"acceptTime"`
	} `json:"result"`
}

// T6 运单路由
type T6 struct {
	Msg    string `json:"msg"`
	Code   string `json:"code"`
	Result struct {
		List []struct {
			AcceptTime    string `json:"acceptTime"`
			Remark        string `json:"remark"`
			AcceptAddress string `json:"acceptAddress"`
			OpCode        string `json:"opCode"`
			MailNo        string `json:"mailNo"`
			ReasonCode    string `json:"reasonCode"`
			ExtendAttach4 string `json:"extendAttach4"`
		} `json:"list"`
	} `json:"result"`
}

func POST(orderNumber, mailNumber string, exInterfaceType int) []byte {
	client := http.Client{}

	s1 := new(T3)
	s1.OrderNumber = orderNumber
	s1.ExInterfaceType = exInterfaceType
	data, _ := json.Marshal(s1)
	timestamp := time.Now().UnixNano() / 1e6
	Sign := getSign(appsecret, fmt.Sprint(timestamp), "", string(data))

	request, _ := http.NewRequest("POST", cxxdurl, bytes.NewBufferString(string(data)))
	request.Header.Add("charset", "UTF-8")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("appKey", appkey)
	request.Header.Add("timestamp", fmt.Sprint(timestamp))
	request.Header.Add("sign", Sign)

	response, err := client.Do(request)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	data, _ = ioutil.ReadAll(response.Body)
	return data
}
