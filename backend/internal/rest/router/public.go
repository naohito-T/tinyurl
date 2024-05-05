package router

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/naohito-T/tinyurl/backend/configs"
	"github.com/naohito-T/tinyurl/backend/internal/infrastructures/slog"
	"github.com/naohito-T/tinyurl/backend/internal/rest/container"
	"github.com/naohito-T/tinyurl/backend/internal/rest/handler"
)

type HealthCheckQuery struct {
	CheckDB bool `query:"q" doc:"Optional database check parameter"`
}

type HealthCheckResponse struct {
	Body struct {
		Message string `json:"message"`
	}
}

type GetTinyURLQuery struct {
	ID string `path:"id" required:"true"`
}

type GetTinyURLResponse struct {
	Body struct {
		ID          string `json:"id" required:"true"`
		OriginalURL string `json:"original_url" required:"true"`
		CreatedAt   string `json:"created_at" required:"true"`
	}
}

type CreateTinyURLBody struct {
	Body struct {
		URL string `json:"url" required:"true" example:"http://example.com" doc:"URL to shorten"`
	}
}

type CreateTinyURLResponse struct {
	Body struct {
		ID string `json:"id"`
	}
}

// 今日の課題
// 1. tinyulrのAPIを作成する（できそう）
// 2. テストを書く
// 3. ドキュメントを書く

// https://tinyurl.com/app/api/url/create"
// NewRouter これもシングルトンにした場合の例が気になる
func NewPublicRouter(app huma.API) {
	h := handler.NewShortURLHandler(container.NewGuestContainer())

	// これ見ていつか修正する https://aws.amazon.com/jp/builders-library/implementing-health-checks/
	huma.Register(app, huma.Operation{
		OperationID: "health",
		Method:      http.MethodGet,
		Path:        configs.Health,
		Summary:     "Health Check",
		Description: "Check the health of the service.",
		Tags:        []string{"Public"},
	}, func(_ context.Context, input *struct {
		HealthCheckQuery
	}) (*HealthCheckResponse, error) {
		slog.NewLogger().Info("Health Check: %v", input.CheckDB)
		return &HealthCheckResponse{
			Body: struct {
				Message string `json:"message"`
			}{
				Message: "ok",
			},
		}, nil
	})

	huma.Register(app, huma.Operation{
		OperationID: "tinyurl",
		Method:      http.MethodGet,
		Path:        configs.GetShortURL,
		Summary:     "Get a original URL",
		Description: "Get a original URL.",
		Tags:        []string{"Public"},
	}, func(ctx context.Context, query *struct {
		GetTinyURLQuery
	}) (*GetTinyURLResponse, error) {
		resp := &GetTinyURLResponse{}
		shortURL, err := h.GetShortURLHandler(ctx, query.ID)
		if err != nil {
			return nil, err
		}
		resp.Body.ID = shortURL.ID
		resp.Body.OriginalURL = shortURL.OriginalURL
		resp.Body.CreatedAt = shortURL.CreatedAt
		return resp, nil
	})

	huma.Register(app, huma.Operation{
		OperationID: "tinyurl",
		Method:      http.MethodPost,
		Path:        configs.CreateShortURL,
		Summary:     "Create a short URL",
		Description: "Create a short URL.",
		Tags:        []string{"Public"},
	}, func(ctx context.Context, body *CreateTinyURLBody) (*CreateTinyURLResponse, error) {
		resp := &CreateTinyURLResponse{}
		shortURL, err := h.CreateShortURLHandler(ctx, body.Body.URL)
		if err != nil {
			return nil, err
		}
		resp.Body.ID = shortURL.ID
		return resp, nil
	})
}
