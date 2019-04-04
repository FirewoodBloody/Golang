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

	orm.RegisterDataBase("default", "oci8", "system/aizhu1219271@127.0.0.1:1521/SF") //连接oracle数据库
}

//查询数据库
func main() {
	//aa := "顺丰快递"
	selects := fmt.Sprintf("SELECT D FroM api WHERE A = '顺丰快递' AND C = '拒收'")
	o := orm.NewOrm()
	var maps []orm.Params
	num, err := o.Raw(selects).Values(&maps)

	if err == nil && num > 0 {
		fmt.Println(len(maps), num)
		fmt.Println(maps[0]["D"])

	}

}
