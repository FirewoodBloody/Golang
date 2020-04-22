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

type Engine struct {
	Engine *xorm.Engine
	Err    error
}

const (
	TimeFormat = "2006-01-02"
	driverName = "oci8"
	dBconnect  = "BLCRM/BLCRM2012@192.168.0.9:1521/BLDB"
	tbMapper   = "BLCRM."
)

var tbMappers core.PrefixMapper

func init() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码
	tbMappers = core.NewPrefixMapper(core.SnakeMapper{}, tbMapper)
}

//初始化化
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

//商品ID查询
func (e *Engine) GetCpId(commodity string) (string, error) {
	//查询商品ID
	var num string
	row, err := e.Engine.Query(fmt.Sprintf("SELECT ID FROM BLCRM.CRM_DAT006 WHERE CPMC = '%v' ", commodity)) //查询商品ID

	if err != nil {
		fmt.Println("00:", err)
		return "", err
	}

	for _, v := range row {
		for _, s := range v {
			if string(s) != "" {
				fmt.Println(string(s), "1err:", err)
				num = string(s)
			}
		}
	}
	fmt.Println(string(num), "2err:", err)
	return num, err
}

//插入购买记录
func (e *Engine) InSert(file *excelize.File, number int) {
	var commodityID string
	num := number
	commodityID, err := e.GetCpId(file.GetCellValue("Sheet2", fmt.Sprintf("P%v", number)))

	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Second * 5)
		return
	} else if commodityID == "" {
		e.Engine.Exec(fmt.Sprintf("INSERT INTO BLCRM.CRM_DAT006(GOODSTYPEID,CPMC,VISIBLED) VALUES('%v','%v','%v')",
			file.GetCellValue("Sheet2", fmt.Sprintf("K%v", number)),
			file.GetCellValue("Sheet2", fmt.Sprintf("P%v", number)), 1))
	}

	commodityID, err = e.GetCpId(file.GetCellValue("Sheet2", fmt.Sprintf("P%v", number)))
	fmt.Println(commodityID, "3err:", err)
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Second * 5)
		return
	} else if commodityID == "" {
		fmt.Println(commodityID, 1111)
		time.Sleep(time.Second * 5)
		return
	}

	fmt.Println(commodityID)

	if num >= 1000 {
		num = number % 1000
	}

	//随机事件数
	timeNum := fmt.Sprintf("%2d%02d%02d%02d%02d%02d%03d\n", 20,
		int(time.Now().Month()), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), num)
	date := convertToFormatDay(file.GetCellValue("Sheet2", fmt.Sprintf("B%v", number)))

	err = e.InSetConsume(timeNum, date, file.GetCellValue("Sheet2", fmt.Sprintf("G%v", number)),
		file.GetCellValue("Sheet2", fmt.Sprintf("R%v", number)),
		file.GetCellValue("Sheet2", fmt.Sprintf("L%v", number)),
		file.GetCellValue("Sheet2", fmt.Sprintf("M%v", number)),
		file.GetCellValue("Sheet2", fmt.Sprintf("N%v", number)),
		file.GetCellValue("Sheet2", fmt.Sprintf("O%v", number)), commodityID)
	if err != nil {
		fmt.Println(err)
		time.Sleep(time.Second * 5)
		return
	}

}

//新建销售记录  随机的时间数  日期 客户的ID  产品ID
func (e *Engine) InSetConsume(TimeNum, Date, KHID, ConsumeID, num, sum, one, typestr, commodityID string) error {
	Sql := fmt.Sprintf("INSERT INTO BLCRM.CRM_DAT007(BILLNO,BILLTYPEID,BILLDATE,BILLFLAG,SALESMAN_ID,OPER_ID,CLIENT_ID,WAREHOUSESID,NEWOROLD)"+
		" VALUES(%s,140001,TO_DATE('%s','YYYY-MM-DD HH24:MI:SS'),1,%v,1973,%v,150001,'老客户')", TimeNum, Date, ConsumeID, KHID)

	_, e.Err = e.Engine.Exec(Sql)
	if e.Err != nil {
		return e.Err
	}

	SqL := fmt.Sprintf("INSERT INTO BLCRM.CRM_DAT009(BILLNO,GOODSID,AMOUNT,SUMPRICE,PRICE,KIND)"+
		" VALUES(%s,%s,%s,%s,%s,'%v')", TimeNum, commodityID, num, sum, one, typestr)

	_, e.Err = e.Engine.Exec(SqL)
	if e.Err != nil {
		return e.Err
	}

	return nil
}

// excel日期字段格式化 yyyy-mm-dd
func convertToFormatDay(excelDaysString string) string {
	// 2006-01-02 距离 1900-01-01的天数
	baseDiffDay := 38719 //在网上工具计算的天数需要加2天，什么原因没弄清楚
	curDiffDay := excelDaysString
	b, _ := strconv.Atoi(curDiffDay)
	// 获取excel的日期距离2006-01-02的天数
	realDiffDay := b - baseDiffDay
	//fmt.Println("realDiffDay:",realDiffDay)
	// 距离2006-01-02 秒数
	realDiffSecond := realDiffDay * 24 * 3600
	//fmt.Println("realDiffSecond:",realDiffSecond)
	// 2006-01-02 15:04:05距离1970-01-01 08:00:00的秒数 网上工具可查出
	baseOriginSecond := 1136185445
	resultTime := time.Unix(int64(baseOriginSecond+realDiffSecond), 0).Format("2006-01-02")
	return resultTime
}

//查询员工和地址
func (e *Engine) SelectAdd(file *excelize.File, number int) (string, string, string) {

	var name, depname, KHID string

	_ = e.NewEngine()
	row, err := e.Engine.Query(fmt.Sprintf("SELECT TO_NCHAR(YWY) FROM CRM_DAT001_VIEW001 WHERE MOBIL = '%v'", file.GetCellValue("Sheet1", fmt.Sprintf("C%v", number))))
	if err != nil {
		fmt.Println(err)
		return name, depname, KHID
	}
	for _, v := range row {
		for _, s := range v {
			if string(s) != "" {
				name = string(s)
			}
		}
	}

	row, err = e.Engine.Query(fmt.Sprintf("SELECT TO_NCHAR(DEPNAME) FROM CRM_DAT001_VIEW001 WHERE MOBIL = '%v'", file.GetCellValue("Sheet1", fmt.Sprintf("C%v", number))))
	if err != nil {
		fmt.Println(err)
		return name, depname, KHID
	}
	for _, v := range row {
		for _, s := range v {
			if string(s) != "" {
				depname = string(s)
			}
		}
	}

	row, err = e.Engine.Query(fmt.Sprintf("SELECT KHID FROM CRM_DAT001 WHERE MOBIL = '%v'", file.GetCellValue("Sheet1", fmt.Sprintf("C%v", number))))
	if err != nil {
		fmt.Println(err)
		return name, depname, KHID
	}
	for _, v := range row {
		for _, s := range v {
			if string(s) != "" {
				KHID = string(s)
			}
		}
	}

	return name, depname, KHID
}

func main() {
	//file, err := excelize.OpenFile("报名表单.xlsx")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//row := file.GetRows("Sheet1")
	//
	//e := new(Engine)
	//err = e.NewEngine()
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//defer e.Engine.Close()
	//
	//for i, _ := range row {
	//	if i == 0 {
	//		continue
	//	}
	//	a, b, C := e.SelectAdd(file, i+1)
	//
	//	file.SetCellValue("Sheet1", fmt.Sprintf("E%v", i+1), a)
	//	file.SetCellValue("Sheet1", fmt.Sprintf("F%v", i+1), b)
	//	file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i+1), C)
	//}
	//
	//file.Save()
	e := Engine{}
	e.NewEngine()
	res, err := e.Engine.Query("SELECT DIZHI FROM BLCRM.CRM_DAT001 WHERE KHID = 7502723")
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range res {
		for _, data := range v {
			fmt.Println(string(data))
		}
	}
}
