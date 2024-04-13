package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/naohito-T/tinyurl/backend/internal/rest/middleware/accesslog"
)

func Middleware(e *echo.Echo) {
	e.Use(accesslog.AccessLog())
}
