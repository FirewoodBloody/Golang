package main

import (
	"Golang/logger"
	"fmt"
	"time"
)

func initLogger(name, level, logPath, logName string) error {
	config := make(map[string]string)
	config["loglevel"] = level
	config["logPath"] = logPath
	config["logName"] = logName
	err := logger.InitLogger(name, config)
	if err != nil {
		return err
	}
	logger.Debug("init logger success!")
	return nil
}

func Run() {
	for {
		logger.Debug("user server running")
		time.Sleep(time.Second)
	}
}

func nowtimeDate() {
	for {
		now := time.Now().Format("2006-01-02")
		time.Sleep(time.Second)
		nows := time.Now().Format("2006-01-02")
		if now != nows {
			logName := fmt.Sprintf("%v_%s", time.Now().Format("2006-01-02"), "log")
			initLogger("console", "debug", "D:/log/", logName)
		}
	}
}
func main() {
	go nowtimeDate()
	logName := fmt.Sprintf("%v_%s", time.Now().Format("2006-01-02"), "log") //定义日志文件名加上date
	initLogger("console", "debug", "D:/log/", logName)
	Run()
	return
}
