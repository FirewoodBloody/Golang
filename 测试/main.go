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

type count_sun struct {
	SumHeji  string
	Counts   string
	Zengsong string
	Tuihuo   string
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

func (e *Engine) Select09(id string) *count_sun {
	as := new(count_sun)

	res1, _ := e.Engine.Query(fmt.Sprintf("SELECT SUM(SUMPRICE) FROM CRM_DAT009_VIEW009 WHERE CLIENT_ID = %v AND KIND = '售' OR CLIENT_ID = %v AND KIND = '兑'", id, id))

	for _, v := range res1 {
		as.SumHeji = string(v["SUM(SUMPRICE)"])
	}

	res2, _ := e.Engine.Query(fmt.Sprintf("SELECT COUNT(SUMPRICE) FROM CRM_DAT009_VIEW009 WHERE CLIENT_ID = %v AND KIND = '售' OR CLIENT_ID = %v AND KIND = '兑'", id, id))
	for _, v := range res2 {
		as.Counts = string(v["COUNT(SUMPRICE)"])
	}

	res3, _ := e.Engine.Query(fmt.Sprintf("SELECT SUM(SUMPRICE) FROM CRM_DAT009_VIEW009 WHERE CLIENT_ID = %v AND KIND = '赠' ", id))
	for _, v := range res3 {
		as.Zengsong = string(v["SUM(SUMPRICE)"])
	}

	res4, _ := e.Engine.Query(fmt.Sprintf("SELECT SUM(SUMPRICE) FROM CRM_DAT009_VIEW009 WHERE CLIENT_ID = %v AND KIND = '退' ", id))
	for _, v := range res4 {
		as.Tuihuo = string(v["SUM(SUMPRICE)"])
	}

	return as
}

func main() {

	file, _ := excelize.OpenFile("./客户模板.xlsx")
	rows := file.GetRows("Sheet1")

	e := Engine{}

	s := 10
	n := 5

	for i := 2; i < len(rows)+1; i++ {

		e.NewEngine()

		res, err := e.Engine.Query(fmt.Sprintf("SELECT KHMC||'',KHID,TO_CHAR(BIRTHDAY,'YYYY-MM-DD'),AGE,SEX,YWY||'',TO_CHAR(XCYY,'YYYY-MM-DD'),RESERVE,SOURCESTR,TO_CHAR(FPRQ,'YYYY-MM-DD'),CONSUME_NATURES,JIFEN,TO_CHAR(SCGMRQ,'YYYY-MM-DD'),MAX_DBGM,TOTALPRICE FROM CRM_DAT001_VIEW001 WHERE KHID = %v ", file.GetCellValue("Sheet1", fmt.Sprintf("A%v", i))))
		//res, err := e.Engine.Query("SELECT KHID,DIZHI||'' FROM BLCRM.CRM_DAT001 WHERE KHID = 7385707")
		if err != nil {
			fmt.Println(err)
		}
		as := e.Select09(file.GetCellValue("Sheet1", fmt.Sprintf("A%v", i)))
		e.Engine.Close()
		for _, v := range res {

			//file.SetCellValue("Sheet1", fmt.Sprintf("A%v", i), string(v["KHID"]))
			//file.SetCellValue("Sheet1", fmt.Sprintf("B%v", i), string(v["KHID"]))
			file.SetCellValue("Sheet1", fmt.Sprintf("C%v", i), string(v["KHMC||''"]))
			file.SetCellValue("Sheet1", fmt.Sprintf("D%v", i), string(v["TO_CHAR(BIRTHDAY,'YYYY-MM-DD')"]))
			file.SetCellValue("Sheet1", fmt.Sprintf("F%v", i), string(v["SEX"]))
			file.SetCellValue("Sheet1", fmt.Sprintf("I%v", i), string(v["YWY||''"]))
			file.SetCellValue("Sheet1", fmt.Sprintf("J%v", i), string(v["TO_CHAR(XCYY,'YYYY-MM-DD')"]))
			file.SetCellValue("Sheet1", fmt.Sprintf("K%v", i), string(v["RESERVE"]))
			//file.SetCellValue("Sheet1", fmt.Sprintf("L%v", i), string(v["SOURCESTR"]))
			file.SetCellValue("Sheet1", fmt.Sprintf("M%v", i), string(v["TO_CHAR(FPRQ,'YYYY-MM-DD')"]))
			file.SetCellValue("Sheet1", fmt.Sprintf("N%v", i), string(v["CONSUME_NATURES"]))
			file.SetCellValue("Sheet1", fmt.Sprintf("O%v", i), string(v["JIFEN"]))
			file.SetCellValue("Sheet1", fmt.Sprintf("Q%v", i), string(v["TO_CHAR(SCGMRQ,'YYYY-MM-DD')"]))
			file.SetCellValue("Sheet1", fmt.Sprintf("S%v", i), string(v["MAX_DBGM"]))
			file.SetCellValue("Sheet1", fmt.Sprintf("T%v", i), string(v["TOTALPRICE"]))

			switch string(v["SOURCESTR"]) {
			case "客户中心未购":

				file.SetCellValue("Sheet1", fmt.Sprintf("L%v", i), "自拓展")

				if as.SumHeji == "" || as.SumHeji == "0" {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "意向客户")
					file.SetCellValue("Sheet1", fmt.Sprintf("N%v", i), "钱币类")
				} else {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "已购客户")
				}

				file.SetCellValue("Sheet1", fmt.Sprintf("T%v", i), as.SumHeji)
				file.SetCellValue("Sheet1", fmt.Sprintf("U%v", i), as.Counts)
				file.SetCellValue("Sheet1", fmt.Sprintf("V%v", i), as.Zengsong)
				file.SetCellValue("Sheet1", fmt.Sprintf("X%v", i), as.Tuihuo)

			case "客户中心已购":
				file.SetCellValue("Sheet1", fmt.Sprintf("L%v", i), "自拓展")
				//as := e.Select09(string(v["KHID"]))
				if as.SumHeji == "" || as.SumHeji == "0" {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "准已购")
					file.SetCellValue("Sheet1", fmt.Sprintf("N%v", i), "钱币类")
				} else {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "已购客户")
				}

				file.SetCellValue("Sheet1", fmt.Sprintf("T%v", i), as.SumHeji)
				file.SetCellValue("Sheet1", fmt.Sprintf("U%v", i), as.Counts)
				file.SetCellValue("Sheet1", fmt.Sprintf("V%v", i), as.Zengsong)
				file.SetCellValue("Sheet1", fmt.Sprintf("X%v", i), as.Tuihuo)
			case "电商已购":
				file.SetCellValue("Sheet1", fmt.Sprintf("L%v", i), "电商")
				file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "天猫")
				//as := e.Select09(string(v["KHID"]))
				if as.SumHeji == "" || as.SumHeji == "0" {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "准已购")
					file.SetCellValue("Sheet1", fmt.Sprintf("N%v", i), "钱币类")
				} else {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "已购客户")
				}

				file.SetCellValue("Sheet1", fmt.Sprintf("T%v", i), as.SumHeji)
				file.SetCellValue("Sheet1", fmt.Sprintf("U%v", i), as.Counts)
				file.SetCellValue("Sheet1", fmt.Sprintf("V%v", i), as.Zengsong)
				file.SetCellValue("Sheet1", fmt.Sprintf("X%v", i), as.Tuihuo)
			case "店面白单来访":
				s++
				file.SetCellValue("Sheet1", fmt.Sprintf("L%v", i), "专卖店来访")
				if s%10 == 1 {
					file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "新华书店大厦店")
				} else if s%10 == 2 {
					file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "新华书店钟楼店")
				} else if s%10 == 3 {
					file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "汉唐小寨店")
				} else if s%10 == 4 {
					file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "汉唐高新店")
				} else if s%10 == 5 {
					file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "卜蜂高新店")
				} else if s%10 == 6 {
					file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "人人乐电子城店")
				} else if s%10 == 7 {
					file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "大雁塔店")
				} else if s%10 == 8 {
					file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "新华书店曲江店")
				} else if s%10 == 9 {
					file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "新华书店四海店")
				} else if s%10 == 0 {
					file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "军区小寨店")
				}

				//as := e.Select09(string(v["KHID"]))
				if as.SumHeji == "" || as.SumHeji == "0" {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "潜在客户")
					//file.SetCellValue("Sheet1", fmt.Sprintf("N%v", i), "钱币类")
				} else {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "已购客户")
				}

				file.SetCellValue("Sheet1", fmt.Sprintf("T%v", i), as.SumHeji)
				file.SetCellValue("Sheet1", fmt.Sprintf("U%v", i), as.Counts)
				file.SetCellValue("Sheet1", fmt.Sprintf("V%v", i), as.Zengsong)
				file.SetCellValue("Sheet1", fmt.Sprintf("X%v", i), as.Tuihuo)
			case "店面来访登记":
				n++
				file.SetCellValue("Sheet1", fmt.Sprintf("L%v", i), "专卖店")
				if n%10 == 1 {
					file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "大厦店")
				} else if n%10 == 2 {
					file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "四海店")
				} else if n%10 == 3 {
					file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "卜蜂店")
				} else if n%10 == 4 {
					file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "军区店")
				} else if n%10 == 0 {
					file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "曲江店")
				}

				as := e.Select09(string(v["KHID"]))
				if as.SumHeji == "" || as.SumHeji == "0" {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "意向客户")
					file.SetCellValue("Sheet1", fmt.Sprintf("N%v", i), "钱币类")
				} else {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "已购客户")
				}

				file.SetCellValue("Sheet1", fmt.Sprintf("T%v", i), as.SumHeji)
				file.SetCellValue("Sheet1", fmt.Sprintf("U%v", i), as.Counts)
				file.SetCellValue("Sheet1", fmt.Sprintf("V%v", i), as.Zengsong)
				file.SetCellValue("Sheet1", fmt.Sprintf("X%v", i), as.Tuihuo)
			case "新媒体进线":
				file.SetCellValue("Sheet1", fmt.Sprintf("L%v", i), "新媒体")
				file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "头条")
				//as := e.Select09(string(v["KHID"]))
				if as.SumHeji == "" || as.SumHeji == "0" {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "意向客户")
					file.SetCellValue("Sheet1", fmt.Sprintf("N%v", i), "钱币类")
				} else {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "已购客户")
				}

				file.SetCellValue("Sheet1", fmt.Sprintf("T%v", i), as.SumHeji)
				file.SetCellValue("Sheet1", fmt.Sprintf("U%v", i), as.Counts)
				file.SetCellValue("Sheet1", fmt.Sprintf("V%v", i), as.Zengsong)
				file.SetCellValue("Sheet1", fmt.Sprintf("X%v", i), as.Tuihuo)
			case "新媒体未妥投":
				file.SetCellValue("Sheet1", fmt.Sprintf("L%v", i), "新媒体")
				file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "头条")
				//as := e.Select09(string(v["KHID"]))
				if as.SumHeji == "" || as.SumHeji == "0" {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "准已购")
					file.SetCellValue("Sheet1", fmt.Sprintf("N%v", i), "钱币类")
				} else {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "已购客户")
				}

				file.SetCellValue("Sheet1", fmt.Sprintf("T%v", i), as.SumHeji)
				file.SetCellValue("Sheet1", fmt.Sprintf("U%v", i), as.Counts)
				file.SetCellValue("Sheet1", fmt.Sprintf("V%v", i), as.Zengsong)
				file.SetCellValue("Sheet1", fmt.Sprintf("X%v", i), as.Tuihuo)
			case "活动引流":
				file.SetCellValue("Sheet1", fmt.Sprintf("L%v", i), "历史活动")
				//file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "头条")
				//as := e.Select09(string(v["KHID"]))
				if as.SumHeji == "" || as.SumHeji == "0" {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "意向客户")
					file.SetCellValue("Sheet1", fmt.Sprintf("N%v", i), "钱币类")
				} else {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "已购客户")
				}

				file.SetCellValue("Sheet1", fmt.Sprintf("T%v", i), as.SumHeji)
				file.SetCellValue("Sheet1", fmt.Sprintf("U%v", i), as.Counts)
				file.SetCellValue("Sheet1", fmt.Sprintf("V%v", i), as.Zengsong)
				file.SetCellValue("Sheet1", fmt.Sprintf("X%v", i), as.Tuihuo)
			case "新媒体已购":
				file.SetCellValue("Sheet1", fmt.Sprintf("L%v", i), "新媒体")
				file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "头条")
				//as := e.Select09(string(v["KHID"]))
				if as.SumHeji == "" || as.SumHeji == "0" {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "准已购")
					file.SetCellValue("Sheet1", fmt.Sprintf("N%v", i), "钱币类")
				} else {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "已购客户")
				}

				file.SetCellValue("Sheet1", fmt.Sprintf("T%v", i), as.SumHeji)
				file.SetCellValue("Sheet1", fmt.Sprintf("U%v", i), as.Counts)
				file.SetCellValue("Sheet1", fmt.Sprintf("V%v", i), as.Zengsong)
				file.SetCellValue("Sheet1", fmt.Sprintf("X%v", i), as.Tuihuo)
			case "业内已购":
				file.SetCellValue("Sheet1", fmt.Sprintf("L%v", i), "公司拓展")
				//file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "头条")
				//as := e.Select09(string(v["KHID"]))
				if as.SumHeji == "" || as.SumHeji == "0" {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "意向客户")
					file.SetCellValue("Sheet1", fmt.Sprintf("N%v", i), "钱币类")
				} else {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "已购客户")
				}

				file.SetCellValue("Sheet1", fmt.Sprintf("T%v", i), as.SumHeji)
				file.SetCellValue("Sheet1", fmt.Sprintf("U%v", i), as.Counts)
				file.SetCellValue("Sheet1", fmt.Sprintf("V%v", i), as.Zengsong)
				file.SetCellValue("Sheet1", fmt.Sprintf("X%v", i), as.Tuihuo)
			case "2020鼠年生肖纪念币兑换":
				file.SetCellValue("Sheet1", fmt.Sprintf("L%v", i), "新媒体")
				file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "头条")
				//as := e.Select09(string(v["KHID"]))
				if as.SumHeji == "" || as.SumHeji == "0" {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "准已购")
					file.SetCellValue("Sheet1", fmt.Sprintf("N%v", i), "钱币类")
				} else {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "已购客户")
				}

				file.SetCellValue("Sheet1", fmt.Sprintf("T%v", i), as.SumHeji)
				file.SetCellValue("Sheet1", fmt.Sprintf("U%v", i), as.Counts)
				file.SetCellValue("Sheet1", fmt.Sprintf("V%v", i), as.Zengsong)
				file.SetCellValue("Sheet1", fmt.Sprintf("X%v", i), as.Tuihuo)
			case "新媒体本地已购":
				file.SetCellValue("Sheet1", fmt.Sprintf("L%v", i), "新媒体")
				file.SetCellValue("Sheet1", fmt.Sprintf("Y%v", i), "头条")
				//as := e.Select09(string(v["KHID"]))
				if as.SumHeji == "" || as.SumHeji == "0" {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "准已购")
					file.SetCellValue("Sheet1", fmt.Sprintf("N%v", i), "钱币类")
				} else {
					file.SetCellValue("Sheet1", fmt.Sprintf("G%v", i), "已购客户")
				}

				file.SetCellValue("Sheet1", fmt.Sprintf("T%v", i), as.SumHeji)
				file.SetCellValue("Sheet1", fmt.Sprintf("U%v", i), as.Counts)
				file.SetCellValue("Sheet1", fmt.Sprintf("V%v", i), as.Zengsong)
				file.SetCellValue("Sheet1", fmt.Sprintf("X%v", i), as.Tuihuo)
			}
			i++

		}
	}
	file.Save()
}
