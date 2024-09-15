package handler

import (
	"context"

	"github.com/naohito-T/tinyurl/backend/domain"
	"github.com/naohito-T/tinyurl/backend/internal/infrastructure"
	"github.com/naohito-T/tinyurl/backend/internal/rest/container"
)

// ポイント1: インターフェースに構造体を埋め込むことはできない。逆に他3つはできる。
// ポイント2: "借り物"のメソッドを自分のものとして使うことができる。
// ポイント3: 他言語の継承とは異なり、埋め込み先のメンバーに影響を与えない。
// ポイント4: 埋め込み元と埋め込み先に同じフィールド名が存在するとき、埋め込み先が優先される。
// 埋め込み(embedding)
type ShortURLHandler struct {
	*container.GuestContainer
	infrastructure.ILogger
}

// IShortURLHandler defines the interface for short URL handler operations.
type IShortURLHandler interface {
	GetShortURLHandler(ctx context.Context, hashID string) (domain.ShortURL, error)
	CreateShortURLHandler(ctx context.Context, originalURL string) (domain.ShortURL, error)
}

// NewShortURLHandler creates a new handler for short URLs.
func NewShortURLHandler(c *container.GuestContainer, logger infrastructure.ILogger) IShortURLHandler {
	return &ShortURLHandler{
		c,
		logger,
	}
}

func (s *ShortURLHandler) GetShortURLHandler(ctx context.Context, hashID string) (domain.ShortURL, error) {
	s.Info("GetShortURLHandler: %v", hashID)
	return s.GetByShortURL(ctx, hashID)
}

func (s *ShortURLHandler) CreateShortURLHandler(ctx context.Context, originalURL string) (domain.ShortURL, error) {
	s.Info("CreateShortURLHandler: %v", originalURL)
	return s.CreateShortURL(ctx, originalURL)
}
