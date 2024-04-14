package usecase

import (
	"github.com/naohito-T/tinyurl/backend/domain"
)

type IShortURLRepo interface {
	GetByShortURL(id string) (domain.ShortURL, error)
	CreateShortURL(params *domain.ShortURL) error
}

type ShortURLUsecase struct {
	shortURLRepo IShortURLRepo
}

func NewURLUsecase(u IShortURLRepo) *ShortURLUsecase {
	return &ShortURLUsecase{
		shortURLRepo: u,
	}
}

func (u *ShortURLUsecase) GetByShortURL(id string) domain.ShortURL {
	shortURL, err := u.shortURLRepo.GetByShortURL(id)
	if err != nil {
		panic(err)
	}
	return shortURL
}

func (u *ShortURLUsecase) CreateShortURL(params *domain.ShortURL) error {
	return u.shortURLRepo.CreateShortURL(params)
}
