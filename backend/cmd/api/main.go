package main

import (
	"github.com/labstack/echo/v4"
	"github.com/naohito-T/tinyurl/backend/internal/rest/middleware"
	"github.com/naohito-T/tinyurl/backend/internal/rest/router"
)

const (
	defaultPort = ":6500"
)

func main() {
	// Echo instance
	e := echo.New()
	middleware.CustomMiddleware(e)
	router.NewRouter(e)
	// e.Startでエラーが発生した場合、Fatalでプログラムを終了する
	e.Logger.Fatal(e.Start(defaultPort))
}
