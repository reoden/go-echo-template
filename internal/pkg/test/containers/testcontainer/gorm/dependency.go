package gorm

import (
	"context"
	"testing"

	"github.com/reoden/go-echo-template/internal/pkg/database/options"
	"github.com/reoden/go-echo-template/internal/pkg/logger"
)

var GormContainerOptionsDecorator = func(t *testing.T, ctx context.Context) interface{} {
	return func(c *options.GormOptions, logger logger.Logger) (*options.GormOptions, error) {
		return NewGormTestContainers(logger).PopulateContainerOptions(ctx, t)
	}
}
