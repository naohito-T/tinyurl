package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naohito-T/tinyurl/backend/domain/customerror"
)

type HealthCheckParams struct {
	// CheckDB *string `query:"check_db" validate:"required"`
	CheckDB *string `query:"check_db"`
}

func HealthHandler(c echo.Context) error {
	h := new(HealthCheckParams)
	c.Logger().Error("Binding...")
	if err := c.Bind(h); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	c.Logger().Error("Validating...")
	if err := c.Validate(h); err != nil {
		return &customerror.ValidationError{Message: "this is wrapped", Err: err}
	}

	if h.CheckDB != nil {
		return c.JSON(http.StatusOK, map[string]string{"message": "check_db"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "ok3"})
}
