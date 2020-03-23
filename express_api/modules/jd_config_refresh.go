package modules

import (
	"encoding/json"
	"fmt"
	"github.com/Unknwon/goconfig"
	"io/ioutil"
	"net/http"
)

//刷新接收参数
type ConfigIni struct {
	Access_token  string `json:"access_token"`  // 接口调用令牌
	Expires_in    int    `json:"expires_in"`    //令牌有效时间, 单位秒
	Refresh_token string `json:"refresh_token"` //用户刷新access_token
	Scope         string `json:"scope"`         //用户授权的作用域，使用逗号（,）分隔
	Open_id       string `json:"open_id"`       //授权用户唯一标识
	Uid           string `json:"uid"`
	Time          int    `json:"time"` //时间戳
	Token_type    string `json:"token_type"`
	Code          int    `json:"code"` //状态码
}

const (
	app_key    = "5BA6F95488F2BA2655367595505F7057" //应用标识
	app_secret = "0053e1814a6345a19d7e06009281d5e9" //应用密钥
)

//读取配置文件信息
func OpenConfig(c *ConfigIni, filename string) {
	cfg, err := goconfig.LoadConfigFile(filename)
	if err != nil {
		fmt.Println("File Open failed:", err)
	}

	//从配置文件中读取所需要的授权口令和刷新授权的access_token
	c.Access_token, _ = cfg.GetValue("JD", "access_token")
	c.Refresh_token, _ = cfg.GetValue("JD", "refresh_token")
	c.Open_id, _ = cfg.GetValue("JD", "open_id")
}

//刷新授权时效
func RefreshKey(c *ConfigIni) {

	//http get 请求刷新，使用 KEY +秘钥 + grant_type + 刷新字符refresh_token
	resp, err := http.Get(fmt.Sprintf("https://open-oauth.jd.com/oauth2/refresh_token?app_key=%v&app_secret=%v&grant_type=refresh_token&refresh_token=%v", app_key, app_secret, c.Refresh_token))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	if err != nil {
		fmt.Println(err)
	}

	//反序列化Josn
	err = json.Unmarshal(data, c)
	if err != nil {
		fmt.Println(err)
	}

}
