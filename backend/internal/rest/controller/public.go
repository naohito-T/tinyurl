package controller

import (
	"context"

	"github.com/naohito-T/tinyurl/backend/configs"
	"github.com/naohito-T/tinyurl/backend/domain"
	"github.com/naohito-T/tinyurl/backend/internal/infrastructure"
	"github.com/naohito-T/tinyurl/backend/internal/usecase"
)

// ポイント1: インターフェースに構造体を埋め込むことはできない。逆に他3つはできる。
// ポイント2: "借り物"のメソッドを自分のものとして使うことができる。
// ポイント3: 他言語の継承とは異なり、埋め込み先のメンバーに影響を与えない。
// ポイント4: 埋め込み元と埋め込み先に同じフィールド名が存在するとき、埋め込み先が優先される。
// 埋め込み(embedding)

type IPublicController interface {
	HealthCheck() (interface{}, error)
	GetShortURL(ctx context.Context, hashID string) (domain.ShortURL, error)
	CreateShortURL(ctx context.Context, originalURL string) (domain.ShortURL, error)
}

type PublicController struct {
	// こっちは構造体
	usecase *usecase.ShortURLUsecase
	env     *configs.AppEnvironment
	// こっちはinterface
	logger infrastructure.ILabelLogger
}

func (c *PublicController) HealthCheck() (interface{}, error) {
	// 実際のヘルスチェックのロジックをここに実装
	return "OK", nil
}

func (c *PublicController) GetShortURL(ctx context.Context, hashID string) (domain.ShortURL, error) {
	// c.logger.Info("GetShortURLHandler: %v", hashID)
	return c.usecase.GetByShortURL(ctx, hashID)
}

func (c *PublicController) CreateShortURL(ctx context.Context, originalURL string) (domain.ShortURL, error) {
	// c.logger.Info("CreateShortURLHandler: %v", originalURL)
	return c.usecase.CreateShortURL(ctx, originalURL)
}

// containerを作りコントローラーを作る。
func NewPublicController(usecase *usecase.ShortURLUsecase,
	env *configs.AppEnvironment,
	logger infrastructure.ILabelLogger) IPublicController {
	return &PublicController{
		usecase: usecase,
		env:     env,
		logger:  logger,
	}
}
