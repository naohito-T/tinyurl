package handler

import (
	"context"

	"github.com/naohito-T/tinyurl/backend/domain"
	"github.com/naohito-T/tinyurl/backend/internal/infrastructures/slog"
	"github.com/naohito-T/tinyurl/backend/internal/rest/container"
)

type ShortURLHandler struct {
	container *container.GuestContainer
}

// IShortURLHandler defines the interface for short URL handler operations.
type IShortURLHandler interface {
	CreateShortURLHandler(ctx context.Context, originalURL string) (domain.ShortURL, error)
}

// NewShortURLHandler creates a new handler for short URLs.
func NewShortURLHandler(c *container.GuestContainer) IShortURLHandler {
	return &ShortURLHandler{
		container: c,
	}
}

func (s *ShortURLHandler) CreateShortURLHandler(ctx context.Context, originalURL string) (domain.ShortURL, error) {
	slog.NewLogger().Info("CreateShortURLHandler: %v", originalURL)
	return s.container.URLUsecase.CreateShortURL(ctx, originalURL)
}
