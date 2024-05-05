package handler

// see: https://go.dev/play/p/TzZE1mdL63_1
// errors.Is()とerrors.As()は、
// errors.Is(err, target)
// ラップされたエラーでも、targetとなるエラーと一致するかどうか、値として判定したい時。
// errors.As(err, target)
// ラップされたエラーでも、targetとなるエラーに代入可能かどうか、型として判定したい時。

import (
	"errors"
	"log"
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
	if err := c.JSON(appErr.Status, map[string]string{"code": appErr.Code, "message": appErr.Message}); err != nil {
		// handle error, e.g., log it or return it
		log.Println("JSON response error:", err)
	}
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
	}
	if errors.Is(err, customerror.WrongEmailVerificationErrorInstance) {
		c.Logger().Error("これがWrongEmailVerificationErrorアプリケーションエラー")
	}

	return &appErr
}
