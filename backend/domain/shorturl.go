package domain

import (
	"crypto/sha1"
	"encoding/base64"
	"strings"
	"time"
)

type ShortURL struct {
	ID          string `json:"id"`
	OriginalURL string `json:"original"`
	CreatedAt   string `json:"created_at"`
}

func GenerateShortURL(originalURL string) *ShortURL {
	hasher := sha1.New()
	hasher.Write([]byte(originalURL))
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	return &ShortURL{
		ID:          strings.TrimRight(sha, "=")[:7],
		OriginalURL: originalURL,
		CreatedAt:   time.Now().Format(time.RFC3339),
	}
}
