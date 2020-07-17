package main

import (
	"fmt"
	"github.com/Luxurioust/excelize"
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
	start_at      = "2020-07-01 00:00:00"
	start_at_type = "%Y-%m-%d %H:%i:%s"
	end_at        = "2020-07-15 23:59:59"
)

type Voice_Download struct {
	start_at  string
	end_at    string
	user_id   string
	user_code string
	user_name string
	times     string
	depname   string
	url       string
	file      string
}

func init() {
	os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8") //修正中文乱码
}

func main() {
	file, _ := excelize.OpenFile("./call_log.xlsx")

	var err error
	V := new(Voice_Download)

	engine, err := xorm.NewEngine(driverName, dBconnect)
	if err != nil {
		fmt.Println(err)
	}

	engine.ShowSQL(true)
	//nomao, _ := engine.Query(fmt.Sprintf("SELECT DATE_FORMAT(cqv.start_at,'%v'), DATE_FORMAT(cqv.end_at,'%v'), cast(cqv.user_id as char), cast(cqv.user_name as char),cast(d.NAME as char),JSON_UNQUOTE(cqv.raw_json -> \"$.record_pbx\"),JSON_UNQUOTE(cqv.raw_json -> \"$.record_file_name\"),cast(cqv.call_duration as char),cast(u.login_name as char) FROM `bl_crm_quality_voice` cqv INNER JOIN bl_users u ON cqv.user_id = u.id INNER JOIN bl_depart d ON u.depart_id = d.id WHERE start_at > '%v' AND start_at < '%v' AND call_duration >= 60 AND call_type = 'extension_outbound' AND LENGTH( raw_json -> \"$.record_file_name\" ) > 5 and d.name = '客户二中心 ';", start_at_type, start_at_type, start_at, end_at))
	nomao, _ := engine.Query(fmt.Sprintf("SELECT DATE_FORMAT(cqv.start_at,'%v'),DATE_FORMAT(cqv.end_at,'%v'),cast(cqv.user_id as char),cast(cqv.user_name as char),cast(d.NAME as char),JSON_UNQUOTE(cqv.raw_json -> \"$.record_pbx\"),JSON_UNQUOTE(cqv.raw_json -> \"$.record_file_name\"),cast(cqv.call_duration as char),cast(u.login_name as char) FROM `bl_crm_quality_voice` cqv INNER JOIN bl_users u ON cqv.user_id = u.id INNER JOIN bl_depart d ON u.depart_id = d.id WHERE start_at > '%v' AND start_at < '%v' AND call_duration >= 90 AND call_type = 'extension_outbound' AND LENGTH( raw_json -> \"$.record_file_name\" ) > 5 and u.login_name in ( '20040013','20040014','20040022','20040029','20040030','20050008','20060018','20060019');", start_at_type, start_at_type, start_at, end_at))

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

			filename := V.user_code + fmt.Sprintf("_%v_", time.Now().Unix()) + V.times + ".mp3"
			err = CreateFile(data, "./"+V.depname+"/"+V.user_name+"/"+filename)
			//fmt.Println("./" + V.depname + "/" + V.user_name + "/" + filename)
			if err != nil {
				fmt.Printf("create failed![%v]", err)
			}
			file.SetCellValue("Sheet1", fmt.Sprintf("A%v", s+1), V.start_at)
			file.SetCellValue("Sheet1", fmt.Sprintf("B%v", s+1), V.end_at)
			file.SetCellValue("Sheet1", fmt.Sprintf("C%v", s+1), V.depname)
			file.SetCellValue("Sheet1", fmt.Sprintf("D%v", s+1), V.user_id)
			file.SetCellValue("Sheet1", fmt.Sprintf("E%v", s+1), V.user_code)
			file.SetCellValue("Sheet1", fmt.Sprintf("F%v", s+1), V.user_name)
			file.SetCellValue("Sheet1", fmt.Sprintf("G%v", s+1), V.times)
			file.SetCellValue("Sheet1", fmt.Sprintf("H%v", s+1), V.url+V.file)
			file.SetCellValue("Sheet1", fmt.Sprintf("I%v", s+1), filename)

		} else {
			filename := V.user_code + fmt.Sprintf("_%v_", time.Now().Unix()) + V.times + ".mp3"
			err = CreateFile(data, "./"+V.depname+"/"+V.user_name+"/"+filename)
			//fmt.Println("./" + V.depname + "/" + V.user_name + "/" + filename)
			if err != nil {
				fmt.Printf("create failed![%v]", err)
			}
			file.SetCellValue("Sheet1", fmt.Sprintf("A%v", s+1), V.start_at)
			file.SetCellValue("Sheet1", fmt.Sprintf("B%v", s+1), V.end_at)
			file.SetCellValue("Sheet1", fmt.Sprintf("C%v", s+1), V.depname)
			file.SetCellValue("Sheet1", fmt.Sprintf("D%v", s+1), V.user_id)
			file.SetCellValue("Sheet1", fmt.Sprintf("E%v", s+1), V.user_code)
			file.SetCellValue("Sheet1", fmt.Sprintf("F%v", s+1), V.user_name)
			file.SetCellValue("Sheet1", fmt.Sprintf("G%v", s+1), V.times)
			file.SetCellValue("Sheet1", fmt.Sprintf("H%v", s+1), V.url+V.file)
			file.SetCellValue("Sheet1", fmt.Sprintf("I%v", s+1), filename)
		}
		s++
		time.Sleep(time.Second * 3)
	}
	file.Save()
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
