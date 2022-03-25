package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"os"
	"xorm.io/xorm"
)

const (
	TimeFormat   = "2006-01-02"
	driverName   = "sql"
	dBconnect    = "root:dcf4aa3f7b982ce4@tcp(192.168.0.11:3306)/bl_crm?charset=utf8"
	NewdBconnect = "bolong:bolong2021!@#@tcp(192.168.0.17:3306)/crm_prod?charset=utf8"

	sss = "%Y-%m-%d %H:%m:%s"
)

type Order struct {
	Result []map[string][]byte
}

type Engine struct {
	Engine *xorm.Engine
	Err    error
}

var tbMappers core.PrefixMapper

func init() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码
}

//////////////////////////////////////////////////////////////////////////////////////////////////
//数据库连接建立
func (e *Engine) NewEngine() error {

	e.Engine, e.Err = xorm.NewEngine(driverName, dBconnect)
	if e.Err != nil {
		return e.Err
	}
	//tbMappe := core.NewPrefixMapper(core.SnakeMapper{}, tbMapper)
	e.Engine.ShowSQL(true)
	e.Engine.SetTableMapper(tbMappers)
	return nil
}

// CRMNewEngine 这个是临时增加 的一个新系统的数据库连接
func (e *Engine) CRMNewEngine() error {

	e.Engine, e.Err = xorm.NewEngine(driverName, NewdBconnect)
	if e.Err != nil {
		return e.Err
	}
	//tbMappe := core.NewPrefixMapper(core.SnakeMapper{}, tbMapper)
	e.Engine.ShowSQL(true)
	e.Engine.SetTableMapper(tbMappers)
	return nil
}

//数据库连接建立
func (e *Engine) Close() {
	e.Engine.Close()
}

// LoginDepartIdPermissions 员工部门迭代查询
func (e *Engine) LoginDepartIdPermissions(login_name string) (string, error) {
	depart_id_str := make(map[int]string, 500) //定义一个用来存储每次查询结果的map
	//查询当前员工的归属部门
	resp, err := e.Engine.Query(fmt.Sprintf("SELECT depart_id FROM bl_users WHERE login_name = '%v'", login_name))
	if err != nil {
		return "", fmt.Errorf("%v，请联系管理员！", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("部门为空！请联系管理员")
	}
	depart_id_str[0] = string(resp[0]["depart_id"])

	//开始查询员工当前部门以及以下所有部门，并进行字符串的拼接
	i := 0
	for {
		resp, err = e.Engine.Query(fmt.Sprintf("SELECT id FROM bl_depart WHERE parent_id in ('%v')", depart_id_str[i]))
		if err != nil {
			return "", fmt.Errorf("%v，请联系管理员！", err)
		}
		if len(resp) == 0 {
			return depart_id_str[0], nil
		}

		for k, v := range resp {
			if k == 0 {
				depart_id_str[i+1] = string(v["id"])
				continue
			}
			depart_id_str[i+1] = depart_id_str[i+1] + "," + string(v["id"])
		}
		depart_id_str[0] = depart_id_str[0] + "," + depart_id_str[i+1]
		i++
	}

}

// NewLoginDepartIdPermissions 员工部门迭代查询
func NewLoginDepartIdPermissions(loginName string) (string, error) {
	b := new(Engine)
	defer b.Close()
	err := b.CRMNewEngine()
	if err != nil {
		return "", err
	}

	departIdStr := make(map[int]string, 500) //定义一个用来存储每次查询结果的map
	//查询当前员工的归属部门
	resp, err := b.Engine.Query(fmt.Sprintf("SELECT organ_id FROM sys_user WHERE staff_no = '%v'", loginName))
	if err != nil {
		return "", fmt.Errorf("%v，请联系管理员！", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("部门为空！请联系管理员")
	}
	departIdStr[0] = "'" + string(resp[0]["organ_id"]) + "'"

	//开始查询员工当前部门以及以下所有部门，并进行字符串的拼接
	i := 0
	for {
		resp, err = b.Engine.Query(fmt.Sprintf("SELECT id FROM sys_organ WHERE parent_id =  (%v)", departIdStr[i]))
		if err != nil {
			return "", fmt.Errorf("%v，请联系管理员！", err)
		}
		if len(resp) == 0 {
			return departIdStr[0], nil
		}

		for k, v := range resp {
			if k == 0 {
				departIdStr[i+1] = "'" + string(v["id"]) + "'"
				continue
			}
			departIdStr[i+1] = departIdStr[i+1] + "," + "'" + string(v["id"]) + "'"
		}
		departIdStr[0] = departIdStr[0] + "," + departIdStr[i+1]
		i++
	}

}

//员工部门查询
func (e *Engine) Login_depart_id(login_name string) (string, error) {
	//查询当前员工的归属部门
	resp, err := e.Engine.Query(fmt.Sprintf("SELECT depart_id FROM bl_users WHERE login_name = '%v'", login_name))
	if err != nil {
		return "", fmt.Errorf("%v，请联系管理员！", err)
	}
	if len(resp) == 0 {
		return "", fmt.Errorf("部门为空！请联系管理员")
	}

	return string(resp[0]["depart_id"]), err

}

//员工职级权限查询
func (e *Engine) Login_position_level(login_name string) (string, error) {
	resp, err := e.Engine.Query(fmt.Sprintf("SELECT position_level FROM bl_users WHERE login_name = '%v'", login_name))
	fmt.Println(len(resp))

	if err != nil {
		return "", fmt.Errorf("%v，请联系管理员！", err)
	}

	if len(resp) == 0 {
		return "1", nil
	}

	return string(resp[0]["position_level"]), nil
}

//报表查询分流
func Select(id, Start_Time, Stop_Time, login_name, client_Models string) ([]map[string][]byte, error) {
	if id == "销售订单明细" {
		e := new(Engine)
		resulist, err := e.Order_select(Start_Time, Stop_Time, login_name, client_Models)
		return resulist, err
	} else if id == "新媒体线上明细" {
		if client_Models == "" {
			e := new(Engine)
			resulist, err := e.NewOrderXsSelect(Start_Time, Stop_Time, login_name, client_Models)
			return resulist, err
		} else {
			e := new(Engine)
			resulist, err := e.Order_xs_select(Start_Time, Stop_Time, login_name, client_Models)
			return resulist, err
		}
	} else if id == "新媒体线上统计" {
		e := new(Engine)
		resulist, err := e.Order_xsTJ_select(Start_Time, Stop_Time, login_name)
		return resulist, err
	} else if id == "新媒体线下二开统计" {
		e := new(Engine)
		resulist, err := e.Order_xxTJ_select(Start_Time, Stop_Time, login_name)
		return resulist, err
	} else if id == "礼品订单明细" {
		e := new(Engine)
		resulist, err := e.Order_gift_select(Start_Time, Stop_Time, login_name, client_Models)
		return resulist, err
	} else if id == "订单出库明细" {
		e := new(Engine)
		resulist, err := e.Order_warehouse_select(Start_Time, Stop_Time)
		return resulist, err
	} else if id == "通话记录统计报表" {
		e := new(Engine)
		resulist, err := e.Call_log_statistics(Start_Time, Stop_Time, login_name, client_Models)
		return resulist, err
	} else if id == "客户积分报表" {
		e := new(Engine)
		resulist, err := e.Client_coin()
		return resulist, err
	} else if id == "客户生日报表" {
		e := new(Engine)
		resulist, err := e.Client_birthday(Start_Time, Stop_Time, login_name, client_Models)
		return resulist, err
	} else if id == "客户名单构成报表" {
		e := new(Engine)
		resulist, err := e.Client_prestigious_university()
		return resulist, err
	}
	return nil, fmt.Errorf("查询报表不存在，请联系管理员！")
}

//销售订单报表
func (e *Engine) Order_select(Start_Time, Stop_Time, login_name, client_Models string) ([]map[string][]byte, error) {
	err := e.NewEngine()
	defer e.Close()
	if err != nil {
		return nil, fmt.Errorf("%v，请联系管理员！", err)
	}

	var resulist []map[string][]byte
	position_level, err := e.Login_position_level(login_name)
	if err != nil {
		return nil, fmt.Errorf("%v，请联系管理员！", err)
	}

	if client_Models == "Staff_TsaveDialog" { //销售
		if position_level <= "1" { //员工
			resulist, _ = e.Engine.Query(fmt.Sprintf("SELECT\n\tmo.created_at AS 订单创建时间,\n\tumc.nickname AS 订单创建人,\n\tumc.login_name AS 订单创建人工号,\n\tdmc.`name` AS 订单创建人部门,\n\tcc.`name` AS 客户姓名,\n\tcc.customer_code AS 客户编码,\n\tmo.order_no AS 系统订单号,\n\tdi2.`name` AS 订单类型,\n\tdi4.`NAME` AS 订单渠道,\n\tdi.`name` AS 订单状态,\n\tmo.performance_at AS 回款日期,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"来源渠道\"' ) AS 线上下单渠道,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) AS 线上下单时间,\n\tumm.nickname AS 线上下单核单人,\n\tumm.login_name AS 线上下单核单人工号,\n\tdmm.`name` AS 线上下单核单人部门,\n\tmo.`name` AS 订单备注,\n\tump.nickname AS 业绩归属人,\n\tump.login_name AS 业绩归属人工号,\n\tdi3.`NAME` AS 员工状态,\n\tdmp.`name` AS 业绩归属人部门,\n\tFORMAT( mo.total_amount / 100,2) AS 订单总金额,\n\tFORMAT( mo.discount_amount / 100 ,2) AS 订单折扣总额,\n\tFORMAT(( mo.total_amount - mo.discount_amount )/ 100 ,2) AS 订单售价总额,\n\tmc.`name` AS 订单商品,\n\tmc.goods_count AS 商品数量,\n\tFORMAT( mc.goods_price / 100 ,2) AS 商品原价,\n\tFORMAT( mc.price_sale / 100 ,2) AS 商品售价,\n\tFORMAT( mc.total_price / 100 ,2) AS 商品总价,\n\tei.ship_channel_name AS 配送方式,\n\tei.ship_channel_no AS 快递单号,\n\tei.last_trace AS 快递状态  \nFROM\n\tbl_mall_order mo\n\tLEFT JOIN bl_mall_cart mc ON mo.id = mc.order_id\n\tLEFT JOIN bl_express_invoice ei ON mo.id = ei.order_id\n\tLEFT JOIN bl_mall_order_media mom ON mo.id = mom.mall_order_id\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 43 ) di ON mo.`status` = di.item_value\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 53 ) di2 ON mo.type = di2.item_value\n\tLEFT JOIN bl_crm_customer cc ON mo.customer_id = cc.id\n\tLEFT JOIN bl_users umc ON mo.create_user_id = umc.id\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 14 ) di3 ON umc.`status` = di3.item_value\n\tLEFT JOIN ( SELECT id,NAME,item_value FROM bl_dict_item WHERE dict_id = 57 ) di4 ON mo.channel_id = di4.item_value\n\tLEFT JOIN bl_depart dmc ON mo.create_user_depart_id = dmc.id\n\tLEFT JOIN bl_users ump ON mo.performance_user_id = ump.id\n\tLEFT JOIN bl_depart dmp ON mo.performance_user_depart_id = dmp.id \n\tLEFT JOIN bl_users umm ON mom.assign_user_id = umm.id\n\tLEFT JOIN bl_depart dmm ON umm.depart_id = dmm.id  \nWHERE\n\tmo.deleted_at IS NULL \n\tAND mo.delete_user_id IS NULL \n\tAND mom.deleted_at IS NULL\n\tAND mom.delete_user_id IS NULL\n\tAND mc.deleted_at IS NULL\n\tAND mc.delete_user_id IS NULL\n\tAND ei.deleted_at IS NULL\n\tAND ei.delete_user_id IS NULL  \n\tAND ((mo.created_at > '%v 00:00:00' AND mo.created_at <= '%v 23:59:59') OR (mo.created_at < '%v 00:00:00' AND mo.performance_at IS NULL AND mo.`status` <= 90) OR (mo.created_at < '%v 00:00:00' AND mo.performance_at >'%v 00:00:00' and mo.performance_at <= '%v 23:59:59' ))\n\tAND di2.item_value = 1\n\tAND ump.login_name = '%v' \nORDER BY\n\tmo.created_at DESC", Start_Time, Stop_Time, Start_Time, Start_Time, Start_Time, Stop_Time, login_name))
		} else if position_level >= "2" { //管理
			depart_id, err := e.LoginDepartIdPermissions(login_name)
			if err != nil {
				return nil, fmt.Errorf("%v，请联系管理员！", err)
			}
			resulist, _ = e.Engine.Query(fmt.Sprintf("SELECT\n\tmo.created_at AS 订单创建时间,\n\tumc.nickname AS 订单创建人,\n\tumc.login_name AS 订单创建人工号,\n\tdmc.`name` AS 订单创建人部门,\n\tcc.`name` AS 客户姓名,\n\tcc.customer_code AS 客户编码,\n\tmo.order_no AS 系统订单号,\n\tdi2.`name` AS 订单类型,\n\tdi4.`NAME` AS 订单渠道,\n\tdi.`name` AS 订单状态,\n\tmo.performance_at AS 回款日期,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"来源渠道\"' ) AS 线上下单渠道,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) AS 线上下单时间,\n\tumm.nickname AS 线上下单核单人,\n\tumm.login_name AS 线上下单核单人工号,\n\tdmm.`name` AS 线上下单核单人部门,\n\tmo.`name` AS 订单备注,\n\tump.nickname AS 业绩归属人,\n\tump.login_name AS 业绩归属人工号,\n\tdi3.`NAME` AS 员工状态,\n\tdmp.`name` AS 业绩归属人部门,\n\tFORMAT( mo.total_amount / 100,2) AS 订单总金额,\n\tFORMAT( mo.discount_amount / 100 ,2) AS 订单折扣总额,\n\tFORMAT(( mo.total_amount - mo.discount_amount )/ 100 ,2) AS 订单售价总额,\n\tmc.`name` AS 订单商品,\n\tmc.goods_count AS 商品数量,\n\tFORMAT( mc.goods_price / 100 ,2) AS 商品原价,\n\tFORMAT( mc.price_sale / 100 ,2) AS 商品售价,\n\tFORMAT( mc.total_price / 100 ,2) AS 商品总价,\n\tei.ship_channel_name AS 配送方式,\n\tei.ship_channel_no AS 快递单号,\n\tei.last_trace AS 快递状态  \nFROM\n\tbl_mall_order mo\n\tLEFT JOIN bl_mall_cart mc ON mo.id = mc.order_id\n\tLEFT JOIN bl_express_invoice ei ON mo.id = ei.order_id\n\tLEFT JOIN bl_mall_order_media mom ON mo.id = mom.mall_order_id\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 43 ) di ON mo.`status` = di.item_value\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 53 ) di2 ON mo.type = di2.item_value\n\tLEFT JOIN bl_crm_customer cc ON mo.customer_id = cc.id\n\tLEFT JOIN bl_users umc ON mo.create_user_id = umc.id\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 14 ) di3 ON umc.`status` = di3.item_value\n\tLEFT JOIN ( SELECT id,NAME,item_value FROM bl_dict_item WHERE dict_id = 57 ) di4 ON mo.channel_id = di4.item_value\n\tLEFT JOIN bl_depart dmc ON mo.create_user_depart_id = dmc.id\n\tLEFT JOIN bl_users ump ON mo.performance_user_id = ump.id\n\tLEFT JOIN bl_depart dmp ON mo.performance_user_depart_id = dmp.id \n\tLEFT JOIN bl_users umm ON mom.assign_user_id = umm.id\n\tLEFT JOIN bl_depart dmm ON umm.depart_id = dmm.id  \nWHERE\n\tmo.deleted_at IS NULL \n\tAND mo.delete_user_id IS NULL \n\tAND mom.deleted_at IS NULL\n\tAND mom.delete_user_id IS NULL\n\tAND mc.deleted_at IS NULL\n\tAND mc.delete_user_id IS NULL\n\tAND ei.deleted_at IS NULL\n\tAND ei.delete_user_id IS NULL \n\tAND ((mo.created_at > '%v 00:00:00' AND mo.created_at <= '%v 23:59:59') OR (mo.created_at < '%v 00:00:00' AND mo.performance_at IS NULL AND mo.`status` <= 90) OR (mo.created_at < '%v 00:00:00' AND mo.performance_at >'%v 00:00:00' and mo.performance_at <= '%v 23:59:59' ))\n\tAND di2.item_value = 1\n\tAND dmp.id IN (%v)  \nORDER BY\n\tmo.created_at DESC", Start_Time, Stop_Time, Start_Time, Start_Time, Start_Time, Stop_Time, depart_id))
		}
	} else if client_Models == "TsaveDialog" || client_Models == "Medium_TsaveDialog" { //无限制
		resulist, _ = e.Engine.Query(fmt.Sprintf("SELECT\n\tmo.created_at AS 订单创建时间,\n\tumc.nickname AS 订单创建人,\n\tumc.login_name AS 订单创建人工号,\n\tdmc.`name` AS 订单创建人部门,\n\tcc.`name` AS 客户姓名,\n\tcc.customer_code AS 客户编码,\n\tmo.order_no AS 系统订单号,\n\tdi2.`name` AS 订单类型,\n\tdi4.`NAME` AS 订单渠道,\n\tdi.`name` AS 订单状态,\n\tmo.performance_at AS 回款日期,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"来源渠道\"' ) AS 线上下单渠道,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) AS 线上下单时间,\n\tumm.nickname AS 线上下单核单人,\n\tumm.login_name AS 线上下单核单人工号,\n\tdmm.`name` AS 线上下单核单人部门,\n\tmo.`name` AS 订单备注,\n\tump.nickname AS 业绩归属人,\n\tump.login_name AS 业绩归属人工号,\n\tdi3.`NAME` AS 员工状态,\n\tdmp.`name` AS 业绩归属人部门,\n\tFORMAT( mo.total_amount / 100,2) AS 订单总金额,\n\tFORMAT( mo.discount_amount / 100 ,2) AS 订单折扣总额,\n\tFORMAT(( mo.total_amount - mo.discount_amount )/ 100 ,2) AS 订单售价总额,\n\tmc.`name` AS 订单商品,\n\tmc.goods_count AS 商品数量,\n\tFORMAT( mc.goods_price / 100 ,2) AS 商品原价,\n\tFORMAT( mc.price_sale / 100 ,2) AS 商品售价,\n\tFORMAT( mc.total_price / 100 ,2) AS 商品总价,\n\tei.ship_channel_name AS 配送方式,\n\tei.ship_channel_no AS 快递单号,\n\tei.last_trace AS 快递状态  \nFROM\n\tbl_mall_order mo\n\tLEFT JOIN bl_mall_cart mc ON mo.id = mc.order_id\n\tLEFT JOIN bl_express_invoice ei ON mo.id = ei.order_id\n\tLEFT JOIN bl_mall_order_media mom ON mo.id = mom.mall_order_id\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 43 ) di ON mo.`status` = di.item_value\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 53 ) di2 ON mo.type = di2.item_value\n\tLEFT JOIN bl_crm_customer cc ON mo.customer_id = cc.id\n\tLEFT JOIN bl_users umc ON mo.create_user_id = umc.id\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 14 ) di3 ON umc.`status` = di3.item_value\n\tLEFT JOIN ( SELECT id,NAME,item_value FROM bl_dict_item WHERE dict_id = 57 ) di4 ON mo.channel_id = di4.item_value\n\tLEFT JOIN bl_depart dmc ON mo.create_user_depart_id = dmc.id\n\tLEFT JOIN bl_users ump ON mo.performance_user_id = ump.id\n\tLEFT JOIN bl_depart dmp ON mo.performance_user_depart_id = dmp.id \n\tLEFT JOIN bl_users umm ON mom.assign_user_id = umm.id\n\tLEFT JOIN bl_depart dmm ON umm.depart_id = dmm.id  \nWHERE\n\tmo.deleted_at IS NULL \n\tAND mo.delete_user_id IS NULL \n\tAND mom.deleted_at IS NULL\n\tAND mom.delete_user_id IS NULL\n\tAND mc.deleted_at IS NULL\n\tAND mc.delete_user_id IS NULL\n\tAND ei.deleted_at IS NULL\n\tAND ei.delete_user_id IS NULL \n\tAND ((mo.created_at > '%v 00:00:00' AND mo.created_at <= '%v 23:59:59') OR (mo.created_at < '%v 00:00:00' AND mo.performance_at IS NULL AND mo.`status` <= 90) OR (mo.created_at < '%v 00:00:00' AND mo.performance_at >'%v 00:00:00' and mo.performance_at <= '%v 23:59:59' ))\n\tAND di2.item_value = 1 \nORDER BY\n\tmo.created_at DESC", Start_Time, Stop_Time, Start_Time, Start_Time, Start_Time, Stop_Time))
	} //} else if client_Models == "Medium_TsaveDialog" { //新媒体
	//	depart_id, err := e.Login_depart_id_permissions(login_name)
	//	if err != nil {
	//		return nil, fmt.Errorf("%v，请联系管理员！", err)
	//	}
	//	resulist, _ = e.Engine.Query(fmt.Sprintf("SELECT\n\tmo.created_at AS 订单创建时间,\n\tumc.nickname AS 订单创建人,\n\tumc.login_name AS 订单创建人工号,\n\tdmc.`name` AS 订单创建人部门,\n\tcc.`name` AS 客户姓名,\n\tcc.customer_code AS 客户编码,\n\tmo.order_no AS 系统订单号,\n\tdi2.`name` AS 订单类型,\n\tdi4.`NAME` AS 订单渠道,\n\tdi.`name` AS 订单状态,\n\tmo.performance_at AS 回款日期,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"来源渠道\"' ) AS 线上下单渠道,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) AS 线上下单时间,\n\tumm.nickname AS 线上下单核单人,\n\tumm.login_name AS 线上下单核单人工号,\n\tdmm.`name` AS 线上下单核单人部门,\n\tmo.`name` AS 订单备注,\n\tump.nickname AS 业绩归属人,\n\tump.login_name AS 业绩归属人工号,\n\tdi3.`NAME` AS 员工状态,\n\tdmp.`name` AS 业绩归属人部门,\n\tFORMAT( mo.total_amount / 100,2) AS 订单总金额,\n\tFORMAT( mo.discount_amount / 100 ,2) AS 订单折扣总额,\n\tFORMAT(( mo.total_amount - mo.discount_amount )/ 100 ,2) AS 订单售价总额,\n\tmc.`name` AS 订单商品,\n\tmc.goods_count AS 商品数量,\n\tFORMAT( mc.goods_price / 100 ,2) AS 商品原价,\n\tFORMAT( mc.price_sale / 100 ,2) AS 商品售价,\n\tFORMAT( mc.total_price / 100 ,2) AS 商品总价,\n\tei.ship_channel_name AS 配送方式,\n\tei.ship_channel_no AS 快递单号,\n\tei.last_trace AS 快递状态  \nFROM\n\tbl_mall_order mo\n\tLEFT JOIN bl_mall_cart mc ON mo.id = mc.order_id\n\tLEFT JOIN bl_express_invoice ei ON mo.id = ei.order_id\n\tLEFT JOIN bl_mall_order_media mom ON mo.id = mom.mall_order_id\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 43 ) di ON mo.`status` = di.item_value\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 53 ) di2 ON mo.type = di2.item_value\n\tLEFT JOIN bl_crm_customer cc ON mo.customer_id = cc.id\n\tLEFT JOIN bl_users umc ON mo.create_user_id = umc.id\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 14 ) di3 ON umc.`status` = di3.item_value\n\tLEFT JOIN ( SELECT id,NAME,item_value FROM bl_dict_item WHERE dict_id = 57 ) di4 ON mo.channel_id = di4.item_value\n\tLEFT JOIN bl_depart dmc ON mo.create_user_depart_id = dmc.id\n\tLEFT JOIN bl_users ump ON mo.performance_user_id = ump.id\n\tLEFT JOIN bl_depart dmp ON mo.performance_user_depart_id = dmp.id \n\tLEFT JOIN bl_users umm ON mom.assign_user_id = umm.id\n\tLEFT JOIN bl_depart dmm ON umm.depart_id = dmm.id  \nWHERE\n\tmo.deleted_at IS NULL \n\tAND mo.delete_user_id IS NULL \n\tAND mom.deleted_at IS NULL\n\tAND mom.delete_user_id IS NULL\n\tAND mc.deleted_at IS NULL\n\tAND mc.delete_user_id IS NULL\n\tAND ei.deleted_at IS NULL\n\tAND ei.delete_user_id IS NULL \n\tAND ((mo.created_at > '%v 00:00:00' AND mo.created_at <= '%v 23:59:59') OR (mo.created_at < '%v 00:00:00' AND mo.performance_at IS NULL AND mo.`status` <= 90) OR (mo.created_at < '%v 00:00:00' AND mo.performance_at >'%v 00:00:00' and mo.performance_at <= '%v 23:59:59' ))\n\tAND di2.item_value = 1 \n\tAND dmp.id IN (%v)   \nORDER BY\n\tmo.created_at DESC", Start_Time, Stop_Time, Start_Time, Start_Time, Start_Time, Stop_Time, depart_id))
	//}
	return resulist, nil
}

//礼品订单报表
func (e *Engine) Order_gift_select(Start_Time, Stop_Time, login_name, client_Models string) ([]map[string][]byte, error) {
	err := e.NewEngine()
	defer e.Close()
	if err != nil {
		return nil, fmt.Errorf("%v，请联系管理员！", err)
	}

	var resulist []map[string][]byte
	position_level, err := e.Login_position_level(login_name)
	if err != nil {
		return nil, fmt.Errorf("%v，请联系管理员！", err)
	}

	if client_Models == "Staff_TsaveDialog" { //销售
		if position_level <= "1" { //员工
			resulist, _ = e.Engine.Query(fmt.Sprintf("SELECT\n\tmo.created_at AS 订单创建时间,\n\tumc.nickname AS 订单创建人,\n\tumc.login_name AS 订单创建人工号,\n\tdmc.`name` AS 订单创建人部门,\n\tcc.`name` AS 客户姓名,\n\tcc.customer_code AS 客户编码,\n\tmo.order_no AS 系统订单号,\n\tdi2.`name` AS 订单类型,\n\tdi4.`NAME` AS 订单渠道,\n\tdi.`name` AS 订单状态,\n\tmo.performance_at AS 回款日期,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"来源渠道\"' ) AS 线上下单渠道,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) AS 线上下单时间,\n\tumm.nickname AS 线上下单核单人,\n\tumm.login_name AS 线上下单核单人工号,\n\tdmm.`name` AS 线上下单核单人部门,\n\tmo.`name` AS 订单备注,\n\tump.nickname AS 业绩归属人,\n\tump.login_name AS 业绩归属人工号,\n\tdi3.`NAME` AS 员工状态,\n\tdmp.`name` AS 业绩归属人部门,\n\tFORMAT( mo.total_amount / 100,2) AS 订单总金额,\n\tFORMAT( mo.discount_amount / 100 ,2) AS 订单折扣总额,\n\tFORMAT(( mo.total_amount - mo.discount_amount )/ 100 ,2) AS 订单售价总额,\n\tmc.`name` AS 订单商品,\n\tmc.goods_count AS 商品数量,\n\tFORMAT( mc.goods_price / 100 ,2) AS 商品原价,\n\tFORMAT( mc.price_sale / 100 ,2) AS 商品售价,\n\tFORMAT( mc.total_price / 100 ,2) AS 商品总价,\n\tei.ship_channel_name AS 配送方式,\n\tei.ship_channel_no AS 快递单号,\n\tei.last_trace AS 快递状态  \nFROM\n\tbl_mall_order mo\n\tLEFT JOIN bl_mall_cart mc ON mo.id = mc.order_id\n\tLEFT JOIN bl_express_invoice ei ON mo.id = ei.order_id\n\tLEFT JOIN bl_mall_order_media mom ON mo.id = mom.mall_order_id\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 43 ) di ON mo.`status` = di.item_value\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 53 ) di2 ON mo.type = di2.item_value\n\tLEFT JOIN bl_crm_customer cc ON mo.customer_id = cc.id\n\tLEFT JOIN bl_users umc ON mo.create_user_id = umc.id\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 14 ) di3 ON umc.`status` = di3.item_value\n\tLEFT JOIN ( SELECT id,NAME,item_value FROM bl_dict_item WHERE dict_id = 57 ) di4 ON mo.channel_id = di4.item_value\n\tLEFT JOIN bl_depart dmc ON mo.create_user_depart_id = dmc.id\n\tLEFT JOIN bl_users ump ON mo.performance_user_id = ump.id\n\tLEFT JOIN bl_depart dmp ON mo.performance_user_depart_id = dmp.id \n\tLEFT JOIN bl_users umm ON mom.assign_user_id = umm.id\n\tLEFT JOIN bl_depart dmm ON umm.depart_id = dmm.id  \n\tWHERE\n\tmo.deleted_at IS NULL \n\tAND mo.delete_user_id IS NULL \n\tAND mom.deleted_at IS NULL\n\tAND mom.delete_user_id IS NULL\n\tAND mc.deleted_at IS NULL\n\tAND mc.delete_user_id IS NULL\n\tAND ei.deleted_at IS NULL\n\tAND ei.delete_user_id IS NULL\n\tAND di2.item_value IN (2,3)\n\tAND (mo.created_at > '%v 00:00:00' AND mo.created_at <= '%v 23:59:59')\n\t AND umc.login_name = '%v' \nORDER BY\n\tmo.created_at DESC", Start_Time, Stop_Time, login_name))
		} else if position_level >= "2" { //管理
			depart_id, err := e.LoginDepartIdPermissions(login_name)
			if err != nil {
				return nil, fmt.Errorf("%v，请联系管理员！", err)
			}
			resulist, _ = e.Engine.Query(fmt.Sprintf("SELECT\n\tmo.created_at AS 订单创建时间,\n\tumc.nickname AS 订单创建人,\n\tumc.login_name AS 订单创建人工号,\n\tdmc.`name` AS 订单创建人部门,\n\tcc.`name` AS 客户姓名,\n\tcc.customer_code AS 客户编码,\n\tmo.order_no AS 系统订单号,\n\tdi2.`name` AS 订单类型,\n\tdi4.`NAME` AS 订单渠道,\n\tdi.`name` AS 订单状态,\n\tmo.performance_at AS 回款日期,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"来源渠道\"' ) AS 线上下单渠道,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) AS 线上下单时间,\n\tumm.nickname AS 线上下单核单人,\n\tumm.login_name AS 线上下单核单人工号,\n\tdmm.`name` AS 线上下单核单人部门,\n\tmo.`name` AS 订单备注,\n\tump.nickname AS 业绩归属人,\n\tump.login_name AS 业绩归属人工号,\n\tdi3.`NAME` AS 员工状态,\n\tdmp.`name` AS 业绩归属人部门,\n\tFORMAT( mo.total_amount / 100,2) AS 订单总金额,\n\tFORMAT( mo.discount_amount / 100 ,2) AS 订单折扣总额,\n\tFORMAT(( mo.total_amount - mo.discount_amount )/ 100 ,2) AS 订单售价总额,\n\tmc.`name` AS 订单商品,\n\tmc.goods_count AS 商品数量,\n\tFORMAT( mc.goods_price / 100 ,2) AS 商品原价,\n\tFORMAT( mc.price_sale / 100 ,2) AS 商品售价,\n\tFORMAT( mc.total_price / 100 ,2) AS 商品总价,\n\tei.ship_channel_name AS 配送方式,\n\tei.ship_channel_no AS 快递单号,\n\tei.last_trace AS 快递状态  \nFROM\n\tbl_mall_order mo\n\tLEFT JOIN bl_mall_cart mc ON mo.id = mc.order_id\n\tLEFT JOIN bl_express_invoice ei ON mo.id = ei.order_id\n\tLEFT JOIN bl_mall_order_media mom ON mo.id = mom.mall_order_id\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 43 ) di ON mo.`status` = di.item_value\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 53 ) di2 ON mo.type = di2.item_value\n\tLEFT JOIN bl_crm_customer cc ON mo.customer_id = cc.id\n\tLEFT JOIN bl_users umc ON mo.create_user_id = umc.id\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 14 ) di3 ON umc.`status` = di3.item_value\n\tLEFT JOIN ( SELECT id,NAME,item_value FROM bl_dict_item WHERE dict_id = 57 ) di4 ON mo.channel_id = di4.item_value\n\tLEFT JOIN bl_depart dmc ON mo.create_user_depart_id = dmc.id\n\tLEFT JOIN bl_users ump ON mo.performance_user_id = ump.id\n\tLEFT JOIN bl_depart dmp ON mo.performance_user_depart_id = dmp.id \n\tLEFT JOIN bl_users umm ON mom.assign_user_id = umm.id\n\tLEFT JOIN bl_depart dmm ON umm.depart_id = dmm.id  \nWHERE\n\tmo.deleted_at IS NULL \n\tAND mo.delete_user_id IS NULL \n\tAND mom.deleted_at IS NULL\n\tAND mom.delete_user_id IS NULL\n\tAND mc.deleted_at IS NULL\n\tAND mc.delete_user_id IS NULL\n\tAND ei.deleted_at IS NULL\n\tAND ei.delete_user_id IS NULL\n\tAND di2.item_value IN (2,3)\n\tAND (mo.created_at > '%v 00:00:00' AND mo.created_at <= '%v 23:59:59')\n\t AND dmc.id IN (%v)  \nORDER BY\n\tmo.created_at DESC", Start_Time, Stop_Time, depart_id))
		}
	} else if client_Models == "TsaveDialog" { //无限制
		resulist, _ = e.Engine.Query(fmt.Sprintf("SELECT\n\tmo.created_at AS 订单创建时间,\n\tumc.nickname AS 订单创建人,\n\tumc.login_name AS 订单创建人工号,\n\tdmc.`name` AS 订单创建人部门,\n\tcc.`name` AS 客户姓名,\n\tcc.customer_code AS 客户编码,\n\tmo.order_no AS 系统订单号,\n\tdi2.`name` AS 订单类型,\n\tdi4.`NAME` AS 订单渠道,\n\tdi.`name` AS 订单状态,\n\tmo.performance_at AS 回款日期,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"来源渠道\"' ) AS 线上下单渠道,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) AS 线上下单时间,\n\tumm.nickname AS 线上下单核单人,\n\tumm.login_name AS 线上下单核单人工号,\n\tdmm.`name` AS 线上下单核单人部门,\n\tmo.`name` AS 订单备注,\n\tump.nickname AS 业绩归属人,\n\tump.login_name AS 业绩归属人工号,\n\tdi3.`NAME` AS 员工状态,\n\tdmp.`name` AS 业绩归属人部门,\n\tFORMAT( mo.total_amount / 100,2) AS 订单总金额,\n\tFORMAT( mo.discount_amount / 100 ,2) AS 订单折扣总额,\n\tFORMAT(( mo.total_amount - mo.discount_amount )/ 100 ,2) AS 订单售价总额,\n\tmc.`name` AS 订单商品,\n\tmc.goods_count AS 商品数量,\n\tFORMAT( mc.goods_price / 100 ,2) AS 商品原价,\n\tFORMAT( mc.price_sale / 100 ,2) AS 商品售价,\n\tFORMAT( mc.total_price / 100 ,2) AS 商品总价,\n\tei.ship_channel_name AS 配送方式,\n\tei.ship_channel_no AS 快递单号,\n\tei.last_trace AS 快递状态  \nFROM\n\tbl_mall_order mo\n\tLEFT JOIN bl_mall_cart mc ON mo.id = mc.order_id\n\tLEFT JOIN bl_express_invoice ei ON mo.id = ei.order_id\n\tLEFT JOIN bl_mall_order_media mom ON mo.id = mom.mall_order_id\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 43 ) di ON mo.`status` = di.item_value\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 53 ) di2 ON mo.type = di2.item_value\n\tLEFT JOIN bl_crm_customer cc ON mo.customer_id = cc.id\n\tLEFT JOIN bl_users umc ON mo.create_user_id = umc.id\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 14 ) di3 ON umc.`status` = di3.item_value\n\tLEFT JOIN ( SELECT id,NAME,item_value FROM bl_dict_item WHERE dict_id = 57 ) di4 ON mo.channel_id = di4.item_value\n\tLEFT JOIN bl_depart dmc ON mo.create_user_depart_id = dmc.id\n\tLEFT JOIN bl_users ump ON mo.performance_user_id = ump.id\n\tLEFT JOIN bl_depart dmp ON mo.performance_user_depart_id = dmp.id \n\tLEFT JOIN bl_users umm ON mom.assign_user_id = umm.id\n\tLEFT JOIN bl_depart dmm ON umm.depart_id = dmm.id  \nWHERE\n\tmo.deleted_at IS NULL \n\tAND mo.delete_user_id IS NULL \n\tAND mom.deleted_at IS NULL\n\tAND mom.delete_user_id IS NULL\n\tAND mc.deleted_at IS NULL\n\tAND mc.delete_user_id IS NULL\n\tAND ei.deleted_at IS NULL\n\tAND ei.delete_user_id IS NULL\n\tAND di2.item_value IN (2,3)\n\tAND (mo.created_at > '%v 00:00:00' AND mo.created_at <= '%v 23:59:59')\n\t ORDER BY\n\tmo.created_at DESC", Start_Time, Stop_Time))
	} else if client_Models == "Medium_TsaveDialog" { //新媒体
		depart_id, err := e.LoginDepartIdPermissions(login_name)
		if err != nil {
			return nil, fmt.Errorf("%v，请联系管理员！", err)
		}
		resulist, _ = e.Engine.Query(fmt.Sprintf("SELECT\n\tmo.created_at AS 订单创建时间,\n\tumc.nickname AS 订单创建人,\n\tumc.login_name AS 订单创建人工号,\n\tdmc.`name` AS 订单创建人部门,\n\tcc.`name` AS 客户姓名,\n\tcc.customer_code AS 客户编码,\n\tmo.order_no AS 系统订单号,\n\tdi2.`name` AS 订单类型,\n\tdi4.`NAME` AS 订单渠道,\n\tdi.`name` AS 订单状态,\n\tmo.performance_at AS 回款日期,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"来源渠道\"' ) AS 线上下单渠道,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) AS 线上下单时间,\n\tumm.nickname AS 线上下单核单人,\n\tumm.login_name AS 线上下单核单人工号,\n\tdmm.`name` AS 线上下单核单人部门,\n\tmo.`name` AS 订单备注,\n\tump.nickname AS 业绩归属人,\n\tump.login_name AS 业绩归属人工号,\n\tdi3.`NAME` AS 员工状态,\n\tdmp.`name` AS 业绩归属人部门,\n\tFORMAT( mo.total_amount / 100,2) AS 订单总金额,\n\tFORMAT( mo.discount_amount / 100 ,2) AS 订单折扣总额,\n\tFORMAT(( mo.total_amount - mo.discount_amount )/ 100 ,2) AS 订单售价总额,\n\tmc.`name` AS 订单商品,\n\tmc.goods_count AS 商品数量,\n\tFORMAT( mc.goods_price / 100 ,2) AS 商品原价,\n\tFORMAT( mc.price_sale / 100 ,2) AS 商品售价,\n\tFORMAT( mc.total_price / 100 ,2) AS 商品总价,\n\tei.ship_channel_name AS 配送方式,\n\tei.ship_channel_no AS 快递单号,\n\tei.last_trace AS 快递状态  \nFROM\n\tbl_mall_order mo\n\tLEFT JOIN bl_mall_cart mc ON mo.id = mc.order_id\n\tLEFT JOIN bl_express_invoice ei ON mo.id = ei.order_id\n\tLEFT JOIN bl_mall_order_media mom ON mo.id = mom.mall_order_id\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 43 ) di ON mo.`status` = di.item_value\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 53 ) di2 ON mo.type = di2.item_value\n\tLEFT JOIN bl_crm_customer cc ON mo.customer_id = cc.id\n\tLEFT JOIN bl_users umc ON mo.create_user_id = umc.id\n\tLEFT JOIN ( SELECT id, NAME, item_value FROM bl_dict_item WHERE dict_id = 14 ) di3 ON umc.`status` = di3.item_value\n\tLEFT JOIN ( SELECT id,NAME,item_value FROM bl_dict_item WHERE dict_id = 57 ) di4 ON mo.channel_id = di4.item_value\n\tLEFT JOIN bl_depart dmc ON mo.create_user_depart_id = dmc.id\n\tLEFT JOIN bl_users ump ON mo.performance_user_id = ump.id\n\tLEFT JOIN bl_depart dmp ON mo.performance_user_depart_id = dmp.id \n\tLEFT JOIN bl_users umm ON mom.assign_user_id = umm.id\n\tLEFT JOIN bl_depart dmm ON umm.depart_id = dmm.id  \nWHERE\n\tmo.deleted_at IS NULL \n\tAND mo.delete_user_id IS NULL \n\tAND mom.deleted_at IS NULL\n\tAND mom.delete_user_id IS NULL\n\tAND mc.deleted_at IS NULL\n\tAND mc.delete_user_id IS NULL\n\tAND ei.deleted_at IS NULL\n\tAND ei.delete_user_id IS NULL\n\tAND di2.item_value IN (2,3)\n\tAND (mo.created_at > '%v 00:00:00' AND mo.created_at <= '%v 23:59:59')\n\t AND dmc.id IN (%v)   \nORDER BY\n\tmo.created_at DESC", Start_Time, Stop_Time, depart_id))
	}
	return resulist, nil
}

//通话记录统计报表
func (e *Engine) Call_log_statistics(Start_Time, Stop_Time, login_name, client_Models string) ([]map[string][]byte, error) {
	err := e.NewEngine()
	defer e.Close()
	if err != nil {
		return nil, fmt.Errorf("%v，请联系管理员！", err)
	}
	var resulist []map[string][]byte
	position_level, err := e.Login_position_level(login_name)
	if client_Models == "Staff_TsaveDialog" { //销售
		if position_level <= "1" { //员工
			resulist, _ = e.Engine.Query(fmt.Sprintf("SELECT\n\td1.NAME AS 部门,\n\td.NAME AS 二级部门,\n\tu.nickname AS 员工姓名,\n\tu.login_name AS 员工工号,\n\tCOUNT(\n\tIF\n\t( cqv.call_duration = 0, cqv.id, NULL )) AS 未接通话总数,\n\tSUM(\n\tIF\n\t( cqv.call_duration = 0, cqv.call_duration, 0 )) AS 未接通时长,\n\tCOUNT(\n\tIF\n\t( cqv.call_duration > 0 && cqv.call_duration < 60, cqv.id, NULL )) AS 无效通话总数,\n\tSUM(\n\tIF\n\t( cqv.call_duration > 0 && cqv.call_duration < 60, cqv.call_duration, 0 )) AS 无效通话时长,\n\tCOUNT(\n\tIF\n\t( cqv.call_duration >= 60 && cqv.call_duration < 180, cqv.id, NULL )) AS 有效通话总数,\n\tSUM(\n\tIF\n\t( cqv.call_duration >= 60 && cqv.call_duration < 180, cqv.call_duration, 0 )) AS 有效通话时长,\n\tCOUNT(\n\tIF\n\t( cqv.call_duration >= 180, cqv.id, NULL )) AS 优质通话总数,\n\tSUM(\n\tIF\n\t( cqv.call_duration >= 180, cqv.call_duration, 0 )) AS 优质通话时长,\n\tCOUNT( cqv.id ) AS 总通话数,\n\tSUM( cqv.call_duration ) AS 通话总时长 \nFROM\n\tbl_crm_quality_voice cqv\n\tINNER JOIN bl_users u ON cqv.user_id = u.id\n\tINNER JOIN bl_depart d ON u.depart_id = d.id\n\tINNER JOIN bl_depart d1 ON d.parent_id = d1.id \nWHERE\n\tstart_at > '%v 00:00:00' \n\tAND start_at < '%v 23:59:59' \n\tAND call_type = 'extension_outbound'\t\n\tAND u.login_name = '%v'\nGROUP BY\n\td1.NAME,\n\td.NAME,\n\tu.nickname,\n\tu.login_name \nORDER BY\n\td1.NAME,\n\td.NAME,\n\tu.nickname,\n\tu.login_name", Start_Time, Stop_Time, login_name))
		} else if position_level >= "2" { //管理
			depart_id, err := e.LoginDepartIdPermissions(login_name)
			if err != nil {
				return nil, fmt.Errorf("%v，请联系管理员！", err)
			}
			resulist, _ = e.Engine.Query(fmt.Sprintf("SELECT\n\td1.NAME AS 部门,\n\td.NAME AS 二级部门,\n\tu.nickname AS 员工姓名,\n\tu.login_name AS 员工工号,\n\tCOUNT(\n\tIF\n\t( cqv.call_duration = 0, cqv.id, NULL )) AS 未接通话总数,\n\tSUM(\n\tIF\n\t( cqv.call_duration = 0, cqv.call_duration, 0 )) AS 未接通时长,\n\tCOUNT(\n\tIF\n\t( cqv.call_duration > 0 && cqv.call_duration < 60, cqv.id, NULL )) AS 无效通话总数,\n\tSUM(\n\tIF\n\t( cqv.call_duration > 0 && cqv.call_duration < 60, cqv.call_duration, 0 )) AS 无效通话时长,\n\tCOUNT(\n\tIF\n\t( cqv.call_duration >= 60 && cqv.call_duration < 180, cqv.id, NULL )) AS 有效通话总数,\n\tSUM(\n\tIF\n\t( cqv.call_duration >= 60 && cqv.call_duration < 180, cqv.call_duration, 0 )) AS 有效通话时长,\n\tCOUNT(\n\tIF\n\t( cqv.call_duration >= 180, cqv.id, NULL )) AS 优质通话总数,\n\tSUM(\n\tIF\n\t( cqv.call_duration >= 180, cqv.call_duration, 0 )) AS 优质通话时长,\n\tCOUNT( cqv.id ) AS 总通话数,\n\tSUM( cqv.call_duration ) AS 通话总时长 \nFROM\n\tbl_crm_quality_voice cqv\n\tINNER JOIN bl_users u ON cqv.user_id = u.id\n\tINNER JOIN bl_depart d ON u.depart_id = d.id\n\tINNER JOIN bl_depart d1 ON d.parent_id = d1.id \nWHERE\n\tstart_at > '%v 00:00:00' \n\tAND start_at < '%v 23:59:59' \n\tAND call_type = 'extension_outbound'\nAND d.id in (%v)\t\nGROUP BY\n\td1.NAME,\n\td.NAME,\n\tu.nickname,\n\tu.login_name \nORDER BY\n\td1.NAME,\n\td.NAME,\n\tu.nickname,\n\tu.login_name", Start_Time, Stop_Time, depart_id))
		}
	} else if client_Models == "TsaveDialog" { //无限制
		resulist, _ = e.Engine.Query(fmt.Sprintf("SELECT\n\td1.NAME AS 部门,\n\td.NAME AS 二级部门,\n\tu.nickname AS 员工姓名,\n\tu.login_name AS 员工工号,\n\tCOUNT(\n\tIF\n\t( cqv.call_duration = 0, cqv.id, NULL )) AS 未接通话总数,\n\tSUM(\n\tIF\n\t( cqv.call_duration = 0, cqv.call_duration, 0 )) AS 未接通时长,\n\tCOUNT(\n\tIF\n\t( cqv.call_duration > 0 && cqv.call_duration < 60, cqv.id, NULL )) AS 无效通话总数,\n\tSUM(\n\tIF\n\t( cqv.call_duration > 0 && cqv.call_duration < 60, cqv.call_duration, 0 )) AS 无效通话时长,\n\tCOUNT(\n\tIF\n\t( cqv.call_duration >= 60 && cqv.call_duration < 180, cqv.id, NULL )) AS 有效通话总数,\n\tSUM(\n\tIF\n\t( cqv.call_duration >= 60 && cqv.call_duration < 180, cqv.call_duration, 0 )) AS 有效通话时长,\n\tCOUNT(\n\tIF\n\t( cqv.call_duration >= 180, cqv.id, NULL )) AS 优质通话总数,\n\tSUM(\n\tIF\n\t( cqv.call_duration >= 180, cqv.call_duration, 0 )) AS 优质通话时长,\n\tCOUNT( cqv.id ) AS 总通话数,\n\tSUM( cqv.call_duration ) AS 通话总时长 \nFROM\n\tbl_crm_quality_voice cqv\n\tINNER JOIN bl_users u ON cqv.user_id = u.id\n\tINNER JOIN bl_depart d ON u.depart_id = d.id\n\tINNER JOIN bl_depart d1 ON d.parent_id = d1.id \nWHERE\n\tstart_at > '%v 00:00:00' \n\tAND start_at < '%v 23:59:59' \n\tAND call_type = 'extension_outbound'\t\nGROUP BY\n\td1.NAME,\n\td.NAME,\n\tu.nickname,\n\tu.login_name \nORDER BY\n\td1.NAME,\n\td.NAME,\n\tu.nickname,\n\tu.login_name", Start_Time, Stop_Time))
	} else if client_Models == "Medium_TsaveDialog" { //新媒体
		depart_id, err := e.LoginDepartIdPermissions(login_name)
		if err != nil {
			return nil, fmt.Errorf("%v，请联系管理员！", err)
		}
		resulist, _ = e.Engine.Query(fmt.Sprintf("SELECT\n\td1.NAME AS 部门,\n\td.NAME AS 二级部门,\n\tu.nickname AS 员工姓名,\n\tu.login_name AS 员工工号,\n\tCOUNT(\n\tIF\n\t( cqv.call_duration = 0, cqv.id, NULL )) AS 未接通话总数,\n\tSUM(\n\tIF\n\t( cqv.call_duration = 0, cqv.call_duration, 0 )) AS 未接通时长,\n\tCOUNT(\n\tIF\n\t( cqv.call_duration > 0 && cqv.call_duration < 60, cqv.id, NULL )) AS 无效通话总数,\n\tSUM(\n\tIF\n\t( cqv.call_duration > 0 && cqv.call_duration < 60, cqv.call_duration, 0 )) AS 无效通话时长,\n\tCOUNT(\n\tIF\n\t( cqv.call_duration >= 60 && cqv.call_duration < 180, cqv.id, NULL )) AS 有效通话总数,\n\tSUM(\n\tIF\n\t( cqv.call_duration >= 60 && cqv.call_duration < 180, cqv.call_duration, 0 )) AS 有效通话时长,\n\tCOUNT(\n\tIF\n\t( cqv.call_duration >= 180, cqv.id, NULL )) AS 优质通话总数,\n\tSUM(\n\tIF\n\t( cqv.call_duration >= 180, cqv.call_duration, 0 )) AS 优质通话时长,\n\tCOUNT( cqv.id ) AS 总通话数,\n\tSUM( cqv.call_duration ) AS 通话总时长 \nFROM\n\tbl_crm_quality_voice cqv\n\tINNER JOIN bl_users u ON cqv.user_id = u.id\n\tINNER JOIN bl_depart d ON u.depart_id = d.id\n\tINNER JOIN bl_depart d1 ON d.parent_id = d1.id \nWHERE\n\tstart_at > '%v 00:00:00' \n\tAND start_at < '%v 23:59:59' \n\tAND call_type = 'extension_outbound'\nAND d.id in (%v)\t\nGROUP BY\n\td1.NAME,\n\td.NAME,\n\tu.nickname,\n\tu.login_name \nORDER BY\n\td1.NAME,\n\td.NAME,\n\tu.nickname,\n\tu.login_name", Start_Time, Stop_Time, depart_id))
	}

	return resulist, nil
}

//新媒体线上明细
func (e *Engine) Order_xs_select(Start_Time, Stop_Time, login_name, client_Models string) ([]map[string][]byte, error) {
	err := e.NewEngine()
	defer e.Close()
	if err != nil {
		return nil, fmt.Errorf("%v，请联系管理员！", err)
	}
	var resulist []map[string][]byte
	if client_Models == "TsaveDialog" || login_name == "20030043" { //无限制
		resulist, _ = e.Engine.Query(fmt.Sprintf("SELECT\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) AS 线上下单时间,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"来源渠道\"' ) AS 线上下单渠道,\n\tcc.customer_code AS 客户编码,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"客户名称\"' ) AS 下单客户姓名,#JSON_UNQUOTE( mom.raw_json -> '$[0].\"客户电话\"' ) AS 下单客户电话,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"商品名称\"' ) AS 下单商品,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"渠道订单号\"' ) AS 渠道订单号,\n  IF(\tmom.`status` = 0,\"已取消\",IF(mom.`status` = 1,\"已创建\",IF(mom.`status` = 10,\"已分配\",IF(mom.`status` = 30,\"核单通过\",IF( mom.`status` = 40, \"核单失败\", \"0\" ))))) AS 核单状态,\n\tdi1.`name` AS 核单备注,\n\td.`name` AS 部门,\n\tu.nickname AS 员工姓名,\n\tu.login_name AS 员工工号,\n\tdi3.`name` AS 员工状态,\n\tmo.order_no AS 系统订单号,\n\tdi2.`name` AS 订单类型,\n\tdi.`name` AS 订单状态,\n\tFORMAT( mo.total_amount / 100 ,2) AS 订单总金额,\n\tFORMAT( mo.discount_amount / 100 ,2) AS 订单折扣总额,\n\tFORMAT(( mo.total_amount - mo.discount_amount )/ 100 ,2) AS 订单售价总额,\n\tmc.`name` AS 订单商品,\n\tmc.goods_count AS 商品数量,\n\tFORMAT( mc.goods_price / 100 ,2) AS 商品原价,\n\tFORMAT( mc.price_sale / 100 ,2) AS 商品售价,\n\tFORMAT( mc.total_price / 100 ,2) AS 商品总价,\n\tei.ship_channel_name AS 配送方式,\n\tei.ship_channel_name AS 配送方式,\n\tei.ship_channel_no AS 快递单号,\n\tei.last_trace AS 快递状态,\n\tmos.receiver_address AS 客户地址\nFROM\n\t`bl_mall_order_media` mom\n\tLEFT JOIN bl_crm_customer cc on cc.id = mom.customer_id\n\tLEFT JOIN bl_express_invoice ei ON mom.mall_order_id = ei.order_id\t\n\tLEFT JOIN (SELECT id,name,item_value FROM bl_dict_item where dict_id = 74) di1 ON mom.verify_order_comment = di1.item_value\n\tLEFT JOIN bl_users u ON mom.assign_user_id = u.id\n\tLEFT JOIN (SELECT id,name,item_value FROM bl_dict_item where dict_id = 14) di3 on u.`status` = di3.item_value\n\tLEFT JOIN bl_mall_order mo ON mom.mall_order_id = mo.id\n\tLEFT JOIN bl_depart d ON u.depart_id = d.id\t\n\tLEFT JOIN (SELECT id,name,item_value FROM bl_dict_item where dict_id = 43) di on mo.`status` = di.item_value\n\tLEFT JOIN (SELECT id,name,item_value FROM bl_dict_item where dict_id = 53) di2 on mo.type = di2.item_value\n\tLEFT JOIN bl_mall_cart mc ON mo.id = mc.order_id\n\tLEFT JOIN bl_mall_order_shipinfo mos on mo.id = mos.order_id\n\tWHERE\n\tmo.deleted_at IS NULL \n\tAND mo.delete_user_id IS NULL \n\tAND mom.deleted_at IS NULL\n\tAND mom.delete_user_id IS NULL\n\tAND mc.deleted_at IS NULL\n\tAND mc.delete_user_id IS NULL\n\tAND ei.deleted_at IS NULL\n\tAND ei.delete_user_id IS NULL \n\tAND  ((JSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) >= '%v 00:00:00' \tAND JSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) <= '%v 23:59:59' )\n\tOR (JSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) < '%v 00:00:00' AND mom.mall_order_id IS NOT NULL AND mo.`status` <= 90 AND mo.performance_at IS NULL) OR (JSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) < '%v 00:00:00' AND mom.mall_order_id IS  NULL AND mom.`status` > 30 ))\n\tORDER BY\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' )\n\t", Start_Time, Stop_Time, Start_Time, Start_Time))
	} else {

		position_level, err := e.Login_position_level(login_name)
		if err != nil {
			return nil, fmt.Errorf("%v，请联系管理员！", err)
		}
		if position_level <= "1" {
			resulist, _ = e.Engine.Query(fmt.Sprintf("SELECT\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) AS 线上下单时间,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"来源渠道\"' ) AS 线上下单渠道,\n\tcc.customer_code AS 客户编码,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"客户名称\"' ) AS 下单客户姓名,#JSON_UNQUOTE( mom.raw_json -> '$[0].\"客户电话\"' ) AS 下单客户电话,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"商品名称\"' ) AS 下单商品,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"渠道订单号\"' ) AS 渠道订单号,\n  IF(\tmom.`status` = 0,\"已取消\",IF(mom.`status` = 1,\"已创建\",IF(mom.`status` = 10,\"已分配\",IF(mom.`status` = 30,\"核单通过\",IF( mom.`status` = 40, \"核单失败\", \"0\" ))))) AS 核单状态,\n\tdi1.`name` AS 核单备注,\n\td.`name` AS 部门,\n\tu.nickname AS 员工姓名,\n\tu.login_name AS 员工工号,\n\tdi3.`name` AS 员工状态,\n\tmo.order_no AS 系统订单号,\n\tdi2.`name` AS 订单类型,\n\tdi.`name` AS 订单状态,\n\tFORMAT( mo.total_amount / 100 ,2) AS 订单总金额,\n\tFORMAT( mo.discount_amount / 100 ,2) AS 订单折扣总额,\n\tFORMAT(( mo.total_amount - mo.discount_amount )/ 100 ,2) AS 订单售价总额,\n\tmc.`name` AS 订单商品,\n\tmc.goods_count AS 商品数量,\n\tFORMAT( mc.goods_price / 100 ,2) AS 商品原价,\n\tFORMAT( mc.price_sale / 100 ,2) AS 商品售价,\n\tFORMAT( mc.total_price / 100 ,2) AS 商品总价,\n\tei.ship_channel_name AS 配送方式,\n\tei.ship_channel_name AS 配送方式,\n\tei.ship_channel_no AS 快递单号,\n\tei.last_trace AS 快递状态,\n\tmos.receiver_address AS 客户地址\nFROM\n\t`bl_mall_order_media` mom\n\tLEFT JOIN bl_crm_customer cc on cc.id = mom.customer_id\n\tLEFT JOIN bl_express_invoice ei ON mom.mall_order_id = ei.order_id\t\n\tLEFT JOIN (SELECT id,name,item_value FROM bl_dict_item where dict_id = 74) di1 ON mom.verify_order_comment = di1.item_value\n\tLEFT JOIN bl_users u ON mom.assign_user_id = u.id\n\tLEFT JOIN (SELECT id,name,item_value FROM bl_dict_item where dict_id = 14) di3 on u.`status` = di3.item_value\n\tLEFT JOIN bl_mall_order mo ON mom.mall_order_id = mo.id\n\tLEFT JOIN bl_depart d ON u.depart_id = d.id\t\n\tLEFT JOIN (SELECT id,name,item_value FROM bl_dict_item where dict_id = 43) di on mo.`status` = di.item_value\n\tLEFT JOIN (SELECT id,name,item_value FROM bl_dict_item where dict_id = 53) di2 on mo.type = di2.item_value\n\tLEFT JOIN bl_mall_cart mc ON mo.id = mc.order_id\n\tLEFT JOIN bl_mall_order_shipinfo mos on mo.id = mos.order_id\n\tWHERE\n\tmo.deleted_at IS NULL \n\tAND mo.delete_user_id IS NULL \n\tAND mom.deleted_at IS NULL\n\tAND mom.delete_user_id IS NULL\n\tAND mc.deleted_at IS NULL\n\tAND mc.delete_user_id IS NULL\n\tAND ei.deleted_at IS NULL\n\tAND ei.delete_user_id IS NULL \n\tAND  ((JSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) >= '%v 00:00:00' \tAND JSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) <= '%v 23:59:59' )\n\tOR (JSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) < '%v 00:00:00' AND mom.mall_order_id IS NOT NULL AND mo.`status` <= 90 AND mo.performance_at IS NULL) OR (JSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) < '%v 00:00:00' AND mom.mall_order_id IS  NULL AND mom.`status` > 30 ))\n\tAND u.login_name = '%v' \n\tORDER BY\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' )\n\t", Start_Time, Stop_Time, Start_Time, Start_Time, login_name))
		} else if position_level >= "2" {
			depart_id, err := e.LoginDepartIdPermissions(login_name)
			if err != nil {
				return nil, fmt.Errorf("%v，请联系管理员！", err)
			}
			resulist, _ = e.Engine.Query(fmt.Sprintf("SELECT\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) AS 线上下单时间,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"来源渠道\"' ) AS 线上下单渠道,\n\tcc.customer_code AS 客户编码,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"客户名称\"' ) AS 下单客户姓名,#JSON_UNQUOTE( mom.raw_json -> '$[0].\"客户电话\"' ) AS 下单客户电话,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"商品名称\"' ) AS 下单商品,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"渠道订单号\"' ) AS 渠道订单号,\n  IF(\tmom.`status` = 0,\"已取消\",IF(mom.`status` = 1,\"已创建\",IF(mom.`status` = 10,\"已分配\",IF(mom.`status` = 30,\"核单通过\",IF( mom.`status` = 40, \"核单失败\", \"0\" ))))) AS 核单状态,\n\tdi1.`name` AS 核单备注,\n\td.`name` AS 部门,\n\tu.nickname AS 员工姓名,\n\tu.login_name AS 员工工号,\n\tdi3.`name` AS 员工状态,\n\tmo.order_no AS 系统订单号,\n\tdi2.`name` AS 订单类型,\n\tdi.`name` AS 订单状态,\n\tFORMAT( mo.total_amount / 100 ,2) AS 订单总金额,\n\tFORMAT( mo.discount_amount / 100 ,2) AS 订单折扣总额,\n\tFORMAT(( mo.total_amount - mo.discount_amount )/ 100 ,2) AS 订单售价总额,\n\tmc.`name` AS 订单商品,\n\tmc.goods_count AS 商品数量,\n\tFORMAT( mc.goods_price / 100 ,2) AS 商品原价,\n\tFORMAT( mc.price_sale / 100 ,2) AS 商品售价,\n\tFORMAT( mc.total_price / 100 ,2) AS 商品总价,\n\tei.ship_channel_name AS 配送方式,\n\tei.ship_channel_name AS 配送方式,\n\tei.ship_channel_no AS 快递单号,\n\tei.last_trace AS 快递状态,\n\tmos.receiver_address AS 客户地址\nFROM\n\t`bl_mall_order_media` mom\n\tLEFT JOIN bl_crm_customer cc on cc.id = mom.customer_id\n\tLEFT JOIN bl_express_invoice ei ON mom.mall_order_id = ei.order_id\t\n\tLEFT JOIN (SELECT id,name,item_value FROM bl_dict_item where dict_id = 74) di1 ON mom.verify_order_comment = di1.item_value\n\tLEFT JOIN bl_users u ON mom.assign_user_id = u.id\n\tLEFT JOIN (SELECT id,name,item_value FROM bl_dict_item where dict_id = 14) di3 on u.`status` = di3.item_value\n\tLEFT JOIN bl_mall_order mo ON mom.mall_order_id = mo.id\n\tLEFT JOIN bl_depart d ON u.depart_id = d.id\t\n\tLEFT JOIN (SELECT id,name,item_value FROM bl_dict_item where dict_id = 43) di on mo.`status` = di.item_value\n\tLEFT JOIN (SELECT id,name,item_value FROM bl_dict_item where dict_id = 53) di2 on mo.type = di2.item_value\n\tLEFT JOIN bl_mall_cart mc ON mo.id = mc.order_id\n\tLEFT JOIN bl_mall_order_shipinfo mos on mo.id = mos.order_id\n\tWHERE\n\tmo.deleted_at IS NULL \n\tAND mo.delete_user_id IS NULL \n\tAND mom.deleted_at IS NULL\n\tAND mom.delete_user_id IS NULL\n\tAND mc.deleted_at IS NULL\n\tAND mc.delete_user_id IS NULL\n\tAND ei.deleted_at IS NULL\n\tAND ei.delete_user_id IS NULL \n\tAND  ((JSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) >= '%v 00:00:00' \tAND JSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) <= '%v 23:59:59' )\n\tOR (JSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) < '%v 00:00:00' AND mom.mall_order_id IS NOT NULL AND mo.`status` <= 90 AND mo.performance_at IS NULL) OR (JSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ) < '%v 00:00:00' AND mom.mall_order_id IS  NULL AND mom.`status` > 30 ))\n\tAND d.id in (%v)\n\tORDER BY\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' )\n\t", Start_Time, Stop_Time, Start_Time, Start_Time, depart_id))
		}
	}

	return resulist, nil
}

// NewOrderXsSelect 新媒体线上明细
func (e *Engine) NewOrderXsSelect(Start_Time, Stop_Time, login_name, client_Models string) ([]map[string][]byte, error) {

	err := e.CRMNewEngine()
	if err != nil {
		return nil, fmt.Errorf("%v，请联系管理员！", err)
	}
	var resulist []map[string][]byte

	resulist, _ = e.Engine.Query(fmt.Sprintf(" SELECT\n\tmmo.buy_at AS 线上下单时间,\n\tmmo.source AS 线上下单渠道,\n\tcc.`code` AS 客户编码,\n\tmmo.customer_name AS 下单客户姓名,\n\tmmo.origin_text  as 下单商品,\n\tCASE\t\t\n\t\tWHEN mmo.`status` = 3 THEN\t\t'核单成功' \n\t\tWHEN mmo.`status` = 4 THEN\t\t'核单失败' \n\t\tWHEN mmo.`status` = 2 THEN\t\t'已分配' \n\t\tELSE '待分配' \n\tEND  AS 核单状态,\n\tbdi.`name` AS 核单备注,\n\tso.`name` AS 部门,\n\tsu.real_name AS 员工姓名,\n\tsu.staff_no AS 员工工号,\n\tmo.order_no AS 系统订单号,\n\tCASE\t\t\n\t\tWHEN mo.`status` = 1  THEN\t\t'新建' \n\t\tWHEN mo.`status` = 2  THEN\t\t'待审核' \n\t\tWHEN mo.`status` = 3  THEN\t\t'代付款' \n\t\tWHEN mo.`status` = 4  THEN\t\t'审核被驳回' \n\t\tWHEN mo.`status` = 5  THEN\t\t'订单部分付款' \n\t\tWHEN mo.`status` = 6  THEN\t\t'待确认发货' \n\t\tWHEN mo.`status` = 7  THEN\t\t'等待备货' \n\t\tWHEN mo.`status` = 8  THEN\t\t'备货中' \n\t\tWHEN mo.`status` = 9  THEN\t\t'订单代发货' \n\t\tWHEN mo.`status` = 10 THEN\t\t'订单已发货' \n\t\tWHEN mo.`status` = 11 THEN\t\t'订单配送中' \n\t\tWHEN mo.`status` = 12 THEN\t\t'订单已签收' \n\t\tWHEN mo.`status` = 13 THEN\t\t'订单未妥投' \n\t\tWHEN mo.`status` = 14 THEN\t\t'商品已退回' \n\t\tWHEN mo.`status` = 15 THEN\t\t'订单售后中' \n\t\tWHEN mo.`status` = 16 THEN\t\t'订单已取消' \n\t\tWHEN mo.`status` = 17 THEN\t\t'订单已完成' \t\n\t\tWHEN mo.`status` = 18 THEN\t\t'售后已完成' \t\n\t\t#ELSE '' \n\tEND  AS 订单状态,\n\tmo.amount / 100 AS 订单应收金额,\n\tmo.discount_amount / 100 AS 订单优惠总额,\n\tmo.actual_amount / 100 AS 订单实收金额,\n\tmo.performance_at as 业绩时间,\n\tmog.sku_name AS 订单商品,\n\tmog.num AS 商品数量,\n\tmog.sale_price / 100 AS 商品原价,\n\tmog.actual_price / 100 AS 商品售价,\n\tmog.total_amount / 100 AS 商品总价,\n\tmo.express AS 配送方式,\n\tmo.logistic_no 快递单号,\n\tmo.logistic_state_desc 快递状态 \nFROM\n\tmall_media_order mmo\n\tLEFT JOIN mall_order mo ON mmo.order_id = mo.id\n\tLEFT JOIN mall_order_goods mog ON mo.id = mog.order_id\n\tLEFT JOIN sys_organ so ON mmo.handler_organ_id = so.id\n\tLEFT JOIN sys_user su ON mmo.handler_id = su.id \n\tLEFT JOIN bas_dictionary_item bdi on bdi.`value` = mmo.note\n\tLEFT JOIN cus_customer cc on cc.id = mmo.customer_id\n\tWHERE\n\t \tmog.`status` != 0 \n\tAND mmo.`status` != 0 and mmo.created_at BETWEEN '%v 00:00:00' AND '%v 23:59:59'\n\nORDER BY\n\tso.`name`,\n\tsu.id,\n\tmo.order_no,\n\tmmo.buy_at,\n\tmmo.source,\n\tmmo.`status`", Start_Time, Stop_Time))

	e.Close()
	//fmt.Println(resulist)
	return resulist, nil
}

//新媒体线上统计
func (e *Engine) Order_xsTJ_select(Start_Time, Stop_Time, login_name string) ([]map[string][]byte, error) {

	err := e.NewEngine()
	defer e.Close()
	if err != nil {
		return nil, fmt.Errorf("%v，请联系管理员！", err)
	}

	position_level, err := e.Login_position_level(login_name)
	if err != nil {
		return nil, fmt.Errorf("%v，请联系管理员！", err)
	}
	if position_level <= "1" {
		return nil, fmt.Errorf("权限不足！")
	}

	resulist, _ := e.Engine.Query(fmt.Sprintf("SELECT\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"来源渠道\"' ) AS 订单渠道,\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"商品名称\"' ) AS 投放产品,\n\tcount( DISTINCT mom.id ) AS 后台订单数,\n\tcount( DISTINCT mom.mall_order_no ) AS 发货数量,\n\tsum(JSON_UNQUOTE( mom.raw_json -> '$[0].\"总金额\"' )) AS 原订单总金额,\n\tsum(  mc.total_price/ 100) AS 成交金额\nFROM\n\t`bl_mall_order_media` mom\n\tLEFT JOIN bl_mall_order mo ON mom.mall_order_id = mo.id\n\tLEFT JOIN bl_mall_cart mc ON mo.id = mc.order_id\n\tLEFT JOIN bl_mall_goods mg ON mc.goods_id = mg.id \nWHERE\n\tDATE_FORMAT( JSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ), '%v' ) >= '%v 00:00:00' \n\tAND DATE_FORMAT(JSON_UNQUOTE( mom.raw_json -> '$[0].\"创建时间\"' ), '%v' ) < '%v 23:59:59' \n\tAND mom.deleted_at IS NULL \nGROUP BY\n\tJSON_UNQUOTE( mom.raw_json -> '$[0].\"来源渠道\"' ),\n\tJSON_UNQUOTE(\n\tmom.raw_json -> '$[0].\"商品名称\"' \n\t)", sss, Start_Time, sss, Stop_Time))

	return resulist, nil
}

//新媒体线下统计
func (e *Engine) Order_xxTJ_select(Start_Time, Stop_Time, login_name string) ([]map[string][]byte, error) {

	err := e.NewEngine()
	defer e.Close()
	if err != nil {
		return nil, fmt.Errorf("%v，请联系管理员！", err)
	}

	position_level, err := e.Login_position_level(login_name)
	if err != nil {
		return nil, fmt.Errorf("%v，请联系管理员！", err)
	}
	if position_level <= "1" {
		return nil, fmt.Errorf("权限不足！")
	}

	depart_id, err := e.LoginDepartIdPermissions(login_name)
	if err != nil {
		return nil, fmt.Errorf("%v，请联系管理员！", err)
	}

	resulist, _ := e.Engine.Query(fmt.Sprintf("SELECT\n\tmc.`name` AS 二开产品,\n\tCOUNT( DISTINCT mo.id ) AS 下单数, \n\tsum( mc.ratio_price / 100 ) AS 增购订单总金额,\n\tsum( mc.total_price / 100 ) AS 增购成交金额 \nFROM\n\t`bl_mall_order` mo\n\tLEFT JOIN bl_mall_cart mc ON mo.id = mc.order_id \nWHERE\n\tmo.created_at >= '%v 00:00:00' \n\tAND mo.created_at < '%v 23:59:59' \n\tAND mo.deleted_at IS NULL \n\tAND mo.id NOT IN ( SELECT mall_order_id FROM bl_mall_order_media WHERE mall_order_id IS NOT NULL ) \n\tand mo.performance_user_depart_id in ('%v')\n\tand mo.`status` > 1\nGROUP BY\n\tmc.`name`", Start_Time, Stop_Time, depart_id))

	return resulist, nil
}

//订单出库记录
func (e *Engine) Order_warehouse_select(Start_Time, Stop_Time string) ([]map[string][]byte, error) {
	err := e.NewEngine()
	defer e.Close()
	if err != nil {
		return nil, fmt.Errorf("%v，请联系管理员！", err)
	}

	resulist, _ := e.Engine.Query(fmt.Sprintf("SELECT\n\teic.created_at AS 出库日期,\n\teic.order_no AS 订单编号,\n\tIF( eic.`status` = 1 , \"订单\" , IF( eic.`status` = 2, \"工单\", \"\" )) AS 类型,\n\tdi.`name` as 订单类型,\n\tu.nickname AS 员工姓名,\n\tu.login_name AS 员工工号,\n\tdi1.`name` AS 员工状态,\n\td.`name` AS 员工部门,\n\teic.goods_name AS 商品名称,\n\teic.goods_count AS 商品数量,\n\tFLOOR(eic.amount/100) AS 商品售价,\n\teic.warehouse_name AS 出货仓库,\n\teic.ship_channel_name AS 配送方式,\n\teic.ship_channel_no AS 快递单号 \nFROM\n\tbl_express_invoice_cart eic\n\tLEFT JOIN bl_mall_order mo ON eic.order_id = mo.id\n\tLEFT JOIN bl_users u ON mo.performance_user_id = u.id\n\tLEFT JOIN (SELECT `name`,item_value FROM bl_dict_item WHERE dict_id = 14) di1 on di1.item_value = u.`status`\n\tLEFT JOIN bl_depart d ON mo.performance_user_depart_id = d.id\n\tLEFT JOIN (SELECT id,name,item_value FROM bl_dict_item where dict_id = 53) di on mo.type = di.item_value\t\nWHERE\n\teic.created_at >= '%v 00:00:00' \n\tAND eic.created_at < '%v 23:59:59'\n\tAND mo.`status` NOT IN (160,150)\n\tAND mo.deleted_at IS NULL \nORDER BY\n\teic.created_at,\n\teic.order_no,\n\teic.goods_name,\n\tu.nickname,\n\td.`name`", Start_Time, Stop_Time))

	return resulist, nil
}

//客户积分查询--暂未设置权限控制
func (e *Engine) Client_coin() ([]map[string][]byte, error) {
	err := e.NewEngine()
	defer e.Close()
	if err != nil {
		return nil, fmt.Errorf("%v，请联系管理员！", err)
	}

	resulist, _ := e.Engine.Query("SELECT\n\td1.NAME AS 部门,\n\td.NAME AS 二级部门,\n\tu.nickname 员工姓名,\n\tu.login_name AS 员工工号,\n\tdi.`name` AS 员工状态,\n\tcc.customer_code AS 客户编码,\n\tcc.NAME AS 客户姓名,\n\tcc.coin AS 客户积分\t\nFROM\n\tbl_crm_customer cc\n\tLEFT JOIN bl_crm_customer_phone ccp ON cc.id = ccp.customer_id\n\tLEFT JOIN bl_crm_customer_append cca ON cc.id = cca.customer_id\n\tLEFT JOIN bl_users u ON cc.user_id = u.id\n\tLEFT JOIN (SELECT `name`,item_value FROM bl_dict_item WHERE dict_id = 14) di on di.item_value = u.`status`\n\tLEFT JOIN bl_depart d ON u.depart_id = d.id\n\tLEFT JOIN bl_depart d1 ON d.parent_id = d1.id \nWHERE\n\tcca.first_category = 1 \n\tAND cc.deleted_at IS NULL \n\tAND cc.coin >0;")

	return resulist, nil
}

//客户生日--暂未设置权限控制
func (e *Engine) Client_birthday(Start_Time, Stop_Time, login_name, client_Models string) ([]map[string][]byte, error) {
	err := e.NewEngine()
	defer e.Close()
	if err != nil {
		return nil, fmt.Errorf("%v，请联系管理员！", err)
	}

	fmt.Println(Start_Time)

	resulist, _ := e.Engine.Query(fmt.Sprintf("SELECT\n\tcc.customer_code AS 客户编码,\n\tcc.`name` AS 客户姓名,\n\tDATE_FORMAT( cc.birthday, '%v' ) AS 客户生日,\n\td.`name` AS 员工部门,\n\tu.nickname AS 员工姓名,\n\tu.login_name AS 员工工号,\n\tdi.`name` AS 员工状态,\n\tFLOOR( mo.sum_amount ) AS 消费合计,\n\tFLOOR( mo.max_amount ) AS 最大单笔消费 \nFROM\n\tbl_crm_customer cc\n\tLEFT JOIN bl_crm_customer_append cca ON cc.id = cca.customer_id\n\tLEFT JOIN bl_users u ON cc.user_id = u.id\n\tLEFT JOIN (SELECT `name`,item_value FROM bl_dict_item WHERE dict_id = 14) di on di.item_value = u.`status`\n\tLEFT JOIN bl_depart d ON u.depart_id = d.id\n\tLEFT JOIN (SELECT\tcustomer_id,\tFLOOR(SUM( DISTINCT pay_amount / 100 )) AS sum_amount,FLOOR(MAX( DISTINCT pay_amount / 100 )) AS max_amount FROM\tbl_mall_order WHERE\tdeleted_at IS NULL \tAND `status` = 90 \tAND type = 1 GROUP BY\tcustomer_id ) AS mo ON mo.customer_id = cc.id \nWHERE\n\tMONTH ( cc.birthday ) = MONTH ( '%v' ) \n\tAND cca.first_category = 1", "%Y-%m-%d", Start_Time))

	return resulist, nil
}

//客户名单构成--暂未设置权限控制
func (e *Engine) Client_prestigious_university() ([]map[string][]byte, error) {
	err := e.NewEngine()
	defer e.Close()
	if err != nil {
		return nil, fmt.Errorf("%v，请联系管理员！", err)
	}

	resulist, _ := e.Engine.Query("SELECT\n\td1.`name` AS 部门,\n\td.`name` AS 二级部门,\n\tu.nickname AS 员工名字,\n\tu.login_name AS 员工工号,\n\tdi.`name` AS 员工状态,\n\tCOUNT(\n\tIF\n\t( cca.first_category = 1, cc.id, NULL )) AS 已购客户,\n\tCOUNT(\n\tIF\n\t( cca.first_category = 2, cc.id, NULL )) AS 准已购客户,\n\tCOUNT(\n\tIF\n\t( cca.first_category = 3, cc.id, NULL )) AS 意向客户,\n\tCOUNT(\n\tIF\n\t( cca.first_category = 4, cc.id, NULL )) AS 潜在客户 \nFROM\n\tbl_crm_customer cc\n\tLEFT JOIN bl_crm_customer_phone ccp ON cc.id = ccp.customer_id\n\tLEFT JOIN bl_crm_customer_append cca ON cc.id = cca.customer_id\n\tLEFT JOIN bl_users u ON cc.user_id = u.id\n\tLEFT JOIN (SELECT `name`,item_value FROM bl_dict_item WHERE dict_id = 14) di on di.item_value = u.`status`\n\tLEFT JOIN bl_depart d ON u.depart_id = d.id\n\tLEFT JOIN bl_depart d1 ON d.parent_id = d1.id \nWHERE\n\tcc.deleted_at IS NULL \nGROUP BY\n\td1.NAME,\n\td.NAME,\n\tu.nickname,\n\tu.login_name ")

	return resulist, nil
}
