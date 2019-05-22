package main

import (
	"fmt"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	"os"
)

//按钮控件结构体
type ControlButton struct {
	button_select   *gtk.Button //查询按钮
	button_register *gtk.Button //登记
	button_add      *gtk.Button //新增
	button_empty    *gtk.Button //清空
}

//文本框控件结构体
type ControlEntry struct {
	entry_select    *gtk.Entry //查询文本框
	entry_name      *gtk.Entry //客户姓名
	entry_number    *gtk.Entry //客户编码
	entry_staff     *gtk.Entry //收藏顾问
	entry_type      *gtk.Entry //客户分类
	entry_cellphone *gtk.Entry //客户电话
}

//标签文本结构体
type ControlLabel struct {
	label_head             *gtk.Label //上半部分标题
	label_select           *gtk.Label //客户查询标题：客户电话
	label_name             *gtk.Label //客户姓名
	label_number           *gtk.Label //客户编码
	label_staff            *gtk.Label //收藏顾问
	label_type             *gtk.Label //客户分类
	label_cellphone        *gtk.Label //客户手机
	label_tail             *gtk.Label //下半部分标题
	label_statistics_type1 *gtk.Label //统计已购类型名称
	label_statistics_type2 *gtk.Label //统计未购类型名称
	label_statistics_type3 *gtk.Label //总计名称
	label_statistics1      *gtk.Label //已购统计
	label_statistics2      *gtk.Label //未购统计
	label_statistics3      *gtk.Label //总计
}

//窗口控件
type Windows struct {
	w             *gtk.Window //win控件
	ControlButton             //按钮
	ControlLabel              //标题文字
	ControlEntry              //文本框

	width, height int
}

//获取窗口，获取控件，并设置属性
func (win *Windows) CreateWin() *gtk.Window {
	//创建builder 创建者
	builder := gtk.NewBuilder()
	builder.AddFromFile("./images/select.glade")

	//获取标签文字控件
	win.label_head = gtk.LabelFromObject(builder.GetObject("label_head"))
	win.label_tail = gtk.LabelFromObject(builder.GetObject("label_tail"))
	win.label_select = gtk.LabelFromObject(builder.GetObject("label_select"))
	win.label_name = gtk.LabelFromObject(builder.GetObject("label_name"))
	win.label_number = gtk.LabelFromObject(builder.GetObject("label_number"))
	win.label_head = gtk.LabelFromObject(builder.GetObject("label_head"))
	win.label_head = gtk.LabelFromObject(builder.GetObject("label_head"))
	win.label_head = gtk.LabelFromObject(builder.GetObject("label_head"))
	win.label_head = gtk.LabelFromObject(builder.GetObject("label_head"))
	win.label_head = gtk.LabelFromObject(builder.GetObject("label_head"))

	//获取文本框
	win.entry_select = gtk.EntryFromObject(builder.GetObject("entry_select"))
	win.entry_name = gtk.EntryFromObject(builder.GetObject("entry_name"))
	win.entry_number = gtk.EntryFromObject(builder.GetObject("entry_number"))
	win.entry_staff = gtk.EntryFromObject(builder.GetObject("entry_staff"))
	win.entry_cellphone = gtk.EntryFromObject(builder.GetObject("entry_cellphone"))
	win.entry_type = gtk.EntryFromObject(builder.GetObject("entry_type"))

	//获取按钮
	win.button_register = gtk.ButtonFromObject(builder.GetObject("button_register"))
	win.button_add = gtk.ButtonFromObject(builder.GetObject("button_add"))
	win.button_select = gtk.ButtonFromObject(builder.GetObject("button_select"))
	win.button_empty = gtk.ButtonFromObject(builder.GetObject("button_empty"))

	//获取窗口控件
	win.w = gtk.WindowFromObject(builder.GetObject("window1"))

	win.width, win.height = 1360, 768
	//设置win属性
	win.w.SetDefaultSize(win.width, win.height) //设置默认窗口大小
	win.w.SetPosition(gtk.WIN_POS_CENTER)       //设置窗口居中显示
	win.w.SetAppPaintable(true)                 //允许窗口绘图
	win.w.SetDecorated(true)                    //去除表框

	//按钮属性
	win.button_add.SetSensitive(false) //按钮开关  false  为关闭按钮不可使用

	//label字体属性

	//添加鼠标事件
	win.w.SetEvents(int(gdk.BUTTON_PRESS_MASK | gdk.BUTTON1_MOTION_MASK))

	return win.w
}

//绘图事件 、绘制背景图片
func DrawingEvents(ctx *glib.CallbackContext) {
	arg := ctx.Data()         //获取用户传递的参数
	win, ok := arg.(*Windows) //类型断言

	if !ok {
		fmt.Println("arg.(*Windows) err")
		return
	}

	//指定窗口绘图区域，在窗口上会绘图
	drawable := win.w.GetWindow().GetDrawable()
	gc := gdk.NewGC(drawable)

	//设置背景图片
	bj, _ := gdkpixbuf.NewPixbufFromFileAtScale("./images/bj.jpg", win.width, win.height, false)
	//bj, _ := gdkpixbuf.NewPixbufFromFile("./images/bj.jpg")
	//画图，画背景图
	//画图
	//bg：需要绘图的 pixbuf，第5、6个参数为画图的起点（相对于窗口而言）
	//第3、4个参数习惯为0，第7、8个参数为-1则按 pixbuf 原来的尺寸绘图
	//gdk.RGB_DITHER_NONE不用抖动，最后2个参数习惯写0
	drawable.DrawPixbuf(gc, bj, 0, 0, 0, 0, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)

	bj.Unref()

}

//事件处理 处理绘图按钮 以及随机抽奖
func (win *Windows) EventProcessing() {
	//窗口关闭程序退出
	win.w.Connect("destroy", func() {
		gtk.MainQuit()
	})

	//改变窗口大小，触发 configure-event ，然后手动刷新绘图区域，否则图片会重叠
	win.w.Connect("configure-event", func() {
		win.w.GetSize(&win.width, &win.height)

		//重新设置按钮 label 字体大小

		//绘图（曝光）事件，其回调函数PaintEvent做绘图操作，把year传递给回调函数 ,刷新背景图片
		win.w.Connect("expose-event", DrawingEvents, win)
		win.w.QueueDraw() //刷新绘图区域

	})
}

func main() {
	//初始化环境
	gtk.Init(&os.Args)

	var Win Windows
	Win.w = Win.CreateWin()
	Win.EventProcessing()
	Win.w.ShowAll()

	gtk.Main()
}
