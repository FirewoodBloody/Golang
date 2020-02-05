package main

import (
	"encoding/json"
	"fmt"
	"github.com/jlaffaye/ftp"
	"github.com/ying32/govcl/vcl"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type crm_dat101 struct {
	Filename string `json:"filename"`
	Shijian  string `json:"shijian"`
}

type Windows struct {
	win *vcl.TForm //窗口

	label_strat *vcl.TLabel //标签
	label_stop  *vcl.TLabel
	label_time  *vcl.TLabel

	time_tedit *vcl.TEdit

	date_strat_label *vcl.TDateTimePicker //开始日期菜单
	date_stop_label  *vcl.TDateTimePicker

	progress_bar *vcl.TProgressBar //进度条

	button *vcl.TButton
}

type file struct {
	Start string `json:"start"`
	Stop  string `json:"stop"`
	Times string `json:"times"`
}

func (w *Windows) init() {
	vcl.Application.Initialize()

	w.win = vcl.Application.CreateForm() //新建窗口
	w.win.SetName("录音下载")                //程序名
	//w.win.SetFormStyle(2)
	w.win.SetHeight(300)     //高
	w.win.SetWidth(400)      //宽
	w.win.ScreenCenter()     //居于当前屏幕中心
	w.win.SetBorderIcons(3)  //设置窗口最大化 最小化 关闭按钮  3代表是 最大化按钮和功能无效
	w.win.Font().SetSize(11) //整体字体大小
	w.win.Font().SetColor(255)
	w.win.Font().SetStyle(16) //字体样式
	w.win.SetColor(16775388)
	//w.win.SetTransparentColor(true)
	//w.win.SetTransparentColorValue(1)

	w.label_strat = vcl.NewLabel(w.win)
	w.label_strat.SetParent(w.win)
	w.label_strat.SetName("开始时间")
	w.label_strat.SetLeft(50) //设置按钮位置  横向
	w.label_strat.SetTop(50)  //设置按钮位置 竖向

	w.label_stop = vcl.NewLabel(w.win)
	w.label_stop.SetParent(w.win)
	w.label_stop.SetName("结束时间")
	w.label_stop.SetLeft(50)
	w.label_stop.SetTop(100)

	w.date_strat_label = vcl.NewDateTimePicker(w.win)
	w.date_strat_label.SetParent(w.win)
	w.date_strat_label.SetLeft(130)
	w.date_strat_label.SetTop(50)

	w.date_stop_label = vcl.NewDateTimePicker(w.win)
	w.date_stop_label.SetParent(w.win)
	w.date_stop_label.SetLeft(130)
	w.date_stop_label.SetTop(100)

	w.button = vcl.NewButton(w.win)
	w.button.SetParent(w.win)
	w.button.SetHeight(50)
	w.button.SetWidth(100)
	w.button.SetTop(150)
	w.button.SetLeft(150)
	w.button.SetName("开始下载")

	w.progress_bar = vcl.NewProgressBar(w.win)
	w.progress_bar.SetParent(w.win)
	w.progress_bar.SetPosition(0)
	w.progress_bar.SetWidth(400)
	w.progress_bar.SetHeight(20)
	w.progress_bar.SetLeft(0)
	w.progress_bar.SetTop(230)

	w.label_time = vcl.NewLabel(w.win)
	w.label_time.SetParent(w.win)
	w.label_time.SetName("有效时长")
	w.label_time.SetLeft(50)
	w.label_time.SetTop(20)

	w.time_tedit = vcl.NewEdit(w.win)
	w.time_tedit.SetParent(w.win)
	w.time_tedit.SetText("100")
	w.time_tedit.SetLeft(130)
	w.time_tedit.SetTop(20)
}

func GetFileName(strat_time, stop_time, text string) ([]crm_dat101, error) {
	//Engine := new(Engine)
	//err := Engine.NewEngine()
	//if err != nil {
	//	return nil, err
	//}
	//
	////fmt.Println("111111")
	//err = Engine.Engine.Where(fmt.Sprintf("SHIJIAN >= TO_DATE('%v 00:00:00', 'YYYY-MM-DD HH24:MI:SS') AND SHIJIAN <= TO_DATE('%v 23:59:59','YYYY-MM-DD HH24:MI:SS') AND SHICHANG >= '%v'", strat_time, stop_time, text)).Find(&Engine.FileName)
	//if err != nil {
	//	return nil, err
	//}
	//Engine.Engine.Close()

	File := new(file)
	File.Start = strat_time
	File.Stop = stop_time
	File.Times = text

	data1, err := json.Marshal(File)

	fileName := new([]crm_dat101)

	bady := url.Values{}
	bady.Add("Customer_number", string(data1))

	resp, err := http.PostForm("http://192.168.0.12:8888/v1/client/Accles", bady)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))
	err = json.Unmarshal(data, &fileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileName)

	return *fileName, nil
}

func (w *Windows) Onclick() {

	w.button.SetOnClick(func(sender vcl.IObject) {

		var C, s int
		//w.win.SetEnabled(false)
		w.button.SetEnabled(false)
		w.date_stop_label.SetEnabled(false)
		w.date_strat_label.SetEnabled(false)

		time_strat := w.date_strat_label.DateTime().Format("2006-01-02")
		time_stop := w.date_stop_label.DateTime().Format("2006-01-02")
		text := w.time_tedit.Text()

		data, err := GetFileName(time_strat, time_stop, text)
		//fmt.Println(len(data))

		if err != nil {
			vcl.ShowMessageFmt("数据库连接出错！")
		}
		//vcl.ShowMessageFmt(fmt.Sprintf("下载总数：%v个", len(data)))
		vcl.MessageDlg(fmt.Sprintf("下载总数：%v个。开始下载...", len(data)), 0, 0)
		//fmt.Println(0)

		go func() {
			for k, v := range data {
				c, err := ftp.Dial("192.168.0.19:21", ftp.DialWithTimeout(15*time.Second)) //创建连接
				if err != nil {
					vcl.ThreadSync(func() {
						vcl.ShowMessageFmt("FTP连接失败！")
					})
					return
				}
				err = c.Login("BOLONG", "131420") //登陆
				if err != nil {
					vcl.ThreadSync(func() {
						vcl.ShowMessageFmt("FTP登陆失败！")
					})
					return
				}

				r, err := c.Retr(fmt.Sprintf("%v/%v", strings.ReplaceAll(v.Shijian, "-", ""), v.Filename))
				//fmt.Println(strings.ReplaceAll(v.Shijian, "-", ""), v.Filename)

				if err != nil {
					s++
					//fmt.Println(err)
					vcl.ThreadSync(func() {
						w.progress_bar.SetPosition(int32(float64(k) / float64(len(data)) * 100))
					})
					continue
				}

				err = PathExists(fmt.Sprintf("D:/Record/%v", strings.ReplaceAll(v.Shijian, "-", "")))

				if err != nil {
					s++
					//fmt.Println(err)
					vcl.ThreadSync(func() {
						w.progress_bar.SetPosition(int32(float64(k) / float64(len(data)) * 100))
					})
					continue
				}

				buf, err := ioutil.ReadAll(r)
				if err != nil {
					s++
					//fmt.Println(err)
					vcl.ThreadSync(func() {
						w.progress_bar.SetPosition(int32(float64(k) / float64(len(data)) * 100))
					})
					continue
				}

				f, err := os.Create(fmt.Sprintf("D:/Record/%v/%v", strings.ReplaceAll(v.Shijian, "-", ""), v.Filename))

				if err != nil {
					s++
					//fmt.Println(err)
					vcl.ThreadSync(func() {
						w.progress_bar.SetPosition(int32(float64(k) / float64(len(data)) * 100))
					})
					continue
					vcl.ShowMessageFmt("创建文件失败！")
				} else {
					_, err = f.Write(buf)
					if err != nil {
						s++
						//fmt.Println(err)
						vcl.ThreadSync(func() {
							w.progress_bar.SetPosition(int32(float64(k) / float64(len(data)) * 100))
						})
						continue
						vcl.ShowMessageFmt("写入文件失败！")
					}

				}
				C++
				vcl.ThreadSync(func() {
					w.progress_bar.SetPosition(int32(float64(k) / float64(len(data)) * 100))
				})
				c.Quit()

			}
			time.Sleep(time.Second)
			vcl.ShowMessageFmt(fmt.Sprintf("成功：%v   失败：%v", C, s))
			vcl.ShowMessageFmt("下载完成！")

			//w.button.SetEnabled(false)
			//w.win.SetEnabled(true)
			w.date_stop_label.SetEnabled(true)
			w.date_strat_label.SetEnabled(true)
			w.button.SetEnabled(true)
			os.Exit(0)
		}()

	})

}

//判断文件是否存在，如不存在创建
func PathExists(path string) error {
	_, err := os.Stat(path)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) { //os.IsNotExist  判断 ERR 这个错误是否是文件不存在
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return err
}

func main() {
	TForm := new(Windows)
	TForm.init()
	TForm.Onclick()

	TForm.win.Show()
	vcl.Application.Run()
}
