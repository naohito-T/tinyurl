package container

import (
	"github.com/naohito-T/tinyurl/backend/configs"
	infra "github.com/naohito-T/tinyurl/backend/internal/infrastructure"
	repo "github.com/naohito-T/tinyurl/backend/internal/repository/dynamo"
	"github.com/naohito-T/tinyurl/backend/internal/usecase"

	"sync"
)

type PublicContainer struct {
	*usecase.ShortURLUsecase
}

var oncePublicContainer = sync.OnceValue(func() *PublicContainer {
	logger := infra.NewLogger()
	env := configs.NewAppEnvironment()
	dynamoRepo := repo.NewShortURLRepository(infra.NewDynamoConnection(logger, env), logger)
	return &PublicContainer{
		usecase.NewURLUsecase(dynamoRepo, logger),
	}
})

func NewPublicContainer() *PublicContainer {
	return oncePublicContainer()
}
