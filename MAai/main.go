package main

import (
	"fmt"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/wendal/go-oci8"
	"os"
	"strconv"
)

const (
	dBconnect  = "BLCRM/BLCRM2012@10.10.57.20:1521/BLDB"
	driverName = "oci8"
	tbMapper   = "BLCRM."
)

var (
	tbMappers core.PrefixMapper
	KHID      int
	Number    int
)

type Engine struct {
	Engine *xorm.Engine
	Err    error
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
	//e.Engine.ShowSQL(true)
	e.Engine.SetTableMapper(tbMappers)
	return nil
}

func (e *Engine) CountNumber() error {
	Query, err := e.Engine.Query(fmt.Sprintf("SELECT COUNT(KHID) FROM BLCRM.CRM_DAT001"))
	if err != nil {
		return nil
	}

	for _, v := range Query {
		for _, k := range v {
			if string(k) != "" {
				KHID, _ = strconv.Atoi(string(k))
			}
		}
	}

	return nil
}

func (e *Engine) SelectId() error {

	Query, err := e.Engine.Query(fmt.Sprintf("SELECT KHID FROM CRM_DAT001 WHERE MOBIL = '17802928284'"))
	if err != nil {
		return err
	}

	for _, v := range Query {
		for _, k := range v {
			if string(k) != "" {
				KHID, _ = strconv.Atoi(string(k))
			}
		}
	}

	return nil
}

func main() {
	Engine := new(Engine)
	Engine.Err = Engine.NewEngine()
	defer Engine.Engine.Close()
	if Engine.Err != nil {
		fmt.Println("1", Engine.Err)
		return
	}

	err := Engine.SelectId()
	if err != nil {
		fmt.Println("2", err)
	}
	fmt.Println(KHID)
}
