package public

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/naohito-T/tinyurl/backend/configs"
)

var HealthAPISchema = &huma.Operation{
		OperationID: "health-check",
		Method:      http.MethodGet,
		Path:        configs.Health,
		Summary:     "Health Check",
		Description: "Check the health of the service.",
		Tags:        []string{"Public"},
		Responses: map[string]*huma.Response{
			"200": {
				Description: "Health check successful",
				Content: map[string]*huma.MediaType{
					"application/json": {
						Schema: &huma.Schema{
							Type: "object",
							Properties: map[string]*huma.Schema{
								"message": {
									Type: "string",
								},
							},
						},
						Example: "{message: ok}",
					},
				},
			},
			"503": {
				Description: "Service unavailable",
				Content: map[string]*huma.MediaType{
					"application/problem+json": {
						Schema: &huma.Schema{
							Type: "object",
							Properties: map[string]*huma.Schema{
								"type": {
									Type:   "string",
									Format: "uri",
								},
								"title": {
									Type: "string",
								},
								"status": {
									Type: "integer",
								},
								"detail": {
									Type: "string",
								},
							},
						},
					},
				},
			},
		},
	}
