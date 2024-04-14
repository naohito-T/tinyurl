package router

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/naohito-T/tinyurl/backend/internal/infrastructures/slog"
	// "github.com/naohito-T/tinyurl/backend/internal/rest/container"
)

// https://tinyurl.com/app/api/url/create"
// NewRouter これもシングルトンにした場合の例が気になる
func NewRouter(e *echo.Echo) {
	// container := container.NewGuestContainer()

	e.GET("/health", hello)
	e.GET("/api/v1/urls/:shortUrl", hello)
	e.POST("/api/v1/urls", hello)
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
