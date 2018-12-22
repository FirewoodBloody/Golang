package logger

import (
	"testing"
)

func TestNewFilelogger(t *testing.T) {
	logger := NewFilelogger(LogLevelDebug, "C:/", "test")
	logger.Debug("user id [%d] is come from china !", 121927)
	logger.Warn("then is warn!")
	logger.Fatal("then is Fatal!")
	logger.Close()
}

func TestNewConsoleLogger(t *testing.T) {
	logger := NewConsoleLogger(LogLevelDebug)
	logger.Debug("user id [%d] is come from china !", 121927)
	logger.Warn("then is warn!")
	logger.Fatal("then is Fatal!")
	logger.Close()
}
