package main

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// Routes
	e.GET("/", hello)
	// e.Startでエラーが発生した場合、Fatalでプログラムを終了する
	e.Logger.Fatal(e.Start(":6500"))
}

	func hello(c echo.Context) error {
	println("Hello, World!")
	// if err := isValid("hello"); err != nil {
	// 	return err
	// }
	isValid("hello")
	return c.String(http.StatusOK, "Hello, World!")
}

func isValid(txt string) error {
	if txt == "" {
		return errors.New("Invalid")
	}
	return nil
}
