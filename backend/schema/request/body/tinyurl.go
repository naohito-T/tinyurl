package body

type CreateTinyURLBody struct {
	Body struct {
		URL string `json:"url" required:"true" example:"http://example.com" doc:"URL to shorten"`
	}
}
