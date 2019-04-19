package models

import (
	"fmt"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/wendal/go-oci8"
	"os"
)

type Kdlyzt struct {
	KDGS   string `xorm:"varchar2(12) 'KDGS'"`
	KDDH   string `xorm:"varchar2(16) pk index unique 'KDDH'"`
	DQZT   string `xorm:"varchar2(12) 'DQZT'"`
	DQZTSJ string `xorm:"datetime  'DQZTSJ'"` //time.Time
	THKDDH string `xorm:"varchar2(16) 'THKDDH'"`
}

type Kdlyxq struct {
	KDDH   string `xorm:"varchar2(16) pk index notnull 'KDDH'"`
	KDZT   string `xorm:"varchar2(128)  notnull 'KDDH'"`
	KDZTSJ string `xorm:"datetime pk index notnull unique 'KDDH'"` //time.Time
}

type Engine struct {
	Engine *xorm.Engine
	Err    error
	GetDb  []Kdlyzt
}

const (
	TimeFormat = "2006-01-02 15:04:05"
	driverName = "oci8"
	dBconnect  = "KD/1219271@192.168.0.9:1521/BLDB"
	tbMapper   = "BLCRM."
)

func init() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码
}

//初始化化
func (e *Engine) NewEngine() error {
	var err error
	e.Engine, err = xorm.NewEngine(driverName, dBconnect)
	if err != nil {
		return err
	}
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, tbMapper)
	e.Engine.ShowSQL(true)
	e.Engine.SetTableMapper(tbMapper)
	return nil
}

//新增无快递状态的快递信息
func (e *Engine) UpDateRefreshZT(setDQZT, setDQZTSJ, setTHKDDH, whereKDDH string) error {
	var UpDateZTSql string
	if len(setTHKDDH) == 0 {
		UpDateZTSql = fmt.Sprintf("UPDATE BLCRM.KDLYZT SET DQZT = '%v', DQZTSJ = TO_DATE('%v','yyyy-MM-dd HH24:mi:ss') WHERE KDDH = %v", setDQZT, setDQZTSJ, whereKDDH)
	} else {
		UpDateZTSql = fmt.Sprintf("UPDATE BLCRM.KDLYZT SET DQZT = '%v', DQZTSJ = TO_DATE('%v','yyyy-MM-dd HH24:mi:ss') , THKDDH = %s WHERE KDDH = %v", setDQZT, setDQZTSJ, setTHKDDH, whereKDDH)
	}
	_, err := e.Engine.Exec(UpDateZTSql)
	if err != nil {
		return err
	}

	return nil
}

//插入路由信息
func (e *Engine) InSetDateXQ(valueKDDH, valueKDZT, valueKDZTSJ string) error {
	var InSetSql string
	InSetSql = fmt.Sprintf("INSERT INTO BLCRM.KDLYXQ VALUES ( %s , '%s' , TO_DATE('%v','yyyy-MM-dd HH24:mi:ss') )", valueKDDH, valueKDZT, valueKDZTSJ)

	_, err := e.Engine.Exec(InSetSql)
	if err != nil {
		return fmt.Errorf("%s : 此记录已存在，跳过，下一条...", InSetSql)
	}
	return nil
}
