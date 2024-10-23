package logger

import (
	"log/slog"
	"os"
)

type FileLogger struct {
	*slog.Logger
	file *os.File
}

func NewFLogger(filePath string) (*FileLogger, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	return &FileLogger{
		Logger: slog.New(slog.NewTextHandler(file, nil)),
		file:   file,
	}, nil
}

func (l *FileLogger) Info(msg string, keysAndValues ...interface{}) {
	l.Logger.Info(msg, keysAndValues...)
}

func (l *FileLogger) Error(msg string, keysAndValues ...interface{}) {
	l.Logger.Error(msg, keysAndValues...)
}

func (l *FileLogger) Close() {
	if l.file != nil {
		l.file.Close()
	}
}
