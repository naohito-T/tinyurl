package main

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humaecho"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/labstack/echo/v4"
	"github.com/naohito-T/tinyurl/backend/configs"
	"github.com/naohito-T/tinyurl/backend/internal/rest/middleware"
	"github.com/naohito-T/tinyurl/backend/internal/rest/router"
)

const (
	defaultPort = ":6500"
)

// Options for the CLI. Pass `--port` or set the `SERVICE_PORT` env var.
type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8888"`
}

func main() {
	cli := humacli.New(func(hooks humacli.Hooks, _ *Options) {
		e := echo.New()
		configs.NewAppEnvironment()
		middleware.CustomMiddleware(e)
		// Create the API
		// e.Startでエラーが発生した場合、Fatalでプログラムを終了する
		// e.Logger.Fatal(e.Start(defaultPort))
		// api := humaecho.New(router.NewRouter(e), huma.DefaultConfig("My API", "1.0.0"))
		humaecho.New(router.NewRouter(e), huma.DefaultConfig("My API", "1.0.0"))

		// Tell the CLI how to start your router.
		hooks.OnStart(func() {
			e.Logger.Fatal(e.Start(defaultPort))
		})
	})
	cli.Run()
}
