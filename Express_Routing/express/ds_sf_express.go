package express

import (
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	xdurl     = "https://ds-api.sf-express.com/externalapi/order/create"
	appkey    = "ExuXufHwRagVqNkCxgSsMT"
	appsecret = "16920210202806206424614043648"
	cxxdurl   = "https://ds-api.sf-express.com/externalapi/query/order"
	ydzt      = "https://ds-api.sf-express.com/externalapi/query/waybill"
	lyzt      = "https://ds-api.sf-express.com/externalapi/query/route"
)

type T struct {
	LoginAccount         string  `json:"loginAccount"`
	CallBackUrl          string  `json:"callBackUrl"`
	OrderNumber          string  `json:"orderNumber"`
	OrderPayMethod       int     `json:"orderPayMethod"`
	ExpressProductCode   string  `json:"expressProductCode"`
	MonthlyCard          string  `json:"monthlyCard"`
	CollectionCardNumber string  `json:"collectionCardNumber"`
	ExpressPayMethod     int     `json:"expressPayMethod"`
	CollectingMoney      int     `json:"collectingMoney"`
	InsuredPrice         float64 `json:"insuredPrice"`
	IndividuationPacking float64 `json:"individuationPacking"`
	IsDoCall             int     `json:"isDoCall"`
	ParcelWeight         float64 `json:"parcelWeight"`
	ParcelQuantity       int     `json:"parcelQuantity"`
	SignBackType         string  `json:"signBackType"`
	OverweightServiceFee float64 `json:"overweightServiceFee"`
	WarehouseName        string  `json:"warehouseName"`
	CustomerName         string  `json:"customerName"`
	PostAge              float64 `json:"postAge"`
	SelfPickup           string  `json:"selfPickup"`
	FreshServices        string  `json:"freshServices"`
	Remark               string  `json:"remark"`
	BuyerMessage         string  `json:"buyerMessage"`
	SellerMessage        string  `json:"sellerMessage"`
	CustomField1         string  `json:"customField1"`
	CustomField2         string  `json:"customField2"`
	FlowDirection        string  `json:"flowDirection"`
	SenderName           string  `json:"senderName"`
	SenderPhone          string  `json:"senderPhone"`
	SenderTelephone      string  `json:"senderTelephone"`
	SenderCompany        string  `json:"senderCompany"`
	SenderProvince       string  `json:"senderProvince"`
	SenderCity           string  `json:"senderCity"`
	SenderDistrict       string  `json:"senderDistrict"`
	SenderAddress        string  `json:"senderAddress"`
	ReceiverName         string  `json:"receiverName"`
	ReceiverPhone        string  `json:"receiverPhone"`
	ReceiverTelephone    string  `json:"receiverTelephone"`
	ReceiverCompany      string  `json:"receiverCompany"`
	ReceiverProvince     string  `json:"receiverProvince"`
	ReceiverCity         string  `json:"receiverCity"`
	ReceiverDistrict     string  `json:"receiverDistrict"`
	ReceiverAddress      string  `json:"receiverAddress"`
	OrderSkus            []struct {
		OrderType      int     `json:"orderType"`
		ProductCode    string  `json:"productCode"`
		ProductName    string  `json:"productName"`
		SkuCode        string  `json:"skuCode"`
		AttributeNames string  `json:"attributeNames"`
		ProductNumber  int     `json:"productNumber"`
		Unit           string  `json:"unit"`
		Price          float64 `json:"price"`
		AdjustAmount   float64 `json:"adjustAmount"`
		SubAmount      float64 `json:"subAmount"`
		Remark         string  `json:"remark"`
		ShortName      string  `json:"shortName"`
		MerchantCode   string  `json:"merchantCode"`
	} `json:"orderSkus"`
}

//sha512加密
//返回十六进制编码
func getSign(appSecret, timestamp, getParameterString, body string) string {
	Str := getParameterString + body + "&appSecret=" + appSecret + "&timestamp=" + timestamp
	hx := sha512.New()
	hx.Write([]byte(Str))
	return hex.EncodeToString(hx.Sum(nil))
}

func Post() []byte {
	body := "{\"loginAccount\":\"SF0294878552\",\"callBackUrl\":\"\",\"orderNumber\":\"DD20210625000099\",\"orderPayMethod\":2,\"expressProductCode\":\"1\",\"monthlyCard\":\"0294878552\",\"collectionCardNumber\":\"0294878552\",\"expressPayMethod\":1,\"collectingMoney\":1,\"insuredPrice\":2.222,\"individuationPacking\":2.232,\"isDoCall\":1,\"parcelWeight\":0.252,\"parcelQuantity\":1,\"signBackType\":\"2\",\"overweightServiceFee\":2.242,\"warehouseName\":\"默认仓库\",\"customerName\":\"默认客户\",\"postAge\":3.25,\"selfPickup\":\"1\",\"freshServices\":\"0\",\"remark\":\"订单备注\",\"buyerMessage\":\"买家留言\",\"sellerMessage\":\"卖家留言\",\"customField1\":\"自定义字段一\",\"customField2\":\"自定义字段二\",\"flowDirection\":\"华北\",\"senderName\":\"何志谋\",\"senderPhone\":\"18620878883\",\"senderTelephone\":\"0755-32516213\",\"senderCompany\":\"寄件人公司\",\"senderProvince\":\"广东省\",\"senderCity\":\"阳江市\",\"senderDistrict\":\"江城区\",\"senderAddress\":\"峰荟花园1B\",\"receiverName\":\"陈月建\",\"receiverPhone\":\"13332973440\",\"receiverTelephone\":\"020-325162216\",\"receiverCompany\":\"收件人公司\",\"receiverProvince\":\"广东省\",\"receiverCity\":\"广州市\",\"receiverDistrict\":\"天河区\",\"receiverAddress\":\"天河地铁站A出口\",\"orderSkus\":[{\"orderType\":1,\"productCode\":\"测试编码\",\"productName\":\"测试商品\",\"skuCode\":\"JBY0002\",\"attributeNames\":\"颜色:红色,尺寸:L\",\"productNumber\":2,\"unit\":\"件\",\"price\":50.55,\"adjustAmount\":60.55,\"subAmount\":63.668,\"remark\":\"aaaa\",\"shortName\":\"商品简称\",\"merchantCode\":\"123456789\"},{\"orderType\":1,\"productCode\":\"测试编码1\",\"productName\":\"测试商品1\",\"skuCode\":\"JBY0002\",\"attributeNames\":\"颜色:绿色,尺寸:L\",\"productNumber\":2,\"unit\":\"件\",\"price\":50.55,\"adjustAmount\":60.55,\"subAmount\":63.668,\"remark\":\"aaaa\",\"shortName\":\"商品简称1\",\"merchantCode\":\"123456788\"}]}"
	timestamp := time.Now().UnixNano() / 1e6
	Sign := getSign(appsecret, fmt.Sprint(timestamp), "", body)

	client := &http.Client{}

	request, err := http.NewRequest("POST", xdurl, bytes.NewBufferString(body))
	request.Header.Add("charset", "UTF-8")
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("appKey", appkey)
	//request.Header.Add("appSecret", appsecret)
	request.Header.Add("timestamp", fmt.Sprint(timestamp))
	request.Header.Add("sign", Sign)
	//处理返回结果
	response, _ := client.Do(request)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)
	return data
}
