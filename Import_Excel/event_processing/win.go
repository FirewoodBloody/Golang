package event_processing

import (
	"Golang/Import_Excel/module"
	"fmt"
	"github.com/Luxurioust/excelize"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	//_ "github.com/wendal/go-oci8"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	dBconnect  = "BLCRM/BLCRM2012@192.168.0.9:1521/BLDB"
	driverName = "oci8"
	tbMapper   = "BLCRM."
)

var (
	tbMappers core.PrefixMapper
	Number    int
)

type Engine struct {
	Engine *xorm.Engine
	Err    error
}

type Market struct {
	TimeNumber          string
	DateTime            string //单据日期
	Job_number          string //工号
	Operator            string //操作员
	Customer_name       string //客户姓名
	Customer_no         string //客户编码
	Customer_phon       string //客户电话
	Customer_site       string //客户地址
	Customer_type       string //客户分类
	Customer_source     string //客户来源
	Warehouse           string //库房
	For_Sale            string //销售类型
	Exchange_shop       string //票号
	Commodity           string //商品ID
	Num                 string //商品数量
	Price_per_uni       string //单价
	Remarks             string //备注
	New_and_old_clients string //新老客户
}

var Import_Type_Name string

func init() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码
	tbMappers = core.NewPrefixMapper(core.SnakeMapper{}, tbMapper)
}

//新建连接
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

//初始化事件
func TheEventInit(t *module.TForm) {
	t.Windows.SetOnClose(func(sender vcl.IObject, action *types.TCloseAction) {
		t.Windows.Close()
		os.Exit(0)
	})

	//	选择导入类型时间触发
	t.RadioButton1.SetOnClick(func(sender vcl.IObject) {
		if t.RadioButton1.Checked() == true {
			t.ProgressBar.SetPosition(0) //清空进度
			t.Open_Import = false

			Import_Type_Name = t.RadioButton1.Name()
			t.ListView.Clear()
			t.ListView.Columns().Clear()

			lvl := t.ListView.Columns().Add()
			lvl.SetCaption("导入结果")
			lvl.SetWidth(80)
			lvl.SetAlignment(types.TaCenter)

			lv1 := t.ListView.Columns().Add()
			lv1.SetCaption("客户姓名")
			lv1.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

			lv1 = t.ListView.Columns().Add()
			lv1.SetCaption("客户电话")
			lv1.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

			lv1 = t.ListView.Columns().Add()
			lv1.SetCaption("客户地址")
			lv1.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

			lv1 = t.ListView.Columns().Add()
			lv1.SetCaption("员工工号")
			lv1.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

			lv1 = t.ListView.Columns().Add()
			lv1.SetCaption("客户分类")
			lv1.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

			lv1 = t.ListView.Columns().Add()
			lv1.SetCaption("来源")
			lv1.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

		}

	})
	//	选择导入类型事件触发
	t.RadioButton2.SetOnClick(func(sender vcl.IObject) {
		if t.RadioButton2.Checked() == true {
			t.ProgressBar.SetPosition(0) //清空进度
			t.Open_Import = false

			Import_Type_Name = t.RadioButton2.Name()
			t.ListView.Clear()
			t.ListView.Columns().Clear()

			//初始化表头
			lvl := t.ListView.Columns().Add()
			lvl.SetCaption("导入结果")
			lvl.SetWidth(80)
			lvl.SetAlignment(types.TaCenter)

			lvl = t.ListView.Columns().Add()
			lvl.SetCaption("单据日期")
			lvl.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

			lvl = t.ListView.Columns().Add()
			lvl.SetCaption("工号")
			lvl.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

			lvl = t.ListView.Columns().Add()
			lvl.SetCaption("客户姓名")
			lvl.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

			lvl = t.ListView.Columns().Add()
			lvl.SetCaption("客户编码")
			lvl.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

			lvl = t.ListView.Columns().Add()
			lvl.SetCaption("客户电话")
			lvl.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

			lvl = t.ListView.Columns().Add()
			lvl.SetCaption("客户地址")
			lvl.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

			lvl = t.ListView.Columns().Add()
			lvl.SetCaption("客户分类")
			lvl.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

			lvl = t.ListView.Columns().Add()
			lvl.SetCaption("客户来源")
			lvl.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

			lvl = t.ListView.Columns().Add()
			lvl.SetCaption("库房")
			lvl.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

			lvl = t.ListView.Columns().Add()
			lvl.SetCaption("类别")
			lvl.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

			lvl = t.ListView.Columns().Add()
			lvl.SetCaption("票号")
			lvl.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

			lvl = t.ListView.Columns().Add()
			lvl.SetCaption("商品ID")
			lvl.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

			lvl = t.ListView.Columns().Add()
			lvl.SetCaption("数量")
			lvl.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

			lvl = t.ListView.Columns().Add()
			lvl.SetCaption("单价")
			lvl.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

			lvl = t.ListView.Columns().Add()
			lvl.SetCaption("备注")
			lvl.SetWidth(t.PanelB.Width() / 10)
			lvl.SetAlignment(types.TaCenter)

		}
	})
	//	选择导入类型事件触发
	t.RadioButton3.SetOnClick(func(sender vcl.IObject) {
		if t.RadioButton3.Checked() == true {
			t.ProgressBar.SetPosition(0) //清空进度
			t.Open_Import = false
			Import_Type_Name = t.RadioButton1.Name()
			t.ListView.Clear()
			t.ListView.Columns().Clear()
			vcl.ShowMessageFmt("正在努力开发中！！！")
		}
	})
	//	选择导入类型事件触发
	t.RadioButton4.SetOnClick(func(sender vcl.IObject) {
		if t.RadioButton4.Checked() == true {
			t.ProgressBar.SetPosition(0) //清空进度
			t.Open_Import = false
			Import_Type_Name = t.RadioButton1.Name()
			t.ListView.Clear()
			t.ListView.Columns().Clear()
			vcl.ShowMessageFmt("正在努力开发中！！！")
		}
	})

	//	打开文件按钮事件获取文件路径
	t.OpenButton.SetOnClick(func(sender vcl.IObject) {
		t.ProgressBar.SetPosition(0) //清空进度
		//去除不完善选项的错误提示
		if t.RadioButton3.Checked() == true {
			vcl.ShowMessageFmt("%v : 正在努力开发中！！！", t.RadioButton3.Name())
			return
		}
		if t.RadioButton4.Checked() == true {
			vcl.ShowMessageFmt("%v : 正在努力开发中！！！", t.RadioButton4.Name())
			return
		}

		//t.OpenTextFileDialog.Execute() //调起打开文件对话框
		//取消当未选择文件或直接关闭穿窗口时的错误提示
		//if t.OpenTextFileDialog.FileName() == "" {
		//	return
		//}

		//if t.OpenTextFileDialog.Execute() {
		//
		//	Data, err := OpenFile(t.OpenTextFileDialog.FileName())
		//	if err != nil {
		//		vcl.ShowMessageFmt("打开文件出错：%v", err)
		//		return
		//	}
		//	fmt.Println(len(Data))
		//	t.Windows.SetEnabled(false)
		//	t.Button.SetEnabled(false)
		//	t.OpenButton.SetEnabled(false)
		//	t.RadioButton1.SetEnabled(false)
		//	t.RadioButton2.SetEnabled(false)
		//	t.RadioButton3.SetEnabled(false)
		//	t.RadioButton4.SetEnabled(false)
		//	t.ListView.SetEnabled(false)
		//	t.PanelA.SetEnabled(false)
		//	t.PanelB.SetEnabled(false)
		//	t.PanelC.SetEnabled(false)
		//	t.ProgressBar.SetEnabled(false)
		//	//len1 := 0
		//	//绘制列表数据
		//	t.ListView.Clear()
		//	//t.ListView.Items().BeginUpdate()
		//	go func() {
		//		for k, v := range Data {
		//			if k == 0 { //对比表头信息
		//				//len1 = len(v)
		//				//fmt.Println(len(v))
		//				for i, d := range v {
		//					//fmt.Println(t.ListView.Column(int32(i) + 1).DisplayName())
		//					if d == t.ListView.Column(int32(i)+1).DisplayName() { //判断导入表格与用户选择的导入类型是否匹配
		//						continue
		//					} else {
		//						vcl.ShowMessageFmt("%v 与 %v 不匹配！", t.OpenTextFileDialog.FileName(), Import_Type_Name)
		//						t.Windows.SetEnabled(true)
		//						t.Button.SetEnabled(true)
		//						t.OpenButton.SetEnabled(true)
		//						t.RadioButton1.SetEnabled(true)
		//						t.RadioButton2.SetEnabled(true)
		//						t.RadioButton3.SetEnabled(true)
		//						t.RadioButton4.SetEnabled(true)
		//						t.ListView.SetEnabled(true)
		//						t.PanelA.SetEnabled(true)
		//						t.PanelB.SetEnabled(true)
		//						t.PanelC.SetEnabled(true)
		//						t.ProgressBar.SetEnabled(true)
		//						return
		//					}
		//				}
		//				continue
		//			}
		//
		//			//绘制数据表数据
		//
		//			item := t.ListView.Items().Add()
		//			//item.SetImageIndex(0)
		//			//item.SetCaption("")
		//
		//			for _, d := range v {
		//				vcl.ThreadSync(func() {
		//					item.SubItems().Add(d)
		//				})
		//			}
		//		}
		//		//t.ListView.Items().EndUpdate()
		//		t.Open_Import = true
		//		t.Windows.SetEnabled(true)
		//		t.Button.SetEnabled(true)
		//		t.OpenButton.SetEnabled(true)
		//		t.RadioButton1.SetEnabled(true)
		//		t.RadioButton2.SetEnabled(true)
		//		t.RadioButton3.SetEnabled(true)
		//		t.RadioButton4.SetEnabled(true)
		//		t.ListView.SetEnabled(true)
		//		t.PanelA.SetEnabled(true)
		//		t.PanelB.SetEnabled(true)
		//		t.PanelC.SetEnabled(true)
		//		t.ProgressBar.SetEnabled(true)
		//	}()
		//}

	})

	//	导入按钮事件，判断进行导入的数据类型
	t.Button.SetOnClick(func(sender vcl.IObject) {

		t.Windows.SetEnabled(false)
		t.Button.SetEnabled(false)
		t.OpenButton.SetEnabled(false)
		t.RadioButton1.SetEnabled(false)
		t.RadioButton2.SetEnabled(false)
		t.RadioButton3.SetEnabled(false)
		t.RadioButton4.SetEnabled(false)
		t.ListView.SetEnabled(false)
		t.PanelA.SetEnabled(false)
		t.PanelB.SetEnabled(false)
		t.PanelC.SetEnabled(false)
		t.ProgressBar.SetEnabled(false)

		//	新建客户事件处理

		if t.RadioButton1.Checked() {
			if !t.Open_Import {
				if t.ListView.Items().Count() == 0 {
					vcl.ShowMessageFmt("请先导入数据！")
					return
				} else {
					vcl.ShowMessageFmt("请勿重复导入！")
					return
				}
			}
			e := new(Engine)
			err := e.NewEngine()
			if err != nil {

				vcl.ShowMessageFmt("数据库链接出错：%v", err)
				t.Windows.SetEnabled(true)
				t.Button.SetEnabled(true)
				t.OpenButton.SetEnabled(true)
				t.RadioButton1.SetEnabled(true)
				t.RadioButton2.SetEnabled(true)
				t.RadioButton3.SetEnabled(true)
				t.RadioButton4.SetEnabled(true)
				t.ListView.SetEnabled(true)
				t.PanelA.SetEnabled(true)
				t.PanelB.SetEnabled(true)
				t.PanelC.SetEnabled(true)
				t.ProgressBar.SetEnabled(true)

				return
			}

			//defer e.Engine.Close()
			//fmt.Println(t.ListView.Items().Count())
			//fmt.Println(t.ListView.Items().Item(int32(1)).SubItems().Strings(1))
			go func() {
				for i := 0; i < int(t.ListView.Items().Count()); i++ {

					_, err = e.InSetClient(t.ListView.Items().Item(int32(i)).SubItems().Strings(0), t.ListView.Items().Item(int32(i)).SubItems().Strings(1), t.ListView.Items().Item(int32(i)).SubItems().Strings(2), t.ListView.Items().Item(int32(i)).SubItems().Strings(3), t.ListView.Items().Item(int32(i)).SubItems().Strings(4), t.ListView.Items().Item(int32(i)).SubItems().Strings(5))

					if err != nil {
						vcl.ThreadSync(func() {
							t.ListView.Items().Item(int32(i)).SetCaption(fmt.Sprintf("%v", err))
							t.ProgressBar.SetPosition(int32(i+1) / t.ListView.Items().Count() * 100)
						})
						continue
					} else {
						vcl.ThreadSync(func() {
							t.ListView.Items().Item(int32(i)).SetCaption("导入成功")
						})
					}

					if i == int(t.ListView.Items().Count()-1) {
						e.Engine.Close()
						t.Open_Import = false
					}

					vcl.ThreadSync(func() {
						t.ProgressBar.SetPosition(int32(i+1) / t.ListView.Items().Count() * 100)
					})

				}
				vcl.ThreadSync(func() {
					vcl.ShowMessageFmt("导入完成！")
				})
			}()

		}

		if t.RadioButton2.Checked() {
			if !t.Open_Import {
				if t.ListView.Items().Count() == 0 {
					vcl.ShowMessageFmt("请先导入数据！")
					return
				} else {
					vcl.ShowMessageFmt("请勿重复导入！")
					return
				}
			}
			e := new(Engine)
			err := e.NewEngine()
			if err != nil {
				vcl.ShowMessageFmt("数据库链接出错：%v", err)
				t.Windows.SetEnabled(true)
				t.Button.SetEnabled(true)
				t.OpenButton.SetEnabled(true)
				t.RadioButton1.SetEnabled(true)
				t.RadioButton2.SetEnabled(true)
				t.RadioButton3.SetEnabled(true)
				t.RadioButton4.SetEnabled(true)
				t.ListView.SetEnabled(true)
				t.PanelA.SetEnabled(true)
				t.PanelB.SetEnabled(true)
				t.PanelC.SetEnabled(true)
				t.ProgressBar.SetEnabled(true)
				return

			}
			Market := new(Market)
			num := 1
			go func() {
				for i := 0; i < int(t.ListView.Items().Count()); i++ {
					Market.TimeNumber = fmt.Sprintf("%2s%02d%02d%02d%02d%02d%03d\n", time.Now().Format("06"),
						int(time.Now().Month()), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), num)
					Market.DateTime = convertToFormatDay(t.ListView.Items().Item(int32(i)).SubItems().Strings(0)) //单据日期
					Market.Job_number = t.ListView.Items().Item(int32(i)).SubItems().Strings(1)                   //工号
					Market.Customer_name = t.ListView.Items().Item(int32(i)).SubItems().Strings(2)                //客户姓名
					Market.Customer_no = t.ListView.Items().Item(int32(i)).SubItems().Strings(3)                  //客户编码
					Market.Customer_phon = t.ListView.Items().Item(int32(i)).SubItems().Strings(4)                //客户电话
					Market.Customer_site = t.ListView.Items().Item(int32(i)).SubItems().Strings(5)                //客户地址
					Market.Customer_type = t.ListView.Items().Item(int32(i)).SubItems().Strings(6)                //客户分类
					Market.Customer_source = t.ListView.Items().Item(int32(i)).SubItems().Strings(7)              //客户来源
					Market.Warehouse = t.ListView.Items().Item(int32(i)).SubItems().Strings(8)                    //库房
					Market.For_Sale = t.ListView.Items().Item(int32(i)).SubItems().Strings(9)                     //类型
					Market.Exchange_shop = t.ListView.Items().Item(int32(i)).SubItems().Strings(10)               //票号
					Market.Commodity = t.ListView.Items().Item(int32(i)).SubItems().Strings(11)                   //商品ID
					Market.Num = t.ListView.Items().Item(int32(i)).SubItems().Strings(12)                         //数量
					Market.Price_per_uni = t.ListView.Items().Item(int32(i)).SubItems().Strings(13)               //单价
					Market.Remarks = t.ListView.Items().Item(int32(i)).SubItems().Strings(14)                     //备注
					Market.Operator = t.Operator

					if Market.Customer_no == "" {
						Market.Customer_no, err = e.InSetClient(Market.Customer_name, Market.Customer_phon, Market.Customer_site, Market.Job_number, Market.Customer_type, Market.Customer_source)
						if Market.Customer_no == "" {
							vcl.ThreadSync(func() {
								t.ListView.Items().Item(int32(i)).SetCaption(fmt.Sprintf("错误：%v", err))
								t.ProgressBar.SetPosition(int32(i+1) / t.ListView.Items().Count() * 100)
							})
							continue
						} else if err != nil {
							Market.New_and_old_clients = "新客户"
						} else {
							Market.New_and_old_clients = "老客户"
						}

					} else {
						Market.New_and_old_clients = "老客户"
					}

					err = e.InSetConsume(Market)

					if err == nil {
						vcl.ThreadSync(func() {
							t.ListView.Items().Item(int32(i)).SetCaption("导入成功")
						})
					} else {
						vcl.ThreadSync(func() {
							t.ListView.Items().Item(int32(i)).SetCaption(fmt.Sprintf("错误：%v", err))
						})

					}

					//提交数据
					if i == int(t.ListView.Items().Count()-1) {
						e.Engine.Close()
						t.Open_Import = false
					}

					//循环数字
					if num == 999 {
						num = 1
					}
					num++
					vcl.ThreadSync(func() {
						t.ProgressBar.SetPosition(int32(i+1) / t.ListView.Items().Count() * 100)
					})
				}

				vcl.ThreadSync(func() {
					vcl.ShowMessageFmt("导入完成！")
				})
			}()

		}

		if t.RadioButton3.Checked() == true {
			vcl.ShowMessageFmt(t.RadioButton3.Name() + ":正在努力开发中！！！")
		}

		if t.RadioButton4.Checked() == true {
			vcl.ShowMessageFmt(t.RadioButton4.Name() + ":正在努力开发中！！！")
		}

		t.Windows.SetEnabled(true)
		t.Button.SetEnabled(true)
		t.OpenButton.SetEnabled(true)
		t.RadioButton1.SetEnabled(true)
		t.RadioButton2.SetEnabled(true)
		t.RadioButton3.SetEnabled(true)
		t.RadioButton4.SetEnabled(true)
		t.ListView.SetEnabled(true)
		t.PanelA.SetEnabled(true)
		t.PanelB.SetEnabled(true)
		t.PanelC.SetEnabled(true)
		t.ProgressBar.SetEnabled(true)
	})

	//自动列表排序事件
	t.ListView.SetOnColumnClick(func(sender vcl.IObject, column *vcl.TListColumn) {
		t.ListView.CustomSort(0, int(column.Index()))

	})

	//自定义排序方法
	t.ListView.SetOnCompare(func(sender vcl.IObject, item1, item2 *vcl.TListItem, data int32, compare *int32) {
		if data == 0 {
			*compare = int32(strings.Compare(item1.Caption(), item2.Caption()))
		} else {
			*compare = int32(strings.Compare(item1.SubItems().Strings(data-1), item2.SubItems().Strings(data-1)))
		}
	})

}

//打开用户指定的文件
func OpenFile(fileName string) ([][]string, error) {
	//打开文件
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		return nil, err
	}

	//读取指定工作薄
	rows := f.GetRows("Sheet1")

	return rows, err
}

//新建客户  - 姓名 电话  地址  来源
func (e *Engine) InSetClient(KHMC string, MOBIL string, DIZHI string, GONGHAO string, TYPEID string, SOURCEID string) (string, error) {

	Query, err := e.Engine.Query(fmt.Sprintf("SELECT KHID FROM BLCRM.CRM_DAT001 WHERE MOBIL = '%v'", MOBIL))

	if err != nil {
		return "", err
	}

	for _, v := range Query {
		for _, k := range v {
			if string(k) != "" {
				return string(k), fmt.Errorf("客户已存在！")
			}
		}
	}

	if KHMC == "" {
		return "", fmt.Errorf("客户姓名为空")
	} else if MOBIL == "" {
		return "", fmt.Errorf("客户电话为空")
	} else if TYPEID == "" {
		return "", fmt.Errorf("客户分类为空")
	} else if SOURCEID == "" {
		return "", fmt.Errorf("客户来源为空")
	}

	Sql := fmt.Sprintf("INSERT INTO BLCRM.CRM_DAT001(KHMC,TYPEID,MOBIL,DIZHI,HUIFANG,HUIFANGJG,SOURCEID,ISVIP,RESERVE,GONGHAO)"+
		" VALUES('%s','%v',%s,'%s',1,5,%v,0,1,'%v')", KHMC, TYPEID, MOBIL, DIZHI, SOURCEID, GONGHAO)

	_, e.Err = e.Engine.Exec(Sql)
	if e.Err != nil {
		return "", e.Err
	}

	Query, err = e.Engine.Query(fmt.Sprintf("SELECT KHID FROM BLCRM.CRM_DAT001 WHERE MOBIL = '%v'", MOBIL))

	if err != nil {
		return "", err
	}

	for _, v := range Query {
		for _, k := range v {
			if string(k) != "" {
				return string(k), nil
			}
		}
	}

	return "", nil
}

//新建销售记录  随机的时间数  日期 客户的ID  产品ID
func (e *Engine) InSetConsume(sell *Market) error {
	if sell.Customer_name == "" {
		return fmt.Errorf("客户姓名为空")
	} else if sell.Customer_no == "" {
		return fmt.Errorf("客户编码为空")
	} else if sell.DateTime == "" {
		return fmt.Errorf("单据日期为空！")
	} else if sell.Job_number == "" {
		return fmt.Errorf("销售员工工号为空！")
	} else if sell.Warehouse == "" {
		return fmt.Errorf("库房为空！")
	} else if sell.For_Sale == "" {
		return fmt.Errorf("销售类型为空！")
	} else if sell.Commodity == "" {
		return fmt.Errorf("商品ID为空！")
	} else if sell.Num == "" {
		return fmt.Errorf("商品数量为空！")
	} else if sell.Price_per_uni == "" {
		return fmt.Errorf("单价为空！")
	}

	// BILLNO 随机时间  BILLTYPEID 导入类型  BILLDATE  单据时间  BILLFLAG 默认1  SALESMAN_ID  员工工号
	// OPER_ID 操作员工号 CLIENT_ID  客户ID  WAREHOUSESID 库房   PIAOHAO  票号   NEWOROLD  新老客户
	Sql := fmt.Sprintf("INSERT INTO BLCRM.CRM_DAT007(BILLNO,BILLTYPEID,BILLDATE,BILLFLAG,SALESMAN_ID,OPER_ID,CLIENT_ID,WAREHOUSESID,PIAOHAO,NEWOROLD)"+
		" VALUES(%s,140001,TO_DATE('%s','YYYY-MM-DD'),1,%v,%v,%v,%v,'%v','%v')", sell.TimeNumber, sell.DateTime, sell.Job_number, sell.Operator,
		sell.Customer_no, sell.Warehouse, sell.Exchange_shop, sell.New_and_old_clients)

	_, e.Err = e.Engine.Exec(Sql)
	if e.Err != nil {
		fmt.Println(e.Err)
		return e.Err
	}

	//  BILLNO 随机时间  GOODSID 工号  AMOUNT 数量  SUMPRICE 总价  PRICE  单价  REMARKS  备注  KIND  销售类型
	Num, _ := strconv.Atoi(sell.Num)
	Price_per_uni, _ := strconv.Atoi(sell.Price_per_uni)
	sum := Num * Price_per_uni
	SqL := fmt.Sprintf("INSERT INTO BLCRM.CRM_DAT009(BILLNO,GOODSID,AMOUNT,SUMPRICE,PRICE,REMARKS,KIND)"+
		" VALUES(%v,%v,%v,%v,%v,'%v','%v')", sell.TimeNumber, sell.Job_number, sell.Num, sum, sell.Price_per_uni, sell.Remarks, sell.For_Sale)

	_, e.Err = e.Engine.Exec(SqL)
	if e.Err != nil {
		return e.Err
	}

	integral := 0
	if sell.New_and_old_clients == "老客户" {
		Query, err := e.Engine.Query(fmt.Sprintf("SELECT JIFEN FROM BLCRM.CRM_DAT001 WHERE KHID = '%v'", sell.Customer_no))

		if err != nil {
			return err
		}

		for _, v := range Query {
			for _, k := range v {
				if string(k) != "" {
					integral, _ = strconv.Atoi(string(k))
					integral = integral + sum/100
				}
			}
		}
	} else {
		integral = sum / 100
	}

	_, e.Err = e.Engine.Exec(fmt.Sprintf("UPDATE BLCRM.CRM_DAT001 SET JIFEN = %v WHERE KHID = %v", integral, sell.Customer_no))

	if e.Err != nil {
		return e.Err
	}

	return nil
}

// excel日期字段格式化 yyyy-mm-dd
func convertToFormatDay(excelDaysString string) string {
	// 2006-01-02 距离 1900-01-01的天数
	baseDiffDay := 38719 //在网上工具计算的天数需要加2天，什么原因没弄清楚
	curDiffDay := excelDaysString
	b, _ := strconv.Atoi(curDiffDay)
	// 获取excel的日期距离2006-01-02的天数
	realDiffDay := b - baseDiffDay
	//fmt.Println("realDiffDay:",realDiffDay)
	// 距离2006-01-02 秒数
	realDiffSecond := realDiffDay * 24 * 3600
	//fmt.Println("realDiffSecond:",realDiffSecond)
	// 2006-01-02 15:04:05距离1970-01-01 08:00:00的秒数 网上工具可查出
	baseOriginSecond := 1136185445
	resultTime := time.Unix(int64(baseOriginSecond+realDiffSecond), 0).Format("2006-01-02")
	return resultTime
}
