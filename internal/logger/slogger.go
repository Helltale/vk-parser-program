package logger

import (
	"log/slog"
	"os"
)

type Slogger struct {
	*slog.Logger
}

func NewSLogger() *Slogger {
	return &Slogger{
		slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}
}

func (l *Slogger) Info(msg string, keysAndValues ...interface{}) {
	l.Logger.Info(msg, keysAndValues...)
}

func (l *Slogger) Error(msg string, keysAndValues ...interface{}) {
	l.Logger.Error(msg, keysAndValues...)
}
