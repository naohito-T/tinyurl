package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naohito-T/tinyurl/backend/domain/customerror"
)

func HealthHandler(c echo.Context) error {
	a := true
	if a {
		return customerror.WrongEmailVerificationErrorInstance
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "ok2"})
}
