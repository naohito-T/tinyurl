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

// golangではFatalは使わない（os.Exit(1)を勝手に使われるため）
// func (l *Logger) Fatal(format string, v ...any) {
// 	l.logger.Fatal(format, v...)
// 	os.Exit(1)
// }
