package sql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"os"
	"strings"
	"time"
	"xorm.io/xorm"
)

type Engine struct {
	Engine *xorm.Engine
	Err    error
}

const (
	TimeFormat = "2006-01-02 15:04:05"
	driverName = "sql"
	dBonnet    = "root:123456@tcp(192.168.0.19:3306)/data_analysis_library?charset=utf8"
	dBonnets   = "root:dcf4aa3f7b982ce4@tcp(192.168.0.11:3306)/bl_crm?charset=utf8"
	sss        = "%Y-%m-%d %H:%m:%s"
)

var tbMappers core.PrefixMapper

func init() {
	err := os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8")
	if err != nil {
		return
	} //修正中文乱码
}

// NewEngine 数据库连接建立
func (e *Engine) NewEngine() error {

	e.Engine, e.Err = xorm.NewEngine(driverName, dBonnet)
	if e.Err != nil {
		return e.Err
	}
	//tbMapped := core.NewPrefixMapper(core.SnakeMapper{}, tbMapper)
	e.Engine.ShowSQL(true)
	e.Engine.SetTableMapper(tbMappers)
	return nil
}

func (e *Engine) NewEngineC() error {

	e.Engine, e.Err = xorm.NewEngine(driverName, dBonnets)
	if e.Err != nil {
		return e.Err
	}
	//tbMapped := core.NewPrefixMapper(core.SnakeMapper{}, tbMapper)
	e.Engine.ShowSQL(true)
	e.Engine.SetTableMapper(tbMappers)
	return nil
}

// Close 数据库连接建立
func (e *Engine) Close() {
	e.Engine.Close()
}

// InsetrData 订单信息写入
// 快递发出订单处理
func (e *Engine) InsetrData(orderOr, expressNumber string) {
	e.NewEngine()
	defer e.Close()
	//开始写入订单数据
	_, err := e.Engine.Query(fmt.Sprintf("INSERT INTO `order_express` (  `order_on`, `Express_numbe` )\nVALUES\n\t( '%v','%v' );", orderOr, expressNumber))
	if err != nil {
		return
	}

}

// UpdateData 更新订单数据
// 退货登记处理
func (e *Engine) UpdateData(orderOr string) bool {
	e.NewEngine()
	defer e.Close()
	//判断订单是否已出库登记
	query, err := e.Engine.Query(fmt.Sprintf("SELECT * FROM `order_express` WHERE order_on = '%v';", orderOr))
	if err != nil {
		return false
	}

	if len(query) == 0 {
		_, _ = e.Engine.Query(fmt.Sprintf("INSERT INTO `order_express` (  `order_on`, `Express_numbe`,`created_at` )\nVALUES\n\t( '%v','%v','%v' );", orderOr, orderOr, time.Now().Format("2005-12-02 15:04:05")))
	}

	//开始更新订单数据
	_, err = e.Engine.Query(fmt.Sprintf("UPDATE order_express \nSET send_back_at = NOW() \nWHERE\n\torder_on = '%v'", orderOr))
	if err != nil {
		return false
	}

	return true
}

func (e *Engine) UpdateDataCrm(orderOr string) {
	e.NewEngineC()
	defer e.Close()
	_, _ = e.Engine.Query(fmt.Sprintf("UPDATE bl_mall_order SET `status` = 140,updated_at = NOW()   WHERE order_no = '%v';", orderOr))

}

// UpdateDataExpress 更新订单快递信息
func (e *Engine) UpdateDataExpress() {
	e.NewEngine()
	defer e.Close()
	//查询需要进行快递查询的订单信息
	query, err := e.Engine.Query("SELECT\n\torder_on\nFROM\n\torder_express \nWHERE\n\tdelete_at IS NULL \n\tAND receiving_at IS NULL \n\tAND send_back_at IS NULL \n\tAND send_back_Express_numbe IS NULL")
	if err != nil {
		return
	}

	//判断快递公司
	var expressType int
	for _, v := range query {
		if !strings.Contains(string(v["order_on"]), "SF") {
			expressType = 1

		} else if !strings.Contains(string(v["order_on"]), "sf") {
			expressType = 1
		} else if !strings.Contains(string(v["order_on"]), "jd") {
			expressType = 2
		} else if !strings.Contains(string(v["order_on"]), "jd") {
			expressType = 2
		}

		//根据快递公司进行对应的快递路查询
		if expressType == 1 {
			//顺丰快递
		} else if expressType == 2 {
			//京东快递
		} else {
			expressType = 0
		}
	}

}

// SelectOrder 查询订单号是否存在
func (e *Engine) SelectOrder(orderOn string) (error, bool) {
	e.NewEngineC()
	defer e.Close()
	query, err := e.Engine.Query(fmt.Sprintf("SELECT * FROM `bl_mall_order` WHERE order_no = '%v';", orderOn))
	if err != nil {
		return err, false
	}

	if len(query) == 0 {
		return nil, false
	}
	return nil, true
}

// DeleteOrder 删除订单
func (e *Engine) DeleteOrder(orderOn string) (error, bool) {
	e.NewEngine()
	defer e.Close()
	_, err := e.Engine.Query(fmt.Sprintf("UPDATE order_express SET delete_at = NOW() WHERE order_on = '%v';", orderOn))
	if err != nil {
		return err, false
	}
	return nil, true
}

// SelectNumber 查询当日操作数量
func (e *Engine) SelectNumber() (cNumber, tNumber int) {
	e.NewEngine()
	defer e.Close()
	query, err := e.Engine.Query(fmt.Sprintf("SELECT id FROM `order_express` WHERE created_at > '%v 00:00:00' AND created_at < '%v 23:59:59' AND delete_at is null", time.Now().Format("2006-01-02"), time.Now().Format("2006-01-02")))
	if err != nil {
		return
	}
	querys, err := e.Engine.Query(fmt.Sprintf("SELECT id FROM `order_express` WHERE send_back_at > '%v 00:00:00' AND send_back_at < '%v 23:59:59' AND delete_at is null", time.Now().Format("2006-01-02"), time.Now().Format("2006-01-02")))
	if err != nil {
		return
	}

	return len(query), len(querys)
}

//查询订单号
//提供数据导出使用
func (e *Engine) SelectOrder_x(statr, stop string) string {
	e.NewEngine()
	defer e.Close()
	query, err := e.Engine.Query(fmt.Sprintf("SELECT order_on FROM `order_express` WHERE send_back_at >= '%v 00:00:00' AND send_back_at <= '%v 23:59:59';", statr, stop))
	if err != nil {
		return ""
	}

	var order string

	for k, v := range query {
		if k == 0 {
			order = "'" + string(v["order_on"]) + "'"
		}
		order = order + ",'" + string(v["order_on"]) + "'"
	}

	return order
}
