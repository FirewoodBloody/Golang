package main

import (
	"fmt"
	"net/http"
	"net/url"
)

type SFPrintf struct {
	//快递面单基础字段信息
	MailNo           string `json:"mailNo"`           //String	是	运单号	mailno
	OrderNo          string `json:"orderNo"`          //String	是	订单号	orderid
	ExpressType      int    `json:"expressType"`      //int	是	快递产品类型 1-标准快递 2-顺丰标快（陆运） 5-顺丰次晨 6-顺丰即日	express_type
	PayMethod        int    `json:"payMethod"`        //int	是	付款方式 1-寄付 2-到付 3-第三方付 1	pay_method
	PayArea          string `json:"payArea"`          //String	是	付款地区 -如第三方支付则必填
	ReturnTrackingNo string `json:"returnTrackingNo"` //String	是	签回单号	return_tracking_no
	ZipCode          string `json:"zipCode"`          //String	是	原寄地代码	origincode
	DestCode         string `json:"destCode"`         //String	是	收件地代码	destcode
	MonthAccount     string `json:"monthAccount"`     //String	是	月结卡号	custid

	//寄件人信息
	DeliverName        string `json:"deliverName"`        //String	是	寄件人姓名	j_contact
	DeliverTel         string `json:"deliverTel"`         //String	是	寄件人电话	j_tel
	DeliverMobile      string `json:"deliverMobile"`      //String	是	寄件人手机	j_mobile
	DeliverCompany     string `json:"deliverCompany"`     //String	是	寄件人公司	j_company
	DeliverProvince    string `json:"deliverProvince"`    //String	是	寄件人省	j_province
	DeliverCounty      string `json:"deliverCounty"`      //String	是	寄件人区	j_county
	DeliverCity        string `json:"deliverCity"`        //String	是	寄件人市	j_city
	DeliverAddress     string `json:"deliverAddress"`     //String	是	寄件人详细地址	j_address
	DeliverShipperCode string `json:"deliverShipperCode"` //String	是	寄件人邮政编码

	//收件人信息

	ConsignerName        string `json:"consignerName"`        //String	是	收件人姓名	d_contact
	ConsignerMobile      string `json:"consignerMobile"`      //String	是	收件人手机	d_mobile
	ConsignerTel         string `json:"consignerTel"`         //String	是	收件人电话	d_tel
	ConsignerCompany     string `json:"consignerCompany"`     //String	是	收件人公司	d_company
	ConsignerProvince    string `json:"consignerProvince"`    //String	是	收件人省	d_province
	ConsignerCity        string `json:"consignerCity"`        //String	是	收件人市	d_city
	ConsignerCounty      string `json:"consignerCounty"`      //String	是	收件人区	d_county
	ConsignerAddress     string `json:"consignerAddress"`     //String	是	收件人详细地址	d_address
	ConsignerShipperCode string `json:"consignerShipperCode"` //String	是	收件人邮政编码

	//logo信息
	Logo         string `json:"logo"`         //String	否
	CustLogo     string `json:"custLogo"`     //String	否
	SftelLogo    string `json:"sftelLogo"`    //String	否
	TopLogo      string `json:"topLogo"`      //String	否
	TopsftelLogo string `json:"topsftelLogo"` //String	否

	//appId
	AppId  string `json:"appId"`  //String	是	顾客编码
	AppKey string `json:"appKey"` //String	是	校验码

	//陆运件标识
	Electric string `json:"electric"` //String	否	E标识（陆运标识）电商特惠、顺丰特惠、电商专配、陆运件 必须打印E标识

	//增值服务
	CodValue        string `json:"codValue"`        //String	否	代收货款金额(元)	COD
	CodMonthAccount string `json:"codMonthAccount"` //String	否	代收货款卡号
	InsureFee       string `json:"insureFee"`       //String	否	保价费用(元)	SINSURE
	InsureValue     string `json:"insureValue"`     //String	否	声明价值(元)	insure

	//加密
	EncryptCustName bool `json:"encryptCustName"` //	String	是	加密寄件及收件联系人 默认不加密
	EncryptMobile   bool `json:"encryptMobile"`   //	String	是	加密电话 默认不加密

	//备注
	ChildRemark          string             `json:"childRemark"`          //String	否	子运单备注信息
	MainRemark           string             `json:"mainRemark"`           //String	否	主运单备注信息
	ReturnTrackingRemark string             `json:"returnTrackingRemark"` //String	否	签回单回单的备注信息
	RlsInfoDtoList       []RlsInfoDtoList   `json:"rlsInfoDtoList"`       // 	List	是	丰密信息	OrderResponse/ rls_info/ rls_detail
	CargoInfoDtoList     []CargoInfoDtoList `json:"cargoInfoDtoList"`     //List	是	托寄物信息

	//费用
	TotalFee string `json:"totalFee"` //String	是	费用合计
}

type RlsInfoDtoList struct {
	QRCode             string `json:"QRCode"`             //String	是	二维码	twoDimensionCode，		自己封装二维码里面		k1,k2,k3,k4,k5,k6,k7,		对应未返回的则留空
	AbFlag             string `json:"abFlag"`             //String	否	AB标	abFlag，如未返回或该值为空，则留空
	CodingMapping      string `json:"codingMapping"`      //String	否	入港映射码	codingMapping,		如未返回或该值为空，则留空
	CodingMappingOut   string `json:"codingMappingOut"`   //String	否	出港映射码	codingMappingOut,		如未返回或该值为空，则留空
	DestRouteLabel     string `json:"destRouteLabel"`     //String	是	目的地路由标签	destRouteLabel，如未返回		或该值为空		则打印目的地城市代码
	DestTeamCode       string `json:"destTeamCode"`       //String	否	目的地单元区域[需要水印的标]	destTeamCode,如未返回		或该值为空，则留空
	PrintIcon          string `json:"printIcon"`          //String	是	打印图标 1重2蟹3鲜4碎	printIcon,如未返回或该值为空，则填‘00000000’
	ProCode            string `json:"proCode"`            //String	是	产品代码	proCode，如未返回或该值为空		按照快件产品类别表中产品代码打印
	SourceTransferCode string `json:"sourceTransferCode"` //String	否	原寄地中转场代码	sourceTransferCode,		如未返回或该值为空，则留空
	WaybillNo          string `json:"waybillNo"`          //String	是	运单号	mailno
	XbFlag             string `json:"xbFlag"`             //String	否	XB标	xbFlag，如未返回或该值为空，则留空
}

type CargoInfoDtoList struct {
	Cargo       string `json:"cargo"`       //String	是	托寄物名称
	CargoCount  int    `json:"cargoCount"`  //int	是	数量
	CargoUnit   string `json:"cargoUnit"`   //String	否	单位
	CargoWeight string `json:"cargoWeight"` //BigDecimal	否	实际重量，单位KG
	Remark      string `json:"remark"`      //remark	String	否	备注	remark
}

const (
	UserId  = "BLWHYSP_tipea" //顾客编码
	UserKey = "eJXycwXzucwcxit4a8WK7b3qGl4UfkB1"

	/*********2联150 丰密面单**************/
	/**调用打印机 不弹出窗口 适用于批量打印【二联单】 **/
	url7 = "http://localhost:4040/sf/waybill/print?type=V2.0.FM_poster_100mm150mm&output=noAlertPrint"

	/** 调用打印机 弹出窗口 可选择份数 适用于单张打印【二联单】**/
	url8 = "http://localhost:4040/sf/waybill/print?type=V2.0.FM_poster_100mm150mm&output=print"

	/**直接输出图片的BASE64编码字符串 可以使用html标签直接转换成图片【二联单】**/
	url9 = "http://localhost:4040/sf/waybill/print?type=V2.0.FM_poster_100mm150mm&output=image"

	/*********3联210 丰密面单**************/
	/**调用打印机 不弹出窗口 适用于批量打印【三联单】**/
	url10 = "http://localhost:4040/sf/waybill/print?type=V3.0.FM_poster_100mm210mm&output=noAlertPrint"

	/**调用打印机 弹出窗口 可选择份数 适用于单张打印【三联单】**/
	url11 = "http://localhost:4040/sf/waybill/print?type=V3.0.FM_poster_100mm210mm&output=print"

	/** 直接输出图片的BASE64编码字符串 可以使用html标签直接转换成图片【三联单】**/
	url12 = "http://localhost:4040/sf/waybill/print?type=V3.0.FM_poster_100mm210mm&output=image"

	/*********2联180 丰密运单**************/
	/**
	 * 调用打印机 不弹出窗口 适用于批量打印【二联单】
	 */
	url13 = "http://localhost:4040/sf/waybill/print?type=V2.8.0.FM_poster_100mm180mm&output=noAlertPrint"

	/**
	 * 调用打印机 弹出窗口 可选择份数 适用于单张打印【二联单】
	 */
	url14 = "http://localhost:4040/sf/waybill/print?type=V2.8.0.FM_poster_100mm180mm&output=print"

	/**
	 * 直接输出图片的BASE64编码字符串 可以使用html标签直接转换成图片【二联单】
	 */
	url15 = "http://localhost:4040/sf/waybill/print?type=V2.8.0.FM_poster_100mm180mm&output=image"
)

func main() {
	requrst := url.Values{}
	requrst.Set("mailNo", "SF7551234567890")
	requrst.Set("orderNo", "DD20200812000120")
	requrst.Set("expressType", "1")
	requrst.Set("payMethod", "1")
	requrst.Set("payArea", "")
	requrst.Set("returnTrackingNo", "")
	requrst.Set("zipCode", "571")
	requrst.Set("destCode", "755")
	requrst.Set("monthAccount", "7550385912")
	requrst.Set("deliverName", "艾丽斯")
	requrst.Set("deliverTel", "")
	requrst.Set("deliverMobile", "15881234567")
	requrst.Set("deliverCompany", "神罗科技集团有限公司")
	requrst.Set("deliverProvince", "浙江省")
	requrst.Set("deliverCounty", "拱墅区")
	requrst.Set("deliverCity", "杭州市")
	requrst.Set("deliverAddress", "舟山东路708号古墩路北（玉泉花园旁）百花苑西区7-2-201室")
	requrst.Set("deliverShipperCode", "")
	requrst.Set("consignerName", "风一样的旭哥")
	requrst.Set("consignerMobile", "15893799999")
	requrst.Set("consignerTel", "")
	requrst.Set("consignerCompany", "神一样的科技")
	requrst.Set("consignerProvince", "广东省")
	requrst.Set("consignerCity", "深圳市")
	requrst.Set("consignerCounty", "南山区")
	requrst.Set("consignerAddress", "学府路软件产业基地2B12楼5200708号")
	requrst.Set("consignerShipperCode", "")
	requrst.Set("logo", "")
	requrst.Set("custLogo", "")
	requrst.Set("sftelLogo", "")
	requrst.Set("topLogo", "")
	requrst.Set("topsftelLogo", "")
	requrst.Set("appId", "BLWHYSP_tipea")
	requrst.Set("appKey", "eJXycwXzucwcxit4a8WK7b3qGl4UfkB1")
	requrst.Set("electric", "")
	requrst.Set("codValue", "999")
	requrst.Set("codMonthAccount", "")
	requrst.Set("insureFee", "")
	requrst.Set("insureValue", "501")
	requrst.Set("encryptCustName", "1")
	requrst.Set("encryptMobile", "1")
	requrst.Set("childRemark", "子单号备注")
	requrst.Set("mainRemark", "这是主面单的备注")
	requrst.Set("returnTrackingRemark", "")
	requrst.Set("rlsInfoDtoList", "[{\"QRCode\":\"{'k1':'755WE','k2':'021WT','k3':'','k4':'T4','k5':'SF7551234567890','k6':'','k7':''}\",\"abFlag\":\"A\",\"codingMapping\":\"F33\",\"codingMappingOut\":\"1A\",\"destRouteLabel\":\"755WE-571A3\",\"destTeamCode\":\"012345678\",\"printIcon\":\"11110000\",\"proCode\":\"T4\",\"sourceTransferCode\":\"021WTF\",\"waybillNo\":\"SF7551234567890\",\"xbFlag\":\"XB\"}]")
	requrst.Set("cargoInfoDtoList", "[{\"cargo\":\"苹果7S\",\"cargoCount\":1,\"cargoUnit\":\"\",\"cargoWeight\":\"\",\"remark\":\"手机贵重物品 小心轻放\"}]")
	requrst.Set("totalFee", "")

	a, _ := http.PostForm(url8, requrst)
	fmt.Printf("%#v", a)
}
