package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/naohito-T/tinyurl/backend/internal/infrastructure"
	"github.com/naohito-T/tinyurl/backend/internal/rest/container"
	"github.com/naohito-T/tinyurl/backend/internal/rest/handler"
	"github.com/naohito-T/tinyurl/backend/schema/api/v1/public"
	"github.com/naohito-T/tinyurl/backend/schema/request/body"
	"github.com/naohito-T/tinyurl/backend/schema/request/query"
	"github.com/naohito-T/tinyurl/backend/schema/response"
)

// 今日の課題
// 1. validationのエラーを返す
// 2. openapiのドキュメントを清書する
// 3. テストを書く
// 4. ドキュメントを書く
func NewPublicRouter(app huma.API, logger infrastructure.ILogger) huma.API {
	h := handler.NewShortURLHandler(container.NewPublicContainer(), infrastructure.NewLogger())

	huma.Register(app, *public.HealthAPISchema, func(_ context.Context, input *struct {
		query.HealthCheckQuery
	}) (*response.HealthCheckResponse, error) {
		logger.Info("Health Check:", "Health Check:", map[string]interface{}{
			"db": input.CheckDB,
		})
		return &response.HealthCheckResponse{
			Body: struct {
				Message string `json:"message"`
			}{
				Message: "ok",
			},
		}, nil
	})

	huma.Register(app, public.TinyURLAPISchema.GET, func(ctx context.Context, input *query.GetTinyURLQuery) (*response.GetTinyURLResponse, error) {
		fmt.Printf("GetInfoTinyURLQuery: %v", input.ID)
		logger.Info("parammm: %v", input.ID)
		shortURL, err := h.GetShortURLHandler(ctx, input.ID)
		logger.Info("Result err GetShortURLHandler: %v", err)
		logger.Info("Result GetShortURLHandler: %v", shortURL.OriginalURL)
		return &response.GetTinyURLResponse{
			Status: http.StatusTemporaryRedirect,
			URL:    shortURL.OriginalURL,
		}, nil
	})

	huma.Register(app, public.TinyURLAPISchema.POST, func(ctx context.Context, body *body.CreateTinyURLBody) (*response.CreateTinyURLResponse, error) {
		resp := &response.CreateTinyURLResponse{}
		shortURL, err := h.CreateShortURLHandler(ctx, body.Body.URL)
		if err != nil {
			return nil, err
		}
		resp.Body.ID = shortURL.ID
		return resp, nil
	})

	huma.Register(app, public.TinyURLInfoAPISchema.GET, func(ctx context.Context, query *query.GetInfoTinyURLQuery) (*response.GetInfoTinyURLResponse, error) {
		resp := &response.GetInfoTinyURLResponse{}
		shortURL, err := h.GetShortURLHandler(ctx, query.ID)
		if err != nil {
			return nil, err
		}
		resp.Body.ID = shortURL.ID
		resp.Body.OriginalURL = shortURL.OriginalURL
		resp.Body.CreatedAt = shortURL.CreatedAt
		return resp, nil
	})

	return app
}
