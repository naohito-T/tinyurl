package main

import (
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humaecho"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/labstack/echo/v4"
	"github.com/naohito-T/tinyurl/backend/configs"
	"github.com/naohito-T/tinyurl/backend/internal/rest/middleware"
	"github.com/naohito-T/tinyurl/backend/internal/rest/router"
	"github.com/spf13/cobra"
)

const (
	defaultPort = ":6500"
)

// ここはCLIからの引数を受け取るための構造体
type Options struct {
	Debug bool   `doc:"Enable debug logging"`
	Host  string `doc:"Hostname to listen on."`
	Port  int    `doc:"Port to listen on." short:"p" default:"8888"`
}

// /api/v1/openapi.yaml
// initHuma: humaのconfigを初期化
func initHuma() huma.Config {
	config := huma.DefaultConfig(configs.OpenAPITitle, configs.OpenAPIVersion)
	config.Servers = []*huma.Server{
		{URL: configs.OpenAPIDocServerPath},
	}

	config.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"bearer": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}
	config.DocsPath = "/docs"
	return config
}

// publicにわける
// user（ログイン必須）
// private（管理者）

func main() {
	var api huma.API

	cli := humacli.New(func(hooks humacli.Hooks, opts *Options) {
		fmt.Printf("Options are debug:%v host:%v port%v\n", opts.Debug, opts.Host, opts.Port)

		e := echo.New()
		// configを初期化
		configs.NewAppEnvironment()
		// ミドルウェアを適用（すべてのリクエストに対して）
		middleware.CustomMiddleware(e)
		// これgroup化したやつをnewUserRouterに渡す必要かも
		api = humaecho.NewWithGroup(e, e.Group("/api/v1"), initHuma())
		router.NewPublicRouter(api)

		// 未定義のルート用のキャッチオールハンドラ
		e.Any("/*", func(c echo.Context) error {
			return c.JSON(http.StatusNotFound, map[string]string{"message": "route_not_found"})
		})

		hooks.OnStart(func() {
			e.Logger.Fatal(e.Start(defaultPort))
		})
	})

	cli.Root().AddCommand(&cobra.Command{
		Use:   "openapi",
		Short: "Print the OpenAPI spec",
		Run: func(_ *cobra.Command, _ []string) {
			b, _ := api.OpenAPI().YAML()
			fmt.Println(string(b))
		},
	})
	cli.Run()
}
