package usecase

import (
	"context"

	"github.com/naohito-T/tinyurl/backend/domain"
)

type IShortURLRepo interface {
	GetByShortURL(ctx context.Context, id string) (domain.ShortURL, error)
	CreateShortURL(ctx context.Context, params *domain.ShortURL) error
}

type ShortURLUsecase struct {
	shortURLRepo IShortURLRepo
}

func NewURLUsecase(u IShortURLRepo) *ShortURLUsecase {
	return &ShortURLUsecase{
		shortURLRepo: u,
	}
}

func (u *ShortURLUsecase) GetByShortURL(ctx context.Context, id string) (domain.ShortURL, error) {
	return u.shortURLRepo.GetByShortURL(ctx, id)
}

func (u *ShortURLUsecase) CreateShortURL(ctx context.Context, params *domain.ShortURL) error {
	return u.shortURLRepo.CreateShortURL(ctx, params)
}
