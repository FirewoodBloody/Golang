package main

import (
	"fmt"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/wendal/go-oci8"
	"os"
	"time"
)

const (
	TimeFormat = "2006-01-02"
	driverName = "oci8"
	dBconnect  = "BLCRM/BLCRM2012@192.168.0.9:1521/BLDB"
	tbMapper   = "BLCRM."
)

type Engine struct {
	Engine *xorm.Engine
	Err    error
	//FileName []crm_dat101
}

var tbMappers core.PrefixMapper

func init() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码
	tbMappers = core.NewPrefixMapper(core.SnakeMapper{}, tbMapper)
}

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

func main() {
	fmt.Printf("%2s", time.Now().Format("06"))
}
