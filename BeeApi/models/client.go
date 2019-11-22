package models

import (
	"fmt"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/wendal/go-oci8"
	"os"
	"strings"
)

type ClientMessage struct {
	ClientId    string
	Name        string
	PhoneMumber string
	Site        string
	Error       string
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

type Engine struct {
	Engine *xorm.Engine
	Err    error
}

func (e *Engine) NewEngine() error {

	e.Engine, e.Err = xorm.NewEngine(driverName, dBconnect)
	if e.Err != nil {
		return e.Err
	}
	//tbMappe := core.NewPrefixMapper(core.SnakeMapper{}, tbMapper)
	e.Engine.ShowSQL(false)
	e.Engine.SetTableMapper(tbMappers)
	return nil
}

func GetClientMessage(id string) map[string]ClientMessage {
	var u = make(map[string]ClientMessage)
	if id == "" {
		return u
	}

	str := strings.Split(id, ",")

	engine := new(Engine)
	err := engine.NewEngine()
	if err != nil {
		fmt.Println(err)
	}
	defer engine.Engine.Close()

	for _, v := range str {

		clientMessage := ClientMessage{}
		clientMessage.ClientId = v

		a, _ := engine.Engine.Query(fmt.Sprintf("SELECT KHMC FROM CRM_DAT001 WHERE KHID = %v", v))
		clientMessage.Name = Strings(a)

		b, _ := engine.Engine.Query(fmt.Sprintf("SELECT MOBIL FROM CRM_DAT001 WHERE KHID = %v", v))
		clientMessage.PhoneMumber = Strings(b)

		c, _ := engine.Engine.Query(fmt.Sprintf("SELECT DIZHI FROM CRM_DAT001 WHERE KHID = %v", v))
		clientMessage.Site = Strings(c)

		if clientMessage.PhoneMumber == "" && clientMessage.Site == "" && clientMessage.Name == "" {
			clientMessage.Error = "The query client does not exist"
		}

		u[v] = clientMessage
	}

	return u
}

func Strings(str []map[string][]byte) (str1 string) {
	for _, v := range str {
		for _, i := range v {
			if string(i) != "" {
				str1 = string(i)
			}
		}
	}

	return str1
}
