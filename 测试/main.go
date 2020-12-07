package main

import (
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	appKey    = "ExmlGKdGnSOzyfyywNqB7R"
	appSecret = "2226200909753282995473596416"
	Url       = "https://ds-sit.sf-express.com/externalapi/order/create/"
)

type ExternalOrderRequestVo struct {
	LoginAccount         string  `json:"loginAccount"`         // 是 代服系统登录账号
	CallBackUrl          string  `json:"callBackUrl"`          // 是 订单下单后异步通知地址
	OrderNumber          string  `json:"orderNumber"`          // 是 订单号（不能重复）
	CustomerName         string  `json:"customerName"`         // 否 客户姓名（商户专业版有，入门版没有）
	Amount               float64 `json:"amount"`               // 否 合计金额，保留两位小数（商户专业版有，入门版没有）
	OrderPayMethod       int     `json:"orderPayMethod"`       // 是 订单付款方式 1:在线支付 2:货到付款 如果填了2，代收金额必须大于0，代收卡号必填
	CollectingMoney      float64 `json:"collectingMoney"`      // 否 代收金额，只保留两位小数
	InsuredPrice         float64 `json:"insuredPrice"`         // 否 保价，只保留两位小数
	IndividuationPacking float64 `json:"individuationPacking"` // 否 包装服务费，只保留两位小数
	IsDoCall             int     `json:"isDoCall"`             //  否 是否通知上门揽收(1是 0否)
	ExpressProductCode   string  `json:"expressProductCode"`   //  是 物流产品编码（3.1.5产品编码映射表）
	MonthlyCard          string  `json:"monthlyCard"`          // 否 月结卡号
	Consignment          string  `json:"consignment"`          // 否 托寄物（入门版有，专业版没有）
	ConsignmentNumber    int     `json:"consignmentNumber"`    //  否 托寄物数量（入门版有且必填，专业版没有）
	ParcelWeight         float64 `json:"parcelWeight"`         // 否 包裹重量，默认0.5，保留2位
	ParcelQuantity       int     `json:"parcelQuantity"`       // 否 包裹数量，默认1
	CollectionCardNumber string  `json:"collectionCardNumber"` // 否 代收卡号
	SignBackType         int     `json:"signBackType"`         // 否 签回单类型 1签名 2盖章 3登记身份证 4身份证复印件
	OverweightServiceFee float64 `json:"overweightServiceFee"` // 否 超长超重服务费，只保留两位小数
	Remark               string  `json:"remark"`               // 否 订单备注
	BuyerMessage         string  `json:"buyerMessage"`         // 否 买家留言
	SellerMessage        string  `json:"sellerMessage"`        // 否 卖家留言
	CustomField1         string  `json:"customField1"`         // 否 自定义字段一
	CustomField2         string  `json:"customField2"`         // 否 自定义字段二
	WarehouseName        string  `json:"warehouseName"`        // 否 仓库名称（商户专业版有，入门版没有）
	FlowDirection        string  `json:"flowDirection"`        // 否 流向
	PostAge              float64 `json:"postAge"`              // 否 邮费，只保留两位小数
	SelfPickup           int     `json:"selfPickup"`           // 否 是否自取(1是 0否)
	FreshServices        int     `json:"freshServices"`        // 否 是否保鲜服务(1是 0否)
	AttributeNames       string  `json:"attributeNames"`       // 否 商品规格（商户专业版没有，入门版有）
	SenderName           string  `json:"senderName"`           // 是 寄件人
	SenderPhone          string  `json:"senderPhone"`          // 是 寄件人手机
	SenderTelephone      string  `json:"senderTelephone"`      // 否 寄件人电话
	SenderCompany        string  `json:"senderCompany"`        // 否 寄件人公司
	SenderProvince       string  `json:"senderProvince"`       // 是 寄件人省 (参照3.1.9行政区规范标准)
	SenderCity           string  `json:"senderCity"`           // 是 寄件人市 (参照3.1.9行政区规范标准)
	SenderDistrict       string  `json:"senderDistrict"`       // 是 寄件人区 (参照3.1.9行政区规范标准)
	SenderAddress        string  `json:"senderAddress"`        // 是 寄件人地址
	ReceiverName         string  `json:"receiverName"`         // 是 收件人
	ReceiverPhone        string  `json:"receiverPhone"`        // 是 收件人手机
	ReceiverTelephone    string  `json:"receiverTelephone"`    // 否 收件人电话
	ReceiverCompany      string  `json:"receiverCompany"`      // 否 收件人公司
	ReceiverProvince     string  `json:"receiverProvince"`     // 是 收件人省 (参照3.1.9行政区规范标准)
	ReceiverCity         string  `json:"receiverCity"`         // 是 收件人市 (参照3.1.9行政区规范标准)
	ReceiverDistrict     string  `json:"receiverDistrict"`     // 是 收件人区 (参照3.1.9行政区规范标准)
	ReceiverAddress      string  `json:"receiverAddress"`      // 是 收件人地址
	ExpressPayMethod     int     `json:"expressPayMethod"`     // 是 物流付款方式 1:寄付月结 2:寄付现结 3:收方付 4:第三方付
	Attach               string  `json:"attach"`               // 否 附加数据
	Order                []Order `json:"orderSkus"`            // 否 订单sku集合

}

type Order struct {
	ProductCode    string  `json:"productCode"`    // 是 商品编号
	ProductName    string  `json:"productName"`    // 是 商品名称
	SkuCode        string  `json:"skuCode"`        // 是 sku编码
	AttributeNames string  `json:"attributeNames"` // 是 商品规格，格式：颜色：红色，尺寸：L码....以此类推
	ProductNumber  int     `json:"productNumber"`  // 是 商品预定数量
	Unit           string  `json:"unit"`           // 否 商品单位
	Price          float64 `json:"price"`          // 是 商品单价
	AdjustAmount   float64 `json:"adjustAmount"`   // 否 调整金额
	SubAmount      float64 `json:"subAmount"`      // 是 小计金额
	Remark         string  `json:"remark"`         // 否 备注
	ShortName      string  `json:"shortName"`      // 否 商品简称
	MerchantCode   string  `json:"merchantCode"`   // 否 商家编码
}

func main() {
	sfOrder := new(ExternalOrderRequestVo)
	sfOrder.LoginAccount = "api001"
	sfOrder.CallBackUrl = "http://61.185.225.118:19374/v1/object/"
	sfOrder.OrderNumber = "DD20200101000111"
	sfOrder.OrderPayMethod = 2
	sfOrder.CollectingMoney = 1980.01
	sfOrder.ExpressProductCode = "1"
	sfOrder.MonthlyCard = "0298956226"
	sfOrder.Consignment = "第一套人民币"
	sfOrder.ConsignmentNumber = 3
	sfOrder.ParcelQuantity = 2
	sfOrder.CollectionCardNumber = sfOrder.MonthlyCard
	sfOrder.Remark = "订单号：DD20200101000111"
	sfOrder.SelfPickup = 1

	sfOrder.SenderName = "李小小"
	sfOrder.SenderPhone = "18620878883"
	sfOrder.SenderProvince = "广东省"
	sfOrder.SenderCity = "深圳市"
	sfOrder.SenderDistrict = "南山区"
	sfOrder.SenderAddress = "软件产业基地1C"

	sfOrder.ReceiverName = "柴雪新"
	sfOrder.ReceiverPhone = "17802928284"
	sfOrder.ReceiverProvince = "陕西省"
	sfOrder.ReceiverCity = "西安市"
	sfOrder.ReceiverDistrict = "雁塔区"
	sfOrder.ReceiverAddress = "旺座国际城E座25层"

	sfOrder.ExpressPayMethod = 1

	sfOrder.Order = make([]Order, 10)

	sfOrder.Order[0].SkuCode = "eb616106"
	sfOrder.Order[0].ProductCode = "123456789"
	sfOrder.Order[0].ProductName = "第一套人民币"
	sfOrder.Order[0].AttributeNames = "套"
	sfOrder.Order[0].ProductNumber = 3
	sfOrder.Order[0].Price = 1980.09
	sfOrder.Order[0].SubAmount = 1980.09

	data, _ := json.Marshal(sfOrder)
	timeUnix := time.Now().UnixNano() / 1e6

	str := Url + string(data) + "&appSecret=" + appSecret + "&timestamp=" + fmt.Sprintf("%v", timeUnix)
	sha := sha512.New()
	sha.Write([]byte(str))
	sign := sha.Sum(nil)
	shastr2 := hex.EncodeToString(sign)

	client := &http.Client{}

	//提交请求
	reqest, err := http.NewRequest("POST", Url, bytes.NewBuffer(data))

	reqest.Header.Add("appKey", appKey)
	reqest.Header.Add("timestamp", fmt.Sprintf("%v", timeUnix))
	reqest.Header.Add("sign", shastr2)

	resp, _ := client.Do(reqest)
	defer resp.Body.Close()
	datas, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("2:", err)
	}
	fmt.Println(string(datas))
}
