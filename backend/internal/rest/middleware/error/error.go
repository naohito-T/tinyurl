package error

// see: https://go.dev/play/p/TzZE1mdL63_1

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naohito-T/tinyurl/backend/domain/customerror"
)

// これはGo言語における型アサーション（type assertion）の一例です。型アサーションは、インターフェースの値が特定の型を持っているかどうかをチェックし、その型の値を取り出すために使用されます。
// このコードの場合、err はエラーを表すインターフェース型です。*echo.HTTPError は echo パッケージの HTTPError 型のポインタです。この行は、err が *echo.HTTPError 型の値を持っているかをチェックし、もしそうであれば he にその値を割り当て、ok に true を設定します。err が *echo.HTTPError 型でなければ、ok は false になり、he は nil になります。
// この型アサーションを使うことで、安全に型変換を行いつつ、エラーチェックを同時に実施できます。
func CustomErrorHandler(err error, c echo.Context) {
	c.Logger().Error("カスタムエラーに入ったよ")
	if errors.Is(err, customerror.WrongEmailVerificationErrorInstance) {
		c.Logger().Error("これがWrongEmailVerificationErrorアプリケーションエラー")
	}
	he, ok := err.(*echo.HTTPError)
	if !ok {
		he = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	// ロギング
	c.Logger().Error(err)

	// エラーレスポンスを送信する例
	if err := c.JSON(he.Code, map[string]string{"message": he.Message.(string)}); err != nil {
		c.Logger().Error(err)
	}
}

// func buildError(err error, code int) *echo.HTTPError {

// }
