package infrastructure

import (
	"log/slog"
	"os"
	"sync"
)

// var once sync.Once

// 効率: ロガーインスタンスが一度だけ作成されるため、メモリ使用量と初期化コストが削減されます。
// 簡潔性: getLogger() を使用することで、コード内のロガーへのアクセスが簡略化されます。
// スレッドセーフ: sync.Once を使用することで、マルチスレッド環境でもロガーが正しく一度だけ初期化されることが保証される。
// GetLogger は、ロガーのインスタンスを取得します。
// func GetLogger() *Logger {
// 	once.Do(func() {
// 		singletonLogger = newLogger()
// 	})
// 	return singletonLogger
// }

type ILabelLogger interface {
	Debug(message string, Detail map[string]interface{})
	Info(message string, Detail map[string]interface{})
	Warn(message string, Detail map[string]interface{}, Error map[string]interface{})
	Error(message string, Detail map[string]interface{}, Error map[string]interface{})
	// Fatal(format string, v ...any)
}

type ILogger interface {
	Debug(kind string, message string, Detail map[string]interface{})
	Info(kind string, message string, Detail map[string]interface{})
	Warn(kind string, message string, Detail map[string]interface{}, Error map[string]interface{})
	Error(kind string, message string, Detail map[string]interface{}, Error map[string]interface{})
	// Fatal(format string, v ...any)
}

type Logger struct {
	logger *slog.Logger
}

var newOnceLogger = sync.OnceValue(func() *Logger {
	// これは1.22.0で追加されたもの
	slog.SetLogLoggerLevel(slog.LevelDebug)
	return &Logger{
		logger: slog.New(slog.NewJSONHandler(os.Stdout, nil)),
	}
})

func NewLogger() *Logger {
	return newOnceLogger()
}

func NewLabelLogger() ILabelLogger {
	return newOnceLogger()
}

func (l *Logger) Debug(kind string, message string, detail map[string]interface{}) {
	l.logger.Debug(kind, message, detail)
}

func (l *Logger) Info(kind string, message string, detail map[string]interface{}) {
	l.logger.Info(kind, message, detail)
}

func (l *Logger) Warn(kind string, message string, detail map[string]interface{}, error map[string]interface{}) {
	l.logger.Warn(kind, message, detail, error)
}

func (l *Logger) Error(kind string, message string, detail map[string]interface{}, error map[string]interface{}) {
	l.logger.Error(kind, message, detail, error)
}

// golangではFatalは使わない（os.Exit(1)を勝手に使われるため）
// func (l *Logger) Fatal(format string, v ...any) {
// 	l.logger.Fatal(format, v...)
// 	os.Exit(1)
// }
