package router

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	Router "github.com/naohito-T/tinyurl/backend/configs"
	"github.com/naohito-T/tinyurl/backend/internal/infrastructures/slog"
	"github.com/naohito-T/tinyurl/backend/internal/rest/handler"
	// "github.com/naohito-T/tinyurl/backend/internal/rest/container"
)

// https://tinyurl.com/app/api/url/create"
// NewRouter これもシングルトンにした場合の例が気になる
func NewRouter(e *echo.Echo) *echo.Echo {
	// container := container.NewGuestContainer()
	e.GET(Router.Health, handler.HealthHandler)
	e.GET(Router.GetShortURL, hello)
	e.POST(Router.CreateShortURL, hello)
	// 未定義のルート用のキャッチオールハンドラ
	e.Any("/*", func(c echo.Context) error {
		return c.JSON(http.StatusNotFound, map[string]string{"message": "route_not_found"})
	})

	return e
}

func hello(c echo.Context) error {
	// {"time":"2024-04-14T02:16:18.08145333Z","level":"INFO","msg":"Hello, World!"}
	slog.NewLogger().Info("Hello, World!")
	if err := isValid("hello"); err != nil {
		return err
	}
	return c.String(http.StatusOK, "Hello, World 2!")
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
