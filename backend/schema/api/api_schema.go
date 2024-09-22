package api

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/naohito-T/tinyurl/backend/configs"

	"sync"
)

var onceHumaConfig = sync.OnceValue(func() huma.Config {
	config := huma.DefaultConfig(configs.OpenAPITitle, configs.OpenAPIVersion)
	// /api/v1/openapi.yaml
	config.Servers = []*huma.Server{
		{URL: configs.OpenAPIDocServerPath, Description: "Local API Server"},
		{URL: configs.OpenAPIDocServerPath, Description: "Dev API Server"},
		{URL: configs.OpenAPIDocServerPath, Description: "Prod API Server"},
	}

	config.Info = &huma.Info{
		Title:       configs.OpenAPITitle,
		Version:     configs.OpenAPIVersion,
		Description: "This is a simple URL shortener service.",
		Contact: &huma.Contact{
			Name:  "naohito-T",
			Email: "naohito.tanaka0523@gmail.com",
			URL:   "https://naohito-t.github.io/",
		},
		License: &huma.License{
			Name: "MIT",
			URL:  "https://opensource.org/licenses/MIT",
		},
	}
	config.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"bearer": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}
	// config.DocsPath = "/docs"
	return config
})

func NewHumaConfig() huma.Config {
	return onceHumaConfig()
}
