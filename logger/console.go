package logger

import (
	"fmt"
	"os"
)

type ConsoleLogger struct {
	level int
}

func NewConsoleLogger(config map[string]string) (LogInterface, error) {
	loglevel, ok := config["loglevel"]
	if !ok {
		err := fmt.Errorf("not found loglevel!")
		return nil, err
	}

	logger := &ConsoleLogger{
		level: GetLevel(loglevel),
	}
	return logger, nil
}

//设置终端打印日志的级别
func (c *ConsoleLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	c.level = level
}

//为struct定义借口方法
func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	logData := LogWrite(os.Stdout, c.level, LogLevelDebug, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s [%s:%d:%s] %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.LineNo, logData.FuncName, logData.Messags)
}

func (c *ConsoleLogger) Trace(format string, args ...interface{}) {
	logData := LogWrite(os.Stdout, c.level, LogLevelTrace, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s [%s:%d:%s] %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.LineNo, logData.FuncName, logData.Messags)

}

func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	logData := LogWrite(os.Stdout, c.level, LogLevelInfo, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s [%s:%d:%s] %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.LineNo, logData.FuncName, logData.Messags)

}

func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	logData := LogWrite(os.Stdout, c.level, LogLevelWarn, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s [%s:%d:%s] %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.LineNo, logData.FuncName, logData.Messags)

}

func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	logData := LogWrite(os.Stdout, c.level, LogLevelError, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s [%s:%d:%s] %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.LineNo, logData.FuncName, logData.Messags)

}

func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	logData := LogWrite(os.Stdout, c.level, LogLevelFatal, format, args...)
	fmt.Fprintf(os.Stdout, "%s %s [%s:%d:%s] %s\n", logData.TimeStr, logData.LevelStr, logData.FileName, logData.LineNo, logData.FuncName, logData.Messags)

}

func (c *ConsoleLogger) Close() {

}
