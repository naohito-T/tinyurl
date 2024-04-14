package router

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// NewRouter これもシングルトンにした場合の例が気になる
func NewRouter(e *echo.Echo) {
	e.GET("/health", health)
	e.GET("/hello", hello)
}

func hello(c echo.Context) error {
	println("Hello, World!")
	if err := isValid("hello"); err != nil {
		return err
	}
	return c.String(http.StatusOK, "Hello, World!")
}

func health(c echo.Context) error {
	println("Hello, World!")
	if err := isValid("hello"); err != nil {
		return err
	}
	return c.String(http.StatusOK, "Hello, World!")
}

func isValid(txt string) error {
	if txt == "" {
		return errors.New("Invalid")
	}
	return nil
}
