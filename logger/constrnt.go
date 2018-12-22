package logger

const (
	LogLevelDebug = iota
	LogLevelWarn
	LogLevelInfo
	LogLevelTrace
	LogLevelError
	LogLevelFatal
)

func GetLevelTest(level int) string {
	switch level {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelWarn:
		return "WARN"
	case LogLevelInfo:
		return "INFO"
	case LogLevelTrace:
		return "TRACE"
	case LogLevelError:
		return "ERROR"
	case LogLevelFatal:
		return "FATAL"
	}
	return "UNKNOWN"
}

func GetLevel(level string) int {
	switch level {
	case "debug", "Debug":
		return LogLevelDebug
	case "warn":
		return LogLevelWarn
	case "info":
		return LogLevelInfo
	case "trace":
		return LogLevelTrace
	case "error":
		return LogLevelError
	case "fatal":
		return LogLevelFatal
	}
	return LogLevelDebug
}
