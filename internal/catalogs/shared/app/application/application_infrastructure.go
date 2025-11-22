package application

import (
	"github.com/reoden/go-echo-template/docs"
	"github.com/reoden/go-echo-template/internal/catalogs/products/contracts"
	"github.com/reoden/go-echo-template/internal/catalogs/products/features/creatingproduct/commands"
	"github.com/reoden/go-echo-template/internal/catalogs/products/features/creatingproduct/dtos"
	"github.com/reoden/go-echo-template/internal/catalogs/products/features/creatingproduct/events"
	dtos2 "github.com/reoden/go-echo-template/internal/catalogs/products/features/gettingproductbyid/dtos"
	"github.com/reoden/go-echo-template/internal/catalogs/products/features/gettingproductbyid/queries"
	"github.com/reoden/go-echo-template/internal/catalogs/shared/behaviours"

	"emperror.dev/errors"
	"github.com/mehdihadeli/go-mediatr"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (a *Application) ConfigInfrastructure() error {
	err := a.configMediator()
	if err != nil {
		return errors.WrapIf(err, "Error in setting mediator handlers")
	}

	a.configSwagger()

	return err
}

func (a *Application) configMediator() error {
	return a.ResolveDependencyFunc(func(productRepository contracts.ProductRepository) error {
		loggerPipeline := &behaviours.RequestLoggerBehaviour{}
		err := mediatr.RegisterRequestPipelineBehaviors(loggerPipeline)
		if err != nil {
			return err
		}

		createProductCommandHandler := commands.NewCreateProductCommandHandler(productRepository)
		err = mediatr.RegisterRequestHandler[*commands.CreateProductCommand, *dtos.CreateProductCommandResponse](
			createProductCommandHandler,
		)
		if err != nil {
			return err
		}

		getByIdQueryHandler := queries.NewGetProductByIdHandler(productRepository)
		err = mediatr.RegisterRequestHandler[*queries.GetProductByIdQuery, *dtos2.GetProductByIdQueryResponse](
			getByIdQueryHandler,
		)
		if err != nil {
			return err
		}

		notificationHandler := events.NewProductCreatedEventHandler()
		err = mediatr.RegisterNotificationHandler[*events.ProductCreatedEvent](notificationHandler)
		if err != nil {
			return err
		}

		return nil
	})
}

func (a *Application) configSwagger() {
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Title = "Catalogs Write-Service Api"
	docs.SwaggerInfo.Description = "Catalogs Write-Service Api."

	a.Echo.GET("/swagger/*", echoSwagger.WrapHandler)
}
