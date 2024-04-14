package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/naohito-T/tinyurl/backend/internal/rest/middleware/accesslog"
	"github.com/naohito-T/tinyurl/backend/internal/rest/middleware/error"
)

func CustomMiddleware(e *echo.Echo) {
	e.Use(accesslog.AccessLog())
	e.HTTPErrorHandler = error.CustomErrorHandler
	e.Use(middleware.Recover())
}
