package main

import (
	"Golang/Express_Routing/models"
	"fmt"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/wendal/go-oci8"
	"os"
)

type Engine struct {
	engine     *xorm.Engine
	driverName string
	dBconnect  string
	err        error
	tbMapper   string
	getDb      []models.Kdlyzt
}

func init() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码
}

//初始化化
func (e *Engine) NewEngine() error {
	var err error
	e.engine, err = xorm.NewEngine(e.driverName, e.dBconnect)
	if err != nil {
		return err
	}
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, e.tbMapper)
	e.engine.ShowSQL(true)
	e.engine.SetTableMapper(tbMapper)
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
	_, err := e.engine.Exec(UpDateZTSql)
	if err != nil {
		return err
	}

	return nil
}

//插入路由信息
func (e *Engine) InSetDateXQ(valueKDDH, valueKDZT, valueKDZTSJ string) error {
	var InSetSql string
	InSetSql = fmt.Sprintf("INSERT INTO BLCRM.KDLYXQ VALUES ( %s , '%s' , TO_DATE('%v','yyyy-MM-dd HH24:mi:ss') )", valueKDDH, valueKDZT, valueKDZTSJ)

	_, err := e.engine.Exec(InSetSql)
	if err != nil {
		return fmt.Errorf("%s : 此记录已存在，跳过，下一条...", InSetSql)
	}
	return nil
}

func main() {
	Engine := Engine{
		driverName: "oci8",
		dBconnect:  "KD/1219271@192.168.0.9:1521/BLDB",
		tbMapper:   "BLCRM.",
	}
	err := Engine.NewEngine()

	if err != nil {
		fmt.Println(err)
	}
	err = Engine.engine.Ping()
	if err != nil {
		fmt.Println("数据库的连接测试", err)
	}

	err = Engine.engine.Where(" DQZT = '已签收'").Find(&Engine.getDb)

	if err != nil {
		fmt.Println("select err:", err)
	}
	defer Engine.engine.Close()

	for _, v := range Engine.getDb {
		fmt.Println(v)
		//InSetSql := fmt.Sprintf("INSERT INTO BLCRM.KDLYXQ VALUES ( %s , '%s' , TO_DATE('%v','yyyy-MM-dd HH24:mi:ss') )", v.KDDH, "这个快递现在已经签收了，你是瞎么！", "2019-03-03 12:12:12")
		err := Engine.InSetDateXQ(v.KDDH, "这个快递现在已经签收了，你是瞎么！", "2019-03-03 12:12:12")
		fmt.Println(err)
	}
}
