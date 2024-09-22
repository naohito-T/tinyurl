package container

import (
	"github.com/naohito-T/tinyurl/backend/configs"
	infra "github.com/naohito-T/tinyurl/backend/internal/infrastructure"
	repo "github.com/naohito-T/tinyurl/backend/internal/repository/dynamo"
	"github.com/naohito-T/tinyurl/backend/internal/rest/controller"
	"github.com/naohito-T/tinyurl/backend/internal/usecase"

	"sync"
)

type container struct {
	public controller.IPublicController
}

// 複数のusecaseをまとめてinjectionする
// containerで once で一度だけ初期化する
var oncePublicContainer = sync.OnceValue(func() *container {
	dynamoRepo := repo.NewShortURLRepository(infra.NewDynamoConnection(infra.InfrastructureLogger, configs.NewAppEnvironment()), infra.RepositoryLogger)

	return &container{
		public: controller.NewPublicController(usecase.NewURLUsecase(dynamoRepo, infra.UsecaseLogger), configs.NewAppEnvironment(), infra.ControllerLogger),
	}
})

func newContainer() *container {
	return oncePublicContainer()
}

var PublicContainer = newContainer().public
