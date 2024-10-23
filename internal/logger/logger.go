package logger

import "github.com/Helltale/vk-parser-program/config"

type LoggerManager interface {
	Info(msg string, keysAndValues ...interface{})
	Error(msg string, keysAndValues ...interface{})
}

type CombinedLogger struct {
	consoleLogger LoggerManager
	fileLogger    LoggerManager
}

func NewCombinedLogger(consoleLogger LoggerManager, fileLogger LoggerManager) *CombinedLogger {
	return &CombinedLogger{
		consoleLogger: consoleLogger,
		fileLogger:    fileLogger,
	}
}

func (l *CombinedLogger) Info(msg string, keysAndValues ...interface{}) {
	l.consoleLogger.Info(msg, keysAndValues...)
	l.fileLogger.Info(msg, keysAndValues...)
}

func (l *CombinedLogger) Error(msg string, keysAndValues ...interface{}) {
	l.consoleLogger.Error(msg, keysAndValues...)
	l.fileLogger.Error(msg, keysAndValues...)
}

func Init(conf *config.Config) (*CombinedLogger, error) {
	slogger := NewSLogger()

	fileLogger, err := NewFLogger(conf.AppLogfile)
	if err != nil {
		slogger.Error("Ошибка создания FileLogger", "error", err)
		return nil, err
	}
	defer fileLogger.Close()

	return NewCombinedLogger(slogger, fileLogger), nil
}
