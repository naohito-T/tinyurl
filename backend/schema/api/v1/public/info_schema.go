package public

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/naohito-T/tinyurl/backend/configs"
)

type tinyURLInfoAPISchema struct {
	GET huma.Operation
}

var TinyURLInfoAPISchema = tinyURLInfoAPISchema{
	GET: huma.Operation{
		OperationID: "info-tinyurl",
		Method:      http.MethodGet,
		Path:        configs.GetOnlyShortURL,
		Summary:     "Get Info tinyurl",
		Description: "Get Info tinyurl",
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
	},
}
