package logger

var log LogInterface

/*
name 的值为：file和console
file为文件，console为终端
*/
func InitLogger(name string, config map[string]string) (err error) {
	if name == "file" {
		log, err = NewFilelogger(config)
	} else if name == "console" {
		log, err = NewConsoleLogger(config)
	}
	return err
}

func Debug(format string, args ...interface{}) {
	log.Debug(format, args...)
}

func Trace(format string, args ...interface{}) {
	log.Trace(format, args...)
}

func Info(format string, args ...interface{}) {
	log.Info(format, args...)
}

func Warn(format string, args ...interface{}) {
	log.Warn(format, args...)
}

func Error(format string, args ...interface{}) {
	log.Error(format, args...)
}

func Fatal(format string, args ...interface{}) {
	log.Fatal(format, args...)
}
