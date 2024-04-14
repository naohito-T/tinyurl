package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naohito-T/tinyurl/backend/internal/rest/middleware"
)

const (
	defaultPort = "6500"
)

func main() {
	// Echo instance
	e := echo.New()
	middleware.CustomMiddleware(e)
	// Routes
	e.GET("/", hello)
	// e.Startでエラーが発生した場合、Fatalでプログラムを終了する
	e.Logger.Fatal(e.Start(defaultPort))
}

func hello(c echo.Context) error {
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
