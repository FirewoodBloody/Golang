package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"time"
)

var url string

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	Win := vcl.Application.CreateForm() //新建窗口
	Win.SetCaption("Update_Install")    //程序名
	//w.Win.SetFormStyle(2)
	Win.SetHeight(30)      //高
	Win.SetWidth(400)      //宽
	Win.ScreenCenter()     //居于当前屏幕中心
	Win.SetBorderIcons(0)  //设置窗口最大化 最小化 关闭按钮  3代表是 最大化按钮和功能无效
	Win.Font().SetSize(11) //整体字体大小
	Win.Font().SetColor(255)
	Win.Font().SetStyle(16) //字体样式
	Win.SetColor(16775388)
	Win.SetEnabled(false)

	TPanel := vcl.NewPanel(Win)
	TPanel.SetAlign(types.AlClient)

	Progress_bar := vcl.NewProgressBar(TPanel)
	Progress_bar.SetParent(Win)
	Progress_bar.SetPosition(0)
	Progress_bar.SetWidth(400)
	Progress_bar.SetHeight(30)
	Win.Show()

	if strconv.IntSize == 64 {
		url = "http://192.168.0.12:8888/file/Download"
	} else if strconv.IntSize == 32 {
		url = "http://192.168.0.12:8888/file/Download"
	}

	exe, _ := http.Get(url)
	defer exe.Body.Close()
	Progress_bar.SetPosition(25)
	time.Sleep(time.Second)
	Progress_bar.SetPosition(55)
	time.Sleep(time.Second)
	Progress_bar.SetPosition(75)
	time.Sleep(time.Second)
	Progress_bar.SetPosition(100)
	data, _ := ioutil.ReadAll(exe.Body)

	file, _ := os.Create("D:\\Program Files (x86)\\数据报表查询\\install.exe")

	file.Write(data)
	file.Close()
	defer os.Exit(0)

	Win.Hide()
	cmd := exec.Command("D:\\Program Files (x86)\\数据报表查询\\install.exe")
	cmd.Start()
	os.Exit(0)
}
