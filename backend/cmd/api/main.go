package main

import (
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humaecho"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/labstack/echo/v4"
	"github.com/naohito-T/tinyurl/backend/configs"
	"github.com/naohito-T/tinyurl/backend/internal/infrastructure"
	"github.com/naohito-T/tinyurl/backend/internal/rest/middleware"
	"github.com/naohito-T/tinyurl/backend/internal/rest/router"
	appSchema "github.com/naohito-T/tinyurl/backend/schema/api"
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

// publicにわける
// public（誰でもアクセス可能）
// user（ログイン必須）
// private（管理者）

func main() {
	var publicAPI huma.API
	var privateAPI huma.API
	var c configs.AppEnvironment
	logger := infrastructure.NewLogger()

	cli := humacli.New(func(hooks humacli.Hooks, opts *Options) {
		e := echo.New()
		c = configs.NewAppEnvironment()

		middleware.CustomMiddleware(e, c)
		public := e.Group("/v1/public")
		private := e.Group("/v1/private")
		publicAPI = router.NewPublicRouter(humaecho.NewWithGroup(e, public, appSchema.NewHumaConfig()), logger)
		privateAPI = router.NewPublicRouter(humaecho.NewWithGroup(e, private, appSchema.NewHumaConfig()), logger)
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
			b, _ := publicAPI.OpenAPI().YAML()
			privateAPI.OpenAPI().YAML()
			fmt.Println(string(b))
		},
	})
	cli.Run()
}
