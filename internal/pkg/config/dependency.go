package config

import (
	"github.com/reoden/go-echo-template/internal/pkg/config/environment"

	"go.uber.org/dig"
)

func AddEnv(container *dig.Container, environments ...environment.Environment) error {
	err := container.Provide(func() environment.Environment {
		return environment.ConfigEnv(environments...)
	})

	return err
}
