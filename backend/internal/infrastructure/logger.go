package infrastructure

import (
	"log/slog"
	"os"
	"sync"
)

type ILogger interface {
	Debug(kind string, message string, detail map[string]interface{})
	Info(kind string, message string, detail map[string]interface{})
	Warn(kind string, message string, detail map[string]interface{}, errorDetail map[string]interface{})
	Error(kind string, message string, detail map[string]interface{}, errorDetail map[string]interface{})
	CustomLabelLoggers(kinds defaultKinds) ILabelLogger
}

type ILabelLogger interface {
	Debug(message string, detail map[string]interface{})
	Info(message string, detail map[string]interface{})
	Warn(message string, detail map[string]interface{}, errorDetail map[string]interface{})
	Error(message string, detail map[string]interface{}, errorDetail map[string]interface{})
}

// Next：構造体でkeyを型定義する。
// 構造体でkeyを渡されなかったらゼロ値で設定される（文字列のため、空文字列）
type defaultKinds struct {
	DebugKind string
	InfoKind  string
	WarnKind  string
	ErrorKind string
}

type labeledLogger struct {
	logger *logger
	kinds  defaultKinds
}

type logger struct {
	rawLogger *slog.Logger
}

var newOnceLogger = sync.OnceValue(func() *logger {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	return &logger{
		rawLogger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
})

func (l *logger) Debug(kind string, message string, detail map[string]interface{}) {
	l.rawLogger.Debug(message, slog.String("kind", kind), slog.Any("detail", detail))
}

func (l *logger) Info(kind string, message string, detail map[string]interface{}) {
	l.rawLogger.Info(message, slog.String("kind", kind), slog.Any("detail", detail))
}

func (l *logger) Warn(kind string, message string, detail map[string]interface{}, errorDetail map[string]interface{}) {
	l.rawLogger.Warn(message, slog.String("kind", kind), slog.Any("detail", detail), slog.Any("error", errorDetail))
}

func (l *logger) Error(kind string, message string, detail map[string]interface{}, errorDetail map[string]interface{}) {
	l.rawLogger.Error(message, slog.String("kind", kind), slog.Any("detail", detail), slog.Any("error", errorDetail))
}

func (l *logger) CustomLabelLoggers(kinds defaultKinds) ILabelLogger {
	if kinds.DebugKind == "" {
		kinds.DebugKind = "debug"
	}
	if kinds.InfoKind == "" {
		kinds.InfoKind = "info"
	}
	if kinds.WarnKind == "" {
		kinds.WarnKind = "warn"
	}
	if kinds.ErrorKind == "" {
		kinds.ErrorKind = "error"
	}

	return &labeledLogger{
		logger: l,
		kinds:  kinds,
	}
}

// これは小文字でもいいかも
func (l *labeledLogger) Debug(message string, detail map[string]interface{}) {
	l.logger.Debug(l.kinds.DebugKind, message, detail)
}

func (l *labeledLogger) Info(message string, detail map[string]interface{}) {
	l.logger.Info(l.kinds.InfoKind, message, detail)
}

func (l *labeledLogger) Warn(message string, detail map[string]interface{}, errorDetail map[string]interface{}) {
	l.logger.Warn(l.kinds.WarnKind, message, detail, errorDetail)
}

func (l *labeledLogger) Error(message string, detail map[string]interface{}, errorDetail map[string]interface{}) {
	l.logger.Error(l.kinds.ErrorKind, message, detail, errorDetail)
}

func NewLogger() ILogger {
	return newOnceLogger()
}

var ErrorLogger = NewLogger().CustomLabelLoggers(defaultKinds{
	WarnKind:  "known_error",
	ErrorKind: "application_error",
})

var RouterLogger = NewLogger().CustomLabelLoggers(defaultKinds{
	DebugKind: "router",
	InfoKind:  "router",
	WarnKind:  "router",
	ErrorKind: "router",
})

var ControllerLogger = NewLogger().CustomLabelLoggers(defaultKinds{
	DebugKind: "controller",
	InfoKind:  "controller",
	WarnKind:  "controller",
	ErrorKind: "controller",
})

var UsecaseLogger = NewLogger().CustomLabelLoggers(defaultKinds{
	DebugKind: "usecase",
	InfoKind:  "usecase",
	WarnKind:  "usecase",
	ErrorKind: "usecase",
})

var RepositoryLogger = NewLogger().CustomLabelLoggers(defaultKinds{
	DebugKind: "repository",
	InfoKind:  "repository",
	WarnKind:  "repository",
	ErrorKind: "repository",
})

var InfrastructureLogger = NewLogger().CustomLabelLoggers(defaultKinds{
	DebugKind: "infrastructure",
	InfoKind:  "infrastructure",
	WarnKind:  "infrastructure",
	ErrorKind: "infrastructure",
})
