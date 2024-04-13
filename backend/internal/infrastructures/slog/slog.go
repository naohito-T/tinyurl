package slog

import (
	"log/slog"
	"os"
)

type ILogger interface {
	Debug(format string, v ...any)
	Info(format string, v ...any)
	Warn(format string, v ...any)
	Error(format string, v ...any)
	// Fatal(format string, v ...any)
}

type Logger struct {
	logger ILogger
}

func NewLogger() *Logger {
	// これは1.22.0で追加されたもの
	// slog.SetLogLoggerLevel(slog.LevelDebug)
	return &Logger{
		logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
}

func (l *Logger) Debug(format string, v ...any) {
	l.logger.Debug(format, v...)
}

func (l *Logger) Info(format string, v ...any) {
	l.logger.Info(format, v...)
}

func (l *Logger) Warn(format string, v ...any) {
	l.logger.Warn(format, v...)
}

func (l *Logger) Error(format string, v ...any) {
	l.logger.Error(format, v...)
}

// func (l *Logger) Fatal(format string, v ...any) {
// 	l.logger.Fatal(format, v...)
// 	os.Exit(1)
// }
