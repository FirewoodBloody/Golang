package logger

import (
	"fmt"
	"os"
)

type Filelogger struct {
	level    int
	logPath  string
	logName  string
	file     *os.File
	warnFile *os.File
}

func NewFilelogger(level int, logPath string, logName string) LogInterface {
	logger := &Filelogger{
		level:   level,
		logName: logName,
		logPath: logPath,
	}
	logger.init()
	return logger
}

func (f *Filelogger) init() {
	filename := fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed ,err:%v", filename, err))
	}
	f.file = file

	//写错误日志和fatal日志的文件
	filename = fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	file, err = os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed ,err:%v", filename, err))
	}
	f.warnFile = file

}

func (f *Filelogger) Debug(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	//str := fmt.Sprintf(format, args...)
	fmt.Fprintf(f.file, format, args...)
}

func (f *Filelogger) Trace(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.file, format, args...)
}

func (f *Filelogger) Info(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.file, format, args...)
}

func (f *Filelogger) Warn(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.file, format, args...)
}

func (f *Filelogger) Error(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.warnFile, format, args...)
}

func (f *Filelogger) Fatal(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	fmt.Fprintf(f.warnFile, format, args...)
}

func (f *Filelogger) Close() {
	f.file.Close()
	f.warnFile.Close()
}

func (f *Filelogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	f.level = level
}
