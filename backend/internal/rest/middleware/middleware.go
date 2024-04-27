package middleware

// see: https://echo.labstack.com/docs/category/middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	// "github.com/naohito-T/tinyurl/backend/internal/rest/middleware/accesslog"
	ehandler "github.com/naohito-T/tinyurl/backend/internal/rest/middleware/error"
	"github.com/naohito-T/tinyurl/backend/internal/rest/middleware/validator"
)

// loggerの考え方
// https://yuya-hirooka.hatenablog.com/entry/2021/10/15/123607
func CustomMiddleware(e *echo.Echo) {
	// Loggerの設定変更
	e.Logger.SetLevel(log.DEBUG) // すべてのログレベルを出力する

	// ミドルウェアとルートの設定
	// e.Use(middleware.Logger()) // ロギングミドルウェアを使う
	e.Validator = validator.NewValidator()
	// e.Use(accesslog.AccessLog())
	// expect this handler is used as fallback unless a more specific is present
	e.Use(middleware.Recover())
	// これでechoのloggerを操作できる。
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	//     Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	// }))
	e.HTTPErrorHandler = ehandler.CustomErrorHandler
}
