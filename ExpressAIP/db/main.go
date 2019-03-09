package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/wendal/go-oci8"
	"os"
)

func init() {
	orm.Debug = true

	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码

	orm.RegisterDriver("oci8", orm.DROracle) //注册驱动

	orm.RegisterDataBase("default", "oci8", "system/aizhu1219271@127.0.0.1:1521/Sf") //连接oracle数据库
}

//查询数据库
func main() {
	//aa := "顺丰快递"
	//selects := fmt.Sprintf("update api set  c = ('已签收')，d = to_date('2018-06-30 23:59:59','yyyy-mm-dd hh24:mi:ss') where b = 753951456852")
	o := orm.NewOrm()
	//var maps []orm.Params
	//num, err := o.Raw(selects).Values(&maps)
	//if err == nil && num > 0 {
	//	fmt.Println(maps[0]["B"], len(maps)) // slene
	//
	//}
	a, err := o.Raw("update api set  c = '已签收' where b = 753951456852").Exec()
	if err != nil {
		fmt.Println(err)
	}

	n, err := a.RowsAffected()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)

}
