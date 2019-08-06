package main

import (
	"fmt"
	"github.com/Luxurioust/excelize"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/wendal/go-oci8"
	"os"
	"strconv"
	"time"
)

const (
	dBconnect  = "BLCRM/BLCRM2012@61.136.101.122:1521/BLSD"
	driverName = "oci8"
	tbMapper   = "BLCRM."
)

var (
	tbMappers core.PrefixMapper
	Number    int
)

type Engine struct {
	Engine   *xorm.Engine
	Err      error
	ClientID string
}

//初始化环境
func init() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码
	tbMappers = core.NewPrefixMapper(core.SnakeMapper{}, tbMapper)
}

//新建连接
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

//统计当前库内客户的数量
func (e *Engine) CountNumber() error {
	Query, err := e.Engine.Query(fmt.Sprintf("SELECT COUNT(KHID) FROM BLCRM.CRM_DAT001"))

	if err != nil {
		return err
	}

	for _, v := range Query {
		for _, k := range v {
			if string(k) != "" {
				Number, _ = strconv.Atoi(string(k))
			}
		}
	}

	return nil
}

//查询客户ID，判断是否存在
func (e *Engine) SelectId(MPno string) error {
	Query, err := e.Engine.Query(fmt.Sprintf("SELECT KHID FROM CRM_DAT001 WHERE MOBIL = '%s'", MPno))

	if err != nil {
		return err
	}

	if len(Query) == 0 {
		e.ClientID = ""
		return nil
	}

	for _, v := range Query {
		for _, k := range v {
			if string(k) != "" {
				e.ClientID = string(k)
			}
		}
	}

	return nil
}

//新建客户  - 姓名 电话  地址  来源
func (e *Engine) InSetClient(KHMC string, MOBIL string, DIZHI string, SOURCEID int) error {

	Sql := fmt.Sprintf("INSERT INTO BLCRM.CRM_DAT001(KHMC,TYPEID,MOBIL,DIZHI,HUIFANG,HUIFANGJG,SOURCEID,ISVIP,RESERVE)"+
		" VALUES('%s',506,%s,'%s',1,5,%d,0,1)", KHMC, MOBIL, DIZHI, SOURCEID)

	_, e.Err = e.Engine.Exec(Sql)
	if e.Err != nil {
		return e.Err
	}
	return nil
}

//新建销售记录  随机的时间数  日期 客户的ID  产品ID
func (e *Engine) InSetConsume(TimeNum, Date, KHID, ConsumeID, num1, num2, num3 string) error {
	Sql := fmt.Sprintf("INSERT INTO BLCRM.CRM_DAT007(BILLNO,BILLTYPEID,BILLDATE,BILLFLAG,SALESMAN_ID,OPER_ID,CLIENT_ID,WAREHOUSESID,NEWOROLD)"+
		" VALUES(%s,140001,TO_DATE('%s','YYYY-MM-DD HH24:MI:SS'),1,1,1,%v,150001,'老客户')", TimeNum, Date, KHID)

	_, e.Err = e.Engine.Exec(Sql)
	if e.Err != nil {
		return e.Err
	}

	SqL := fmt.Sprintf("INSERT INTO BLCRM.CRM_DAT009(BILLNO,GOODSID,AMOUNT,SUMPRICE,PRICE,KIND)"+
		" VALUES(%s,%s,%s,%s,%s,'兑')", TimeNum, ConsumeID, num1, num2, num3)

	_, e.Err = e.Engine.Exec(SqL)
	if e.Err != nil {
		return e.Err
	}

	return nil
}

func main() {
	num := 1
	Engine := new(Engine)
	err := Engine.NewEngine()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer Engine.Engine.Close()

	f, err := excelize.OpenFile("./123.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	rows := f.GetRows("Sheet1")
	for i, data := range rows {
		if i == 0 {
			continue
		}

		err := Engine.SelectId(data[4])
		if err != nil {
			fmt.Println("客户信息查询失败：", err, data[4])
			continue
		}
		if Engine.ClientID == "" {
			continue
			err := Engine.InSetClient(data[3], data[4], data[5], 60004)
			if err != nil {
				fmt.Println("创建客户信息失败：", data[4])
			}
			err = Engine.SelectId(data[4])
			if err != nil {
				fmt.Println("客户信息查询失败2：", err, data[4])
				continue
			}
		}

		timeNum := fmt.Sprintf("%2d%02d%02d%02d%02d%02d%03d\n", 19,
			int(time.Now().Month()), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), num)
		err = Engine.InSetConsume(timeNum, fmt.Sprintf("%s-%0s-%s", data[0], data[1], data[2]), Engine.ClientID, data[6], data[7], data[8], data[9])
		if err != nil {
			fmt.Println("购买记录写入失败：3", err, data[4])
			continue
		}
		num++
		if num == 999 {
			num = 1
		}

	}
}

//func main() {
//	f, _ := excelize.OpenFile("./123.xlsx")
//
//	rows := f.GetRows("Sheet1")
//
//	for _, v := range rows {
//		fmt.Println(v)
//		time.Sleep(time.Second * 2)
//	}
//}
