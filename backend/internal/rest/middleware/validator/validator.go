package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// CustomValidator カスタムバリデータの定義
type CustomValidator struct {
	validator *validator.Validate
}

// NewValidator はカスタムバリデータのインスタンスを作成します
func NewValidator() echo.Validator {
	return &CustomValidator{validator: validator.New()}
}

// Validate はインタフェースを受け取り、構造体の検証を行います
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
