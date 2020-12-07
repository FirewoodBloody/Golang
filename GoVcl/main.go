package main

import (
	"fmt"
	"github.com/FirewoodBloody/Golang/Express_Routing/express"
	"github.com/FirewoodBloody/Golang/express_api/modules"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"strconv"
	"strings"
)

var (
	getExp map[string]string
)

type TForm struct {
	win      *vcl.TForm
	wins     *vcl.TForm
	combobox *vcl.TComboBox
	button   *vcl.TButton
	memo     *vcl.TMemo
	grid     *vcl.TStringGrid
	grids    *vcl.TStringGrid
	listView *vcl.TListView
}

//初始化Tform
func (t *TForm) Init() {
	vcl.Application.Initialize()         //初始化环境
	t.win = vcl.Application.CreateForm() //新建窗口
	t.win.SetCaption("快递查询")             //程序名
	t.win.SetHeight(668)                 //高
	t.win.SetWidth(1160)                 //宽
	t.win.ScreenCenter()                 //居于当前屏幕中心
	t.win.SetBorderIcons(3)              //设置窗口最大化 最小化 关闭按钮  3代表是 最大化按钮和功能无效
	t.win.Font().SetSize(11)             //整体字体大小
	t.win.Font().SetStyle(16)            //字体样式

	t.combobox = vcl.NewComboBox(t.win) //新建下拉选项
	t.combobox.SetParent(t.win)         //指定父窗口
	t.combobox.SetHeight(30)            //设置高度
	t.combobox.SetWidth(150)            //设置宽度
	t.combobox.SetLeft(50)              //设置按钮位置  横向
	t.combobox.SetTop(60)               //设置按钮位置 竖向
	//设置下拉选项内容
	t.combobox.AddItem("顺丰快递", nil)
	t.combobox.AddItem("京东快递", nil)
	t.combobox.AddItem("圆通快递", nil)
	t.combobox.AddItem("中通快递", nil)
	t.combobox.AddItem("邮政快递", nil)
	t.combobox.AddItem("申通快递", nil)
	t.combobox.AddItem("天天快递", nil)
	t.combobox.AddItem("EMS", nil)

	t.wins = vcl.Application.CreateForm() //新建窗口  用于进行窗口弹出
	t.wins.SetCaption("快递路由详情")           //指定窗口名称
	t.wins.SetHeight(768)                 //高度
	t.wins.SetWidth(920)                  //跨密度
	t.wins.ScreenCenter()                 //居于当前屏幕中心
	t.wins.SetBorderIcons(1)              //设置窗口最大化 最小化 关闭按钮  1代表是 只有关闭按钮生效
	t.wins.Font().SetSize(10)             //设置窗口字体大小
	t.wins.Font().SetStyle(16)            //设置窗口字体样式

	t.button = vcl.NewButton(t.win) //新建按钮
	t.button.SetLeft(150)           //设置按钮位置  横向
	t.button.SetTop(450)            //设置按钮位置 竖向
	t.button.SetHeight(40)          //按钮的高度
	t.button.SetWidth(50)           //按钮的宽度
	t.button.SetCaption("查询")       //按钮显示文字  名称
	t.button.SetParent(t.win)       //设置父容器

	t.memo = vcl.NewMemo(t.win) //新建多行文本框
	t.memo.SetParent(t.win)     //设置父容器
	t.memo.SetWidth(150)        //文本的宽
	t.memo.SetHeight(300)       //文本的高
	t.memo.SetLeft(50)          //设置位置  横向
	t.memo.SetTop(100)          //设置位置 竖向
	t.memo.Font().SetStyle(16)

	//view自动列表
	t.listView = vcl.NewListView(t.win)
	t.listView.SetParent(t.win)
	//t.listView.SetAlign(types.AlTop)
	t.listView.SetRowSelect(true)
	t.listView.SetReadOnly(false)
	t.listView.SetViewStyle(types.VsReport)
	t.listView.SetGridLines(true)
	t.listView.SetWidth(900)
	t.listView.SetHeight(500)
	t.listView.SetLeft(220)
	t.listView.SetTop(50)
	t.listView.Font().SetStyle(16)
	t.listView.Font().SetSize(11)

	lvl := t.listView.Columns().Add()
	lvl.SetCaption("序号")
	lvl.SetWidth(50)

	lvl = t.listView.Columns().Add()
	lvl.SetCaption("快递单号")
	lvl.SetWidth(100)

	lvl = t.listView.Columns().Add()
	lvl.SetCaption("当前所在地")
	lvl.SetWidth(100)

	lvl = t.listView.Columns().Add()
	lvl.SetCaption("动态时间")
	lvl.SetWidth(130)

	lvl = t.listView.Columns().Add()
	lvl.SetCaption("快递状态")
	lvl.SetWidth(500)

	t.grid = vcl.NewStringGrid(t.win) //新建自动列表
	t.grid.SetParent(t.win)           //设置父容器
	t.grid.SetLeft(220)               //设置位置  横向
	t.grid.SetTop(50)                 //设置位置 竖向
	t.grid.SetHeight(550)             //列表高度
	t.grid.SetWidth(900)              //列表宽度
	t.grid.SetColCount(5)             //设置列的个数
	t.grid.SetRowCount(2)             //设置行的个数
	t.grid.Font().SetSize(9)          //设置自动列表字体大小
	t.grid.Font().SetStyle(16)        //设置自动列表字体样式

	//设置列宽
	t.grid.SetColWidths(0, 30) //设置每列的宽度  单位是PX像素
	t.grid.SetColWidths(1, 100)
	t.grid.SetColWidths(2, 100)
	t.grid.SetColWidths(3, 130)
	t.grid.SetColWidths(4, 500)

	//设置标题名
	t.grid.SetCells(0, int32(0), "序号")
	t.grid.SetCells(1, int32(0), "快递单号")
	t.grid.SetCells(2, int32(0), "当前所在地")
	t.grid.SetCells(3, int32(0), "动态时间")
	t.grid.SetCells(4, int32(0), "快递状态")

	t.grids = vcl.NewStringGrid(t.wins)
	t.grids.SetParent(t.wins)
	t.grids.SetLeft(2)     //设置位置  横向
	t.grids.SetTop(1)      //设置位置 竖向
	t.grids.SetHeight(720) //列表高度
	t.grids.SetWidth(900)  //列表宽度
	t.grids.SetColWidths(0, 30)
	t.grids.SetColWidths(1, 100)
	t.grids.SetColWidths(2, 100)
	t.grids.SetColWidths(3, 130)
	t.grids.SetColWidths(4, 500)
	t.grids.SetCells(0, int32(0), "序号")
	t.grids.SetCells(1, int32(0), "快递单号")
	t.grids.SetCells(2, int32(0), "当前所在地")
	t.grids.SetCells(3, int32(0), "动态时间")
	t.grids.SetCells(4, int32(0), "快递状态")
	t.grids.SetColCount(5) //设置列的个数
	t.grids.Font().SetSize(9)
	t.grids.Font().SetStyle(16)
}

//事件处理
func (t *TForm) incident() {

	//自动列表内鼠标双击事件处理
	//t.grid.SetOnDblClick(func(sender vcl.IObject) {
	//
	//
	//})
	//自动列表双击事件
	t.listView.SetOnDblClick(func(sender vcl.IObject) {

		if t.listView.ItemIndex() == -1 {
			return
		}
		str := t.listView.Selected().SubItems().Strings(0)

		if t.combobox.Text() == "顺丰快递" {
			datas, _ := express.SfCreateData(str)
			if len(datas.Body.RouteResponse.Route) == 0 {
				return
			}

			t.grids.SetRowCount(int32(len(datas.Body.RouteResponse.Route) + 1)) //设置行的个数

			for k, v := range datas.Body.RouteResponse.Route {
				t.grids.SetCells(0, int32(k+1), strconv.Itoa(k+1))
				t.grids.SetCells(1, int32(k+1), datas.Body.RouteResponse.Mailno)
				t.grids.SetCells(2, int32(k+1), v.Accept_Address)
				t.grids.SetCells(3, int32(k+1), v.Accept_Time)
				t.grids.SetCells(4, int32(k+1), v.Remark)
			}
		} else if t.combobox.Text() == "京东快递" {
			datas := modules.SelectData(str)
			if len(datas.Data) == 0 {
				return
			}

			t.grids.SetRowCount(int32(len(datas.Data) + 1)) //设置行的个数
			for k, v := range datas.Data {
				t.grids.SetCells(0, int32(k+1), strconv.Itoa(k+1))
				t.grids.SetCells(1, int32(k+1), v.WaybillCode)
				t.grids.SetCells(2, int32(k+1), v.OpeTitle)
				t.grids.SetCells(3, int32(k+1), v.OpeTime)
				t.grids.SetCells(4, int32(k+1), v.OpeRemark)
			}

		} else {
			datas, _ := express.KdnExpressInformation(getExp[t.combobox.Text()], str)
			if len(datas.Traces) == 0 {
				return
			}

			t.grids.SetRowCount(int32(len(datas.Traces) + 1)) //设置行的个数

			for k, v := range datas.Traces {
				t.grids.SetCells(0, int32(k+1), strconv.Itoa(k+1))
				t.grids.SetCells(1, int32(k+1), datas.LogisticCode)
				t.grids.SetCells(2, int32(k+1), "")
				t.grids.SetCells(3, int32(k+1), v.AcceptTime)
				t.grids.SetCells(4, int32(k+1), v.AcceptStation)
			}
		}

		//显示弹出的窗口
		t.wins.Show()
	})

	//查询按钮事件
	t.button.SetOnClick(func(sender vcl.IObject) {
		t.listView.Clear()
		data := strings.Split(t.memo.Text(), "\r\n")

		types := t.combobox.Text()
		if types != "顺丰快递" && types != "京东快递" {
			types = getExp[types]
		}
		//grid.SetCells(0, int32(0), "")
		//grid.SetCells(1, int32(0), "")
		//grid.SetCells(2, int32(0), "")
		//grid.SetCells(3, int32(0), "")
		//grid.SetCells(4, int32(0), "")
		//t.grid.SetRowCount(int32(len(data) + 1)) //重设行数
		//t.listView.Items().Update()
		//t.listView.Items().EndUpdate()
		for i, v := range data {
			if v == "\t" || v == " " || v == "/" || v == "\t\r" || v == "" {
				continue
			}

			str, err := IFStr(v, types)

			if err != nil {
				vcl.ShowMessage(fmt.Sprintf("%v - NO(%v)：%v", v, i+1, err))
				return
			}

			go func(i int) {

				if types == "顺丰快递" {

					datas, _ := modules.SfCreateData(str)

					if len(datas.Body.RouteResponse.Route) == 0 {
						vcl.ThreadSync(func() {
							item := t.listView.Items().Add()
							item.SetImageIndex(0)
							item.SetCaption(fmt.Sprintf("%3d", i))
							item.SubItems().Add(str)
							item.SubItems().Add("")
							item.SubItems().Add("")
							item.SubItems().Add("查询失败（无路由信息）或单号错误！")
						})
					} else {
						vcl.ThreadSync(func() {
							item := t.listView.Items().Add()
							item.SetImageIndex(0)
							item.SetCaption(fmt.Sprintf("%3d", i))
							item.SubItems().Add(datas.Body.RouteResponse.Mailno)
							item.SubItems().Add(datas.Body.RouteResponse.Route[len(datas.Body.RouteResponse.Route)-1].Accept_Address)
							item.SubItems().Add(datas.Body.RouteResponse.Route[len(datas.Body.RouteResponse.Route)-1].Accept_Time)
							item.SubItems().Add(datas.Body.RouteResponse.Route[len(datas.Body.RouteResponse.Route)-1].Remark)
						})
					}
				} else if types == "京东快递" {
					datas := modules.SelectData(str)

					if len(datas.Data) == 0 {
						vcl.ThreadSync(func() {
							item := t.listView.Items().Add()
							item.SetImageIndex(0)
							item.SetCaption(fmt.Sprintf("%3d", i))
							item.SubItems().Add(str)
							item.SubItems().Add("")
							item.SubItems().Add("")
							item.SubItems().Add("查询失败（无路由信息）或单号错误！")
						})
					} else {
						vcl.ThreadSync(func() {
							item := t.listView.Items().Add()
							item.SetImageIndex(0)
							item.SetCaption(fmt.Sprintf("%3d", i))
							item.SubItems().Add(str)
							item.SubItems().Add(datas.Data[len(datas.Data)-1].OpeTitle)
							item.SubItems().Add(datas.Data[len(datas.Data)-1].OpeTime)
							item.SubItems().Add(datas.Data[len(datas.Data)-1].OpeRemark)
						})
					}
				} else {
					datas, _ := express.KdnExpressInformation(types, str)

					if len(datas.Traces) == 0 {
						vcl.ThreadSync(func() {
							item := t.listView.Items().Add()
							item.SetImageIndex(0)
							item.SetCaption(fmt.Sprintf("%3d", i))
							item.SubItems().Add(str)
							item.SubItems().Add("")
							item.SubItems().Add("")
							item.SubItems().Add("查询失败（无路由信息）或单号错误！")
						})
					} else {
						vcl.ThreadSync(func() {
							item := t.listView.Items().Add()
							item.SetImageIndex(0)
							item.SetCaption(fmt.Sprintf("%3d", i))
							item.SubItems().Add(datas.LogisticCode)
							item.SubItems().Add("")
							item.SubItems().Add(datas.Traces[len(datas.Traces)-1].AcceptTime)
							item.SubItems().Add(datas.Traces[len(datas.Traces)-1].AcceptStation)
						})
					}
				}
			}(i + 1)
		}
		//t.listView.Items().EndUpdate()

	})

	//自动列表排序事件
	t.listView.SetOnColumnClick(func(sender vcl.IObject, column *vcl.TListColumn) {
		t.listView.CustomSort(0, int(column.Index()))

	})

	//自定义排序方法
	t.listView.SetOnCompare(func(sender vcl.IObject, item1, item2 *vcl.TListItem, data int32, compare *int32) {
		if data == 0 {
			*compare = int32(strings.Compare(item1.Caption(), item2.Caption()))
		} else {
			*compare = int32(strings.Compare(item1.SubItems().Strings(data-1), item2.SubItems().Strings(data-1)))
		}
	})
}

func IFStr(str, types string) (string, error) {
	if str == "" {
		return "", fmt.Errorf("str is nil!")
	}

	if types == "顺丰快递" {

		if len(strings.Replace(str, "\t", "", -1)) != len(str) {
			str = strings.Replace(str, "\t", "", -1)
		} else if len(strings.Replace(str, " ", "", -1)) != len(str) {
			str = strings.Replace(str, " ", "", -1)
		} else if len(strings.Replace(str, "/", "", -1)) != len(str) {
			str = strings.Replace(str, "/", "", -1)
		}

	} else {
		if len(strings.Replace(str, "\t", "", -1)) != len(str) {
			str = strings.Replace(str, "\t", "", -1)
		} else if len(strings.Replace(str, " ", "", -1)) != len(str) {
			str = strings.Replace(str, " ", "", -1)
		} else if len(strings.Replace(str, "/", "", -1)) != len(str) {
			str = strings.Replace(str, "/", "", -1)
		}
	}
	return strings.ToUpper(str), nil
}

func main() {

	getExp = make(map[string]string, 10)
	getExp["圆通快递"] = "YTO"
	getExp["中通快递"] = "ZTO"
	getExp["申通快递"] = "STO"
	getExp["天天快递"] = "HHTT"
	getExp["EMS"] = "EMS"
	getExp["邮政快递"] = "YZPY"

	Form := new(TForm)
	Form.Init()
	Form.incident()

	Form.win.Show()
	Form.grid.Hide()
	vcl.Application.Run()
}
