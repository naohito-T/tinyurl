package security

import (
	"github.com/labstack/echo/v4"
)

func AttachSecurity(e *echo.Echo) {
	// サーバー起動時のバナーを非表示にする
	e.HideBanner = true
	e.HidePort = true
}
