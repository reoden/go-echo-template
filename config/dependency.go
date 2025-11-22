package config

import (
	"github.com/reoden/go-echo-template/internal/pkg/config/environment"

	"go.uber.org/dig"
)

func AddAppConfig(container *dig.Container) error {
	err := container.Provide(func(environment environment.Environment) (*Config, error) {
		return NewAppConfig(environment)
	})

	return err
}
