package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/naohito-T/tinyurl/backend/configs"
	"github.com/naohito-T/tinyurl/backend/domain"
	"github.com/naohito-T/tinyurl/backend/internal/rest/container"
	"github.com/naohito-T/tinyurl/backend/internal/rest/handler"
)

// huma.Register(app, huma.Operation{
// 	OperationID: "omittable",
// 	Method:      http.MethodPost,
// 	Path:        "/omittable",
// 	Summary:     "Omittable / nullable example",
// }, func(ctx context.Context, input *struct {
// 	Body *struct {
// 		Name OmittableNullable[string] `json:"name,omitempty" maxLength:"10"`
// 	}
// }) (*MyResponse, error) {
// 	resp := &MyResponse{}
// 	if input.Body == nil {
// 		resp.Body.Message = "Body was not sent"
// 	} else if !input.Body.Name.Sent {
// 		resp.Body.Message = "Name was omitted from the request"
// 	} else if input.Body.Name.Null {
// 		resp.Body.Message = "Name was set to null"
// 	} else {
// 		resp.Body.Message = "Name was set to: " + input.Body.Name.Value
// 	}
// 	return resp, nil
// })

type FilterOrQuery struct {
	FilterID int64  `query:"filter_id" doc:"filter_id and query are mutually exclusive"`
	Query    string `query:"query" doc:"filter_id and query are mutually exclusive"`
}

func (f *FilterOrQuery) Resolve(ctx huma.Context) []error {
	if f.FilterID != 0 && f.Query != "" {
		return []error{&huma.ErrorDetail{
			Message:  "Cannot pass both filter_id and query at the same time",
			Location: "query",
			Value:    fmt.Sprintf("filter_id:%d query:%s", f.FilterID, f.Query),
		}}
	}
	return nil
}

// type HealthCheckParams struct {
// 	CheckDB string `json:"check_db"`
// }

// GreetingOutput represents the greeting operation response.
type CreateTinyURLBody struct {
	Body struct {
		URL string `json:"url" example:"http://example.com" doc:"URL to shorten"`
	}
}

type HealthCheckQuery struct {
	Body struct {
		Message string `json:"message,omitempty" example:"Hello, world!" doc:"Greeting message"`
	}
}

// type HealthCheckResponse struct {
// 	Body struct {
// 		Message string `json:"message,omitempty" example:"Hello, world!" doc:"Greeting message"`
// 	}
// }

// HealthCheckParams はヘルスチェックのリクエストパラメータを定義します。
type HealthCheckParams struct {
	CheckDB *string `query:"check_db" doc:"Optional database check parameter"`
}

// HealthCheckResponse はヘルスチェックのレスポンスを定義します。
type HealthCheckResponse struct {
	Message string `json:"message"`
}

// type OmittableNullable[T any] struct {
// 	Sent  bool
// 	Null  bool
// 	Value T
// }

// UnmarshalJSON unmarshals this value from JSON input.
// func (o *OmittableNullable[T]) UnmarshalJSON(b []byte) error {
// 	if len(b) > 0 {
// 		o.Sent = true
// 		if bytes.Equal(b, []byte("null")) {
// 			o.Null = true
// 			return nil
// 		}
// 		return json.Unmarshal(b, &o.Value)
// 	}
// 	return nil
// }

// Schema returns a schema representing this value on the wire.
// It returns the schema of the contained type.
// func (o OmittableNullable[T]) Schema(r huma.Registry) *huma.Schema {
// 	return r.Schema(reflect.TypeOf(o.Value), true, "")
// }

// type MyResponse struct {
// 	Body struct {
// 		Message string `json:"message"`
// 	}
// }

// 今日の課題
// 1. tinyulrのAPIを作成する
// 2. テストを書く
// 3. ドキュメントを書く

// https://tinyurl.com/app/api/url/create"
// NewRouter これもシングルトンにした場合の例が気になる
func NewPublicRouter(app huma.API) {
	h := handler.NewShortURLHandler(container.NewGuestContainer())
	huma.Register(app, huma.Operation{
		OperationID: "health",
		Method:      http.MethodGet,
		Path:        configs.Health,
		Summary:     "Health Check",
		Description: "Check the health of the service.",
		Tags:        []string{"Public"},
	}, func(_ context.Context, input *struct {
		FilterOrQuery
	}) (*struct{}, error) {
		fmt.Printf("Got filter_id:%d query:%s\n", input.FilterID, input.Query)
		return nil, nil
	})

	// オリジナルのURLは必須
	// ここではすぐにhandlerへ渡す

	huma.Register(app, huma.Operation{
		OperationID: "tinyurl",
		Method:      http.MethodPost,
		Path:        configs.CreateShortURL,
		Summary:     "Create a short URL",
		Description: "Create a short URL.",
		Tags:        []string{"Public"},
	}, func(ctx context.Context, body *CreateTinyURLBody) (*domain.ShortURL, error) {
		shortURL, err := h.CreateShortURLHandler(ctx, body.Body.URL)
		if err != nil {
			return nil, err
		}
		return &shortURL, nil
	})

	// e.GET(Router.GetShortURL, hello)
	// e.POST(Router.CreateShortURL, hello)
}
