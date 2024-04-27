package configs

import (
	"fmt"
	"sync"

	"github.com/kelseyhightower/envconfig"
)

type AppEnvironment struct {
	Stage string `default:"local"`
}

var newOnceLogger = sync.OnceValue(func() *AppEnvironment {
	var ae AppEnvironment
	if err := envconfig.Process("", &ae); err != nil {
		panic(fmt.Sprintf("Failed to process environment config: %v", err))
	}
	return &ae
})

// SEE: https://pkg.go.dev/github.com/kelseyhightower/envconfig
func NewAppEnvironment() *AppEnvironment {
	return newOnceLogger()
}
