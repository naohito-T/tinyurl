package container

import (
	"github.com/naohito-T/tinyurl/backend/configs"
	infra "github.com/naohito-T/tinyurl/backend/internal/infrastructure"
	repo "github.com/naohito-T/tinyurl/backend/internal/repository/dynamo"
	"github.com/naohito-T/tinyurl/backend/internal/usecase"

	"sync"
)

type GuestContainer struct {
	*usecase.ShortURLUsecase
}

var onceGuestContainer = sync.OnceValue(func() *GuestContainer {
	logger := infra.NewLogger()
	env := configs.NewAppEnvironment()
	dynamoRepo := repo.NewShortURLRepository(infra.NewDynamoConnection(logger, env), logger)
	return &GuestContainer{
		usecase.NewURLUsecase(dynamoRepo, logger),
	}
})

func NewGuestContainer() *GuestContainer {
	return onceGuestContainer()
}
