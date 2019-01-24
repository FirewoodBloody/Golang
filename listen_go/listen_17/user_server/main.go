package main

import "Golang/listen_go/listen_17/log"

func main() {
	//file := log.NewFilelog("maim.go")
	//file.LogDebug("log debug")
	//file.LogWarn("logwarn")

	//console := log.NewConsileLog("xxx")
	//console.LogWarn("then is logwarn")
	//console.LogDebug("then is logdebug")

	logs := log.NewFilelog("xxx")
	logs.LogDebug("then is logdebug")
	logs.LogWarn("then is logwarn")
}
