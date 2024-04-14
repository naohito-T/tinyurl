package container

import (
	"github.com/naohito-T/tinyurl/backend/internal/infrastructures/dynamo"
	"github.com/naohito-T/tinyurl/backend/internal/infrastructures/slog"
	repoDynamo "github.com/naohito-T/tinyurl/backend/internal/repository/dynamo"
	"github.com/naohito-T/tinyurl/backend/internal/usecase"

	"sync"
)

type GuestContainer struct {
	URLUsecase *usecase.ShortURLUsecase
}

var newGuestContainer = sync.OnceValue(func() *GuestContainer {
	dynamoRepo := repoDynamo.NewShortURLRepository(dynamo.NewDynamoConnection(), slog.NewLogger())
	return &GuestContainer{
		URLUsecase: usecase.NewURLUsecase(dynamoRepo),
	}
})

func NewGuestContainer() *GuestContainer {
	return newGuestContainer()
}
