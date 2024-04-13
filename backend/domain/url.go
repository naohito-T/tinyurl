package domain

type URL struct {
	ID          int    `json:"id"`
	ShortURL    string `json:"short_url"`
	OriginalURL string `json:"original"`
	CreatedAt   string `json:"created_at"`
}
