package module

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type Users struct {
	Win           []Win  `json:"win"`
	Err           string `json:"err"`
	Client_Models string `json:"client_models"`
}

type Win struct {
	Rd string `json:"rd"`
}

type Windows struct {
	LoginTForm LoginTForm

	Win      *vcl.TForm   //窗口
	Button   *vcl.TButton //下载
	listView *vcl.TListView

	TPanel1 *vcl.TPanel
	TPanel2 *vcl.TPanel

	LoginName     string
	Client_Models string
}

var (
	isLogin   bool
	tbMappers core.PrefixMapper
)

type LoginTForm struct {
	WinLogin *vcl.TForm
	TEdit1   *vcl.TEdit
	TEdit2   *vcl.TEdit
	TButton1 *vcl.TButton
	TButton2 *vcl.TButton
	TLabel1  *vcl.TLabel
	TLabel2  *vcl.TLabel
}

type Statu struct {
	Err    string `json:"err"`
	Status bool   `json:"status"`
}

func init() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码
}

//登陆窗口控件初始化
func (l *LoginTForm) Init() {
	l.WinLogin = vcl.Application.CreateForm() //新建窗口
	l.WinLogin.SetCaption("电话量查询")            //程序名
	//w.Win.SetFormStyle(2)
	l.WinLogin.SetHeight(300)     //高
	l.WinLogin.SetWidth(400)      //宽
	l.WinLogin.ScreenCenter()     //居于当前屏幕中心
	l.WinLogin.SetBorderIcons(3)  //设置窗口最大化 最小化 关闭按钮  3代表是 最大化按钮和功能无效
	l.WinLogin.Font().SetSize(11) //整体字体大小
	l.WinLogin.Font().SetColor(255)
	l.WinLogin.Font().SetStyle(16) //字体样式
	l.WinLogin.SetColor(16775388)

	l.TButton1 = vcl.NewButton(l.WinLogin)
	l.TButton1.SetParent(l.WinLogin)
	l.TButton1.SetCaption("登陆")
	l.TButton1.SetLeft(80)
	l.TButton1.SetTop(200)

	l.TButton2 = vcl.NewButton(l.WinLogin)
	l.TButton2.SetParent(l.WinLogin)
	l.TButton2.SetCaption("取消")
	l.TButton2.SetLeft(240)
	l.TButton2.SetTop(200)

	l.TEdit1 = vcl.NewEdit(l.WinLogin)
	l.TEdit1.SetParent(l.WinLogin)
	l.TEdit1.SetTop(80)
	l.TEdit1.SetLeft(165)
	l.TEdit1.SetWidth(100)

	l.TEdit2 = vcl.NewEdit(l.WinLogin)
	l.TEdit2.SetParent(l.WinLogin)
	l.TEdit2.SetTop(125)
	l.TEdit2.SetLeft(165)
	l.TEdit2.SetPasswordChar(42)
	l.TEdit2.SetWidth(100)

	l.TLabel1 = vcl.NewLabel(l.WinLogin)
	l.TLabel1.SetParent(l.WinLogin)
	l.TLabel1.SetCaption("工号")
	l.TLabel1.SetLeft(110)
	l.TLabel1.SetTop(83)

	l.TLabel2 = vcl.NewLabel(l.WinLogin)
	l.TLabel2.SetParent(l.WinLogin)
	l.TLabel2.SetCaption("密码")
	l.TLabel2.SetLeft(110)
	l.TLabel2.SetTop(128)

}

//工作窗口控件初始化
func (w *Windows) Init() {
	w.Win = vcl.Application.CreateForm() //新建窗口
	w.Win.SetCaption("电话量查询")            //程序名
	//w.Win.SetFormStyle(2)
	w.Win.SetHeight(700)     //高
	w.Win.SetWidth(1300)     //宽
	w.Win.ScreenCenter()     //居于当前屏幕中心
	w.Win.SetBorderIcons(3)  //设置窗口最大化 最小化 关闭按钮  3代表是 最大化按钮和功能无效
	w.Win.Font().SetSize(11) //整体字体大小
	w.Win.Font().SetColor(0)
	w.Win.Font().SetStyle(16) //字体样式
	w.Win.SetColor(16775388)
	//w.Win.SetTransparentColor(true)
	//w.Win.SetTransparentColorValue(1)

	w.TPanel1 = vcl.NewPanel(w.Win)
	w.TPanel1.SetParent(w.Win)
	w.TPanel1.SetHeight(600)
	w.TPanel1.SetWidth(1300)
	w.TPanel1.SetAlign(types.AlTop)

	w.TPanel2 = vcl.NewPanel(w.Win)
	w.TPanel2.SetParent(w.Win)
	w.TPanel2.SetHeight(100)
	w.TPanel2.SetWidth(1300)
	w.TPanel2.SetAlign(types.AlBottom)

	w.Button = vcl.NewButton(w.Win)
	w.Button.SetParent(w.TPanel2)
	//w.Button.SetHeight(50)
	//w.Button.SetWidth(100)
	w.Button.SetLeft(w.TPanel2.Width() - w.Button.Width() - 20)
	w.Button.SetTop(w.TPanel2.Height() - w.Button.Height() - 40)
	w.Button.SetCaption("刷新")

	w.listView = vcl.NewListView(w.Win)
	w.listView.SetParent(w.TPanel1)
	w.listView.SetAlign(types.AlTop)
	w.listView.SetRowSelect(true)
	w.listView.SetReadOnly(false)
	w.listView.SetViewStyle(types.VsReport)
	w.listView.SetGridLines(true)
	//w.listView.SetWidth(1400)
	w.listView.SetHeight(597)
	w.listView.SetLeft(w.TPanel1.Width() - w.listView.Width() - 3)
	w.listView.SetTop(w.TPanel1.Height() - w.listView.Height() - 3)
	w.listView.Font().SetStyle(16)
	w.listView.Font().SetSize(10)
	w.listView.SetViewStyle(types.VsReport)

	lvl := w.listView.Columns().Add()
	lvl.SetCaption("序号")
	lvl.SetWidth(30)

	lvl = w.listView.Columns().Add()
	lvl.SetCaption("部门")
	lvl.SetWidth(85)

	lvl = w.listView.Columns().Add()
	lvl.SetCaption("二级部门")
	lvl.SetWidth(90)

	lvl = w.listView.Columns().Add()
	lvl.SetCaption("员工姓名")
	lvl.SetWidth(90)

	lvl = w.listView.Columns().Add()
	lvl.SetCaption("员工工号")
	lvl.SetWidth(90)

	lvl = w.listView.Columns().Add()
	lvl.SetCaption("未接通话总数")
	lvl.SetWidth(90)

	lvl = w.listView.Columns().Add()
	lvl.SetCaption("未接通时长")
	lvl.SetWidth(90)

	lvl = w.listView.Columns().Add()
	lvl.SetCaption("无效通话总数")
	lvl.SetWidth(90)

	lvl = w.listView.Columns().Add()
	lvl.SetCaption("无效通话时长")
	lvl.SetWidth(90)

	lvl = w.listView.Columns().Add()
	lvl.SetCaption("有效通话总数")
	lvl.SetWidth(90)

	lvl = w.listView.Columns().Add()
	lvl.SetCaption("有效通话时长")
	lvl.SetWidth(90)

	lvl = w.listView.Columns().Add()
	lvl.SetCaption("优质通话总数")
	lvl.SetWidth(90)

	lvl = w.listView.Columns().Add()
	lvl.SetCaption("优质通话时长")
	lvl.SetWidth(90)

	lvl = w.listView.Columns().Add()
	lvl.SetCaption("总通话数")
	lvl.SetWidth(95)

	lvl = w.listView.Columns().Add()
	lvl.SetCaption("通话总时长")
	lvl.SetWidth(90)

}

//工作窗口事件
func (w *Windows) Onclick() {

	//	查询下载按钮
	w.Button.SetOnClick(func(sender vcl.IObject) {
		w.listView.Clear()
		w.Win.SetEnabled(false)
		w.Button.SetEnabled(false)
		go func() { Call_log_statistics(w) }()
	})

	w.Win.SetOnCloseQuery(func(sender vcl.IObject, canClose *bool) {
		if *canClose {
			*canClose = vcl.MessageDlg("是否退出?", types.MtConfirmation, types.MbYes, types.MbNo) == types.MrYes
		}
		if *canClose {
			os.Exit(0)
		}
	})

	w.listView.SetOnCustomDrawSubItem(func(sender *vcl.TListView, item *vcl.TListItem, subItem int32, state types.TCustomDrawStage, defaultDraw *bool) {
		if subItem == 11 || subItem == 12 {
			if a, _ := strconv.Atoi(item.SubItems().Strings(10)); a >= 8 {
				sender.Canvas().Brush().SetColor(colors.ClGreen)
				sender.Canvas().Font().SetColor(colors.ClBlack)
			} else {
				sender.Canvas().Font().SetColor(colors.ClRed)
			}
		}

		//sender.Canvas().Font().SetColor(colors.ClBlack)
		//fmt.Println(item.SubItems().Strings(subItem - 1))

	})

	//自动列表排序事件
	w.listView.SetOnColumnClick(func(sender vcl.IObject, column *vcl.TListColumn) {
		w.listView.CustomSort(0, int(column.Index()))

	})

	//自定义排序方法
	w.listView.SetOnCompare(func(sender vcl.IObject, item1, item2 *vcl.TListItem, data int32, compare *int32) {
		if data == 0 {
			*compare = int32(strings.Compare(item1.Caption(), item2.Caption()))
		} else {
			*compare = int32(strings.Compare(item1.SubItems().Strings(data-1), item2.SubItems().Strings(data-1)))
		}
	})

}

//登录窗口事件
func (w *Windows) LoginOnclick() {
	//登录按钮
	w.LoginTForm.TButton1.SetOnClick(func(sender vcl.IObject) {
		w.LoginTForm.WinLogin.SetEnabled(false)
		w.LoginTForm.TButton1.SetEnabled(false)
		w.LoginTForm.TButton2.SetEnabled(false)

		if w.LoginTForm.TEdit1.Text() == "" || w.LoginTForm.TEdit2.Text() == "" {
			vcl.ShowMessage("账号或密码不能为空！")
			w.LoginTForm.WinLogin.SetEnabled(true)
			w.LoginTForm.TButton1.SetEnabled(true)
			w.LoginTForm.TButton2.SetEnabled(true)
			return
		}

		bady := url.Values{}
		bady.Add("Login_Name", w.LoginTForm.TEdit1.Text())
		bady.Add("Password", w.LoginTForm.TEdit2.Text())

		resp, err := http.PostForm(Url, bady)
		if err != nil {
			vcl.ShowMessage(fmt.Sprintf("%v：请联系管理员！", err))
			w.LoginTForm.WinLogin.SetEnabled(true)
			w.LoginTForm.TButton1.SetEnabled(true)
			w.LoginTForm.TButton2.SetEnabled(true)
			return
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)

		Stat := new(Statu)
		err = json.Unmarshal(data, Stat)
		if err != nil {
			vcl.ShowMessage(fmt.Sprintf("%v：请联系管理员！", err))
			w.LoginTForm.WinLogin.SetEnabled(true)
			w.LoginTForm.TButton1.SetEnabled(true)
			w.LoginTForm.TButton2.SetEnabled(true)
			return
		}

		if Stat.Status {

			w.LoginName = w.LoginTForm.TEdit1.Text()

			bady := url.Values{}
			bady.Add("Client_Models", "Client_Models")
			bady.Add("Login_Name", w.LoginTForm.TEdit1.Text())
			resp, err := http.PostForm(Url1, bady)
			if err != nil {
				vcl.ShowMessage(fmt.Sprintf("%v", err))
			}
			defer resp.Body.Close()
			data, err := ioutil.ReadAll(resp.Body)

			rds := Users{}

			err = json.Unmarshal(data, &rds)
			if err != nil {
				vcl.ShowMessage(fmt.Sprintf("%v", err))
			}

			w.Client_Models = rds.Client_Models

			w.Win.Show()
			w.LoginTForm.WinLogin.Hide()
			w.listView.Clear()
			w.Win.SetEnabled(false)
			w.Button.SetEnabled(false)
			go func() { Call_log_statistics(w) }()
			return
		} else {
			vcl.ShowMessage(fmt.Sprintf("%v", Stat.Err))
			w.LoginTForm.WinLogin.SetEnabled(true)
			w.LoginTForm.TButton1.SetEnabled(true)
			w.LoginTForm.TButton2.SetEnabled(true)
			return
		}

	})

	w.LoginTForm.TEdit2.SetOnKeyDown(func(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
		if *key == 13 {
			w.LoginTForm.TButton1.Click()
		}
	})

	w.LoginTForm.TEdit1.SetOnKeyDown(func(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
		if *key == 13 {
			w.LoginTForm.TButton1.Click()
		}
	})

	w.LoginTForm.WinLogin.SetOnKeyDown(func(sender vcl.IObject, key *types.Char, shift types.TShiftState) {
		if *key == 13 {
			w.LoginTForm.TButton1.Click()
		}
	})

	//取消按钮退出
	w.LoginTForm.TButton2.SetOnClick(func(sender vcl.IObject) {
		os.Exit(0)
	})

	w.LoginTForm.WinLogin.SetOnClose(func(sender vcl.IObject, action *types.TCloseAction) {
		os.Exit(0)
	})
}
