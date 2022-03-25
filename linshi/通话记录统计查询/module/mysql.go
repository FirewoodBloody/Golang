package module

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ying32/govcl/vcl"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var (
	URL = "http://61.185.225.118:19374/v1/object/" //请求数据
	//URL  = "http://61.185.225.118:19374/v1/object/"            //请求数据
	Url  = "http://61.185.225.118:19374/v1/user/"              //数据报表类型
	Url2 = "http://61.185.225.118:19374/v1/user/Version"       //版本
	Url1 = "http://61.185.225.118:19374/v1/user/Client_Models" //版本 和 数据报表类型
)

//通话记录统计
func Call_log_statistics(w *Windows) {
	bady := url.Values{}
	bady.Add("Mysql_Select", "通话记录统计报表")
	bady.Add("Start_Time", time.Now().Format("2006-01-02"))
	bady.Add("Stop_Time", time.Now().Format("2006-01-02"))
	bady.Add("Client_Models", w.Client_Models)
	bady.Add("Login_Name", w.LoginName)

	resp, err := http.PostForm(URL, bady)
	if err != nil {
		vcl.ShowMessage(fmt.Sprintf("%v", err))
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)

	var resulist []map[string][]byte
	err = json.Unmarshal(data, &resulist)
	if err != nil {
		vcl.ShowMessage(fmt.Sprintf("%v", err))
	}

	for k, v := range resulist {
		vcl.ThreadSync(func() {
			items := w.listView.Items().Add()
			items.SetImageIndex(0)
			items.SetCaption(fmt.Sprintf("%3d", k+1))
			items.SubItems().Add(string(v["部门"]))
			items.SubItems().Add(string(v["二级部门"]))
			items.SubItems().Add(string(v["员工姓名"]))
			items.SubItems().Add(string(v["员工工号"]))
			items.SubItems().Add(string(v["未接通话总数"]))
			items.SubItems().Add(string(v["未接通时长"]))
			items.SubItems().Add(string(v["无效通话总数"]))
			items.SubItems().Add(string(v["无效通话时长"]))
			items.SubItems().Add(string(v["有效通话总数"]))
			items.SubItems().Add(string(v["有效通话时长"]))
			items.SubItems().Add(string(v["优质通话总数"]))
			items.SubItems().Add(string(v["优质通话时长"]))
			items.SubItems().Add(string(v["总通话数"]))
			items.SubItems().Add(string(v["通话总时长"]))
		})

	}

	w.Button.SetEnabled(true)
	w.Win.SetEnabled(true)
	//os.Exit(0)
}
