package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	Router "github.com/naohito-T/tinyurl/backend/configs"
	// "github.com/naohito-T/tinyurl/backend/internal/rest/container"
)

type HealthCheckParams struct {
	CheckDB string `json:"check_db"`
}

// type GreetingOutput struct {
// 	Greeting    string `json:"greeting"`
// 	Suffix      string `json:"suffix"`
// 	Length      int    `json:"length"`
// 	ContentType string `json:"content_type"`
// 	Num         int    `json:"num"`
// }

type HealthCheckParams2 struct {
	Body struct {
		// Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

// GreetingOutput represents the greeting operation response.
type GreetingOutput3 struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

// https://tinyurl.com/app/api/url/create"
// NewRouter これもシングルトンにした場合の例が気になる
func NewPublicRouter(app huma.API) {
	// container := container.NewGuestContainer()

	huma.Register(app, huma.Operation{
		OperationID: "health",
		Method:      http.MethodGet,
		Path:        Router.Health,
		Summary:     "Health Check",
		Description: "Check the health of the service.",
		Tags:        []string{"Greetings"},
	}, func(_ context.Context, _ *HealthCheckParams) (*HealthCheckParams2, error) {
		resp := &HealthCheckParams2{Body: struct {
			Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
		}{Message: "ok"}}
		return resp, nil
	})

	huma.Register(app, huma.Operation{
		OperationID: "get-greeting",
		Method:      http.MethodGet,
		Path:        "/greeting/{name}",
		Summary:     "Get a greeting",
		Description: "Get a greeting for a person by name.",
		Tags:        []string{"Greetings"},
	}, func(_ context.Context, input *struct {
		Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
	}) (*GreetingOutput3, error) {
		resp := &GreetingOutput3{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
		return resp, nil
	})

	// e.GET(Router.Health, health)
	// e.GET(Router.GetShortURL, hello)
	// e.POST(Router.CreateShortURL, hello)

}

// func hello(c echo.Context) error {
// 	// {"time":"2024-04-14T02:16:18.08145333Z","level":"INFO","msg":"Hello, World!"}
// 	slog.NewLogger().Info("Hello, World!")
// 	if err := isValid("hello"); err != nil {
// 		return err
// 	}
// 	return c.String(http.StatusOK, "Hello, World 2!")
// }

// func health(c echo.Context) error {
// 	println("Hello, World!")
// 	if err := isValid("hello"); err != nil {
// 		return err
// 	}
// 	return c.String(http.StatusOK, "Hello, World!")
// }

// func isValid(txt string) error {
// 	if txt == "" {
// 		return errors.New("Invalid")
// 	}
// 	return nil
// }
