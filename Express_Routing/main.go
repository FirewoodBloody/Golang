package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/wendal/go-oci8"
	"os"
)

func init() {
	orm.Debug = true
	// 自动建表
	//orm.RunSyncdb("default", false, true)

	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码

	orm.RegisterDriver("oci8", orm.DROracle) //注册驱动

	orm.RegisterDataBase("default", "oci8", "KD/123@192.168.0.9:1521/BLDB") //连接oracle数据库
}

func main() {

	selects := fmt.Sprintf("SELECT KDDH FROM BLCRM.KDLYZT WHERE KDGS = '%v'", "圆通快递")
	updates := fmt.Sprintf("UPDATE BLCRM.KDLYZT SET DQZT = '已签收' WHERE KDDH = 12345679")
	o := orm.NewOrm()

	var maps []orm.Params
	num, err := o.Raw(selects).Values(&maps)
	fmt.Println("num:", num)
	if err == nil && num != 0 {
		for i := 0; i < int(num); i++ {
			fmt.Println(maps[i]["KDDH"])

			err := o.Begin()

			if err != nil {
				os.Exit(1)
			}
			_, err = o.Raw(updates).Exec()
			if err != nil {
				o.Commit()
			} else {
				o.Rollback()
			}

		}
	}

}
