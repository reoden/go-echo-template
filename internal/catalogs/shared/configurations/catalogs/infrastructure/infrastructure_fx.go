package infrastructure

import (
	"github.com/reoden/go-echo-template/internal/pkg/core"
	"github.com/reoden/go-echo-template/internal/pkg/grpc"
	"github.com/reoden/go-echo-template/internal/pkg/health"
	customEcho "github.com/reoden/go-echo-template/internal/pkg/http/customecho"
	"github.com/reoden/go-echo-template/internal/pkg/mongodb"
	"github.com/reoden/go-echo-template/internal/pkg/otel/metrics"
	"github.com/reoden/go-echo-template/internal/pkg/otel/tracing"
	"github.com/reoden/go-echo-template/internal/pkg/redis"

	"github.com/go-playground/validator"
	"go.uber.org/fx"
)

// https://pmihaylov.com/shared-components-go-microservices/
var Module = fx.Module(
	"infrastructurefx",
	// Modules
	core.Module,
	customEcho.Module,
	grpc.Module,
	mongodb.Module,
	redis.Module,
	//rabbitmq.ModuleFunc(
	//	func(v *validator.Validate, l logger.Logger, tracer tracing.AppTracer) configurations.RabbitMQConfigurationBuilderFuc {
	//		return func(builder configurations.RabbitMQConfigurationBuilder) {
	//			rabbitmq2.ConfigProductsRabbitMQ(builder, l, v, tracer)
	//		}
	//	},
	//),
	health.Module,
	tracing.Module,
	metrics.Module,

	// Other provides
	fx.Provide(validator.New),
)
