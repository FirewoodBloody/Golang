//京东快递下单文档

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

const (
	Url1          = "https://api.jd.com/routerjson"
	JdLoginId1    = "029K708772"
	App_key1      = "5BA6F95488F2BA2655367595505F7057" //应用标识
	App_secret1   = "0053e1814a6345a19d7e06009281d5e9" //应用密钥
	access_token1 = "c5ec850f6c2a4c7288b51d7910df3673yzg1"
	method1       = "jingdong.ldop.waybill.receive"
)

//用户验证信息
type PublicParameters struct {
	Method       string `json:"method"`       //API接口名称
	Access_token string `json:"access_token"` //采用OAuth授权方式是必填参数
	App_key      string `json:"app_key"`      //应用的app_key
	//签名算法
	//将所有请求参数按照字母先后顺序排列 access_token,app_key,method,timestamp,v,360buy_param_json ，
	//排序为360buy_param_json,access_token,app_key,method,timestamp,v
	//把所有参数名和参数值进行拼接，例如：360buy_param_jsonxxxaccess_tokenxxxapp_keyxxxmethodxxxxxxtimestampxxxxxxvx
	//把appSecret夹在字符串（上一步拼接串）的两端，例如：appSecret+XXXX+appSecret
	//使用MD5进行加密，再转化成大写
	Sign      string `json:"sign"`      //签名算法：详见“签名算法”描述
	Timestamp string `json:"timestamp"` //时间戳，格式为yyyy-MM-dd HH:mm:ss，例如：2019-05-01 00:00:00。京东API服务端允许客户端请求时间误差为10分钟
	//Format         string `json:"format"`    //暂时只支持json
	V              string `json:"v"` //
	Buy_param_json string `json:"360buy_param_json"`
}

//请求快递信息 data数据
type Format struct {
	SalePlat     string `json:"salePlat"`     //销售平台（非 JD 商城请填： 0030001）
	CustomerCode string `json:"customerCode"` // 商家编码
	OrderId      string `json:"orderId"`      //订单号

	//ThrOrderId   string `json:"thrOrderId"`
	//销售平台订单号(例如京东订单号或天猫订单号等等。总长度不要超过100。如果有多个单号，用英文,间隔。例如：7898675,7898676)

	SenderName     string `json:"senderName"`     //寄件人姓名，说明：不能为生僻字
	SenderAddress  string `json:"senderAddress"`  //寄件人地址，说明：不能为生僻字
	SenderTel      string `json:"senderTel"`      //寄件人电话
	SenderMobile   string `json:"senderMobile"`   //寄件人手机(寄件人电话、手机至少有一个不为空)
	SenderPostcode string `json:"senderPostcode"` //寄件人邮编，长度：6位
	ReceiveName    string `json:"receiveName"`    //收件人名称，说明：不能为生僻字
	ReceiveAddress string `json:"receiveAddress"` //收件人地址，说明：不能为生僻字
	SenderCompany  string `json:"senderCompany"`  //寄件人公司，长度：50，说明：不能为生僻字
	ReceiveCompany string `json:"receiveCompany"` //收件人公司，长度：50，说明：不能为生僻字

	//非必填项
	Province   string `json:"province"`   //	收件人省
	City       string `json:"city"`       //	收件人市
	County     string `json:"county"`     //	收件人县
	Town       string `json:"town"`       //收件人镇
	ProvinceId int    `json:"provinceId"` //	收件人省编码
	CityId     int    `json:"cityId"`     //收件人市编码
	CountyId   int    `json:"countyId"`   //	收件人县编码
	TownId     int    `json:"townId"`     //收件人镇编码
	SiteType   int    `json:"siteType"`   //站点类型
	SiteId     int    `json:"siteId"`     //	站点编码
	SiteName   string `json:"siteName"`   //	站点名称
	//非必填项

	ReceiveTel    string `json:"receiveTel"`    //收件人电话
	ReceiveMobile string `json:"receiveMobile"` //	收件人手机号(收件人电话、手机至少有一个不为空)
	//Postcode      string `json:"postcode"` //收件人邮编，长度：6

	//包裹信息
	PackageCount int `json:"packageCount"` //包裹数(大于0，小于1000)
	Weight       int `json:"weight"`       //重量(单位：kg，保留小数点后两位)

	//快件体积
	//非必填
	//VloumLong   int `json:"vloumLong"`   //包裹长(单位：cm,保留小数点后两位)
	//VloumWidth  int `json:"vloumWidth"`  //包裹宽(单位：cm，保留小数点后两位)
	//VloumHeight int `json:"vloumHeight"` //包裹高(单位：cm，保留小数点后两位)
	//非必填
	Vloumn int `json:"vloumn"` //体积(单位：cm3，保留小数点后两位)

	//商品信息
	Description string `json:"description"` //商品描述、
	Goods       string `json:"goods"`       //托寄物名称，长度：200，说明：为避免托运物品在铁路、航空安检中被扣押，请务必下传商品类目或名称，例如：服装、3C等
	GoodsCount  int    `json:"goodsCount"`  //寄托物数量

	//快递信息
	CollectionValue      int    `json:"collectionValue"`      //是否代收货款(是：1，否：0。不填或者超出范围，默认是0)
	CollectionMoney      int    `json:"collectionMoney"`      //代收货款金额(保留小数点后两位)
	GuaranteeValue       int    `json:"guaranteeValue"`       //是否保价(是：1，否：0。不填或者超出范围，默认是0)
	GuaranteeValueAmount int    `json:"guaranteeValueAmount"` //保价金额(保留小数点后两位)
	SignReturn           int    `json:"signReturn"`           //签单返还(签单返还类型：0不返单，1普通返单，2校验身份返单，3电子签返单，4电子返单+普通返单)
	Aging                int    `json:"aging"`                //时效(普通：1，工作日：2，非工作日：3，晚间：4。O2O一小时达：5。O2O定时达：6。不填或者超出范围，默认是1)
	TransType            int    `json:"transType"`            //运输类型(陆运：1，航空：2。不填或者超出范围，默认是1)
	Remark               string `json:"remark"`               //运单备注，长度：20,说明：打印面单时备注内容也会显示在快递面单上

	//非必填项
	//GoodsType         int       `json:"goodsType"`         //配送业务类型（ 1:普通，3:填仓，4:特配，6:控温，7:冷藏，8:冷冻，9:深冷）默认是1
	//OrderType         int       `json:"orderType"`         //运单类型。(普通外单：0，O2O外单：1)默认为0
	//ShopCode          string    `json:"shopCode"`          //门店编码(只O2O运单需要传，普通运单不需要传)
	//OrderSendTime     string    `json:"orderSendTime"`     //预约配送时间（格式：yyyyMMdd HH:mm:ss）
	//WarehouseCode     string    `json:"warehouseCode"`     //发货仓编码
	//AreaProvId        int       `json:"areaProvId"`        //接货省ID
	//AreaCityId        int       `json:"areaCityId"`        //接货市ID
	//ShipmentStartTime time.Time `json:"shipmentStartTime"` //配送起始时间
	//ShipmentEndTime   time.Time `json:"shipmentEndTime"`   //配送结束时间
	//Idint             string    `json:"idint"`             //身份证号
	//AddedService      string    `json:"addedService"`      //扩展服务
	//ExtendField1      string    `json:"extendField1"`      //扩展字段1
	//ExtendField2      string    `json:"extendField2"`      //扩展字段2
	//ExtendField3      string    `json:"extendField3"`      //扩展字段3
	//ExtendField4      int       `json:"extendField4"`      //扩展字段4
	//ExtendField5      int       `json:"extendField5"`      //扩展字段5
	//
	//PromiseTimeType     int       `json:"promiseTimeType"`     //产品类型（1：特惠送 2：特快送 4：城际闪送 5：同城当日达 6：次晨达 7：微小件 8: 生鲜专送 16：生鲜特快 17、生鲜特惠 21：特惠小包）
	//Freight             int       `json:"freight"`             //运费
	//PickUpStartTime     time.Time `json:"pickUpStartTime"`     //预约取件开始时间
	//PickUpEndTime       time.Time `json:"pickUpEndTime"`       //预约取件结束时间
	//UnpackingInspection string    `json:"unpackingInspection"` //开箱验货标识
	//BoxCode             []string  `json:"boxCode"`             //商家箱号,多个箱号请用逗号分隔，例如三个箱号为：a123,b456,c789
	//FileUrl             string    `json:"fileUrl"`             //文件url
}

//响应参数
type Parameters struct {
	Jingdong_ldop_waybill_receive_responce Jingdong_ldop_waybill_receive_responce `json:"jingdong_ldop_waybill_receive_responce"`
}

//响应参数
type Jingdong_ldop_waybill_receive_responce struct {
	Code                    string                  `json:"code"`
	ReceiveorderInfo_Result ReceiveorderInfo_Result `json:"receiveorderinfo_result"`
}

//响应参数
type ReceiveorderInfo_Result struct {
	ResultCode      int           `json:"resultCode"`      //结果编码
	ResultMessage   string        `json:"resultMessage"`   //结果描述
	OrderId         string        `json:"orderId"`         //订单号
	DeliveryId      string        `json:"deliveryId"`      //运单号
	PromiseTimeType int           `json:"promiseTimeType"` //产品类型（1：特惠送 2：特快送 4：城际闪送 5：同城当日达 6：次晨达 7：微小件 8: 生鲜专送 16：生鲜特快 17:生鲜特惠）
	PreSortResult   PreSortResult `json:"preSortResult"`   //预分拣结果：面单打印信息
	TransType       int           `json:"transType"`       //运输类型（1：陆运 2：航空）
	NeedRetry       bool          `json:"needRetry"`       //是否需要重试

}

//面单打印信息
type PreSortResult struct {
	SiteId                 int    `json:"siteId"`                 //站点id
	SiteName               string `json:"siteName"`               //站点名称
	Road                   string `json:"road"`                   //路区
	SlideNo                string `json:"slideNo"`                //滑道号
	SourceSortCenterId     int    `json:"sourceSortCenterId"`     //始发分拣中心id
	SourceSortCenterName   string `json:"sourceSortCenterName"`   //始发分拣中心名称
	SourceCrossCode        string `json:"sourceCrossCode"`        //始发道口号
	SourceTabletrolleyCode string `json:"sourceTabletrolleyCode"` //始发笼车号
	TargetSortCenterId     int    `json:"targetSortCenterId"`     //目的分拣中心id
	TargetSortCenterName   string `json:"targetSortCenterName"`   //目的分拣中心名称
	TargetTabletrolleyCode string `json:"targetTabletrolleyCode"` //目的笼车号
	Aging                  int    `json:"aging"`                  //时效
	AgingName              string `json:"agingName"`              //时效名称
	SiteType               int    `json:"siteType"`               //站点类型
	IsHideName             int    `json:"isHideName"`             //是否隐藏姓名
	IsHideContractints     int    `json:"isHideContractints"`     //是否隐藏联系方式

}

func (p *PublicParameters) GetOrders() *Parameters {
	str := fmt.Sprintf("%v360buy_param_json%vaccess_token%vapp_key%vmethod%vtimestamp%vv%v%v", App_secret1, p.Buy_param_json, p.Access_token, p.App_key, p.Method, p.Timestamp, p.V, App_secret1)
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

	//urlstr := fmt.Sprintf("%v?v=%v&method=%v&app_key=%v&access_token=%v&360buy_param_json=%v&timestamp=%v&sign=%v", Url1, p.V, p.Method, p.App_key, p.Access_token,	p.buy_param_json, p.Timestamp, p.Sign)
	//HTTP POST 请求方式
	resp, err := http.PostForm(Url1, requestParameters)

	if err != nil {
		fmt.Println("1:", err)
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("2:", err)
	}
	fmt.Println(string(data))
	dataStruct := &Parameters{}

	err = json.Unmarshal(data, dataStruct) //返回数据新兴JSON反序列化
	if err != nil {
		fmt.Println("3:", err)
	}

	return dataStruct
}

func FormatSet(f *Format) ([]byte, error) {
	f.SalePlat = "0030001"
	f.CustomerCode = JdLoginId1
	f.OrderId = "951753856651"
	f.SenderName = "柴雪新"
	f.SenderAddress = "陕西省西安市雁塔区唐延路旺座国际城E座"
	f.SenderMobile = "17802928284"
	f.ReceiveName = "柴雪新"
	f.ReceiveAddress = "陕西省西安市雁塔区唐延路旺座国际城E座"
	f.ReceiveMobile = "17802928284"
	f.PackageCount = 1
	f.Weight = 1
	f.Vloumn = 1
	f.Goods = "上线测试，请取消"
	f.GoodsCount = 1

	return json.Marshal(f) //JSO序列化
}

//公共参数，用户授权
func (p *PublicParameters) SetUserLogin() *Parameters {
	p.App_key = App_key1
	p.V = "2.0" //版本
	p.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	p.Access_token = access_token1
	p.Method = method1

	f := new(Format)
	data, err := FormatSet(f)
	if err != nil {
		fmt.Println("4:", err)
	}
	fmt.Println(string(data))
	p.Buy_param_json = string(data)

	return p.GetOrders()

}
