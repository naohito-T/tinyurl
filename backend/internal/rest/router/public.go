package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/naohito-T/tinyurl/backend/configs"
	"github.com/naohito-T/tinyurl/backend/internal/infrastructure"
	"github.com/naohito-T/tinyurl/backend/internal/rest/container"
	"github.com/naohito-T/tinyurl/backend/internal/rest/handler"
	"github.com/naohito-T/tinyurl/backend/schema/api"
)

type HealthCheckQuery struct {
	CheckDB bool `query:"q" doc:"Optional DynamoDB check parameter"`
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
	Status int
	Url    string `header:"Location"`
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

type GetInfoTinyURLQuery struct {
	ID string `path:"id" required:"true"`
}

type GetInfoTinyURLResponse struct {
	Body struct {
		ID          string `json:"id" required:"true"`
		OriginalURL string `json:"original_url" required:"true"`
		CreatedAt   string `json:"created_at" required:"true"`
	}
}

// 今日の課題
// 1. validationのエラーを返す
// 2. openapiのドキュメントを清書する
// 3. テストを書く
// 4. ドキュメントを書く

// https://tinyurl.com/app/api/url/create"
// NewRouter これもシングルトンにした場合の例が気になる
func NewPublicRouter(app huma.API, logger infrastructure.ILogger) huma.API {
	h := handler.NewShortURLHandler(container.NewGuestContainer(), infrastructure.NewLogger())

	// これ見ていつか修正する https://aws.amazon.com/jp/builders-library/implementing-health-checks/
	// dynamoDBのヘルスチェックはない（SELECT 1 とかできない）
	// publicに開放しているapiのため、レートリミットとかの縛りは必要。
	huma.Register(app, *api.GetHealthAPISchema(), func(_ context.Context, input *struct {
		HealthCheckQuery
	}) (*HealthCheckResponse, error) {
		logger.Info("Health Check:", input.CheckDB)
		return &HealthCheckResponse{
			Body: struct {
				Message string `json:"message"`
			}{
				Message: "ok",
			},
		}, nil
	})

	huma.Register(app, huma.Operation{
		OperationID: "get-tinyurl-with-redirect",
		Method:      http.MethodGet,
		Path:        configs.GetShortURL,
		Summary:     "Redirect to original URL",
		Tags:        []string{"Public"},
		Parameters: []*huma.Param{
			{
				Name:        "id",
				In:          "path",
				Description: "ID of the short URL",
				Required:    true,
				Schema: &huma.Schema{
					Type: "string",
				},
			},
		},
		Responses: map[string]*huma.Response{
			"301": {
				Description: "Redirect to original URL",
				Headers: map[string]*huma.Header{
					"Location": {
						Description: "Location of the original URL",
						Schema: &huma.Schema{
							Type:   "string",
							Format: "uri",
						},
					},
				},
			},
			"404": {
				Description: "Short URL not found",
				Content: map[string]*huma.MediaType{
					"text/plain": {
						Schema: &huma.Schema{
							Type: "string",
						},
					},
				},
			},
		},
	}, func(ctx context.Context, input *GetInfoTinyURLQuery) (*GetTinyURLResponse, error) {
		fmt.Printf("GetInfoTinyURLQuery: %v", input.ID)
		logger.Info("parammm: %v", input.ID)
		shortURL, err := h.GetShortURLHandler(ctx, input.ID)
		logger.Info("Result err GetShortURLHandler: %v", err)              // null
		logger.Info("Result GetShortURLHandler: %v", shortURL.OriginalURL) // https://example.com/
		return &GetTinyURLResponse{
			Status: http.StatusTemporaryRedirect,
			Url:    shortURL.OriginalURL,
		}, nil
	})

	huma.Register(app, huma.Operation{
		OperationID: "create-tinyurl",
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

	huma.Register(app, huma.Operation{
		OperationID: "info-tinyurl",
		Method:      http.MethodGet,
		Path:        configs.GetOnlyShortURL,
		Summary:     "Get Info tinyurl",
		Description: "Get Info tinyurl",
		Tags:        []string{"Public"},
	}, func(ctx context.Context, query *struct {
		GetInfoTinyURLQuery
	}) (*GetInfoTinyURLResponse, error) {
		resp := &GetInfoTinyURLResponse{}
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
