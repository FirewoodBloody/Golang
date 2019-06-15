package models

//获取access_token POST 授权ID和秒
type Access struct {
	AppId     int    `json:"appId"`     //开发者ID
	AppSecret string `json:"appSecret"` //开发者秘钥
}

//接收授权响应
type AccessToken struct {
	ErrCode int      `json:"errCode"` //返回码。 200 代表成功，其他表示具体错误信息
	ErrMsg  string   `json:"errMsg"`  //响应信息。OK代表成功，其他表示具体错误信息
	Data    struct { //具体业务信息
		AccessToken string `json:"accessToken"` //调用业务接口时必须要传入的  授权值
		ExpiresIn   int    `json:"expiresIn"`   //授权值的剩余有效时间(秒)、在有效期内调用业务接口时会自动续期
	}
}
