package error

// see: https://go.dev/play/p/TzZE1mdL63_1

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/naohito-T/tinyurl/backend/domain/customerror"
)

// これはGo言語における型アサーション（type assertion）の一例です。型アサーションは、インターフェースの値が特定の型を持っているかどうかをチェックし、その型の値を取り出すために使用されます。
// このコードの場合、err はエラーを表すインターフェース型です。*echo.HTTPError は echo パッケージの HTTPError 型のポインタです。この行は、err が *echo.HTTPError 型の値を持っているかをチェックし、もしそうであれば he にその値を割り当て、ok に true を設定します。err が *echo.HTTPError 型でなければ、ok は false になり、he は nil になります。
// この型アサーションを使うことで、安全に型変換を行いつつ、エラーチェックを同時に実施できます。
func CustomErrorHandler(err error, c echo.Context) {
	c.Logger().Error("カスタムエラーに入ったよ")
	// ロギング
	// c.Logger().Error(err)
	appErr := buildError(err, c)
	c.JSON(appErr.Status, map[string]string{"code": appErr.Code, "message": appErr.Message})
}

func buildError(err error, c echo.Context) *customerror.ApplicationError {
	appErr := customerror.ApplicationError{
		Status:  customerror.UnexpectedCode.Status,
		Code:    customerror.UnexpectedCode.Code,
		Message: customerror.UnexpectedCode.Message,
	}
	if errors.Is(err, customerror.WrongEmailVerificationErrorInstance) {
		c.Logger().Error("これがWrongEmailVerificationErrorアプリケーションエラー")
		// return &customerror.ApplicationError{}
	}
	return &appErr
}
