package main

import (
	"fmt"
	"github.com/FirewoodBloody/logs"
	"time"
)

func initLogger(name, level, logPath, logName string) error {
	config := make(map[string]string)
	config["loglevel"] = level
	config["logPath"] = logPath
	config["logName"] = logName
	config["logSplitSize"] = "100000"
	err := logs.InitLogger(name, config)
	if err != nil {
		return err
	}
	logs.Debug("init logs success!")
	return nil
}

func Run() {
	for {
		logs.Debug("user server running")
		logs.Fatal("then is Fatal!")
		//time.Sleep(time.Second)
	}
}

func nowtimeDate() {
	for {
		now := time.Now().Format("2006-01-02")
		time.Sleep(time.Second)
		nows := time.Now().Format("2006-01-02")
		if now != nows {
			logName := fmt.Sprintf("%v_%s", time.Now().Format("2006-01-02"), "log")
			initLogger("file", "debug", "D:/log/", logName)
		}
	}
}
func main() {
	//go nowtimeDate()
	//logName := fmt.Sprintf("%v_%s", time.Now().Format("2006-01-02"), "log") //定义日志文件名加上date
	initLogger("file", "debug", "D:/log/", "111")
	Run()
	return
}
