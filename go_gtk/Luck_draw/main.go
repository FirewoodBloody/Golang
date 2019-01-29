package main

import "C"
import (
	"fmt"
	"github.com/FirewoodBloody/PacketProup/ini_config"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	StartLabel = "S t a r t" //按钮显示内容
	StopLabel  = "S t o p"   //按钮显示内容
)

var Conf Config

type Config struct {
	ServerConf ServerConfig `ini:"server"`
}
type ServerConfig struct {
	SoftwareNam string `ini:"softwarename"`
	LabelName   string `ini:"labelname"`
	IconImage   string `ini:"iconimage"`
	FileName    string `ini:"filename"`
}

//控件结构体
type Control struct {
	window *gtk.Window
	label  *gtk.Label
	button *gtk.Button
}

type Attribute struct {
	w, h      int //窗口与背景的宽度和高度
	labelTime int //
}

type AnnualMeeting struct {
	Control
	Attribute

	ZJName []string

	Name   map[int]string //存储员工姓名
	NameId int

	LabelTimeId int
}

//处理员工姓名文件，并将信息存储到 name 中
func (year *AnnualMeeting) ReadFilename() {
	data, err := ioutil.ReadFile(Conf.ServerConf.FileName)
	if err != nil {
		fmt.Println(err)
	}
	year.NameId = 0
	year.Name = make(map[int]string, 1000)
	str := strings.Split(string(data), "\n")
	for k, v := range str {

		if str[k] == "" || str[k] == "\r\n" || len(v) < 3 {
			continue
		}
		if k != 0 {
			year.NameId++
		}
		year.Name[year.NameId] = v
	}

}

//按钮事件处理 ,
func (year *AnnualMeeting) ButtonEvents() {

	year.button.Connect("clicked", func() {
		buttonLabel := year.button.GetLabel()
		if buttonLabel == StartLabel {
			//重新设置按钮 label 字体大小
			year.button.SetLabel(StopLabel)
			year.button.SetLabelFontSize(int(year.w / 3 / 3 / 6))
			//定时器
			//开始滚动label标签内容，每个10毫秒进行一次切换
			year.LabelTimeId = 0
			year.labelTime = glib.TimeoutAdd(1, func() bool {
				year.LabelTimeId++
				if year.LabelTimeId > year.NameId {
					year.LabelTimeId = 0
				}
				if year.button.GetLabel() == StopLabel {
					year.label.SetText(year.Name[year.LabelTimeId])
				}
				return true

			})

		} else if buttonLabel == StopLabel {
			//重新设置按钮 label 字体大小
			year.button.SetLabel(StartLabel)
			year.button.SetLabelFontSize(int(year.w / 3 / 3 / 6))
			if year.NameId == 0 {
				return
			}
			//生成随机数，进行中奖人员抽取
			rand.Seed(time.Now().Unix()) //初始化时间戳

			a := rand.Intn(year.NameId)
			year.label.SetText(year.Name[a])
			year.Name[a] = year.Name[year.NameId]

			year.NameId--

		}
	})

}

//绘图事件 、绘制背景图片
func DrawingEvents(ctx *glib.CallbackContext) {
	arg := ctx.Data()                //获取用户传递的参数
	year, ok := arg.(*AnnualMeeting) //类型断言

	if !ok {
		fmt.Println("arg.(*AnnualMeeting) err")
		return
	}

	//指定窗口绘图区域，在窗口上会绘图
	drawable := year.window.GetWindow().GetDrawable()
	gc := gdk.NewGC(drawable)

	//设置背景图片
	bj, _ := gdkpixbuf.NewPixbufFromFileAtScale("./image/bj.jpg", year.w, year.h, false)
	//画图，画背景图
	//画图
	//bg：需要绘图的 pixbuf，第5、6个参数为画图的起点（相对于窗口而言）
	//第3、4个参数习惯为0，第7、8个参数为-1则按 pixbuf 原来的尺寸绘图
	//gdk.RGB_DITHER_NONE不用抖动，最后2个参数习惯写0
	drawable.DrawPixbuf(gc, bj, 0, 0, 0, 0, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)

	bj.Unref()

}

//获取窗口，获取控件，设置属性
func (year *AnnualMeeting) CreateWindow() *gtk.Window {
	//新建 builder 创建者
	builder := gtk.NewBuilder()
	builder.AddFromFile("./image/window.glade")

	//获取控件
	year.window = gtk.WindowFromObject(builder.GetObject("window1"))
	year.label = gtk.LabelFromObject(builder.GetObject("label1"))
	year.button = gtk.ButtonFromObject(builder.GetObject("button1"))

	year.w, year.h = 1200, 800
	//设置窗口属性
	year.window.SetDefaultSize(year.w, year.h)  //设置窗口大小
	year.window.SetPosition(gtk.WIN_POS_CENTER) //设置窗口居中显示
	year.window.SetAppPaintable(true)           //允许窗口绘图
	year.window.SetDecorated(true)              //去除表框
	year.window.SetTitle(Conf.ServerConf.SoftwareNam)
	year.window.SetIconFromFile(Conf.ServerConf.IconImage)

	//设置按钮属性
	year.button.SetCanFocus(false) //去掉按钮上的聚焦框
	year.button.SetLabel(StartLabel)
	year.button.SetLabelFontSize(int(year.w / 3 / 3 / 6))

	//设置 label 属性、字体大小
	year.label.SetText(Conf.ServerConf.LabelName)
	year.label.ModifyFontSize(int(year.w / 3 / 6))

	//绘图（曝光）事件，其回调函数PaintEvent做绘图操作，把year传递给回调函数
	//year.window.Connect("expose-event", DrawingEvents, year)

	//添加鼠标事件
	year.window.SetEvents(int(gdk.BUTTON_PRESS_MASK | gdk.BUTTON1_MOTION_MASK))

	return year.window
}

//事件处理 处理绘图按钮 以及随机抽奖
func (year *AnnualMeeting) EventProcessing() {
	//窗口关闭程序退出
	year.window.Connect("destroy", func() {
		gtk.MainQuit()
	})

	//改变窗口大小，触发 configure-event ，然后手动刷新绘图区域，否则图片会重叠
	year.window.Connect("configure-event", func() {
		year.window.GetSize(&year.w, &year.h)

		//重新设置按钮 label 字体大小
		year.button.SetLabelFontSize(int(year.w / 3 / 3 / 6))
		year.label.ModifyFontSize(int(year.w / 3 / 6))

		//绘图（曝光）事件，其回调函数PaintEvent做绘图操作，把year传递给回调函数 ,刷新背景图片
		year.window.Connect("expose-event", DrawingEvents, year)
		year.window.QueueDraw() //刷新绘图区域

	})

	//按钮事件
	year.ButtonEvents()

	//处理员工姓名文件
	year.ReadFilename()
}

func main() {
	//初始化环境
	gtk.Init(&os.Args)

	var year AnnualMeeting

	err := ini_config.UnMarshalFile("./config/config.ini", &Conf)
	if err != nil {
		fmt.Println(err)
	}

	window := year.CreateWindow() //获取窗口和控件
	year.EventProcessing()        //事件处理

	window.ShowAll()

	gtk.Main()
}
