//圆通下单接口

package modules

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//<!—订单基本信息-->
type RequestOrder struct {
	ClientID           string `xml:"clientID"`           //渠道编码（电商标识，由圆通人员给出）
	LogisticProviderID string `xml:"logisticProviderID"` //物流公司ID（YTO）
	CustomerId         string `xml:"customerId"`         //客户标识（COD业务，且有多个仓发货则不能为空，请填写分仓号）（可选）
	TxLogisticID       string `xml:"txLogisticID"`       //物流号
	//TradeNo            string `xml:"tradeNo"`              //业务交易号（可选）
	//MailNo             string `xml:"mailNo"`               //物流运单号（可选）
	//TotalServiceFee    string `xml:"totalServiceFee"`      //总服务费[COD]（可选）
	//CodSplitFee        string `xml:"codSplitFee"`          //物流公司分润[COD]（可选）
	OrderType   int `xml:"orderType"`   //订单类型(0-COD,1-普通订单,3-退货单)
	ServiceType int `xml:"serviceType"` //服务类型(1-上门揽收, 2-次日达 4-次晨达 8-当日达,0-自己联系)。（数据库未使用）（目前暂未使用默认为0）
	//Flag        int `xml:"flag"`        //订单flag标识（暂未使用）
	//<!—发货方信息-->
	Sender Sender `xml:"sender"`
	//<!--收货方信息-->
	Receiver Receiver `xml:"receiver"`
	//<!--物流公司上门取货时间段-->
	//SendStartTime time.Time `xml:"sendStartTime"` //（可选）物流公司上门取货时间段，通过“yyyy-MM-dd HH:mm： ss.S z”格式化，本文中所有时间格式相同。
	//SendEndTime   time.Time `xml:"sendEndTime"`   //（可选）物流公司上门取货时间段，通过“yyyy-MM-dd HH:mm： ss.S z”格式化，本文中所有时间格式相同。
	//<!--商品信息-->
	//GoodsValue     float64 `xml:"goodsValue"`     //（可选）商品金额，包括优惠和运费，但无服务费
	//ItemsValue     float64 `xml:"itemsValue"`     //（可选）总费用
	Items Items `xml:"items"` //<!--商品信息-->
	//insuranceValue float64 `xml:"insuranceValue"` //保价金额（可选）
	Special int `xml:"special"` //商品类型（保留字段，暂时不用，默认填0）
	//remark  string `xml:"remark"`  //备注（可选）
}

//<!—发货方信息-->
type Sender struct {
	Name     string `xml:"name"`     //用户姓名
	PostCode string `xml:"postCode"` //用户邮编（如果没有可以填默认的0
	//Phone    string `xml:"phone"`    //用户电话，包括区号、电话号码及分机号，中间用“-”分隔；（可选）
	Mobile  string `xml:"mobile"`  //用户移动电话（可选）
	Prov    string `xml:"prov"`    // 用户所在省
	City    string `xml:"city"`    //用户所在市、县（区），市和区中间用“,”分隔；注意有些市下面是没有区
	Address string `xml:"address"` //用户详细地址
}

//<!--收货方信息-->
type Receiver struct {
	Name     string `xml:"name"`     //用户姓名
	PostCode string `xml:"postCode"` //用户邮编（如果没有可以填默认的0
	//Phone    string `xml:"phone"`    //用户电话，包括区号、电话号码及分机号，中间用“-”分隔；（可选）
	Mobile  string `xml:"mobile"`  //用户移动电话（可选）
	Prov    string `xml:"prov"`    // 用户所在省
	City    string `xml:"city"`    //用户所在市、县（区），市和区中间用“,”分隔；注意有些市下面是没有区
	Address string `xml:"address"` //用户详细地址
}

//<!--商品信息-->
type Item struct {
	ItemName string `xml:"itemName"` //商品名称（可填默认的0）
	Number   int    `xml:"number"`   //商品数量（可填默认的0）
	//ItemValue string //商品单价（两位小数）
}

//<!--商品信息-->
type Items struct {
	Item []Item `xml:"item"`
}

const (
	Url = "http://opentestapi.yto.net.cn/service/order_create/v1/D2nvLq"
)

func Create_application() {
	RequestOrder := RequestOrder{}
	RequestOrder.ClientID = "K21000119"
	RequestOrder.CustomerId = "K21000119"
	RequestOrder.LogisticProviderID = "YTO"
	RequestOrder.TxLogisticID = "LP07082300225709"
	RequestOrder.OrderType = 1
	RequestOrder.ServiceType = 0
	RequestOrder.Sender.Name = "柴雪新"
	RequestOrder.Sender.PostCode = "0"
	RequestOrder.Sender.Mobile = "17802928284"
	RequestOrder.Sender.Prov = "陕西省"
	RequestOrder.Sender.City = "西安市,雁塔区"
	RequestOrder.Sender.Address = "唐延路1号旺座国际城E做25楼"
	RequestOrder.Receiver.Name = "柴"
	RequestOrder.Receiver.PostCode = "0"
	RequestOrder.Receiver.Mobile = "17802928282"
	RequestOrder.Receiver.Prov = "陕西省"
	RequestOrder.Receiver.City = "西安市,未央区"
	RequestOrder.Receiver.Address = "唐延路1号旺座国际城E做25楼"
	RequestOrder.Special = 0
	RequestOrder.Items.Item = make([]Item, 1) //初始化
	RequestOrder.Items.Item[0].ItemName = "盛世中华"
	RequestOrder.Items.Item[0].Number = 1

	date, _ := xml.Marshal(RequestOrder)

	_ = url.QueryEscape(string(date))

	dataStr := string(date) + "u2Z1F7Fh"

	md5Key := md5.New()
	md5Key.Write([]byte(dataStr))

	//xmlKey := url.QueryEscape(base64.StdEncoding.EncodeToString(md5Key.Sum(nil)[:]))
	xmlKey := base64.StdEncoding.EncodeToString(md5Key.Sum(nil)[:])

	requestParameters := url.Values{}
	requestParameters.Set("logistics_interface", string(date))
	requestParameters.Add("data_digest", xmlKey)
	requestParameters.Add("clientId", "K21000119")
	requestParameters.Add("type", "online")

	resp, err := http.PostForm(Url, requestParameters)

	//req, err := http.NewRequest("POST", url2, strings.NewReader(requestParameters.Encode()))

	if err != nil {
		fmt.Println(err)
	}

	//伪装头部
	//req.Header.Set("Accept", "application/json, text/javascript, */*; q=0.01")
	//req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	//req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.6,en;q=0.4")
	//req.Header.Add("Connection", "keep-alive")
	//req.Header.Add("Content-Length", "25")
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	//req.Header.Add("Cookie", "user_trace_token=20170425200852-dfbddc2c21fd492caac33936c08aef7e; LGUID=20170425200852-f2e56fe3-29af-11e7-b359-5254005c3644; showExpriedIndex=1; showExpriedCompanyHome=1; showExpriedMyPublish=1; hasDeliver=22; index_location_city=%E5%85%A8%E5%9B%BD; JSESSIONID=CEB4F9FAD55FDA93B8B43DC64F6D3DB8; TG-TRACK-CODE=search_code; SEARCH_ID=b642e683bb424e7f8622b0c6a17ffeeb; Hm_lvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1493122129,1493380366; Hm_lpvt_4233e74dff0ae5bd0a3d81c6ccf756e6=1493383810; _ga=GA1.2.1167865619.1493122129; LGSID=20170428195247-32c086bf-2c09-11e7-871f-525400f775ce; LGRID=20170428205011-376bf3ce-2c11-11e7-8724-525400f775ce; _putrc=AFBE3C2EAEBB8730")
	//req.Header.Add("Host", "www.lagou.com")
	//req.Header.Add("Origin", "https://www.lagou.com")
	//req.Header.Add("Referer", "https://www.lagou.com/jobs/list_python?labelWords=&fromSearch=true&suginput=")
	//req.Header.Add("X-Anit-Forge-Code", "0")
	//req.Header.Add("X-Anit-Forge-Token", "None")
	//req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	//req.Header.Add("X-Requested-With", "XMLHttpRequest")

	//client := &http.Client{}

	//resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}
	//读取返回值
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(result))
}

func main() {

	//requestParameters := url.Values{}
	//requestParameters.Set("logistics_interface", "%3CRequestOrder%3E%3CagencyFund%3E0.0%3C%2FagencyFund%3E%3CclientID%3ETEST%3C%2FclientID%3E%3CcodSplitFee%3E1.0%3C%2FcodSplitFee%3E%3CcustomerId%3ETEST%3C%2FcustomerId%3E%3Cflag%3E1%3C%2Fflag%3E%3CgoodsValue%3E1.0%3C%2FgoodsValue%3E%3CinsuranceValue%3E1.0%3C%2FinsuranceValue%3E%3Citems%3E%3Citem%3E%3CitemName%3E%E5%95%86%E5%93%81%3C%2FitemName%3E%3CitemValue%3E0.0%3C%2FitemValue%3E%3Cnumber%3E2%3C%2Fnumber%3E%3C%2Fitem%3E%3C%2Fitems%3E%3ClogisticProviderID%3EYTO%3C%2FlogisticProviderID%3E%3CorderType%3E1%3C%2ForderType%3E%3Creceiver%3E%3Cname%3E%E6%94%B6%E4%BB%B6%E4%BA%BA%E5%A7%93%E5%90%8D%3C%2Fname%3E%3CpostCode%3E0%3C%2FpostCode%3E%3Cphone%3E021-12345678%3C%2Fphone%3E%3Cmobile%3E18112345678%3C%2Fmobile%3E%3Cprov%3E%E4%B8%8A%E6%B5%B7%E5%B8%82%3C%2Fprov%3E%3Ccity%3E%E4%B8%8A%E6%B5%B7%E5%B8%82%2C%E9%9D%92%E6%B5%A6%E5%8C%BA%3C%2Fcity%3E%3Caddress%3E%E5%8D%8E%E5%BE%90%E5%85%AC%E8%B7%AF3029%E5%BC%8428%E5%8F%B7%3C%2Faddress%3E%3C%2Freceiver%3E%3CsendEndTime%3E2019-12-30+12%3A12%3A33%3C%2FsendEndTime%3E%3CsendStartTime%3E2019-12-30+12%3A12%3A33%3C%2FsendStartTime%3E%3Csender%3E%3Cname%3E%E5%8F%91%E4%BB%B6%E4%BA%BA%E5%A7%93%E5%90%8D%3C%2Fname%3E%3CpostCode%3E526238%3C%2FpostCode%3E%3Cphone%3E021-12345678%3C%2Fphone%3E%3Cmobile%3E18112345678%3C%2Fmobile%3E%3Cprov%3E%E4%B8%8A%E6%B5%B7%E5%B8%82%3C%2Fprov%3E%3Ccity%3E%E4%B8%8A%E6%B5%B7%E5%B8%82%2C%E9%9D%92%E6%B5%A6%E5%8C%BA%3C%2Fcity%3E%3Caddress%3E%E5%8D%8E%E5%BE%90%E5%85%AC%E8%B7%AF3029%E5%BC%8428%E5%8F%B7%3C%2Faddress%3E%3C%2Fsender%3E%3CserviceType%3E1%3C%2FserviceType%3E%3Cspecial%3E1%3C%2Fspecial%3E%3CtotalServiceFee%3E1.0%3C%2FtotalServiceFee%3E%3CtxLogisticID%3ELP07082300225709%3C%2FtxLogisticID%3E%3C%2FRequestOrder%3E")
	//requestParameters.Set("data_digest", "pPoLaCHWASHGiD68X84duA%3D%3D")
	//requestParameters.Set("clientId", "TEST")
	//requestParameters.Set("type", "online")
	//resp, err := http.PostForm("http://opentestapi.yto.net.cn/service/order_create/v1/D2nvLq", requestParameters)
	//if err != nil {
	//	return
	//}
	//defer resp.Body.Close()
	//data, err := ioutil.ReadAll(resp.Body)
	//
	//fmt.Println(string(data))
	Create_application()

}
