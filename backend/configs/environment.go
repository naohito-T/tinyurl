package configs

import (
	"fmt"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type AppEnvironment struct {
	stage             string `default:"local"`
	tinyURLCollection string `default:"offline-tinyurls"`
}

// コピーで返す場合
// 構造体のコピーを返すだけなので生成も一度だけでコストは低い。
// ただし、構造体が大きい場合はコピーのコストがかかるので注意。
// 参照で返す場合
// 構造体のコピーを作成せずに参照を返すため、コストは低い。
var newOnceEnvironment = sync.OnceValue(func() *AppEnvironment {
	var appEnvironment AppEnvironment
	if err := envconfig.Process("", &appEnvironment); err != nil {
		panic(fmt.Sprintf("Failed to process environment config: %v", err))
	}

	return &appEnvironment
})

func (a *AppEnvironment) IsTest() bool {
	return a.stage == "test"
}

func (a *AppEnvironment) IsLocal() bool {
	return a.stage == "local"
}

func (a *AppEnvironment) IsDev() bool {
	return a.stage == "dev"
}

func (a *AppEnvironment) IsProd() bool {
	return a.stage == "prod"
}

func (a *AppEnvironment) GetTinyURLCollectionName() string {
	return a.tinyURLCollection
}

// SEE: https://pkg.go.dev/github.com/kelseyhightower/envconfig
func NewAppEnvironment() *AppEnvironment {
	return newOnceEnvironment()
}
