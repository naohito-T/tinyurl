package usecase

import (
	"context"

	"github.com/naohito-T/tinyurl/backend/domain"
	"github.com/naohito-T/tinyurl/backend/internal/infrastructure"
)

type IShortURLRepo interface {
	Get(ctx context.Context, hashURL string) (domain.ShortURL, error)
	Put(ctx context.Context, shortURL *domain.ShortURL) (domain.ShortURL, error)
}

type ShortURLUsecase struct {
	shortURLRepo IShortURLRepo
	logger       infrastructure.ILabelLogger
}

func (u *ShortURLUsecase) GetByShortURL(ctx context.Context, hashID string) (domain.ShortURL, error) {
	// u.logger.Info("GetByShortURL: %v", hashID)
	return u.shortURLRepo.Get(ctx, hashID)
}

func (u *ShortURLUsecase) CreateShortURL(ctx context.Context, originalURL string) (domain.ShortURL, error) {
	// u.logger.Info("CreateShortURL: %v", originalURL)
	return u.shortURLRepo.Put(ctx, domain.GenerateShortURL(originalURL))
}

func (u *ShortURLUsecase) Search(ctx context.Context, originalURL string) (domain.ShortURL, error) {
	// u.logger.Info("GetByOriginalURL: %v", originalURL)
	return u.shortURLRepo.Get(ctx, originalURL)
}

func NewURLUsecase(u IShortURLRepo, logger infrastructure.ILabelLogger) *ShortURLUsecase {
	return &ShortURLUsecase{
		shortURLRepo: u,
		logger:       logger,
	}
}
