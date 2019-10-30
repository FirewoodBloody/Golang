package main

import (
	"database/sql"
	"fmt"
	_ "github.com/wendal/go-oci8"
	"log"
	"os"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
	driverName = "oci8"
	dBconnect  = "BLCRM/BLCRM2012@192.168.0.9:1521/BLDB"
	tbMapper   = "BLCRM."
)

//var tbMappers core.PrefixMapper
//
//type Engine struct {
//	Engine *xorm.Engine
//	Err    error
//	//	GetDb  []Kdlyzt
//}

func init() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码
	//tbMappers = core.NewPrefixMapper(core.SnakeMapper{}, tbMapper)
}

//初始化化
//func (e *Engine) NewEngine() error {
//
//	e.Engine, e.Err = xorm.NewEngine(driverName, dBconnect)
//	if e.Err != nil {
//		return e.Err
//	}
//	//tbMappe := core.NewPrefixMapper(core.SnakeMapper{}, tbMapper)
//	//e.Engine.ShowSQL(true)
//	e.Engine.SetTableMapper(tbMappers)
//	return nil
//}

var KHID, KHMC, TYPEID, LN1, LN2, BIRTHDAY, AGE, SEX, MOBIL, TEL, FAX, GONGHAO, DEPNAME, DIZHI, VIP_JBSTR, CONSUME_NATURES, SOURCESTR string

func main() {

	////f, err := excelize.OpenFile("./123.xlsx")
	//
	//Engine := new(Engine)
	//err := Engine.NewEngine()
	//if err != nil {
	//	fmt.Println(1, err)
	//}
	//
	//Query, err := Engine.Engine.Query(fmt.Sprintf("SELECT KHID,KHMC,TYPEID,LN1,LN2,BIRTHDAY,AGE,SEX,MOBIL,TEL,FAX,GONGHAO," +
	//	"DEPNAME,DIZHI,VIP_JBSTR,CONSUME_NATURES,SOURCESTR FROM CRM_DAT001_VIEW001"))
	//if err != nil {
	//	fmt.Println(2, err)
	//}
	//
	//if len(Query) == 0 {
	//	fmt.Println(3, err)
	//}
	//
	//for _, v := range Query {
	//
	//	fmt.Println(v["AGE"])
	//}

	db, err := sql.Open("oci8", dBconnect)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("SELECT KHID,KHMC,TYPEID,LN1,LN2,BIRTHDAY,AGE,SEX,MOBIL,TEL,FAX,GONGHAO," +
		"DEPNAME,DIZHI,VIP_JBSTR,CONSUME_NATURES,SOURCESTR FROM CRM_DAT001_VIEW001 WHERE TYPEID = ANY(501,502,503,504,505)")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&KHID, &KHMC, &TYPEID, &LN1, &LN2, &BIRTHDAY, &AGE, &SEX, &MOBIL, &TEL, &FAX, &GONGHAO, &DEPNAME, &DIZHI, &VIP_JBSTR, &CONSUME_NATURES, &SOURCESTR)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(KHID, KHMC, TYPEID, LN1, LN2, BIRTHDAY, AGE, SEX, MOBIL, TEL, FAX, GONGHAO, DEPNAME, DIZHI, VIP_JBSTR, CONSUME_NATURES, SOURCESTR)
	}
	rows.Close()
}
