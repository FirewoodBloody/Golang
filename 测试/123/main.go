package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {

	requestParameters := url.Values{}
	requestParameters.Set("order_id", "DD20200704002791")
	requestParameters.Set("id", "121012")
	requestParameters.Set("ExpressType", "1")
	//UrlA := "http://192.168.0.8:30001/api/admin/mallorder/expressBill&order_id=DD20200704002791&id=121012&ExpressType=1"
	//response, err := http.Get(UrlA)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//body, err := ioutil.ReadAll(response.Body)

	resp, err := http.PostForm("http://192.168.0.8:30001/api/admin/mallorder/expressBill", requestParameters)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))

	if err != nil {

		fmt.Println(err)
		return
	}

}
