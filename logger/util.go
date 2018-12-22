package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
)

type LogData struct {
	Messags      string
	TimeStr      string
	LevelStr     string
	FileName     string
	FuncName     string
	LineNo       int
	WarnAndFaral bool
}

//获取当前错误的文件名，函数以及行数
func GetlineInfo() (fileName string, funcName string, lineNo int) {
	pc, file, line, ok := runtime.Caller(4)
	if ok {
		fileName = file
		funcName = runtime.FuncForPC(pc).Name()
		lineNo = line
	}
	return
}

//定义日志写入方法
func LogWrite(file *os.File, now_level int, level int, format string, args ...interface{}) *LogData {
	if now_level > level {
		return nil
	}

	now := time.Now().Format("2006-01-02 15:04:05.999")           //获取当前操作的时间，并进行格式化
	levelStr := GetLevelTest(level)                               //判断当前运行级别string形式
	fileName, funcName, lineNo := GetlineInfo()                   //获取当前运行的文件以及错误行号和函数
	fileName, funcName = path.Base(fileName), path.Base(funcName) //截取文件文件名
	msg := fmt.Sprintf(format, args...)                           //格式化用户的输入
	LogData := &LogData{
		Messags:      msg,
		TimeStr:      now,
		LevelStr:     levelStr,
		FileName:     fileName,
		FuncName:     funcName,
		LineNo:       lineNo,
		WarnAndFaral: false,
	}
	if level == LogLevelError || level == LogLevelFatal || level == LogLevelWarn {
		LogData.WarnAndFaral = true
	}
	return LogData
	//fmt.Fprintf(file, "%s [ %s ] ( %s:%d:%v ) %s\n", now, levelStr, fileName, lineNo, funcName, msg) //写入日志
}
