package error

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CustomErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if !ok {
		he = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	// ロギング
	c.Logger().Error(err)

	// エラータイプによって条件分岐
	if he.Code == http.StatusNotFound {
		// リダイレクトを行う例
		if err := c.Redirect(http.StatusTemporaryRedirect, "/some-fallback-url"); err != nil {
			c.Logger().Error(err)
		}
	} else {
		// エラーレスポンスを送信する例
		if err := c.JSON(he.Code, map[string]string{"message": he.Message.(string)}); err != nil {
			c.Logger().Error(err)
		}
	}
}
