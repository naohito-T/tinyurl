package usecase

import (
	"context"

	"github.com/naohito-T/tinyurl/backend/domain"
	"github.com/naohito-T/tinyurl/backend/internal/infrastructures/slog"
)

type IShortURLRepo interface {
	Get(ctx context.Context, hashURL string) (domain.ShortURL, error)
	Put(ctx context.Context, shortURL *domain.ShortURL) (domain.ShortURL, error)
}

type ShortURLUsecase struct {
	// ここでIShortURLRepoを埋め込むことで、UsecaseがRepositoryを知っている
	// 	はい、その通りです。IShortURLRepo インターフェースは、Get と Create という二つのメソッドを定義しており、このインターフェースを実装するどのクラスも、これらのメソッドを具体的に実装する必要があります。そして、ShortURLUsecase の中で shortURLRepo.Create(originalURL) を呼び出すことによって、このインターフェースを満たす具体的な実装に対して処理を委譲しています。
	// ここでのポイントは、ShortURLUsecase クラスが IShortURLRepo インターフェースの具体的な実装に依存していないということです。この設計により、IShortURLRepo の実装を変更しても、ShortURLUsecase クラスを修正する必要がなくなります。つまり、データアクセス層の実装が変わっても、ビジネスロジック層は影響を受けないという設計原則（オープン/クローズド原則）に従っています。
	shortURLRepo IShortURLRepo
}

func NewURLUsecase(u IShortURLRepo) *ShortURLUsecase {
	return &ShortURLUsecase{
		shortURLRepo: u,
	}
}

func (u *ShortURLUsecase) GetByShortURL(ctx context.Context, hashID string) (domain.ShortURL, error) {
	slog.NewLogger().Info("GetByShortURL: %v", hashID)
	return u.shortURLRepo.Get(ctx, hashID)
}

func (u *ShortURLUsecase) CreateShortURL(ctx context.Context, originalURL string) (domain.ShortURL, error) {
	slog.NewLogger().Info("CreateShortURL: %v", originalURL)
	return u.shortURLRepo.Put(ctx, domain.GenerateShortURL(originalURL))
}
