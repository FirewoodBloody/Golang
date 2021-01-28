package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math"
	"strconv"
	"time"
)

//	附加
var (
	orderNumber int    //初始、校验 订单数
	orderCount  int    //订单数
	priceSum    int    //订单计算金额
	orderSum    int    //订单计算金额
	userCoinSum int    //积分抵扣金额
	remark      string //错误信息
)

//订单主信息
type order struct {
	orderId        int //订单ID
	orderType      int //订单类型
	discountAmount int //订单折扣价
	totalAmount    int //订单总金额
	useCoin        int //积分抵扣
	useCoinRatio   int //积分抵扣比例
	orderCart      []orderCart
}

//订单购物车信息
type orderCart struct {
	cartId     int //购物车ID
	goodsId    int //商品ID
	goodsSkuId int //商品规格ID
	goodsCount int //商品数量
	goodsPrice int //商品原价
	priceSale  int //商品售价
	totalPrice int //商品总价
	goodsRatio int //折扣比例
	orderGoodsSku
}

//商品规格信息
type orderGoodsSku struct {
	goodsId    int //商品ID
	goodsSkuId int //商品规格ID
	priceSale  int //售价
	priceRaw   int //原价
}

// 监控订单新增数据，以每六分钟进行一次循环 ；
// 校验订单 总金额 折扣金额 积分抵扣金额 商品原价 售价 商品规格；
// 修正订单错误数据；
func CheckOrderData() {
	e := new(Engine)
	orderNumber, _ = e.orderCountNumber()

	for {
		orderCount, _ = e.orderCountNumber()
		if orderNumber == orderCount { //判断订单表订单数量是否新增
			time.Sleep(time.Second * 360)
			continue
		}
		var ord []order
		ord = make([]order, orderCount-orderNumber)
		for i := 0; i < orderCount-orderNumber; i++ {
			ord[i], e.Err = e.orderParticulars(i) //查询新增订单的订单信息
			if e.Err != nil {
				continue
			}
		}
		orderNumber = orderCount //如新增则将新的数量赋值用来保存校验订单数量的变量

		for i := 0; i < len(ord); i++ {

			//校验订单类型
			if ord[i].orderType != 1 {
				continue
			}
			priceSum = 0 //初始化订单售价总价
			orderSum = 0 //初始化订单总价
			remark = ""
			userCoinSum = ord[i].useCoin * ord[i].useCoinRatio //计算订单积分抵扣金额
			for k := 0; k < len(ord[i].orderCart); k++ {

				//商品为空
				if ord[i].orderCart[k].goodsId == 0 {
					continue
				}

				//商品规格为空
				if ord[i].orderCart[k].goodsSkuId == 0 {
					fmt.Println("订单ID:", ord[i].orderId, " 商品ID:", ord[i].orderCart[k].cartId, " 商品规格ID为空")
					continue
				}

				/*检验购物车商品原价 商品规格原价是否一致*/
				if ord[i].orderCart[k].orderGoodsSku.priceRaw != ord[i].orderCart[k].goodsPrice {
					//更新商品原价 (商品原价 = 商品规格原价)
					//_ = e.NewEngine()
					//e.Engine.Query(fmt.Sprintf("UPDATE bl_mall_cart SET goods_price = %v WHERE id = %v", ord[i].orderCart[k].orderGoodsSku.priceRaw, ord[i].orderCart[k].cartId))
					//e.Close()
					remark += "\"orderCart\":\"检验购物车商品原价 商品规格原价 失败\","
					fmt.Printf("UPDATE bl_mall_cart SET goods_price = %v WHERE id = %v\n", ord[i].orderCart[k].orderGoodsSku.priceRaw, ord[i].orderCart[k].cartId)

					/*校验购物车商品售价 原价是否一致*/
					if ord[i].orderCart[k].goodsPrice == ord[i].orderCart[k].priceSale {
						//更新商品售价(商品原价 = 商品规格原价 = 商品售价)
						//_ = e.NewEngine()
						//e.Engine.Query(fmt.Sprintf("UPDATE bl_mall_cart SET price_sale = %v WHERE id = %v", ord[i].orderCart[k].orderGoodsSku.priceRaw, ord[i].orderCart[k].cartId))
						//e.Close()
						remark += "," + "\"priceSale\":\"校验购物车商品售价 原价是否一致\""
						fmt.Printf("UPDATE bl_mall_cart SET price_sale = %v WHERE id = %v\n", ord[i].orderCart[k].orderGoodsSku.priceRaw, ord[i].orderCart[k].cartId)

						/*校验购物车商品总价 是否一致*/
						if ord[i].orderCart[k].orderGoodsSku.priceRaw*ord[i].orderCart[k].goodsCount != ord[i].orderCart[k].totalPrice {
							//更新商品总价
							//_ = e.NewEngine()
							//e.Engine.Query(fmt.Sprintf("UPDATE bl_mall_cart SET total_price = %v WHERE id = %v", ord[i].orderCart[k].orderGoodsSku.priceRaw*ord[i].orderCart[k].goodsCount, ord[i].orderCart[k].cartId))
							//e.Close()
							remark += "," + "totalPrice：\"校验购物车商品总价 是否一致\""
							fmt.Printf("UPDATE bl_mall_cart SET total_price = %v WHERE id = %v\n", ord[i].orderCart[k].orderGoodsSku.priceRaw*ord[i].orderCart[k].goodsCount, ord[i].orderCart[k].cartId)
						}

						/*校验商品折扣是否一致*/
						if ord[i].orderCart[k].goodsRatio != 100 {
							//更新商品折扣比例
							//_ = e.NewEngine()
							//e.Engine.Query(fmt.Sprintf("UPDATE bl_mall_cart SET goods_ratio = 100 WHERE id = %v", ord[i].orderCart[k].cartId))
							//e.Close()
							remark += "," + "goodsRatio：\"校验商品折扣是否一致\""
							fmt.Printf("UPDATE bl_mall_cart SET goods_ratio = 100 WHERE id = %v\n", ord[i].orderCart[k].cartId)
						}

						priceSum += ord[i].orderCart[k].goodsCount * ord[i].orderCart[k].orderGoodsSku.priceRaw //计算累计商品售价金额 (商品规格原价 * 商品数量)
						orderSum += ord[i].orderCart[k].goodsCount * ord[i].orderCart[k].orderGoodsSku.priceRaw //计算订单总金额
						continue
					}

					/*校验购物车商品总价 是否一致*/
					if ord[i].orderCart[k].goodsCount*ord[i].orderCart[k].priceSale != ord[i].orderCart[k].totalPrice {
						//更新商品总价
						//_ = e.NewEngine()
						//e.Engine.Query(fmt.Sprintf("UPDATE bl_mall_cart SET total_price = %v WHERE id = %v", ord[i].orderCart[k].goodsCount*ord[i].orderCart[k].priceSale, ord[i].orderCart[k].cartId))
						//e.Close()
						remark += "," + "\"totalPrice\":\"校验购物车商品总价 是否一致\""
						fmt.Printf("UPDATE bl_mall_cart SET total_price = %v WHERE id = %v\n", ord[i].orderCart[k].goodsCount*ord[i].orderCart[k].priceSale, ord[i].orderCart[k].cartId)
					}

					/*校验商品折扣是否一致*/
					if int(math.Floor(float64(ord[i].orderCart[k].priceSale/ord[i].orderCart[k].orderGoodsSku.priceRaw*100)+0.5)) != ord[i].orderCart[k].goodsRatio {
						//更新商品折扣比例
						//_ = e.NewEngine()
						//e.Engine.Query(fmt.Sprintf("UPDATE bl_mall_cart SET goods_ratio = %v WHERE id = %v", int(math.Floor(float64(ord[i].orderCart[k].priceSale/ord[i].orderCart[k].orderGoodsSku.priceRaw*100)+0.5)), ord[i].orderCart[k].cartId))
						//e.Close()
						remark += "," + "\"totalPrice\":\"校验商品折扣是否一致\""
						fmt.Printf("UPDATE bl_mall_cart SET goods_ratio = %v WHERE id = %v\n", int(math.Floor(float64(ord[i].orderCart[k].priceSale/ord[i].orderCart[k].orderGoodsSku.priceRaw*100)+0.5)), ord[i].orderCart[k].cartId)
					}
					priceSum += ord[i].orderCart[k].goodsCount * ord[i].orderCart[k].priceSale              //计算累计商品售价金额
					orderSum += ord[i].orderCart[k].goodsCount * ord[i].orderCart[k].orderGoodsSku.priceRaw //计算订单总金额
					continue
				}

				priceSum += ord[i].orderCart[k].goodsCount * ord[i].orderCart[k].priceSale //计算累计商品售价金额
				orderSum += ord[i].orderCart[k].goodsCount * ord[i].orderCart[k].priceRaw  //计算订单总金额
			}

			/*校验订单主表总金额*/
			fmt.Println(ord[i].totalAmount, priceSum, ord[i].discountAmount, userCoinSum)
			if ord[i].totalAmount != priceSum+ord[i].discountAmount+userCoinSum {
				//更新订单总金额 订单积分 订单折扣价
				//_ = e.NewEngine()
				//e.Engine.Query(fmt.Sprintf("UPDATE bl_mall_order SET total_amount = %v,discount_amount = %v,discount_ratio = %v WHERE id = %v", orderSum, orderSum-priceSum, orderSum-priceSum, ord[i].orderId))
				//e.Close()
				if remark == "" {
					remark += "\"totalAmount\":\"校验订单主表总金额\""
				} else {
					remark += "," + "\"totalAmount\":\"校验订单主表总金额\""
				}
				fmt.Printf("UPDATE bl_mall_order SET total_amount = %v,discount_amount = %v,discount_ratio = %v WHERE id = %v\n", orderSum, orderSum-priceSum, orderSum-priceSum, ord[i].orderId)
			}

			if remark != "" {
				_ = e.NewEngine()
				_, _ = e.Engine.Query(fmt.Sprintf("INSERT INTO bl_mall_order_detection_log (create_at,order_id,remark) VALUES (NOW(),%v,'{%v}')", ord[i].orderId, remark))
				e.Close()
			}
		}
		time.Sleep(time.Second * 360) //
	}
}

//每日凌晨四点进行新媒体客户回收计划运行
func RecoveryOfTheCustomer() {
	e := new(Engine)
	_ = e.NewEngine()
	//查询回收名单归属部门
	departId, err := e.Login_depart_id_permissions("33")
	if err != nil {
		return
	}
	e.Close()

	//开始进行每日回收计划
	for {
		//查询需要回收的所有客户
		_ = e.NewEngine()
		result, err := e.Engine.Query(fmt.Sprintf("SELECT \ncc.id AS id,\ncca.first_category AS first_category,\ncc.user_id AS user_id\nFROM\nbl_crm_customer cc \nLEFT JOIN bl_crm_customer_append cca on cca.customer_id = cc.id\nLEFT JOIN bl_users ccu on ccu.id = cc.user_id\nWHERE\ncc.depart_id in (%v)", departId))
		if err != nil {
			fmt.Println(err)
		}
		e.Close()
		//开始执行回收计划
		//每日凌晨4点进行回收计划进行
		//依据客户的当前情况进行客户的回收保留
		for _, v := range result {
			_ = e.NewEngine()

			//尅是执行回收计划；判断客户的一级分类，根据一级分类的类型进行不同方式的回收和保留

			//客户的一级分类为 已购客户
			if string(v["first_category"]) == "1" {
				//判断当前客户在当前归属员工名下 有没有 历史消费（新媒体线上除外）单笔大于等于 10000 元的消费记录
				orderCounts, _ := e.Engine.Query(fmt.Sprintf("SELECT\nCOUNT(DISTINCT id) AS count_id\nFROM\nbl_mall_order\nWHERE\ncustomer_id = %v\nAND performance_user_id = %v\nAND deleted_at IS NULL\nAND `status` = 90\nAND performance_at IS NOT NULL\nAND pay_amount >= 1000000\nAND channel_id < 3", string(v["id"]), string(v["user_id"])))
				//当前客户在当前归属员工名下历史消费记录有单笔大于等于 10000 万 时，跳过此客户回收 保留给当前员工
				if string(orderCounts[0]["count_id"]) != "0" {
					//跳过当前客户回收计划循环
					e.Close()
					continue
				}

				//当前客户在当前归属员工名下历史消费记录没有有单笔大于等于 10000 万 时，进行此客户的回收计划
				hour, _ := time.ParseDuration("-168h")
				knows := time.Now().Add(hour)
				Month := fmt.Sprintf("%02d", int(knows.Month()))

				hour, _ = time.ParseDuration("-336h")
				know := time.Now().Add(hour)
				//依据客户是否为当月引进新 并且历史首次消费的客户
				orderCounts, _ = e.Engine.Query(fmt.Sprintf("SELECT\nCOUNT(DISTINCT id) AS count_id\nFROM\nbl_mall_order\nWHERE\ncustomer_id = %v\nAND performance_user_id = %v\nAND deleted_at IS NULL\nAND ((\n\t\t\t`status` = 90 \n\t\t\tAND type = 1 \n\t\t\tAND performance_at IS NOT NULL \n\t\t\tAND performance_at > '%v-%v-01 00:00:00' \n\t\t\tAND performance_at < '%v 23:59:59' \n\t\t) \n\tOR ( `status` = 90 AND performance_at IS NULL AND created_at < '%v 00:00:00'))", string(v["id"]), string(v["user_id"]), knows.Year(), Month, knows.Format("2006-01-06"), know.Format("2006-01-06")))
				//判断如果客户在员工名下当月消费记录是否为0
				if string(orderCounts[0]["count_id"]) == "0" {
					//判断客户当月在其他员工名下的消费记录
					orderCounts, _ = e.Engine.Query(fmt.Sprintf("SELECT\nCOUNT(DISTINCT id) AS count_id\nFROM\nbl_mall_order\nWHERE\ncustomer_id = %v\nAND deleted_at IS NULL\nAND ((\n\t\t\t`status` = 90 \n\t\t\tAND type = 1 \n\t\t\tAND performance_at IS NOT NULL \n\t\t\tAND performance_at > '%v-%v-01 00:00:00' \n\t\t\tAND performance_at < '%v 23:59:59' \n\t\t) \n\tOR ( `status` = 90 AND performance_at IS NULL AND created_at < '%v 00:00:00'))", string(v["id"]), knows.Year(), Month, knows.Format("2006-01-06"), know.Format("2006-01-06")))
					//当月消费记录为空
					if string(orderCounts[0]["count_id"]) == "0" {
						//没有购买记录的客户
						//判断历史是否存在消费记录
						orderCounts, _ = e.Engine.Query(fmt.Sprintf("SELECT\nCOUNT(DISTINCT id) AS count_id\nFROM\nbl_mall_order\nWHERE\ncustomer_id = %v\nAND deleted_at IS NULL\nAND `status` = 90\nAND type = 1\nAND performance_at IS NOT NULL\nAND performance_at < '%v-%v-01 00:00:00'", string(v["id"]), knows.Year(), fmt.Sprintf("%02d", int(knows.Month()))))
						if string(orderCounts[0]["count_id"]) == "0" {
							//历史不存在购买记录 回收纸系统工程师名下
							_, _ = e.Engine.Query(fmt.Sprintf("UPDATE bl_crm_customer SET user_id = 1427,assign_at = '%v' WHERE id = %v", time.Now().Format("2006-01-06 15:04:05"), string(v["id"])))
						} else {
							//历史存在购买记录回收至 新媒体新进已购名单
							_, _ = e.Engine.Query(fmt.Sprintf("UPDATE bl_crm_customer SET user_id = 2168,assign_at = '%v' WHERE id = %v", time.Now().Format("2006-01-06 15:04:05"), string(v["id"])))
						}
					} else {
						//当月消费记录不为空
						//判断历史是否存在消费记录
						orderCounts, _ = e.Engine.Query(fmt.Sprintf("SELECT\nCOUNT(DISTINCT id) AS count_id\nFROM\nbl_mall_order\nWHERE\ncustomer_id = %v\nAND deleted_at IS NULL\nAND `status` = 90\nAND type = 1\nAND performance_at IS NOT NULL\nAND performance_at < '%v-%v-01 00:00:00'", string(v["id"]), knows.Year(), fmt.Sprintf("%02d", int(knows.Month()))))
						if string(orderCounts[0]["count_id"]) == "0" {
							//历史不存在购买记录 回收至新媒体新进已购库
							_, _ = e.Engine.Query(fmt.Sprintf("UPDATE bl_crm_customer SET user_id = 2156,assign_at = '%v' WHERE id = %v", time.Now().Format("2006-01-06 15:04:05"), string(v["id"])))
						} else {
							//历史存在购买记录回收至 新媒体新进已购名单
							_, _ = e.Engine.Query(fmt.Sprintf("UPDATE bl_crm_customer SET user_id = 2168,assign_at = '%v' WHERE id = %v", time.Now().Format("2006-01-06 15:04:05"), string(v["id"])))
						}
					}
				} else {
					//当月客户在当前员工名下消费不为空
					//判断历史是否有消费记录
					//当月消费记录不为空
					//判断历史是否存在消费记录
					orderCounts, _ = e.Engine.Query(fmt.Sprintf("SELECT\nCOUNT(DISTINCT id) AS count_id\nFROM\nbl_mall_order\nWHERE\ncustomer_id = %v\nAND deleted_at IS NULL\nAND `status` = 90\nAND type = 1\nAND performance_at IS NOT NULL\nAND performance_at < '%v-%v-01 00:00:00'", string(v["id"]), knows.Year(), fmt.Sprintf("%02d", int(knows.Month()))))
					if string(orderCounts[0]["count_id"]) == "0" {
						//历史不存在购买记录 回收至新媒体新进已购库
						_, _ = e.Engine.Query(fmt.Sprintf("UPDATE bl_crm_customer SET user_id = 2156,assign_at = '%v' WHERE id = %v", time.Now().Format("2006-01-06 15:04:05"), string(v["id"])))
					} else {
						//历史存在购买记录回收至 新媒体新进已购名单
						_, _ = e.Engine.Query(fmt.Sprintf("UPDATE bl_crm_customer SET user_id = 2168,assign_at = '%v' WHERE id = %v", time.Now().Format("2006-01-06 15:04:05"), string(v["id"])))
					}

				}

			} else if string(v["first_category"]) == "2" {
				hour, _ := time.ParseDuration("-720h")
				knows := time.Now().Add(hour)
				//未购客户（准已购）是暂时暂时性的，有客户会进行变化
				//准已购：凡是有订单的但是没有成交订单的客户定义为准已购
				//首先需要排除目前一致的有订单未完结的客户

				//查询客户是否有未完结的订单
				orderCounts, _ := e.Engine.Query(fmt.Sprintf("SELECT\nCOUNT(DISTINCT id) AS orderCount\nFROM\nbl_mall_order\nWHERE\ndeleted_at IS NULL\nAND type = 1\nAND `status` < 90 \nAND `status` > 1\nAND customer_id = %v\nAND user_id = %v", string(v["id"]), string(v["user_id"])))
				if string(orderCounts[0]["orderCount"]) != "0" {
					//当前客户 在当前员工 名下存在未完成的订单，跳过当前客户回收计划
					e.Close()
					continue
				} else {
					//当前客户 在 当前员工名下 不存在未完成的订单
					//准已购客户回收规则是 30 天
					//查询当前客户 30 天 是否存在 已完成未成交的订单
					orderCounts, _ := e.Engine.Query(fmt.Sprintf("SELECT\nCOUNT(DISTINCT id) AS orderCount\nFROM\nbl_mall_order\nWHERE\ndeleted_at IS NULL\nAND type = 1\nAND (`status` > 90 OR `status` = 1)\nAND customer_id = %v\nAND performance_at < '%v'", string(v["id"]), knows.Format("2006-01-06")))
					if string(orderCounts[0]["orderCount"]) == "0" {
						//判断是否为假单客户
						orderCounts, _ = e.Engine.Query(fmt.Sprintf("SELECT\nCOUNT(DISTINCT id) AS orderCount \nFROM\n\tbl_mall_order_media \nWHERE\n\tdeleted_at IS NULL \n\tAND `status` = 40 \n\tAND ((verify_order_comment = 50 ) OR ( verify_order_comment != 50 AND remark LIKE '%v' )) \n\tAND customer_id = %v", "%假单%", string(v["id"])))
						if string(orderCounts[0]["orderCount"]) == "0" {
							//判断客户是否在当前员工名下存在核单数据
							orderCounts, _ = e.Engine.Query(fmt.Sprintf("SELECT\nCOUNT(DISTINCT id) AS orderCount \nFROM\n\tbl_mall_order_media \nWHERE\n`status` > 20\nAND deleted_at IS NULL\nAND customer_id =%v \nAND assign_user_id = %v", string(v["id"]), string(v["user_id"])))
							if string(orderCounts[0]["orderCount"]) == "0" {
								//30天前当前客户存在已完成的为成交订单，并且非假单客户；回收名单至 新媒体未妥投
								_, _ = e.Engine.Query(fmt.Sprintf("UPDATE bl_crm_customer SET user_id = 2105,assign_at = '%v' WHERE id = %v", time.Now().Format("2006-01-06 15:04:05"), string(v["id"])))

							} else {
								//如果当前客户 不是假单 当前员工名下存在未完成的新媒体订单，不进行回收
								e.Close()
								continue
							}
						} else {
							//假单客户；回收名单至 假单名单库
							_, _ = e.Engine.Query(fmt.Sprintf("UPDATE bl_crm_customer SET user_id = 2132,assign_at = '%v' WHERE id = %v", time.Now().Format("2006-01-06 15:04:05"), string(v["id"])))
						}
					} else {
						//判断是否为假单客户
						orderCounts, _ = e.Engine.Query(fmt.Sprintf("SELECT\nCOUNT(DISTINCT id) AS orderCount \nFROM\n\tbl_mall_order_media \nWHERE\n\tdeleted_at IS NULL \n\tAND `status` = 40 \n\tAND ((verify_order_comment = 50 ) OR ( verify_order_comment != 50 AND remark LIKE '%v' )) \n\tAND customer_id = %v", "%假单%", string(v["id"])))
						if string(orderCounts[0]["orderCount"]) == "0" {
							//判断客户是否在当前员工名下存在核单数据
							orderCounts, _ = e.Engine.Query(fmt.Sprintf("SELECT\nCOUNT(DISTINCT id) AS orderCount \nFROM\n\tbl_mall_order_media \nWHERE\n`status` > 20\nAND deleted_at IS NULL\nAND customer_id =%v \nAND assign_user_id = %v", string(v["id"]), string(v["user_id"])))
							if string(orderCounts[0]["orderCount"]) == "0" {
								//30天前当前客户存在已完成的为成交订单，并且非假单客户；回收名单至 新媒体未妥投
								_, _ = e.Engine.Query(fmt.Sprintf("UPDATE bl_crm_customer SET user_id = 2105,assign_at = '%v' WHERE id = %v", time.Now().Format("2006-01-06 15:04:05"), string(v["id"])))

							} else {
								//如果当前客户 不是假单 当前员工名下存在未完成的新媒体订单，不进行回收
								e.Close()
								continue
							}
						} else {
							//假单客户；回收名单至 假单名单库
							_, _ = e.Engine.Query(fmt.Sprintf("UPDATE bl_crm_customer SET user_id = 2132,assign_at = '%v' WHERE id = %v", time.Now().Format("2006-01-06 15:04:05"), string(v["id"])))
						}
					}
				}
			} else {
				_, _ = e.Engine.Query(fmt.Sprintf("UPDATE bl_crm_customer SET user_id = 1427,assign_at = '%v' WHERE id = %v", time.Now().Format("2006-01-06 15:04:05"), string(v["id"])))
			}
			e.Close()
		}

		//查询需要回收的管理员名下所有客户
		_ = e.NewEngine()
		result, err = e.Engine.Query("SELECT\nid \nFROM\n\tbl_crm_customer \nWHERE\n\tuser_id = 1 \n\tAND deleted_at IS NULL \n\tAND id IN (\n\tSELECT\n\t\tcustomer_id \n\tFROM\n\t\tbl_mall_order_media \nWHERE\n\t`status` >= 30)")
		if err != nil {
			fmt.Println(err)
		}
		e.Close()

		//回收管理员名下名单
		for _, v := range result {
			_ = e.NewEngine()
			//判断是否为假单客户
			orderCounts, _ := e.Engine.Query(fmt.Sprintf("SELECT\nCOUNT(DISTINCT id) AS orderCount \nFROM\n\tbl_mall_order_media \nWHERE\n\tdeleted_at IS NULL \n\tAND `status` = 40 \n\tAND ((verify_order_comment = 50 ) OR ( verify_order_comment != 50 AND remark LIKE '%v' )) \n\tAND customer_id = %v", "%假单%", string(v["id"])))
			if string(orderCounts[0]["orderCount"]) == "0" {
				//30天前当前客户存在已完成的为成交订单，并且非假单客户；回收名单至 新媒体未妥投
				_, _ = e.Engine.Query(fmt.Sprintf("UPDATE bl_crm_customer SET user_id = 2105,assign_at = '%v' WHERE id = %v", time.Now().Format("2006-01-06 15:04:05"), string(v["id"])))

			} else {
				//假单客户；回收名单至 假单名单库
				_, _ = e.Engine.Query(fmt.Sprintf("UPDATE bl_crm_customer SET user_id = 2132,assign_at = '%v' WHERE id = %v", time.Now().Format("2006-01-06 15:04:05"), string(v["id"])))
			}
			e.Close()
		}

		//定时器
		now := time.Now()                                                                    //获取当前时间，放到now里面，要给next用
		next := now.Add(time.Hour * 24)                                                      //通过now偏移24小时
		next = time.Date(next.Year(), next.Month(), next.Day(), 4, 0, 0, 0, next.Location()) //获取下一个凌晨四点的日期
		t := time.NewTimer(next.Sub(now))                                                    //计算当前时间到凌晨的时间间隔，设置一个定时器
		<-t.C
	}
}

//查询订单有总数
func (e *Engine) orderCountNumber() (OrderNumber int, err error) {
	_ = e.NewEngine()
	defer e.Close()
	//查询订单总数  count
	results, err := e.Engine.Query("SELECT COUNT(id) FROM bl_mall_order")
	if err != nil {
		return 0, fmt.Errorf("查询订单总数出错，err：%v", err)
	}

	return strconv.Atoi(string(results[0]["COUNT(id)"]))
}

// 查询新增订单信息主信息
func (e *Engine) orderParticulars(i int) (o order, err error) {

	_ = e.NewEngine()
	defer e.Close()
	//查询订单信息
	results, err := e.Engine.Query("SELECT MAX(id) FROM bl_mall_order")
	if err != nil || len(results) == 0 {
		return o, fmt.Errorf("查询订单总数出错，err：%v", err)
	}
	//订单ID
	o.orderId, _ = strconv.Atoi(string(results[0]["MAX(id)"]))
	o.orderId = o.orderId - i

	//订单主表信息
	results, err = e.Engine.Query(fmt.Sprintf("SELECT use_coin,discount_amount,total_amount,type FROM bl_mall_order WHERE id = %v", o.orderId))
	if err != nil || len(results) == 0 {
		return o, fmt.Errorf("查询订单总数出错，err：%v", err)
	}
	o.discountAmount, _ = strconv.Atoi(string(results[0]["discount_amount"])) // 订单折扣价
	o.useCoin, _ = strconv.Atoi(string(results[0]["use_coin"]))               //订单积分抵扣
	o.totalAmount, _ = strconv.Atoi(string(results[0]["total_amount"]))       //订单总价
	o.orderType, _ = strconv.Atoi(string(results[0]["type"]))                 //订单类型

	//订单附表信息
	results, err = e.Engine.Query(fmt.Sprintf("SELECT coin_consume FROM bl_mall_order_append WHERE order_id = %v", o.orderId))
	if err != nil || len(results) == 0 {
		return o, fmt.Errorf("查询订单总数出错，err：%v", err)
	}
	o.useCoinRatio, _ = strconv.Atoi(string(results[0]["coin_consume"])) //订单积分折扣比例

	//订单购物车信息
	results, err = e.Engine.Query(fmt.Sprintf("SELECT id,goods_id,goods_sku_id,goods_count,goods_price,total_price,price_sale,goods_ratio FROM bl_mall_cart WHERE goods_type = 1 AND order_id =  %v", o.orderId))
	if err != nil || len(results) == 0 {
		return o, fmt.Errorf("查询订单总数出错，err：%v", err)
	}
	o.orderCart = make([]orderCart, len(results))
	for k, v := range results {
		//fmt.Println(v)
		o.orderCart[k].cartId, _ = strconv.Atoi(string(v["id"]))               //购物车ID
		o.orderCart[k].goodsId, _ = strconv.Atoi(string(v["goods_id"]))        //购物车商品ID
		o.orderCart[k].goodsSkuId, _ = strconv.Atoi(string(v["goods_sku_id"])) //商品规格ID
		o.orderCart[k].goodsCount, _ = strconv.Atoi(string(v["goods_count"]))  //商品数量
		o.orderCart[k].goodsPrice, _ = strconv.Atoi(string(v["goods_price"]))  //商品原价
		o.orderCart[k].priceSale, _ = strconv.Atoi(string(v["price_sale"]))    //商品售价
		o.orderCart[k].totalPrice, _ = strconv.Atoi(string(v["total_price"]))  //商品总价
		o.orderCart[k].goodsRatio, _ = strconv.Atoi(string(v["goods_ratio"]))  //商品总价
	}

	//订单购物车商品规格信息
	for k, v := range o.orderCart {
		results, err = e.Engine.Query(fmt.Sprintf("SELECT id,price_sale,price_raw,goods_id FROM `bl_mall_goods_sku` WHERE id =  %v", v.goodsSkuId))
		if err != nil || len(results) == 0 {
			return o, fmt.Errorf("查询订单总数出错，err：%v", err)
		}
		o.orderCart[k].orderGoodsSku.goodsSkuId, _ = strconv.Atoi(string(results[0]["id"]))        //商品规格ID
		o.orderCart[k].orderGoodsSku.goodsId, _ = strconv.Atoi(string(results[0]["goods_id"]))     //商品ID
		o.orderCart[k].orderGoodsSku.priceSale, _ = strconv.Atoi(string(results[0]["price_sale"])) //售价
		o.orderCart[k].orderGoodsSku.priceRaw, _ = strconv.Atoi(string(results[0]["price_raw"]))   //原价
	}

	return o, nil
}
