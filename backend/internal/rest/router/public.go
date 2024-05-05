package router

import (
	"context"
	"fmt"
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
	Status int
	Url    string `header:"Location"`
}

// これはcore routerにアクセスしたい場合の例
// func (m *GetTinyURLQuery) Resolve(ctx huma.Context) []error {
// 	h := handler.NewShortURLHandler(container.NewGuestContainer())
// 	id := ctx.Param("id")
// 	slog.NewLogger().Info("parammm: %v", id)
// 	shortURL, err := h.GetShortURLHandler(ctx.Context(), id)
// 	if err != nil {
// 		slog.NewLogger().Info("Error retrieving URL: %v", err)
// 		ctx.SetStatus(http.StatusNotFound)
// 		ctx.BodyWriter().Write([]byte("Short URL not found"))
// 		return []error{err} // ここで処理を終了させる
// 	}

// 	// エラーがなければリダイレクト処理を行う
// 	redirectURL := shortURL.OriginalURL
// 	ctx.SetStatus(http.StatusMovedPermanently)
// 	ctx.SetHeader("Location", redirectURL)
// 	// 正常なリダイレクト処理後は、ここで処理を終了
// 	return nil
// }

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
// 1. tinyulrのAPIを作成する（できそう）
// 2. テストを書く
// 3. ドキュメントを書く

// https://tinyurl.com/app/api/url/create"
// NewRouter これもシングルトンにした場合の例が気になる
func NewPublicRouter(app huma.API) huma.API {
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
		slog.NewLogger().Info("parammm: %v", input.ID)
		shortURL, err := h.GetShortURLHandler(ctx, input.ID)
		slog.NewLogger().Info("Result err GetShortURLHandler: %v", err)              // null
		slog.NewLogger().Info("Result GetShortURLHandler: %v", shortURL.OriginalURL) // https://example.com/
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
