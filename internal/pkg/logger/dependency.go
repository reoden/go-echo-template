package logger

import (
	"github.com/reoden/go-echo-template/internal/pkg/config/environment"
	"github.com/reoden/go-echo-template/internal/pkg/logger/config"
	"github.com/reoden/go-echo-template/internal/pkg/logger/zap"

	"go.uber.org/dig"
)

func AddLogger(container *dig.Container) error {
	err := container.Provide(func(environment environment.Environment) (*config.LogOptions, error) {
		return config.ProvideLogConfig(environment)
	})
	if err != nil {
		return err
	}

	err = container.Provide(func(opts *config.LogOptions, environment environment.Environment) Logger {
		return zap.NewZapLogger(opts, environment)
	})
	if err != nil {
		return err
	}

	return err
}
