package domain

type ShortURL struct {
	ID          string `json:"id"`
	OriginalURL string `json:"original"`
	CreatedAt   string `json:"created_at"`
}
