package handler

import "github.com/labstack/echo/v4"

type Handler interface {
	HealthHandler(c echo.Context) error
}
