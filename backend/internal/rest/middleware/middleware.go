package middleware

// see: https://echo.labstack.com/docs/category/middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	// "github.com/naohito-T/tinyurl/backend/internal/rest/middleware/accesslog"
	"github.com/naohito-T/tinyurl/backend/configs"
	"github.com/naohito-T/tinyurl/backend/internal/rest/controller"
	"github.com/naohito-T/tinyurl/backend/internal/rest/middleware/validator"
)

// loggerの考え方
// https://yuya-hirooka.hatenablog.com/entry/2021/10/15/123607
func CustomMiddleware(e *echo.Echo, c *configs.AppEnvironment) {
	// echo.Loggerの設定変更
	if c.IsLocal() {
		e.Logger.SetLevel(log.DEBUG)
	} else {
		e.Logger.SetLevel(log.INFO)
	}
	// ミドルウェアとルートの設定
	// e.Use(middleware.Logger()) // ロギングミドルウェアを使う
	e.Validator = validator.NewValidator()
	// e.Use(accesslog.AccessLog())
	e.Use(middleware.Recover())
	// これでechoのloggerを操作できる。
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	//     Format: "time=${time_rfc3339_nano}, method=${method}, uri=${uri}, status=${status}\n",
	// }))
	e.HTTPErrorHandler = controller.CustomErrorHandler
}
