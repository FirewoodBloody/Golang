package main

import (
	"Golang/Express_Routing/models"
	"fmt"
	"github.com/FirewoodBloody/PacketProup/logs"
)

var (
	logLevel = make(map[string]string, 10)
	logger   logs.LogInterface
)

func logsInit() {
	logLevel["loglevel"] = "info"
	logLevel["logPath"] = "./logs"
	logLevel["logName"] = "access"
	logger, err := logs.InitLogger("file", logLevel)

	if err != nil {
		fmt.Println(err)
		logger.Error("打开日志文件错误：", err)
	}
}

func main() {
	logsInit()

	engine := models.Engine{}
	engine.Err = engine.NewEngine()
	if engine.Err == nil {
		logger.Error("创建engine连接错误：", engine.Err)
	}

	engine.Err = engine.Engine.Ping()
	if engine.Err != nil {
		logger.Error("建立数据库连接失败：", engine.Err)
	}

}
