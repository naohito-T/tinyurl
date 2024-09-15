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

var newOnceLogger = sync.OnceValue(func() AppEnvironment {
	var ae AppEnvironment
	if err := envconfig.Process("", &ae); err != nil {
		panic(fmt.Sprintf("Failed to process environment config: %v", err))
	}
	return ae
})

// SEE: https://pkg.go.dev/github.com/kelseyhightower/envconfig
func NewAppEnvironment() AppEnvironment {
	return newOnceLogger()
}

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
