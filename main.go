package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {

	bady := url.Values{}
	bady.Add("Customer_number", "80")

	resp, err := http.PostForm("http://2shoucang.f3322.net:8888/v1/client/Accles", bady)

	fmt.Println(err)

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	fmt.Println(err)

	fmt.Println(1, string(data))

	time.Sleep(time.Second * 20)

}
