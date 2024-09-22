package controller

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

	"github.com/aws/smithy-go"
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
	var ve validator.ValidationErrors
	var oe *smithy.OperationError
	appErr := customerror.ApplicationError{
		Status:  customerror.UnexpectedCode.Status,
		Code:    customerror.UnexpectedCode.Code,
		Message: customerror.UnexpectedCode.Message,
	}

	if errors.As(err, &ve) {
		c.Logger().Error("errors.As():Failed at バリデーションエラー発生")
		return &customerror.ApplicationError{
			Status:  http.StatusBadRequest,
			Code:    "ValidationError",
			Message: "Input validation failed",
		}
	}

	// aws-sdk-go-v2のエラー処理
	if errors.As(err, &oe) {
		log.Printf("failed to call service: %s, operation: %s, error: %v", oe.Service(), oe.Operation(), oe.Unwrap())
	}

	if errors.Is(err, customerror.WrongEmailVerificationErrorInstance) {
		c.Logger().Error("これがWrongEmailVerificationErrorアプリケーションエラー")
	}

	return &appErr
}
