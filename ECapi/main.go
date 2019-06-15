package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Access struct {
	AppId     int    `json:"appId"`
	AppSecret string `json:"appSecret"`
}

type AccessToken struct {
	ErrCode int    `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
	Data    struct {
		AccessToken string `json:"accessToken"`
		ExpiresIn   int    `json:"expiresIn"`
	}
}

const (
	appId     = 489030333769973760
	appSecret = "AOWiiXXD1EWr0tjYADi"
	accessUrl = "https://open.workec.com/auth/accesstoken"
)

func main() {

	access := Access{
		AppId:     appId,
		AppSecret: appSecret,
	}

	data, err := json.Marshal(access)
	if err != nil {
		fmt.Println(err)
	}
	client := &http.Client{}
	rep, err := http.NewRequest("POST", accessUrl, bytes.NewReader(data))
	if err != nil {
		fmt.Println(err)
	}
	rep.Header.Add("cache-control", "no-cache")
	rep.Header.Add("content-type", "application/json")
	resp, err := client.Do(rep)

	defer resp.Body.Close()
	dataa, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(dataa))
	accessToken := &AccessToken{}

	err = json.Unmarshal(dataa, accessToken)

	fmt.Println(accessToken)

}
