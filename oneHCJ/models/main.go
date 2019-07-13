package main

import (
	"fmt"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/mattn/go-gtk/gdk"
	"github.com/mattn/go-gtk/gdkpixbuf"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
	_ "github.com/wendal/go-oci8"
	"os"
	"strconv"
	"time"
)

//按钮控件结构体
type ControlButton struct {
	button_select   *gtk.Button //查询按钮
	button_register *gtk.Button //登记
	button_add      *gtk.Button //新增
	//button_empty    *gtk.Button //清空
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

	entrySelectLabel, entryNameLabel, entryCellphoneLabel, entryStaffLabel string

	CustomerInformation CustomerInformation
}

type CustomerInformation struct {
	name      string
	number    string
	staff     string
	types     string
	cellphone string
}

type CrmDat001 struct {
	Name      string `xorm:"varchar2(15) 'KHMC'"`
	Number    string `xorm:"number(18) 'KHID'"`
	Staff     string `xorm:"number(15) 'GONGHAO'"`
	Types     string `xorm:"varchar2(32) 'TYPEID'"`
	Cellphone string `xorm:"varchar2(12) 'MOBIL'"`
}

var tbMappers core.PrefixMapper

func init() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码
	tbMappers = core.NewPrefixMapper(core.SnakeMapper{}, "BLCRM.")
}

//获取窗口，获取控件，并设置属性
func (win *Windows) CreateWin() *gtk.Window {
	//创建builder 创建者
	builder := gtk.NewBuilder()
	builder.AddFromFile("./images/select.glade")

	//获取窗口控件
	win.w = gtk.WindowFromObject(builder.GetObject("window1"))

	//获取标签文字控件
	win.label_head = gtk.LabelFromObject(builder.GetObject("label_head"))
	win.label_tail = gtk.LabelFromObject(builder.GetObject("label_tail"))
	win.label_select = gtk.LabelFromObject(builder.GetObject("label_select"))
	win.label_name = gtk.LabelFromObject(builder.GetObject("label_name"))
	win.label_number = gtk.LabelFromObject(builder.GetObject("label_number"))
	win.label_staff = gtk.LabelFromObject(builder.GetObject("label_staff"))
	win.label_type = gtk.LabelFromObject(builder.GetObject("label_type"))
	win.label_cellphone = gtk.LabelFromObject(builder.GetObject("label_cellphone"))
	win.label_statistics_type1 = gtk.LabelFromObject(builder.GetObject("label_statistics_type1"))
	win.label_statistics_type2 = gtk.LabelFromObject(builder.GetObject("label_statistics_type2"))
	win.label_statistics_type3 = gtk.LabelFromObject(builder.GetObject("label_statistics_type3"))
	win.label_statistics1 = gtk.LabelFromObject(builder.GetObject("label_statistics1"))
	win.label_statistics2 = gtk.LabelFromObject(builder.GetObject("label_statistics2"))
	win.label_statistics3 = gtk.LabelFromObject(builder.GetObject("label_statistics3"))

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
	//win.button_empty = gtk.ButtonFromObject(builder.GetObject("button_empty"))

	win.width, win.height = 1360, 768

	//设置win属性
	win.w.SetDefaultSize(win.width, win.height) //设置默认窗口大小
	win.w.SetPosition(gtk.WIN_POS_CENTER)       //设置窗口居中显示
	win.w.SetAppPaintable(true)                 //允许窗口绘图
	win.w.SetDecorated(true)                    //去除表框
	win.w.SetTitle("来访登记管理")                    //设置标题
	win.w.SetIconFromFile("./images/bl.jpg")    //设置陈旭图标

	//按钮属性
	win.button_add.SetSensitive(false)      //按钮开关  false  为关闭按钮不可使用
	win.button_register.SetSensitive(false) //按钮开关  false  为关闭按钮不可使用
	win.button_select.SetLabelFontSize(10)
	win.button_add.SetLabelFontSize(10)
	win.button_register.SetLabelFontSize(10)

	//设置 标签 属性 label字体属性
	win.label_head.ModifyFontSize(22)
	win.label_tail.ModifyFontSize(22)
	win.label_select.ModifyFontSize(11)
	win.label_name.ModifyFontSize(11)
	win.label_number.ModifyFontSize(11)
	win.label_staff.ModifyFontSize(11)
	win.label_type.ModifyFontSize(11)
	win.label_cellphone.ModifyFontSize(11)
	win.label_statistics_type1.ModifyFontSize(14)
	win.label_statistics_type2.ModifyFontSize(14)
	win.label_statistics_type3.ModifyFontSize(14)
	win.label_statistics1.ModifyFontSize(16)
	win.label_statistics2.ModifyFontSize(16)
	win.label_statistics3.ModifyFontSize(16)

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
	//bj, _ := gdkpixbuf.NewPixbufFromFile("./images/1bj.jpg")
	//画图，画背景图
	//画图
	//bg：需要绘图的 pixbuf，第5、6个参数为画图的起点（相对于窗口而言）
	//第3、4个参数习惯为0，第7、8个参数为-1则按 pixbuf 原来的尺寸绘图
	//gdk.RGB_DITHER_NONE不用抖动，最后2个参数习惯写0
	drawable.DrawPixbuf(gc, bj, 0, 0, 0, 0, -1, -1, gdk.RGB_DITHER_NONE, 0, 0)

	bj.Unref()

}

//清空数据
func (win *Windows) empty() {
	win.entry_name.SetText("")
	win.entry_number.SetText("")
	win.entry_staff.SetText("")
	win.entry_type.SetText("")
	win.entry_cellphone.SetText("")
}

//事件处理 处理绘图按钮 按钮事件处理
func (win *Windows) EventProcessing() {
	//窗口关闭程序退出
	win.w.Connect("destroy", func() {
		glib.TimeoutRemove(0)
		gtk.MainQuit()
	})

	_ = glib.TimeoutAdd(5000, func() bool {
		date := time.Now().Format("2006-01-02")
		Engine, _ := xorm.NewEngine("oci8", "KD/1219271@192.168.0.9:1521/BLDB")

		sql1 := fmt.Sprintf("SELECT COUNT(IFS) FROM BLCRM.KHDJ WHERE IFS = '老' and RIQI = TO_DATE('%v','YYYY-MM-DD')", date)
		sql2 := fmt.Sprintf("SELECT COUNT(IFS) FROM BLCRM.KHDJ WHERE IFS = '新' and RIQI = TO_DATE('%v','YYYY-MM-DD')", date)
		sql := fmt.Sprintf("SELECT COUNT(IFS) FROM BLCRM.KHDJ WHERE  RIQI = TO_DATE('%v','YYYY-MM-DD')", date)
		A, _ := Engine.Query(sql1)
		B, _ := Engine.Query(sql2)
		C, _ := Engine.Query(sql)
		_ = Engine.Close()
		for _, v := range A {
			for _, i := range v {
				if string(i) != "" {
					win.label_statistics1.SetText(string(i))
				}
			}
		}

		for _, v := range B {
			for _, i := range v {
				if string(i) != "" {
					win.label_statistics2.SetText(string(i))
				}
			}
		}

		for _, v := range C {
			for _, i := range v {
				if string(i) != "" {
					win.label_statistics3.SetText(string(i))
				}
			}
		}

		return true
	})

	//改变窗口大小，触发 configure-event ，然后手动刷新绘图区域，否则图片会重叠
	win.w.Connect("configure-event", func() {
		win.w.GetSize(&win.width, &win.height)

		//重新设置按钮 label 字体大小

		//绘图（曝光）事件，其回调函数PaintEvent做绘图操作，把year传递给回调函数 ,刷新背景图片
		win.w.Connect("expose-event", DrawingEvents, win)
		win.w.QueueDraw() //刷新绘图区域

	})

	//查询按钮事件处理
	win.button_select.Connect("clicked", func() {
		//获取查询按钮所填写的查询条件数据：电话号
		win.entrySelectLabel = win.entry_select.GetText()

		_, err := strconv.Atoi(win.entrySelectLabel)

		if len(win.entrySelectLabel) != 11 || err != nil { //判断输入信息是否为纯数字
			//新建消息对话框，提示对话框
			dialog := gtk.NewMessageDialog(
				win.button_select.GetTopLevelAsWindow(), //指定父窗口)
				gtk.DIALOG_MODAL,                        //模态对话框
				gtk.MESSAGE_QUESTION,                    //指定对话框类型
				gtk.BUTTONS_OK,                          //默认按钮
				"您输入的号码有误，请重新输入！")                       //设置内容

			dialog.SetTitle("警告！") //对话框设置标题

			dialog.Run()     //运行对话框
			dialog.Destroy() //销毁对话框

			//清空文本内容
			win.empty()

		} else {
			err := win.QueryCustomerInformation()
			if err != nil {
				dialog := gtk.NewMessageDialog( //新建消息对话框，提示对话框
					win.button_select.GetTopLevelAsWindow(), //指定父窗口)
					gtk.DIALOG_MODAL,                        //模态对话框
					gtk.MESSAGE_QUESTION,                    //指定对话框类型
					gtk.BUTTONS_OK,                          //默认按钮
					err.Error())                             //设置内容

				dialog.SetTitle("错误！") //对话框设置标题

				dialog.Run()     //运行对话框
				dialog.Destroy() //销毁对话框

				//清空文本内容
				win.empty()
			} else {
				if win.CustomerInformation.name == "" {
					dialog := gtk.NewMessageDialog( //新建消息对话框，提示对话框
						win.button_select.GetTopLevelAsWindow(), //指定父窗口)
						gtk.DIALOG_MODAL,                        //模态对话框
						gtk.MESSAGE_QUESTION,                    //指定对话框类型
						gtk.BUTTONS_YES_NO,                      //默认按钮
						"客户信息不存在，是否新增？")                         //设置内容

					dialog.SetTitle("提示！") //对话框设置标题

					flag := dialog.Run() //运行对话框
					if flag == gtk.RESPONSE_YES {
						win.button_add.SetSensitive(true) //查询无此数据，打开新增按钮

						//打开行编辑
						win.entry_name.SetEditable(true)
						win.entry_staff.SetEditable(true)
						win.entry_cellphone.SetEditable(true)

						dialog.Destroy()
						//清空文本内容
						win.empty()
					} else if flag == gtk.RESPONSE_NO {
						dialog.Destroy() //销毁对话框
						//清空文本内容
						win.empty()
					} else {
						dialog.Destroy() //销毁对话框
						//清空文本内容
						win.empty()
					}
				} else {
					win.entry_name.SetText(win.CustomerInformation.name)
					win.entry_number.SetText(win.CustomerInformation.number)
					win.entry_staff.SetText(win.CustomerInformation.staff)
					win.entry_type.SetText(win.CustomerInformation.types)
					win.entry_cellphone.SetText(win.CustomerInformation.cellphone)
					win.button_register.SetSensitive(true) //打开登记按钮
				}
			}
		}

	})

	//登记按钮事件处理
	win.button_register.Connect("clicked", func() {
		err := Insert(win.CustomerInformation.name, win.CustomerInformation.number, win.CustomerInformation.cellphone, win.CustomerInformation.staff, win.CustomerInformation.types, "老")
		if err != nil {
			dialog := gtk.NewMessageDialog( //新建消息对话框，提示对话框
				win.button_select.GetTopLevelAsWindow(), //指定父窗口)
				gtk.DIALOG_MODAL,                        //模态对话框
				gtk.MESSAGE_QUESTION,                    //指定对话框类型
				gtk.BUTTONS_OK,                          //默认按钮
				err.Error())                             //设置内容

			dialog.SetTitle("错误！") //对话框设置标题

			dialog.Run()     //运行对话框
			dialog.Destroy() //销毁对话框
			//清空文本内容
			win.button_register.SetSensitive(false) //关闭登记按钮
			win.empty()
		} else {
			dialog := gtk.NewMessageDialog( //新建消息对话框，提示对话框
				win.button_select.GetTopLevelAsWindow(), //指定父窗口)
				gtk.DIALOG_MODAL,                        //模态对话框
				gtk.MESSAGE_QUESTION,                    //指定对话框类型
				gtk.BUTTONS_OK,                          //默认按钮
				"登记成功！")                                 //设置内容

			dialog.SetTitle("提示！") //对话框设置标题

			dialog.Run()     //运行对话框
			dialog.Destroy() //销毁对话框
			//清空文本内容
			win.button_register.SetSensitive(false) //关闭登记按钮
			win.empty()
		}

	})

	//处理新增按钮事件
	win.button_add.Connect("clicked", func() {
		//获取文本框数据
		win.entryNameLabel = win.entry_name.GetText()
		win.entryCellphoneLabel = win.entry_cellphone.GetText()
		win.entryStaffLabel = win.entry_staff.GetText()

		err := Insert(win.entryNameLabel, "", win.entryCellphoneLabel, win.entryStaffLabel, "", "新")
		if err != nil {
			dialog := gtk.NewMessageDialog( //新建消息对话框，提示对话框
				win.button_select.GetTopLevelAsWindow(), //指定父窗口)
				gtk.DIALOG_MODAL,                        //模态对话框
				gtk.MESSAGE_QUESTION,                    //指定对话框类型
				gtk.BUTTONS_OK,                          //默认按钮
				err.Error())                             //设置内容

			dialog.SetTitle("错误！") //对话框设置标题

			dialog.Run()     //运行对话框
			dialog.Destroy() //销毁对话框
		} else {
			dialog := gtk.NewMessageDialog( //新建消息对话框，提示对话框
				win.button_select.GetTopLevelAsWindow(), //指定父窗口)
				gtk.DIALOG_MODAL,                        //模态对话框
				gtk.MESSAGE_QUESTION,                    //指定对话框类型
				gtk.BUTTONS_OK,                          //默认按钮
				"登记成功！")                                 //设置内容

			dialog.SetTitle("提示！") //对话框设置标题

			dialog.Run()     //运行对话框
			dialog.Destroy() //销毁对话框
		}

		//清空文本内容
		win.empty()

		win.button_add.SetSensitive(false)
	})
}

//查询客户信息
func (win *Windows) QueryCustomerInformation() error {
	datas := new(CrmDat001)
	Engine, err := xorm.NewEngine("oci8", "BLCRM/BLCRM2012@192.168.0.9:1521/BLDB")
	if err != nil {
		return fmt.Errorf("查询数据失败！")
	}
	//Engine.ShowSQL(true)
	Engine.SetTableMapper(tbMappers)
	defer Engine.Close()

	sql := fmt.Sprintf("MOBIL = '%s'", win.entrySelectLabel)
	has, err := Engine.Where(sql).Get(datas)

	if err != nil && !has {
		return fmt.Errorf("查询数据失败！！")
	}

	if !has {
		win.CustomerInformation.name = ""
		win.CustomerInformation.number = ""
		win.CustomerInformation.cellphone = ""
		win.CustomerInformation.staff = ""
		win.CustomerInformation.types = ""
	} else {
		win.CustomerInformation.name = datas.Name
		win.CustomerInformation.number = datas.Number
		win.CustomerInformation.cellphone = datas.Cellphone
		if datas.Staff != "" {
			sql := fmt.Sprintf("SELECT NAME FROM CRM_SYS02 WHERE NO = '%v'", datas.Staff)
			maps, _ := Engine.Query(sql)
			for _, v := range maps {
				for _, i := range v {
					if string(i) != "" {
						win.CustomerInformation.staff = string(i)
					}
				}
			}
		} else {
			win.CustomerInformation.staff = ""
		}
		if datas.Types != "" {
			sql := fmt.Sprintf("SELECT LN1 FROM CRM_DAT011 WHERE TYPEID = '%v'", datas.Types)
			maps, _ := Engine.Query(sql)
			for _, v := range maps {
				for _, i := range v {
					if string(i) != "" {
						win.CustomerInformation.types = string(i)
					}
				}
			}
		} else {
			win.CustomerInformation.types = ""
		}
	}

	return nil
}

func Insert(Name, number, Cellphone, Staff, Type, IF string) error {
	Engine, err := xorm.NewEngine("oci8", "KD/1219271@192.168.0.9:1521/BLDB")
	if err != nil {
		return fmt.Errorf("登记失败，请重试！")
	}
	defer Engine.Close()
	date := time.Now().Format("2006-01-02")

	SqlIsert := fmt.Sprintf("INSERT INTO BLCRM.KHDJ VALUES(TO_DATE('%v','YYYY-MM-DD'),'%v','%v','%v','%v','%v','%v')", date, Name, number, Cellphone, Staff, Type, IF)
	_, err = Engine.Exec(SqlIsert)
	if err != nil {
		return fmt.Errorf("客户已登记！")
	}
	return nil
}

func updateshu(win Windows) {
	for {
		date := time.Now().Format("2006-01-02")
		Engine, _ := xorm.NewEngine("oci8", "KD/1219271@192.168.0.9:1521/BLDB")

		sql1 := fmt.Sprintf("SELECT COUNT(IFS) FROM BLCRM.KHDJ WHERE IFS = '老' and RIQI = TO_DATE('%v','YYYY-MM-DD')", date)
		sql2 := fmt.Sprintf("SELECT COUNT(IFS) FROM BLCRM.KHDJ WHERE IFS = '新' and RIQI = TO_DATE('%v','YYYY-MM-DD')", date)
		sql := fmt.Sprintf("SELECT COUNT(IFS) FROM BLCRM.KHDJ WHERE  RIQI = TO_DATE('%v','YYYY-MM-DD')", date)
		A, _ := Engine.Query(sql1)
		B, _ := Engine.Query(sql2)
		C, _ := Engine.Query(sql)

		for _, v := range A {
			for _, i := range v {
				if string(i) != "" {
					win.label_statistics1.SetText(string(i))
				}
			}
		}

		for _, v := range B {
			for _, i := range v {
				if string(i) != "" {
					win.label_statistics2.SetText(string(i))
				}
			}
		}

		for _, v := range C {
			for _, i := range v {
				if string(i) != "" {
					win.label_statistics3.SetText(string(i))
				}
			}
		}

		time.Sleep(time.Second * 3)
	}
}

func main() {
	//初始化环境
	gtk.Init(&os.Args)

	Win := Windows{}

	Win.w = Win.CreateWin()
	Win.EventProcessing()
	go updateshu(Win)
	Win.w.ShowAll()

	gtk.Main()
}
