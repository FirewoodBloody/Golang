package module

import (
	"encoding/json"
	"fmt"
	"github.com/FirewoodBloody/Golang/Express_Routing/express"
	"github.com/FirewoodBloody/Golang/express_api/modules"
	"github.com/Luxurioust/excelize"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ying32/govcl/vcl"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var (
	URL  = "http://192.168.0.12:8888/v1/object/"            //请求数据
	Url  = "http://192.168.0.12:8888/v1/user/"              //数据报表类型
	Url2 = "http://192.168.0.12:8888/v1/user/Version"       //版本
	Url1 = "http://192.168.0.12:8888/v1/user/Client_Models" //版本 和 数据报表类型
)

//销售订单明细
func Order_select(w *Windows, f *excelize.File) {
	bady := url.Values{}
	bady.Add("Mysql_Select", w.Taskdlg.RadioButton().Caption())
	bady.Add("Start_Time", w.Date_strat_label.DateTime().Format("2006-01-02"))
	bady.Add("Stop_Time", w.Date_stop_label.DateTime().Format("2006-01-02"))
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
	time.Sleep(time.Second * 5)

	for k, v := range resulist {
		if k == 0 {
			f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+1), "订单创建时间")
			f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+1), "订单创建人")
			f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), "订单创建人工号")
			f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+1), "订单创建人部门")
			f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+1), "客户姓名")
			f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+1), "客户编码")
			f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+1), "系统订单号")
			f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+1), "订单类型")
			f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+1), "订单渠道")
			f.SetCellValue("Sheet1", fmt.Sprintf("J%v", k+1), "订单状态")
			f.SetCellValue("Sheet1", fmt.Sprintf("K%v", k+1), "回款日期")
			f.SetCellValue("Sheet1", fmt.Sprintf("L%v", k+1), "线上下单渠道")
			f.SetCellValue("Sheet1", fmt.Sprintf("M%v", k+1), "线上下单时间")
			f.SetCellValue("Sheet1", fmt.Sprintf("N%v", k+1), "线上下单核单人")
			f.SetCellValue("Sheet1", fmt.Sprintf("O%v", k+1), "线上下单核单人工号")
			f.SetCellValue("Sheet1", fmt.Sprintf("P%v", k+1), "线上下单核单人部门")
			f.SetCellValue("Sheet1", fmt.Sprintf("Q%v", k+1), "订单备注")
			f.SetCellValue("Sheet1", fmt.Sprintf("R%v", k+1), "业绩归属人")
			f.SetCellValue("Sheet1", fmt.Sprintf("S%v", k+1), "业绩归属人工号")
			f.SetCellValue("Sheet1", fmt.Sprintf("T%v", k+1), "员工状态")
			f.SetCellValue("Sheet1", fmt.Sprintf("U%v", k+1), "业绩归属人部门")
			f.SetCellValue("Sheet1", fmt.Sprintf("V%v", k+1), "订单总金额")
			f.SetCellValue("Sheet1", fmt.Sprintf("W%v", k+1), "订单折扣总额")
			f.SetCellValue("Sheet1", fmt.Sprintf("X%v", k+1), "订单售价总额")
			f.SetCellValue("Sheet1", fmt.Sprintf("Y%v", k+1), "订单商品")
			f.SetCellValue("Sheet1", fmt.Sprintf("Z%v", k+1), "商品数量")
			f.SetCellValue("Sheet1", fmt.Sprintf("AA%v", k+1), "商品原价")
			f.SetCellValue("Sheet1", fmt.Sprintf("AB%v", k+1), "商品售价")
			f.SetCellValue("Sheet1", fmt.Sprintf("AC%v", k+1), "商品总价")
			f.SetCellValue("Sheet1", fmt.Sprintf("AD%v", k+1), "配送方式")
			f.SetCellValue("Sheet1", fmt.Sprintf("AE%v", k+1), "快递单号")
			f.SetCellValue("Sheet1", fmt.Sprintf("AF%v", k+1), "快递状态")
		}
		f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+2), v["订单创建时间"])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+2), v["订单创建人"])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+2), v["订单创建人工号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+2), v["订单创建人部门"])
		f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+2), v["客户姓名"])
		f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+2), v["客户编码"])
		f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+2), v["系统订单号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+2), v["订单类型"])
		f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+2), v["订单渠道"])

		f.SetCellValue("Sheet1", fmt.Sprintf("J%v", k+2), v["订单状态"])
		f.SetCellValue("Sheet1", fmt.Sprintf("K%v", k+2), v["回款日期"])
		f.SetCellValue("Sheet1", fmt.Sprintf("L%v", k+2), v["线上下单渠道"])
		f.SetCellValue("Sheet1", fmt.Sprintf("M%v", k+2), v["线上下单时间"])
		f.SetCellValue("Sheet1", fmt.Sprintf("N%v", k+2), v["线上下单核单人"])
		f.SetCellValue("Sheet1", fmt.Sprintf("O%v", k+2), v["线上下单核单人工号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("P%v", k+2), v["线上下单核单人部门"])
		f.SetCellValue("Sheet1", fmt.Sprintf("Q%v", k+2), v["订单备注"])
		f.SetCellValue("Sheet1", fmt.Sprintf("R%v", k+2), v["业绩归属人"])
		f.SetCellValue("Sheet1", fmt.Sprintf("S%v", k+2), v["业绩归属人工号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("T%v", k+2), v["员工状态"])
		f.SetCellValue("Sheet1", fmt.Sprintf("U%v", k+2), v["业绩归属人部门"])
		f.SetCellValue("Sheet1", fmt.Sprintf("V%v", k+2), v["订单总金额"])
		f.SetCellValue("Sheet1", fmt.Sprintf("W%v", k+2), v["订单折扣总额"])
		f.SetCellValue("Sheet1", fmt.Sprintf("X%v", k+2), v["订单售价总额"])
		f.SetCellValue("Sheet1", fmt.Sprintf("Y%v", k+2), v["订单商品"])
		f.SetCellValue("Sheet1", fmt.Sprintf("Z%v", k+2), v["商品数量"])
		f.SetCellValue("Sheet1", fmt.Sprintf("AA%v", k+2), v["商品原价"])
		f.SetCellValue("Sheet1", fmt.Sprintf("AB%v", k+2), v["商品售价"])
		f.SetCellValue("Sheet1", fmt.Sprintf("AC%v", k+2), v["商品总价"])
		f.SetCellValue("Sheet1", fmt.Sprintf("AD%v", k+2), v["配送方式"])
		f.SetCellValue("Sheet1", fmt.Sprintf("AE%v", k+2), v["快递单号"])
		if string(v["订单状态"]) == "订单配送中" || string(v["订单状态"]) == "订单未妥投" {
			if string(v["快递单号"]) != "" {
				if string(v["配送方式"]) == "顺丰快递" {
					data, _ := express.SfCreateData(string(v["快递单号"]))
					if len(data.Body.RouteResponse.Route) != 0 {
						//fmt.Println(data.Body.RouteResponse.Route[len(data.Body.RouteResponse.Route)-1].Remark)
						f.SetCellValue("Sheet1", fmt.Sprintf("AF%v", k+2), data.Body.RouteResponse.Route[len(data.Body.RouteResponse.Route)-1].Remark)
					}
				} else if string(v["配送方式"]) == "京东快递" {
					data := modules.SelectData(string(v["快递单号"]))
					if len(data.Data) != 0 {
						//fmt.Println(data.Data[len(data.Data)-1].OpeRemark)
						f.SetCellValue("Sheet1", fmt.Sprintf("AF%v", k+2), data.Data[len(data.Data)-1].OpeRemark)
					}
				}
			}
		} else {
			f.SetCellValue("Sheet1", fmt.Sprintf("AF%v", k+2), v["快递状态"])
		}
		vcl.ThreadSync(func() {
			w.Progress_bar.SetPosition(int32(float64(k+1) / float64(len(resulist)) * 100))
		})
	}

	vcl.ThreadSync(func() {
		w.Progress_bar.SetPosition(100)
		f.SaveAs(w.TsaveDialog.FileName())
		time.Sleep(time.Second * 1)
		vcl.ShowMessageFmt("保存完成！")
	})

	w.Button.SetEnabled(true)
	w.Win.SetEnabled(true)
	//os.Exit(0)
}

//新媒体线上明细
func Order_xs_select(w *Windows, f *excelize.File) {
	bady := url.Values{}
	bady.Add("Mysql_Select", w.Taskdlg.RadioButton().Caption())
	bady.Add("Start_Time", w.Date_strat_label.DateTime().Format("2006-01-02"))
	bady.Add("Stop_Time", w.Date_stop_label.DateTime().Format("2006-01-02"))
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
		if k == 0 {
			f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+1), "线上下单时间")
			f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+1), "线上下单渠道")
			f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), "客户编码")
			f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+1), "下单客户姓名")
			f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+1), "下单商品")
			f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+1), "渠道订单号")
			f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+1), "核单状态")
			f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+1), "核单备注")
			f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+1), "部门")
			f.SetCellValue("Sheet1", fmt.Sprintf("J%v", k+1), "员工姓名")
			f.SetCellValue("Sheet1", fmt.Sprintf("K%v", k+1), "员工工号")
			f.SetCellValue("Sheet1", fmt.Sprintf("L%v", k+1), "员工状态")
			f.SetCellValue("Sheet1", fmt.Sprintf("M%v", k+1), "系统订单号")
			f.SetCellValue("Sheet1", fmt.Sprintf("N%v", k+1), "订单类型")
			f.SetCellValue("Sheet1", fmt.Sprintf("O%v", k+1), "订单状态")
			f.SetCellValue("Sheet1", fmt.Sprintf("P%v", k+1), "订单总金额")
			f.SetCellValue("Sheet1", fmt.Sprintf("Q%v", k+1), "订单折扣总额")
			f.SetCellValue("Sheet1", fmt.Sprintf("R%v", k+1), "订单售价总额")
			f.SetCellValue("Sheet1", fmt.Sprintf("S%v", k+1), "订单商品")
			f.SetCellValue("Sheet1", fmt.Sprintf("T%v", k+1), "商品数量")
			f.SetCellValue("Sheet1", fmt.Sprintf("U%v", k+1), "商品原价")
			f.SetCellValue("Sheet1", fmt.Sprintf("V%v", k+1), "商品售价")
			f.SetCellValue("Sheet1", fmt.Sprintf("W%v", k+1), "商品总价")
			f.SetCellValue("Sheet1", fmt.Sprintf("X%v", k+1), "配送方式")
			f.SetCellValue("Sheet1", fmt.Sprintf("Y%v", k+1), "快递单号")
			f.SetCellValue("Sheet1", fmt.Sprintf("X%v", k+1), "快递状态")
			f.SetCellValue("Sheet1", fmt.Sprintf("AA%v", k+1), "客户地址")
		}
		f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+2), v["线上下单时间"])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+2), v["线上下单渠道"])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+2), v["客户编码"])
		f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+2), v["下单客户姓名"])
		f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+2), v["下单商品"])
		f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+2), v["渠道订单号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+2), v["核单状态"])
		f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+2), v["核单备注"])
		f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+2), v["部门"])
		f.SetCellValue("Sheet1", fmt.Sprintf("J%v", k+2), v["员工姓名"])
		f.SetCellValue("Sheet1", fmt.Sprintf("K%v", k+2), v["员工工号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("L%v", k+2), v["员工状态"])
		f.SetCellValue("Sheet1", fmt.Sprintf("M%v", k+2), v["系统订单号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("N%v", k+2), v["订单类型"])
		f.SetCellValue("Sheet1", fmt.Sprintf("O%v", k+2), v["订单状态"])
		f.SetCellValue("Sheet1", fmt.Sprintf("P%v", k+2), v["订单总金额"])
		f.SetCellValue("Sheet1", fmt.Sprintf("Q%v", k+2), v["订单折扣总额"])
		f.SetCellValue("Sheet1", fmt.Sprintf("R%v", k+2), v["订单售价总额"])
		f.SetCellValue("Sheet1", fmt.Sprintf("S%v", k+2), v["订单商品"])
		f.SetCellValue("Sheet1", fmt.Sprintf("T%v", k+2), v["商品数量"])
		f.SetCellValue("Sheet1", fmt.Sprintf("U%v", k+2), v["商品原价"])
		f.SetCellValue("Sheet1", fmt.Sprintf("V%v", k+2), v["商品售价"])
		f.SetCellValue("Sheet1", fmt.Sprintf("W%v", k+2), v["商品总价"])
		f.SetCellValue("Sheet1", fmt.Sprintf("X%v", k+2), v["配送方式"])
		f.SetCellValue("Sheet1", fmt.Sprintf("Y%v", k+2), v["快递单号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("Z%v", k+2), v["客户地址"])
		if string(v["订单状态"]) == "订单配送中" || string(v["订单状态"]) == "订单未妥投" {
			if string(v["快递单号"]) != "" {
				if string(v["配送方式"]) == "顺丰快递" {
					data, _ := express.SfCreateData(string(v["快递单号"]))
					if len(data.Body.RouteResponse.Route) != 0 {
						f.SetCellValue("Sheet1", fmt.Sprintf("AA%v", k+2), data.Body.RouteResponse.Route[len(data.Body.RouteResponse.Route)-1].Remark)
					}
				} else if string(v["配送方式"]) == "京东快递" {
					data := modules.SelectData(string(v["快递单号"]))
					if len(data.Data) != 0 {
						f.SetCellValue("Sheet1", fmt.Sprintf("AA%v", k+2), data.Data[len(data.Data)-1].OpeRemark)
					}
				}
			}
		} else {
			f.SetCellValue("Sheet1", fmt.Sprintf("AA%v", k+2), v["快递状态"])
		}
		vcl.ThreadSync(func() {
			w.Progress_bar.SetPosition(int32(float64(k+1) / float64(len(resulist)) * 100))
		})
	}
	vcl.ThreadSync(func() {
		w.Progress_bar.SetPosition(100)
		f.SaveAs(w.TsaveDialog.FileName())
		time.Sleep(time.Second * 1)
		vcl.ShowMessageFmt("保存完成！")
	})

	w.Button.SetEnabled(true)
	w.Win.SetEnabled(true)
	//os.Exit(0)
}

//新媒体线上统计
func Order_xsTJ_select(w *Windows, f *excelize.File) {
	bady := url.Values{}
	bady.Add("Mysql_Select", w.Taskdlg.RadioButton().Caption())
	bady.Add("Start_Time", w.Date_strat_label.DateTime().Format("2006-01-02"))
	bady.Add("Stop_Time", w.Date_stop_label.DateTime().Format("2006-01-02"))
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

	time.Sleep(time.Second * 5)

	for k, v := range resulist {
		if k == 0 {
			f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+1), "订单渠道")
			f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+1), "投放产品")
			f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), "后台订单数")
			f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+1), "发货数量")
			f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+1), "原订单总金额")
			f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+1), "成交金额")

		}
		f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+2), v["订单渠道"])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+2), v["投放产品"])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+2), v["后台订单数"])
		f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+2), v["发货数量"])
		f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+2), v["原订单总金额"])
		f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+2), v["成交金额"])
		vcl.ThreadSync(func() {
			w.Progress_bar.SetPosition(int32(float64(k+1) / float64(len(resulist)) * 100))
		})
	}
	vcl.ThreadSync(func() {
		w.Progress_bar.SetPosition(100)
		f.SaveAs(w.TsaveDialog.FileName())
		time.Sleep(time.Second * 1)
		vcl.ShowMessageFmt("保存完成！")
	})

	w.Button.SetEnabled(true)
	w.Win.SetEnabled(true)
	//os.Exit(0)
}

//新媒体线下统计
func Order_xxTJ_select(w *Windows, f *excelize.File) {

	bady := url.Values{}
	bady.Add("Mysql_Select", w.Taskdlg.RadioButton().Caption())
	bady.Add("Start_Time", w.Date_strat_label.DateTime().Format("2006-01-02"))
	bady.Add("Stop_Time", w.Date_stop_label.DateTime().Format("2006-01-02"))
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
	time.Sleep(time.Second * 5)

	for k, v := range resulist {
		if k == 0 {
			f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+1), "二开产品")
			f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+1), "下单数")
			f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), "增购订单总金额")
			f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+1), "增购成交金额")
		}
		f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+2), v["二开产品"])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+2), v["下单数"])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+2), v["增购订单总金额"])
		f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+2), v["增购成交金额"])
		vcl.ThreadSync(func() {
			w.Progress_bar.SetPosition(int32(float64(k+1) / float64(len(resulist)) * 100))
		})
	}
	vcl.ThreadSync(func() {
		w.Progress_bar.SetPosition(100)
		f.SaveAs(w.TsaveDialog.FileName())
		time.Sleep(time.Second * 1)
		vcl.ShowMessageFmt("保存完成！")
	})

	w.Button.SetEnabled(true)
	w.Win.SetEnabled(true)
	//os.Exit(0)
}

//礼品订单明细
func Order_gift_select(w *Windows, f *excelize.File) {
	bady := url.Values{}
	bady.Add("Mysql_Select", w.Taskdlg.RadioButton().Caption())
	bady.Add("Start_Time", w.Date_strat_label.DateTime().Format("2006-01-02"))
	bady.Add("Stop_Time", w.Date_stop_label.DateTime().Format("2006-01-02"))
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
	time.Sleep(time.Second * 5)

	for k, v := range resulist {
		if k == 0 {
			f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+1), "订单创建时间")
			f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+1), "订单创建人")
			f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), "订单创建人工号")
			f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+1), "订单创建人部门")
			f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+1), "客户姓名")
			f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+1), "客户编码")
			f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+1), "系统订单号")
			f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+1), "订单类型")
			f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+1), "订单状态")
			f.SetCellValue("Sheet1", fmt.Sprintf("J%v", k+1), "回款日期")
			f.SetCellValue("Sheet1", fmt.Sprintf("K%v", k+1), "线上下单渠道")
			f.SetCellValue("Sheet1", fmt.Sprintf("L%v", k+1), "线上下单时间")
			f.SetCellValue("Sheet1", fmt.Sprintf("M%v", k+1), "线上下单核单人")
			f.SetCellValue("Sheet1", fmt.Sprintf("N%v", k+1), "线上下单核单人工号")
			f.SetCellValue("Sheet1", fmt.Sprintf("O%v", k+1), "线上下单核单人部门")
			f.SetCellValue("Sheet1", fmt.Sprintf("P%v", k+1), "订单备注")
			f.SetCellValue("Sheet1", fmt.Sprintf("Q%v", k+1), "业绩归属人")
			f.SetCellValue("Sheet1", fmt.Sprintf("R%v", k+1), "业绩归属人工号")
			f.SetCellValue("Sheet1", fmt.Sprintf("S%v", k+1), "员工状态")
			f.SetCellValue("Sheet1", fmt.Sprintf("T%v", k+1), "业绩归属人部门")
			f.SetCellValue("Sheet1", fmt.Sprintf("U%v", k+1), "订单总金额")
			f.SetCellValue("Sheet1", fmt.Sprintf("V%v", k+1), "订单折扣总额")
			f.SetCellValue("Sheet1", fmt.Sprintf("W%v", k+1), "订单售价总额")
			f.SetCellValue("Sheet1", fmt.Sprintf("X%v", k+1), "订单商品")
			f.SetCellValue("Sheet1", fmt.Sprintf("Y%v", k+1), "商品数量")
			f.SetCellValue("Sheet1", fmt.Sprintf("Z%v", k+1), "商品原价")
			f.SetCellValue("Sheet1", fmt.Sprintf("AA%v", k+1), "商品售价")
			f.SetCellValue("Sheet1", fmt.Sprintf("AB%v", k+1), "商品总价")
			f.SetCellValue("Sheet1", fmt.Sprintf("AC%v", k+1), "配送方式")
			f.SetCellValue("Sheet1", fmt.Sprintf("AD%v", k+1), "快递单号")
			f.SetCellValue("Sheet1", fmt.Sprintf("AE%v", k+1), "快递状态")
		}
		f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+2), v["订单创建时间"])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+2), v["订单创建人"])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+2), v["订单创建人工号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+2), v["订单创建人部门"])
		f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+2), v["客户姓名"])
		f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+2), v["客户编码"])
		f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+2), v["系统订单号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+2), v["订单类型"])
		f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+2), v["订单状态"])
		f.SetCellValue("Sheet1", fmt.Sprintf("J%v", k+2), v["回款日期"])
		f.SetCellValue("Sheet1", fmt.Sprintf("K%v", k+2), v["线上下单渠道"])
		f.SetCellValue("Sheet1", fmt.Sprintf("L%v", k+2), v["线上下单时间"])
		f.SetCellValue("Sheet1", fmt.Sprintf("M%v", k+2), v["线上下单核单人"])
		f.SetCellValue("Sheet1", fmt.Sprintf("N%v", k+2), v["线上下单核单人工号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("O%v", k+2), v["线上下单核单人部门"])
		f.SetCellValue("Sheet1", fmt.Sprintf("P%v", k+2), v["订单备注"])
		f.SetCellValue("Sheet1", fmt.Sprintf("Q%v", k+2), v["业绩归属人"])
		f.SetCellValue("Sheet1", fmt.Sprintf("R%v", k+2), v["业绩归属人工号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("S%v", k+2), v["员工状态"])
		f.SetCellValue("Sheet1", fmt.Sprintf("T%v", k+2), v["业绩归属人部门"])
		f.SetCellValue("Sheet1", fmt.Sprintf("U%v", k+2), v["订单总金额"])
		f.SetCellValue("Sheet1", fmt.Sprintf("V%v", k+2), v["订单折扣总额"])
		f.SetCellValue("Sheet1", fmt.Sprintf("W%v", k+2), v["订单售价总额"])
		f.SetCellValue("Sheet1", fmt.Sprintf("X%v", k+2), v["订单商品"])
		f.SetCellValue("Sheet1", fmt.Sprintf("Y%v", k+2), v["商品数量"])
		f.SetCellValue("Sheet1", fmt.Sprintf("Z%v", k+2), v["商品原价"])
		f.SetCellValue("Sheet1", fmt.Sprintf("AA%v", k+2), v["商品售价"])
		f.SetCellValue("Sheet1", fmt.Sprintf("AB%v", k+2), v["商品总价"])
		f.SetCellValue("Sheet1", fmt.Sprintf("AC%v", k+2), v["配送方式"])
		f.SetCellValue("Sheet1", fmt.Sprintf("AD%v", k+2), v["快递单号"])
		if string(v["订单状态"]) == "订单配送中" || string(v["订单状态"]) == "订单未妥投" {
			if string(v["快递单号"]) != "" {
				if string(v["配送方式"]) == "顺丰快递" {
					data, _ := express.SfCreateData(string(v["快递单号"]))
					if len(data.Body.RouteResponse.Route) != 0 {
						//fmt.Println(data.Body.RouteResponse.Route[len(data.Body.RouteResponse.Route)-1].Remark)
						f.SetCellValue("Sheet1", fmt.Sprintf("AE%v", k+2), data.Body.RouteResponse.Route[len(data.Body.RouteResponse.Route)-1].Remark)
					}
				} else if string(v["配送方式"]) == "京东快递" {
					data := modules.SelectData(string(v["快递单号"]))
					if len(data.Data) != 0 {
						//fmt.Println(data.Data[len(data.Data)-1].OpeRemark)
						f.SetCellValue("Sheet1", fmt.Sprintf("AE%v", k+2), data.Data[len(data.Data)-1].OpeRemark)
					}
				}
			}
		} else {
			f.SetCellValue("Sheet1", fmt.Sprintf("AE%v", k+2), v["快递状态"])
		}
		vcl.ThreadSync(func() {
			w.Progress_bar.SetPosition(int32(float64(k+1) / float64(len(resulist)) * 100))
		})
	}

	vcl.ThreadSync(func() {
		w.Progress_bar.SetPosition(100)
		f.SaveAs(w.TsaveDialog.FileName())
		time.Sleep(time.Second * 1)
		vcl.ShowMessageFmt("保存完成！")
	})

	w.Button.SetEnabled(true)
	w.Win.SetEnabled(true)
	//os.Exit(0)
}

//订单出库明细
func Order_warehouse_select(w *Windows, f *excelize.File) {
	bady := url.Values{}
	bady.Add("Mysql_Select", w.Taskdlg.RadioButton().Caption())
	bady.Add("Start_Time", w.Date_strat_label.DateTime().Format("2006-01-02"))
	bady.Add("Stop_Time", w.Date_stop_label.DateTime().Format("2006-01-02"))
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
	time.Sleep(time.Second * 5)

	for k, v := range resulist {
		if k == 0 {
			f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+1), "出库日期")
			f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+1), "订单编号")
			f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), "类型")
			f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+1), "订单类型")
			f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+1), "员工姓名")
			f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+1), "员工工号")
			f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+1), "员工状态")
			f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+1), "员工部门")
			f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+1), "商品名称")
			f.SetCellValue("Sheet1", fmt.Sprintf("J%v", k+1), "商品数量")
			f.SetCellValue("Sheet1", fmt.Sprintf("K%v", k+1), "商品售价")
			f.SetCellValue("Sheet1", fmt.Sprintf("L%v", k+1), "出货仓库")
			f.SetCellValue("Sheet1", fmt.Sprintf("M%v", k+1), "配送方式")
			f.SetCellValue("Sheet1", fmt.Sprintf("N%v", k+1), "快递单号")
		}
		f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+2), v["出库日期"])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+2), v["订单编号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+2), v["类型"])
		f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+2), v["订单类型"])
		f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+2), v["员工姓名"])
		f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+2), v["员工工号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+2), v["员工状态"])
		f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+2), v["员工部门"])
		f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+2), v["商品名称"])
		f.SetCellValue("Sheet1", fmt.Sprintf("J%v", k+2), v["商品数量"])
		f.SetCellValue("Sheet1", fmt.Sprintf("K%v", k+2), v["商品售价"])
		f.SetCellValue("Sheet1", fmt.Sprintf("L%v", k+2), v["出货仓库"])
		f.SetCellValue("Sheet1", fmt.Sprintf("M%v", k+2), v["配送方式"])
		f.SetCellValue("Sheet1", fmt.Sprintf("N%v", k+2), v["快递单号"])
		vcl.ThreadSync(func() {
			w.Progress_bar.SetPosition(int32(float64(k+1) / float64(len(resulist)) * 100))
		})
	}

	vcl.ThreadSync(func() {
		w.Progress_bar.SetPosition(100)
		f.SaveAs(w.TsaveDialog.FileName())
		time.Sleep(time.Second * 1)
		vcl.ShowMessageFmt("保存完成！")
	})

	w.Button.SetEnabled(true)
	w.Win.SetEnabled(true)
	//os.Exit(0)
}

//通话记录统计
func Call_log_statistics(w *Windows, f *excelize.File) {
	bady := url.Values{}
	bady.Add("Mysql_Select", w.Taskdlg.RadioButton().Caption())
	bady.Add("Start_Time", w.Date_strat_label.DateTime().Format("2006-01-02"))
	bady.Add("Stop_Time", w.Date_stop_label.DateTime().Format("2006-01-02"))
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
	time.Sleep(time.Second * 5)

	for k, v := range resulist {
		if k == 0 {
			f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+1), "部门")
			f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+1), "二级部门")
			f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), "员工姓名")
			f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+1), "员工工号")
			f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+1), "未接通话总数")
			f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+1), "未接通时长")
			f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+1), "无效通话总数")
			f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+1), "无效通话时长")
			f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+1), "有效通话总数")
			f.SetCellValue("Sheet1", fmt.Sprintf("J%v", k+1), "有效通话时长")
			f.SetCellValue("Sheet1", fmt.Sprintf("K%v", k+1), "优质通话总数")
			f.SetCellValue("Sheet1", fmt.Sprintf("L%v", k+1), "优质通话时长")
			f.SetCellValue("Sheet1", fmt.Sprintf("M%v", k+1), "总通话数")
			f.SetCellValue("Sheet1", fmt.Sprintf("N%v", k+1), "通话总时长")
		}
		f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+2), v["部门"])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+2), v["二级部门"])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+2), v["员工姓名"])
		f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+2), v["员工工号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+2), v["未接通话总数"])
		f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+2), v["未接通时长"])
		f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+2), v["无效通话总数"])
		f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+2), v["无效通话时长"])
		f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+2), v["有效通话总数"])
		f.SetCellValue("Sheet1", fmt.Sprintf("J%v", k+2), v["有效通话时长"])
		f.SetCellValue("Sheet1", fmt.Sprintf("K%v", k+2), v["优质通话总数"])
		f.SetCellValue("Sheet1", fmt.Sprintf("L%v", k+2), v["优质通话时长"])
		f.SetCellValue("Sheet1", fmt.Sprintf("M%v", k+2), v["总通话数"])
		f.SetCellValue("Sheet1", fmt.Sprintf("N%v", k+2), v["通话总时长"])
		vcl.ThreadSync(func() {
			w.Progress_bar.SetPosition(int32(float64(k+1) / float64(len(resulist)) * 100))
		})
	}

	vcl.ThreadSync(func() {
		w.Progress_bar.SetPosition(100)
		f.SaveAs(w.TsaveDialog.FileName())
		time.Sleep(time.Second * 1)
		vcl.ShowMessageFmt("保存完成！")
	})

	w.Button.SetEnabled(true)
	w.Win.SetEnabled(true)
	//os.Exit(0)
}

//客户当前积分
func Client_coin(w *Windows, f *excelize.File) {
	vcl.ShowMessage("查询数据为当前数据！")
	bady := url.Values{}
	bady.Add("Mysql_Select", w.Taskdlg.RadioButton().Caption())
	bady.Add("Start_Time", w.Date_strat_label.DateTime().Format("2006-01-02"))
	bady.Add("Stop_Time", w.Date_stop_label.DateTime().Format("2006-01-02"))
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
	time.Sleep(time.Second * 5)

	for k, v := range resulist {
		if k == 0 {
			f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+1), "部门")
			f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+1), "二级部门")
			f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), "员工姓名")
			f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+1), "员工工号")
			f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+1), "员工状态")
			f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+1), "客户编码")
			f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+1), "客户姓名")
			f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+1), "客户积分")

		}
		f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+2), v["部门"])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+2), v["二级部门"])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+2), v["员工姓名"])
		f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+2), v["员工工号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+2), v["员工状态"])
		f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+2), v["客户编码"])
		f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+2), v["客户姓名"])
		f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+2), v["客户积分"])

		vcl.ThreadSync(func() {
			w.Progress_bar.SetPosition(int32(float64(k+1) / float64(len(resulist)) * 100))
		})
	}

	vcl.ThreadSync(func() {
		w.Progress_bar.SetPosition(100)
		f.SaveAs(w.TsaveDialog.FileName())
		time.Sleep(time.Second * 1)
		vcl.ShowMessageFmt("保存完成！")
	})

	w.Button.SetEnabled(true)
	w.Win.SetEnabled(true)
	//os.Exit(0)
}

//客户生日
func Client_birthday(w *Windows, f *excelize.File) {
	if w.Date_strat_label.DateTime().Format("2006-01") != w.Date_stop_label.DateTime().Format("2006-01") {
		vcl.ShowMessage("请选择同一月份进行查询！")
		w.Button.SetEnabled(true)
		w.Win.SetEnabled(true)
		return
	}
	bady := url.Values{}
	bady.Add("Mysql_Select", w.Taskdlg.RadioButton().Caption())
	bady.Add("Start_Time", w.Date_strat_label.DateTime().Format("2006-01-02"))
	bady.Add("Stop_Time", w.Date_stop_label.DateTime().Format("2006-01-02"))
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
	time.Sleep(time.Second * 5)

	for k, v := range resulist {
		if k == 0 {
			f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+1), "客户编码")
			f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+1), "客户姓名")
			f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), "客户生日")
			f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+1), "员工部门")
			f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+1), "员工姓名")
			f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+1), "员工工号")
			f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+1), "员工状态")
			f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+1), "消费合计")
			f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+1), "最大单笔消费")

		}
		f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+2), v["客户编码"])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+2), v["客户姓名"])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+2), v["客户生日"])
		f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+2), v["员工部门"])
		f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+2), v["员工姓名"])
		f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+2), v["员工工号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+2), v["员工状态"])
		f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+2), v["消费合计"])
		f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+2), v["最大单笔消费"])

		vcl.ThreadSync(func() {
			w.Progress_bar.SetPosition(int32(float64(k+1) / float64(len(resulist)) * 100))
		})
	}

	vcl.ThreadSync(func() {
		w.Progress_bar.SetPosition(100)
		f.SaveAs(w.TsaveDialog.FileName())
		time.Sleep(time.Second * 1)
		vcl.ShowMessageFmt("保存完成！")
	})

	w.Button.SetEnabled(true)
	w.Win.SetEnabled(true)
	//os.Exit(0)
}

//客户名单构成报表
func Client_prestigious_university(w *Windows, f *excelize.File) {
	vcl.ShowMessage("查询数据为当前数据！")
	bady := url.Values{}
	bady.Add("Mysql_Select", w.Taskdlg.RadioButton().Caption())
	bady.Add("Start_Time", w.Date_strat_label.DateTime().Format("2006-01-02"))
	bady.Add("Stop_Time", w.Date_stop_label.DateTime().Format("2006-01-02"))
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
	time.Sleep(time.Second * 5)

	for k, v := range resulist {
		if k == 0 {
			f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+1), "部门")
			f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+1), "二级部门")
			f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+1), "员工名字")
			f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+1), "员工工号")
			f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+1), "员工状态")
			f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+1), "已购客户")
			f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+1), "准已购客户")
			f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+1), "意向客户")
			f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+1), "潜在客户")

		}
		f.SetCellValue("Sheet1", fmt.Sprintf("A%v", k+2), v["部门"])
		f.SetCellValue("Sheet1", fmt.Sprintf("B%v", k+2), v["二级部门"])
		f.SetCellValue("Sheet1", fmt.Sprintf("C%v", k+2), v["员工名字"])
		f.SetCellValue("Sheet1", fmt.Sprintf("D%v", k+2), v["员工工号"])
		f.SetCellValue("Sheet1", fmt.Sprintf("E%v", k+2), v["员工状态"])
		f.SetCellValue("Sheet1", fmt.Sprintf("F%v", k+2), v["已购客户"])
		f.SetCellValue("Sheet1", fmt.Sprintf("G%v", k+2), v["准已购客户"])
		f.SetCellValue("Sheet1", fmt.Sprintf("H%v", k+2), v["意向客户"])
		f.SetCellValue("Sheet1", fmt.Sprintf("I%v", k+2), v["潜在客户"])

		vcl.ThreadSync(func() {
			w.Progress_bar.SetPosition(int32(float64(k+1) / float64(len(resulist)) * 100))
		})
	}

	vcl.ThreadSync(func() {
		w.Progress_bar.SetPosition(100)
		f.SaveAs(w.TsaveDialog.FileName())
		time.Sleep(time.Second * 1)
		vcl.ShowMessageFmt("保存完成！")
	})

	w.Button.SetEnabled(true)
	w.Win.SetEnabled(true)
	//os.Exit(0)
}
