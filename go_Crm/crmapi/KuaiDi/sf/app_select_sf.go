package sf

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/url"
	"time"
)

const (

	//客户校验码    使用顺丰分配的客户校验码
	checkWord = "v1v0WHMrcgbKwyactWkgocsx1rNtt5iK"
)

func Md5(msgData string) url.Values {
	//业务报文  去报文中的msgData（业务数据报文）
	//msgData := "{\"language\":\"zh-CN\",\"orderId\":\"QIAO-20200618-004\"}"
	timestamp := time.Now().Unix()
	//将业务报文+时间戳+校验码组合成需加密的字符串(注意顺序)
	toVerifyText := msgData + fmt.Sprintf("%v", timestamp) + checkWord

	//因业务报文中可能包含加号、空格等特殊字符，需要urlEnCode处理
	//toVerifyText := url.eURLEncoder.encode(toVerifyText,"UTF-8")

	escapeUrl := url.QueryEscape(toVerifyText)

	//进行Md5加密
	md5Key := md5.New()
	md5Key.Write([]byte(escapeUrl))
	//通过BASE64生成数字签名
	msgDigest := base64.StdEncoding.EncodeToString(md5Key.Sum(nil))

	//配置请求参数,
	requestParameters := url.Values{}
	requestParameters.Set("Content-type", "application/x-www-form-urlencoded;charset=UTF-8")
	requestParameters.Set("partnerID", "BLWHYJb9V_KD100")
	requestParameters.Set("requestID", "00017E4C3D81543FA1F0480564E0F03F")
	requestParameters.Set("serviceCode", "EXP_RECE_SEARCH_ROUTES")
	requestParameters.Set("timestamp", fmt.Sprintf("%v", timestamp))
	requestParameters.Set("msgDigest", msgDigest)
	//requestParameters.Set("accessToken", "JhpuwRiK6Mts6GtBLJbfkNjAO5dU2pch")
	requestParameters.Set("msgData", msgData)

	return requestParameters
}
