package contracts

import (
	"context"

	"github.com/reoden/go-echo-template/internal/pkg/config/environment"
	"github.com/reoden/go-echo-template/internal/pkg/logger"

	"go.uber.org/fx"
)

type Application interface {
	Container
	RegisterHook(function interface{})
	Run()
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Wait() <-chan fx.ShutdownSignal
	Logger() logger.Logger
	Environment() environment.Environment
}
