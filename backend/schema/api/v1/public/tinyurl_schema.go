package public

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/naohito-T/tinyurl/backend/configs"
)

type tinyURLAPISchema struct {
	GET huma.Operation
	POST huma.Operation
}

var TinyURLAPISchema = tinyURLAPISchema{
	GET: huma.Operation{
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
			"500": {
				Description: "Failed to Get short URL",
				Content: map[string]*huma.MediaType{
					"text/plain": {
						Schema: &huma.Schema{
							Type: "string",
						},
					},
				},
			},
		},
	},
	POST: huma.Operation{
		OperationID: "create-tinyurl",
		Method:      http.MethodPost,
		Path:        configs.CreateShortURL,
		Summary:     "Create a short URL",
		Description: "Create a short URL.",
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
			"201": {
				Description: "Created short URL",
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
			"500": {
				Description: "Failed to create short URL",
				Content: map[string]*huma.MediaType{
					"text/plain": {
						Schema: &huma.Schema{
							Type: "string",
						},
					},
				},
			},
		},
	},
}
