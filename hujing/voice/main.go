package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"os"
	"time"
	"xorm.io/xorm"
)

const (
	TimeFormat    = "2006-01-02"
	driverName    = "mysql"
	dBconnect     = "root:dcf4aa3f7b982ce4@tcp(192.168.0.11:3306)/bl_crm?charset=utf8"
	start_at      = "2020-08-14 00:00:00"
	start_at_type = "%Y-%m-%d %H:%i:%s"
	end_at        = "2020-08-14 23:59:59"
)

type Voice_Download struct {
	start_at      string
	end_at        string
	user_id       string
	user_code     string
	user_name     string
	times         string
	depname       string
	url           string
	file          string
	customer_name string
}

func init() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码
}

func main() {
	//file, _ := excelize.OpenFile("./call_log.xlsx")

	var err error
	V := new(Voice_Download)

	engine, err := xorm.NewEngine(driverName, dBconnect)
	if err != nil {
		fmt.Println(err)
	}

	engine.ShowSQL(true)
	nomao, _ := engine.Query(fmt.Sprintf("SELECT cqv.customer_name,DATE_FORMAT(cqv.start_at,'%v'), DATE_FORMAT(cqv.end_at,'%v'), cast(cqv.user_id as char), cast(cqv.user_name as char),cast(d.NAME as char),JSON_UNQUOTE(cqv.raw_json -> \"$.record_pbx\"),JSON_UNQUOTE(cqv.raw_json -> \"$.record_file_name\"),cast(cqv.call_duration as char),cast(u.login_name as char) FROM `bl_crm_quality_voice` cqv INNER JOIN bl_users u ON cqv.user_id = u.id INNER JOIN bl_depart d ON u.depart_id = d.id WHERE start_at > '%v' AND start_at < '%v' AND call_duration >= 180 AND call_type = 'extension_outbound' AND LENGTH( raw_json -> \"$.record_file_name\" ) > 5 ;", start_at_type, start_at_type, start_at, end_at))
	//nomao, _ := engine.Query(fmt.Sprintf("SELECT DATE_FORMAT(cqv.start_at,'%v'),DATE_FORMAT(cqv.end_at,'%v'),cast(cqv.user_id as char),cast(cqv.user_name as char),cast(d.NAME as char),JSON_UNQUOTE(cqv.raw_json -> \"$.record_pbx\"),JSON_UNQUOTE(cqv.raw_json -> \"$.record_file_name\"),cast(cqv.call_duration as char),cast(u.login_name as char) FROM `bl_crm_quality_voice` cqv INNER JOIN bl_users u ON cqv.user_id = u.id INNER JOIN bl_depart d ON u.depart_id = d.id WHERE start_at > '%v' AND start_at < '%v' AND call_duration >= 180 AND call_type = 'extension_outbound' AND LENGTH( raw_json -> \"$.record_file_name\" ) > 5 and u.login_name in ('13020015',\n'19100014',\n'18080007',\n'18060008',\n'18050012',\n'19070009',\n'20040016',\n'20030030');", start_at_type, start_at_type, start_at, end_at))
	//nomao, _ := engine.Query(fmt.Sprintf("SELECT DATE_FORMAT(cqv.start_at,'%v'),DATE_FORMAT(cqv.end_at,'%v'),cast(cqv.user_id as char),cast(cqv.user_name as char),cast(d.NAME as char),JSON_UNQUOTE(cqv.raw_json -> \"$.record_pbx\"),JSON_UNQUOTE(cqv.raw_json -> \"$.record_file_name\"),cast(cqv.call_duration as char),cast(u.login_name as char) FROM `bl_crm_quality_voice` cqv INNER JOIN bl_users u ON cqv.user_id = u.id INNER JOIN bl_depart d ON u.depart_id = d.id WHERE start_at > '%v' AND start_at < '%v' AND call_duration >= 180 AND call_type = 'extension_outbound' AND LENGTH( raw_json -> \"$.record_file_name\" ) > 5 and  cqv.called_no in ('13701004505',\n'13067778130',\n'13922832177',\n'13877334652',\n'13902607011',\n'15325359156',\n'18261941949',\n'15255524768',\n'18602928025',\n'13895053019',\n'13308409869');", start_at_type, start_at_type, start_at, end_at))

	engine.Close()
	fmt.Println(len(nomao))
	s := 0
	for _, v := range nomao {

		V.start_at = string(v["DATE_FORMAT(cqv.start_at,'%Y-%m-%d %H:%i:%s')"])
		V.end_at = string(v["DATE_FORMAT(cqv.end_at,'%Y-%m-%d %H:%i:%s')"])
		V.user_id = string(v["cast(cqv.user_id as char)"])
		V.user_name = string(v["cast(cqv.user_name as char)"])
		V.user_code = string(v["cast(u.login_name as char)"])
		V.depname = string(v["cast(d.NAME as char)"])
		V.times = string(v["cast(cqv.call_duration as char)"])
		V.url = string(v["JSON_UNQUOTE(cqv.raw_json -> \"$.record_pbx\")"])
		V.file = string(v["JSON_UNQUOTE(cqv.raw_json -> \"$.record_file_name\")"])
		V.customer_name = string(v["customer_name"])
		//fmt.Println(V)
		//if V.depname != "客户二中心" {
		//	continue
		//}
		//http请求录音文件下载
		data, err := FileGet(V.url, V.file)
		if err != nil {
			fmt.Println(err)
		}
		// 判断文件夹是否存在
		status, _ := PathExists("./" + V.depname + "/" + V.user_name)
		if !status {
			_ = os.MkdirAll("./"+V.depname+"/"+V.user_name, os.ModePerm)

			filename := V.user_name + "-" + V.customer_name + fmt.Sprintf("_%v_", time.Now().Unix()) + V.times + ".mp3"
			err = CreateFile(data, "./"+V.depname+"/"+V.user_name+"/"+filename)
			//fmt.Println("./" + V.depname + "/" + V.user_name + "/" + filename)
			if err != nil {
				fmt.Printf("create failed![%v]", err)
			}
			//file.SetCellValue("Sheet1", fmt.Sprintf("A%v", s+1), V.start_at)
			//file.SetCellValue("Sheet1", fmt.Sprintf("B%v", s+1), V.end_at)
			//file.SetCellValue("Sheet1", fmt.Sprintf("C%v", s+1), V.depname)
			//file.SetCellValue("Sheet1", fmt.Sprintf("D%v", s+1), V.user_id)
			//file.SetCellValue("Sheet1", fmt.Sprintf("E%v", s+1), V.user_code)
			//file.SetCellValue("Sheet1", fmt.Sprintf("F%v", s+1), V.user_name)
			//file.SetCellValue("Sheet1", fmt.Sprintf("G%v", s+1), V.times)
			//file.SetCellValue("Sheet1", fmt.Sprintf("H%v", s+1), V.url+V.file)
			//file.SetCellValue("Sheet1", fmt.Sprintf("I%v", s+1), filename)

		} else {
			filename := V.user_name + "-" + V.customer_name + fmt.Sprintf("_%v_", time.Now().Unix()) + V.times + ".mp3"
			err = CreateFile(data, "./"+V.depname+"/"+V.user_name+"/"+filename)
			//fmt.Println("./" + V.depname + "/" + V.user_name + "/" + filename)
			if err != nil {
				fmt.Printf("create failed![%v]", err)
			}
			//file.SetCellValue("Sheet1", fmt.Sprintf("A%v", s+1), V.start_at)
			//file.SetCellValue("Sheet1", fmt.Sprintf("B%v", s+1), V.end_at)
			//file.SetCellValue("Sheet1", fmt.Sprintf("C%v", s+1), V.depname)
			//file.SetCellValue("Sheet1", fmt.Sprintf("D%v", s+1), V.user_id)
			//file.SetCellValue("Sheet1", fmt.Sprintf("E%v", s+1), V.user_code)
			//file.SetCellValue("Sheet1", fmt.Sprintf("F%v", s+1), V.user_name)
			//file.SetCellValue("Sheet1", fmt.Sprintf("G%v", s+1), V.times)
			//file.SetCellValue("Sheet1", fmt.Sprintf("H%v", s+1), V.url+V.file)
			//file.SetCellValue("Sheet1", fmt.Sprintf("I%v", s+1), filename)
		}
		s++
		time.Sleep(time.Second * 1)
	}
	//file.Save()
}

// 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, err
	}

	return false, err
}

//http请求录音文件下载
func FileGet(Url, fileName string) ([]byte, error) {
	UrlA := Url + fileName
	response, err := http.Get(UrlA)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

//存储录音文件
func CreateFile(data []byte, filename string) error {
	f, err := os.Create(filename)

	if err != nil {
		return err
	} else {
		_, err = f.Write(data)
		if err != nil {
			return err
		}
	}
	return nil
}
