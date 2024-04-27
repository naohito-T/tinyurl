package error

// see: https://go.dev/play/p/TzZE1mdL63_1

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/naohito-T/tinyurl/backend/domain/customerror"
)

func CustomErrorHandler(err error, c echo.Context) {
	c.Logger().Error("カスタムエラーに入ったよ")
	// ロギング
	// c.Logger().Error(err)
	appErr := buildError(err, c)
	c.JSON(appErr.Status, map[string]string{"code": appErr.Code, "message": appErr.Message})
}

func buildError(err error, c echo.Context) *customerror.ApplicationError {
	c.Logger().Error("ビルドエラー実施中")
	appErr := customerror.ApplicationError{
		Status:  customerror.UnexpectedCode.Status,
		Code:    customerror.UnexpectedCode.Code,
		Message: customerror.UnexpectedCode.Message,
	}
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		c.Logger().Error("errors.As():Failed at バリデーションエラー発生")
		return &customerror.ApplicationError{
			Status:  http.StatusBadRequest,
			Code:    "ValidationError",
			Message: "Input validation failed",
		}
	} else {
		c.Logger().Error("判定なし。")
	}

	if errors.Is(err, customerror.WrongEmailVerificationErrorInstance) {
		c.Logger().Error("これがWrongEmailVerificationErrorアプリケーションエラー")
	}

	return &appErr
}
