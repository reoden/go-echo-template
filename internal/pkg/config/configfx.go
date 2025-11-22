package config

import (
	"github.com/reoden/go-echo-template/internal/pkg/config/environment"

	"go.uber.org/fx"
)

var Module = fx.Module(
	"configfx",
	fx.Provide(func() environment.Environment {
		return environment.ConfigEnv()
	}),
)

var ModuleFunc = func(e environment.Environment) fx.Option {
	return fx.Module(
		"configfx",
		fx.Provide(func() environment.Environment {
			return environment.ConfigEnv()
		}),
	)
}
