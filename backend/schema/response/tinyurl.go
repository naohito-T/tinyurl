package response

type GetTinyURLResponse struct {
	Status int
	URL    string `header:"Location"`
}

type CreateTinyURLResponse struct {
	Body struct {
		ID string `json:"id"`
	}
}

type GetInfoTinyURLResponse struct {
	Body struct {
		ID          string `json:"id" required:"true"`
		OriginalURL string `json:"original_url" required:"true"`
		CreatedAt   string `json:"created_at" required:"true"`
	}
}
