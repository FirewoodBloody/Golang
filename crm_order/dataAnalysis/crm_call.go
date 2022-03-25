package dataAnalysis

import (
	"fmt"
	"strconv"
	"time"
)

func CrmCall() {
	//定时器    首次启动时间设定
	now := time.Now()                                                                   //获取当前时间，放到now里面，要给next用
	next := time.Date(now.Year(), now.Month(), now.Day(), 23, 10, 0, 0, now.Location()) //获取晚上23点的日期
	t := time.NewTimer(next.Sub(now))                                                   //计算当前时间到凌晨的时间间隔，设置一个定时器
	<-t.C

	b := new(BlCrm)
	c := new(Call)
	var types, is_effective_call, is_standard_call, duration int

	for {

		time1 := time.Now().Format(TimeFormat)
		//
		//time1 := "2021-03-" + fmt.Sprintf("%02d", i)

		b.NewEngine()
		c.NewEngine_Call()

		b.Engine.Quote("UPDATE bl_crm_customer_phone \nSET pure_number = phone \nWHERE\n\tpure_number IS NULL")

		b.Engine.Query(fmt.Sprintf("UPDATE `bl_crm_quality_voice` cq\n\tLEFT JOIN bl_crm_customer_phone ccp ON ccp.phone = cq.called_no\n\tLEFT JOIN bl_crm_customer cc on cc.id = ccp.customer_id\n\tSET cq.customer_id = cc.id,cq.customer_code=cc.customer_code,cq.customer_name=cc.`name`\n\tWHERE DATE_FORMAT( start_at, '%v' ) = '%v' \n\tAND cq.call_type = 'extension_outbound'", "%Y-%m-%d", time1))

		resulist, _ := b.Engine.Query(fmt.Sprintf("SELECT\n\td1.NAME AS 部门,\n\td.NAME AS 二级部门,\n\tcqv.user_name AS 员工姓名,\n\tu.login_name AS  员工工号,\n\tcc.id as 客户ID,\n\tcc.`name` AS 客户姓名,\n\tcc.customer_code AS 客户编号,\n\t cqv.start_at AS 开始时间,\n\tcqv.end_at AS 结束时间,\n\tcqv.call_duration AS  通话时长\n\t\nFROM\n\t`bl_crm_quality_voice` cqv\n\tLEFT JOIN bl_crm_customer cc on cqv.customer_id = cc.id\n\tLEFT JOIN bl_users u ON cqv.user_id = u.id\n\tLEFT JOIN bl_depart d ON u.depart_id = d.id\n\tLEFT JOIN bl_depart d1 ON d.parent_id = d1.id \nWHERE\n\tcqv.start_at > '%v 00:00:00' \n\tAND cqv.start_at < '%v 23:59:59' \n\tAND cqv.call_type = 'extension_outbound'", time1, time1))
		b.Close()
		for k, v := range resulist {
			duration, _ = strconv.Atoi(string(v["通话时长"]))
			if duration >= 30 {
				types = 1
				is_effective_call = 1
			} else {
				types = 3
				is_effective_call = 0
			}
			if duration >= 180 {
				is_standard_call = 1
			}

			//通过员工ID查询员工姓名
			user_resulist, _ := c.Engine.Query(fmt.Sprintf("SELECT\n\tdepart_v1_name,\n\tdepart_v2_name,\n\tnickname \nFROM\n\t`users` \nWHERE\n\tlogin_name = '%v'", string(v["员工工号"])))

			if len(user_resulist) == 0 {
				b.NewEngine()
				user_resulist, b.Err = b.Engine.Query(fmt.Sprintf("SELECT\nnickname,\n`status` \nFROM\n\tbl_users \nWHERE\n\tlogin_name =  '%v'", string(v["员工工号"])))
				b.Close()
				if len(user_resulist) == 0 {
					continue
				}
				_, _ = c.Engine.Query(fmt.Sprintf("INSERT INTO users (nickname,login_name,status) VALUES('%v','%v',%v)", string(user_resulist[0]["nickname"]), string(v["员工工号"]), string(user_resulist[0]["status"])))
			}

			c.Engine.Query(fmt.Sprintf("INSERT INTO `data_analysis_library`.`quality_voice` (\n\t`stat_at`,\n\t`end_at`,\n\t`customer_id`,\n\t`customer_code`,\n\t`customer_name`,\n\t`nickname`,\n\t`login_name`,\n\t`call_duration`,\n\t`type`,\n\t`is_effective_call`,\n\t`is_standard_call` \n, `depart_v1_name`, `depart_v2_name`)  VALUES ('%v','%v',%v,'%v','%v','%v','%v',%v,%v,%v,%v,'%v','%v');", string(v["开始时间"]), string(v["结束时间"]), string(v["客户ID"]), string(v["客户编号"]), string(v["客户姓名"]), string(v["员工姓名"]), string(v["员工工号"]), duration, types, is_effective_call, is_standard_call, string(user_resulist[0]["depart_v1_name"]), string(user_resulist[0]["depart_v2_name"])))

			if k%500 == 0 {
				c.Close()
				time.Sleep(time.Second * 2)
				c.NewEngine_Call()
			}
		}

		//定时器
		b.Close()
		c.Close()

		now := time.Now()                                                                      //获取当前时间，放到now里面，要给next用
		next := now.Add(time.Hour * 24)                                                        //通过now偏移24小时
		next = time.Date(next.Year(), next.Month(), next.Day(), 23, 10, 0, 0, next.Location()) //获取下一个晚上23点的日期
		t := time.NewTimer(next.Sub(now))                                                      //计算当前时间到凌晨的时间间隔，设置一个定时器
		<-t.C

	}
}
