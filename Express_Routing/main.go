package main

import (
	"Golang/Express_Routing/express"
	"Golang/Express_Routing/models"
	"fmt"
	"github.com/FirewoodBloody/PacketProup/logs"
	"strconv"
	"strings"
	"time"
)

var (
	logLevel = make(map[string]string, 10)
	logger   logs.LogInterface
	getExp   map[string]string
)

func init() {
	getExp = make(map[string]string, 10)
	getExp["圆通快递"] = "YTO"
	getExp["中通快递"] = "ZTO"
	getExp["EMS"] = "EMS"
	getExp["邮政快递"] = "YZPY"
}

func getCode(Code string) string {
	switch Code {
	case "50", "54":
		return "上门收件"
	case "70", "77", "33", "99":
		return "问题件"
	case "40", "47", "204":
		return "正在派送"
	case "80", "8000":
		return "已签收"
	case "130":
		return "顺丰站"
	case "648":
		return "快递退回"
	default:
		return "转运中"
	}
}

func UpdateState(engine models.Engine) {

	for {
		engine.Err = engine.NewEngine()
		if engine.Err != nil {
			logger.Error("创建engine连接错误：", engine.Err)
		}
		engine.Err = engine.Engine.Ping()
		if engine.Err != nil {
			logger.Error("建立数据库连接失败：", engine.Err)
		}
		var num, Jnum int
		maps, err := engine.Engine.Query("select count(*) from blcrm.kdlyzt")
		fmt.Println(err)
		if err != nil {
			continue
		}

		for _, v := range maps {
			for _, i := range v {
				if string(i) != "" {
					num, _ = strconv.Atoi(string(i))
					fmt.Println(string(i))
				}
			}
		}
		if num == Jnum {
			continue
		}

		Jnum = num
		engine.Err = engine.Engine.Where("DQZT IS NULL").Find(&engine.GetDb)
		if engine.Err != nil {
			continue
		}

		for _, v := range engine.GetDb {
			engine.Err = IFUpdate(v, engine)
			if engine.Err != nil {
				continue
			}

		}

	}
}

func main() {

	engine := models.Engine{}

	go UpdateState(engine)

	for {
		var hH, mM, sS int
		var sleep int
		hH = time.Now().Hour()
		mM = time.Now().Minute()
		sS = time.Now().Second()
		if hH >= 8 || hH < 20 {
			if mM == 0 {
				engine.Err = engine.NewEngine()
				if engine.Err != nil {
					continue
				}
				engine.Err = engine.Engine.Ping()
				if engine.Err != nil {
					continue
				}

				engine.Err = engine.Engine.Where("WHERE DQZT != '已签收'").Find(&engine.GetDb)
				if engine.Err != nil {
					continue
				}
				engine.Err = engine.Engine.Close()
				if engine.Err != nil {
					continue
				}
				for _, v := range engine.GetDb {
					engine.Err = IFUpdate(v, engine)
					if engine.Err != nil {
						continue
					}
				}

			} else {
				sleep = (60-mM)*60 - sS
				time.Sleep(time.Second * time.Duration(sleep))
			}

		} else {
			if hH >= 20 {
				sleep = (24-hH+8)*360 - mM*60 - sS
			} else if hH < 8 {
				sleep = 8 - hH - mM*60 - sS
			}
			time.Sleep(time.Second * time.Duration(sleep))
		}
	}
}

func IFUpdate(v models.Kdlyzt, engine models.Engine) error {
	engine.Err = engine.NewEngine()
	if engine.Err != nil {
		return engine.Err
	}
	engine.Err = engine.Engine.Ping()
	if engine.Err != nil {
		return engine.Err
	}
	defer engine.Engine.Close()

	if v.THKDDH == "" {
		if v.KDGS == "顺丰快递" {
			data, err := express.SfCreateData(v.KDDH)
			if err != nil {
				return err
			}

			go func() {
				for i := 0; i < len(data.Body.RouteResponse.Route); i++ {
					engine.Err = engine.NewEngine()
					if engine.Err != nil {
						continue
					}
					engine.Err = engine.Engine.Ping()
					if engine.Err != nil {
						continue
					}
					engine.Err = engine.InSetDateXQ(data.Body.RouteResponse.Mailno, data.Body.RouteResponse.Route[i].Remark, data.Body.RouteResponse.Route[i].Accept_Time)
					if engine.Err != nil {
						continue
					}
				}
			}()

			if data.Body.RouteResponse.Mailno == v.KDDH {
				if data.Body.RouteResponse.Route[len(data.Body.RouteResponse.Route)-1].Accept_Time == v.DQZTSJ {
					return fmt.Errorf("当前记录已是最新记录！")
				}
				if data.Body.RouteResponse.Route[len(data.Body.RouteResponse.Route)-1].Opcode == "648" {
					v.DQZT = getCode(data.Body.RouteResponse.Route[len(data.Body.RouteResponse.Route)-1].Opcode)
					v.DQZTSJ = data.Body.RouteResponse.Route[len(data.Body.RouteResponse.Route)-1].Accept_Time
					str := strings.Split(data.Body.RouteResponse.Route[len(data.Body.RouteResponse.Route)-1].Remark, " ")
					if len(str) > 0 {
						v.THKDDH = str[len(str)-1]
					}
					engine.Err = engine.UpDateRefreshZT(v.DQZT, v.DQZTSJ, v.THKDDH, v.KDDH)
					if engine.Err != nil {
						return engine.Err
					}
				} else {
					v.DQZT = getCode(data.Body.RouteResponse.Route[len(data.Body.RouteResponse.Route)-1].Opcode)
					v.DQZTSJ = data.Body.RouteResponse.Route[len(data.Body.RouteResponse.Route)-1].Accept_Time
					engine.Err = engine.UpDateRefreshZT(v.DQZT, v.DQZTSJ, v.THKDDH, v.KDDH)
					if engine.Err != nil {
						return engine.Err
					}
				}

			} else {
				return fmt.Errorf("查询获取单号不一致，跳过！")
			}
		} else {
			data, err := express.KdnExpressInformation(getExp[v.KDGS], v.KDDH)
			if err != nil {
				return err
			}

			if !data.Success {
				return fmt.Errorf("快递鸟查询出错，Err ")
			}

			go func() {
				for i := 0; i < len(data.Traces); i++ {
					engine.Err = engine.NewEngine()
					if engine.Err != nil {
						continue
					}
					engine.Err = engine.Engine.Ping()
					if engine.Err != nil {
						continue
					}
					engine.Err = engine.InSetDateXQ(data.LogisticCode, data.Traces[i].AcceptStation, data.Traces[i].AcceptTime)
					if engine.Err != nil {
						continue
					}
				}
			}()

			switch data.State {
			case "2":
				v.DQZT = "转运中"
				v.DQZTSJ = data.Traces[len(data.Traces)-1].AcceptTime
			case "3":
				v.DQZT = "已签收"
				v.DQZTSJ = data.Traces[len(data.Traces)-1].AcceptTime
			case "4":
				v.DQZT = "问题件"
				v.DQZTSJ = data.Traces[len(data.Traces)-1].AcceptTime

			}

			engine.Err = engine.UpDateRefreshZT(v.DQZT, v.DQZTSJ, "", v.KDDH)
			if engine.Err != nil {

				return engine.Err
			}
		}
	} else {
		if v.KDGS == "顺丰快递" {
			switch len(v.THKDDH) {
			case 12:
				v.THKDDH = v.THKDDH
			case 11:
				v.THKDDH = fmt.Sprintf("0%s", v.THKDDH)
			case 10:
				v.THKDDH = fmt.Sprintf("00%s", v.THKDDH)
			}
			data, err := express.SfCreateData(v.THKDDH)
			if err != nil {
				return err
			}

			go func() {
				for i := 0; i < len(data.Body.RouteResponse.Route); i++ {
					engine.Err = engine.NewEngine()
					if engine.Err != nil {
						continue
					}
					engine.Err = engine.Engine.Ping()
					if engine.Err != nil {
						continue
					}
					engine.Err = engine.InSetDateXQ(data.Body.RouteResponse.Mailno, data.Body.RouteResponse.Route[i].Remark, data.Body.RouteResponse.Route[i].Accept_Time)
					if engine.Err != nil {
						continue
					}
				}
			}()

			if data.Body.RouteResponse.Mailno == v.THKDDH {
				if data.Body.RouteResponse.Route[len(data.Body.RouteResponse.Route)-1].Accept_Time == v.DQZTSJ {
					return fmt.Errorf("当前记录已是最新记录！")
				}

				v.DQZT = getCode(data.Body.RouteResponse.Route[len(data.Body.RouteResponse.Route)-1].Opcode)
				v.DQZTSJ = data.Body.RouteResponse.Route[len(data.Body.RouteResponse.Route)-1].Accept_Time
				engine.Err = engine.UpDateRefreshZT(v.DQZT, v.DQZTSJ, v.THKDDH, v.KDDH)
				if engine.Err != nil {
					return engine.Err
				}
			} else {
				return fmt.Errorf("查询获取单号不一致，跳过！")
			}

		} else {
			data, err := express.KdnExpressInformation(getExp[v.KDGS], v.THKDDH)
			if err != nil {
				return err
			}

			if !data.Success {
				return fmt.Errorf("快递鸟查询出错，！")
			}

			go func() {
				for i := 0; i < len(data.Traces); i++ {
					engine.Err = engine.NewEngine()
					if engine.Err != nil {
						continue
					}
					engine.Err = engine.Engine.Ping()
					if engine.Err != nil {
						continue
					}
					engine.Err = engine.InSetDateXQ(data.LogisticCode, data.Traces[i].AcceptStation, data.Traces[i].AcceptTime)
					if engine.Err != nil {
						continue
					}
				}
			}()

			switch data.State {
			case "2":
				v.DQZT = "转运中"
				v.DQZTSJ = data.Traces[len(data.Traces)-1].AcceptTime
			case "3":
				v.DQZT = "已签收"
				v.DQZTSJ = data.Traces[len(data.Traces)-1].AcceptTime
			case "4":
				v.DQZT = "问题件"
				v.DQZTSJ = data.Traces[len(data.Traces)-1].AcceptTime
			default:
				v.DQZT = "转运中"
				v.DQZTSJ = data.Traces[len(data.Traces)-1].AcceptTime
			}

			engine.Err = engine.UpDateRefreshZT(v.DQZT, v.DQZTSJ, v.THKDDH, v.KDDH)
			if engine.Err != nil {
				return engine.Err
			}
		}
	}
	return nil
}
