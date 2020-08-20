package module

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TForm struct {
	Windows      *vcl.TForm     //win 窗口
	PanelA       *vcl.TPanel    //左侧容器
	PanelB       *vcl.TPanel    //右侧容器
	PanelC       *vcl.TPanel    //下册容器
	ListView     *vcl.TListView //自动列表
	Button       *vcl.TButton   //导出功能按钮
	OpenButton   *vcl.TButton   //导出功能按钮
	RadioButton1 *vcl.TRadioButton
	RadioButton2 *vcl.TRadioButton
	RadioButton3 *vcl.TRadioButton
	RadioButton4 *vcl.TRadioButton
	//OpenTextFileDialog *vcl.TOpenTextFileDialog
	ProgressBar *vcl.TProgressBar //进度条
	TLabel      *vcl.TLabel
	Operator    string
	Open_Import bool //打开文件次数
}

func (t *TForm) Init() {

	//新建窗口
	t.WindowsInit()            //初始化窗口环境
	t.PanelInit()              //初始化容器环境
	t.ListViewInit()           //初始化自动列表环境
	t.ButtonInit()             //初始化导出按钮
	t.RadioButtonSet()         //列表
	t.TOpenTextFileDialogSet() //打开文件的隐藏按钮
	t.ProgressBarSet()         //进度条
	t.LabelSet()

}

//窗口属性初始化
func (t *TForm) WindowsInit() {
	//新建窗口
	t.Windows = vcl.Application.CreateForm()    //创建床就
	t.Windows.SetPosition(types.PoScreenCenter) //
	t.Windows.SetName("信息导入")                   //程序名
	t.Windows.SetHeight(668)                    //高
	t.Windows.SetWidth(1160)                    //宽
	t.Windows.ScreenCenter()                    //居于当前屏幕中心
	//t.Windows.SetBorderIcons(3)                 //设置窗口最大化 最小化 关闭按钮  3代表是 最大化按钮和功能无效
	t.Windows.Font().SetSize(11)  //整体字体大小
	t.Windows.Font().SetStyle(16) //字体样式
	//t.Windows.SetTransparentColor(true)
	//t.Windows.SetTransparentColorValue(colors.ClGreenyellow)
}

//容器布局，属性初始化
func (t *TForm) PanelInit() {
	t.PanelA = vcl.NewPanel(t.Windows) //新建容器（左侧），并指定父窗口
	//t.PanelA.SetCaption("左")
	t.PanelA.SetParentBackground(false)
	t.PanelA.SetColor(0xF5FFFA)     //设置容器的颜色
	t.PanelA.SetParent(t.Windows)   //指定父窗口
	t.PanelA.SetWidth(200)          //设置容器额宽度
	t.PanelA.SetAlign(types.AlLeft) //设置容器的位置
	t.PanelA.Font().SetSize(11)
	t.PanelA.Font().SetStyle(2)

	t.PanelB = vcl.NewPanel(t.Windows) //新建容器（右侧），并指定父窗口
	//t.PanelB.SetCaption("右")
	t.PanelB.SetParentBackground(false)
	t.PanelB.SetColor(0xF8F8FF)       //设置容器的颜色
	t.PanelB.SetParent(t.Windows)     //指定父窗口
	t.PanelB.SetWidth(900)            //设置容器额宽度
	t.PanelB.SetAlign(types.AlClient) //设置容器的位置

	t.PanelC = vcl.NewPanel(t.Windows) //新建容器（下侧），并指定父窗口
	//t.PanelC.SetCaption("右")
	t.PanelC.SetParentBackground(false)
	t.PanelC.SetColor(0xF0FFFF)       //设置容器的颜色
	t.PanelC.SetParent(t.PanelB)      //指定父窗口
	t.PanelC.SetHeight(30)            //设置容器额高度
	t.PanelC.SetAlign(types.AlBottom) //设置容器的位置
}

//自动列表属性初始化化
func (t *TForm) ListViewInit() {
	t.ListView = vcl.NewListView(t.Windows)
	t.ListView.SetParent(t.PanelB)
	//t.ListView.SetWidth(t.PanelB.Width() + 200)
	t.ListView.SetAlign(types.AlClient)
	t.ListView.SetRowSelect(true)
	t.ListView.SetReadOnly(false)
	t.ListView.SetViewStyle(types.VsReport)
	t.ListView.SetGridLines(true)
	t.ListView.Font().SetStyle(16)
	t.ListView.Font().SetSize(11)
	//t.ListView.EnableAlign()
	//t.ListView.Realign()
	//t.ListView.SetFlatScrollBars(true)
	//t.ListView.Scro

	//初始化表头
	lvl := t.ListView.Columns().Add()
	lvl.SetCaption("导入结果")
	lvl.SetWidth(80)
	lvl.SetAlignment(types.TaCenter)

	lvl = t.ListView.Columns().Add()
	lvl.SetCaption("单据日期")
	lvl.SetWidth(t.PanelB.Width() / 10)
	lvl.SetAlignment(types.TaCenter)

	lvl = t.ListView.Columns().Add()
	lvl.SetCaption("工号")
	lvl.SetWidth(t.PanelB.Width() / 10)
	lvl.SetAlignment(types.TaCenter)

	lvl = t.ListView.Columns().Add()
	lvl.SetCaption("客户姓名")
	lvl.SetWidth(t.PanelB.Width() / 10)
	lvl.SetAlignment(types.TaCenter)

	lvl = t.ListView.Columns().Add()
	lvl.SetCaption("客户编码")
	lvl.SetWidth(t.PanelB.Width() / 10)
	lvl.SetAlignment(types.TaCenter)

	lvl = t.ListView.Columns().Add()
	lvl.SetCaption("客户电话")
	lvl.SetWidth(t.PanelB.Width() / 10)
	lvl.SetAlignment(types.TaCenter)

	lvl = t.ListView.Columns().Add()
	lvl.SetCaption("客户地址")
	lvl.SetWidth(t.PanelB.Width() / 10)
	lvl.SetAlignment(types.TaCenter)

	lvl = t.ListView.Columns().Add()
	lvl.SetCaption("客户分类")
	lvl.SetWidth(t.PanelB.Width() / 10)
	lvl.SetAlignment(types.TaCenter)

	lvl = t.ListView.Columns().Add()
	lvl.SetCaption("客户来源")
	lvl.SetWidth(t.PanelB.Width() / 10)
	lvl.SetAlignment(types.TaCenter)

	lvl = t.ListView.Columns().Add()
	lvl.SetCaption("库房")
	lvl.SetWidth(t.PanelB.Width() / 10)
	lvl.SetAlignment(types.TaCenter)

	lvl = t.ListView.Columns().Add()
	lvl.SetCaption("类别")
	lvl.SetWidth(t.PanelB.Width() / 10)
	lvl.SetAlignment(types.TaCenter)

	lvl = t.ListView.Columns().Add()
	lvl.SetCaption("票号")
	lvl.SetWidth(t.PanelB.Width() / 10)
	lvl.SetAlignment(types.TaCenter)

	lvl = t.ListView.Columns().Add()
	lvl.SetCaption("商品ID")
	lvl.SetWidth(t.PanelB.Width() / 10)
	lvl.SetAlignment(types.TaCenter)

	lvl = t.ListView.Columns().Add()
	lvl.SetCaption("数量")
	lvl.SetWidth(t.PanelB.Width() / 10)
	lvl.SetAlignment(types.TaCenter)

	lvl = t.ListView.Columns().Add()
	lvl.SetCaption("单价")
	lvl.SetWidth(t.PanelB.Width() / 10)
	lvl.SetAlignment(types.TaCenter)

	lvl = t.ListView.Columns().Add()
	lvl.SetCaption("备注")
	lvl.SetWidth(t.PanelB.Width() / 10)
	lvl.SetAlignment(types.TaCenter)

}

//导出功能按钮属性设置
//打开文件按钮属性
func (t *TForm) ButtonInit() {
	t.Button = vcl.NewButton(t.Windows) //开始导入按钮创建
	t.Button.SetParent(t.PanelC)
	t.Button.SetName("开始导入")         // 名字
	t.Button.SetAlign(types.AlRight) //位置
	//t.Button.SetLeft(t.PanelC.Width())
	//t.Button.SetTop((t.PanelC.Height() - t.Button.Height()) / 2)

	t.OpenButton = vcl.NewButton(t.Windows) //打开文件按钮创建
	t.OpenButton.SetParent(t.PanelC)
	t.OpenButton.SetName("打开文件")
	t.OpenButton.SetAlign(types.AlLeft) //位置
	//t.Button1.SetLeft(t.PanelC.Width() - t.Button.Width() - t.Button.Width() - 10)

}

//选择列表
func (t *TForm) RadioButtonSet() {
	//新建单选按钮
	t.RadioButton1 = vcl.NewRadioButton(t.Windows)
	t.RadioButton2 = vcl.NewRadioButton(t.Windows)
	t.RadioButton3 = vcl.NewRadioButton(t.Windows)
	t.RadioButton4 = vcl.NewRadioButton(t.Windows)

	//指定父容器
	t.RadioButton1.SetParent(t.PanelA)
	t.RadioButton2.SetParent(t.PanelA)
	t.RadioButton3.SetParent(t.PanelA)
	t.RadioButton4.SetParent(t.PanelA)

	//设定横向距离
	t.RadioButton1.SetLeft(30)
	t.RadioButton2.SetLeft(30)
	t.RadioButton3.SetLeft(30)
	t.RadioButton4.SetLeft(30)

	//设定竖向距离
	t.RadioButton1.SetTop(250)
	t.RadioButton2.SetTop(280)
	t.RadioButton3.SetTop(310)
	t.RadioButton4.SetTop(340)

	//设定名字内容
	t.RadioButton1.SetName("客户信息导入")
	t.RadioButton2.SetName("销售记录导入")
	t.RadioButton3.SetName("礼品记录导入")
	t.RadioButton4.SetName("邮寄记录导入")

	t.RadioButton2.SetChecked(true)

	//t.RadioButton3.SetEnabled(false)
	//t.RadioButton4.SetEnabled(false)
}

//打开文件-属性设置-操作打开文件的隐藏按钮
func (t *TForm) TOpenTextFileDialogSet() {
	//创建按钮
	//t.OpenTextFileDialog = vcl.NewOpenTextFileDialog(t.Windows)
	//指定父容器
	//t.OpenTextFileDialog.SetName("打开文件")
	//t.OpenTextFileDialog.Execute()
}

//进度条
//设置属性
//初始化
func (t *TForm) ProgressBarSet() {
	//创建进度条
	t.ProgressBar = vcl.NewProgressBar(t.Windows)
	t.ProgressBar.SetParent(t.PanelC)
	t.ProgressBar.SetPosition(0)           //初始化进度值
	t.ProgressBar.SetAlign(types.AlClient) //填充容器

}

//登陆人标签提示
func (t *TForm) LabelSet() {
	t.TLabel = vcl.NewLabel(t.Windows)
	t.TLabel.SetParent(t.PanelA)
	t.TLabel.SetAlign(types.AlBottom)
	t.TLabel.SetAlignment(types.TfAllowDialogCancellation)
}
