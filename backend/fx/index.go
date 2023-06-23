package fx

import (
	"go.uber.org/fx"
	"itl/config"
	"itl/controller"
	"itl/db"
	"itl/gin"
	"itl/interceptor"
	"itl/repository"
	"itl/runner"
	"itl/service"
)

func Index(testing bool) fx.Option {

	options := []fx.Option{
		fx.Provide(config.NewHttpServerConfig),
		fx.Provide(config.NewPostgresConfig),
		db.Index(),
		repository.Index(),
		interceptor.Index(),
		service.Index(),
		controller.Index(),
		fx.Provide(gin.NewEngine),
	}

	if !testing {
		options = append(options, runner.Index())
	}

	return fx.Module("itl", options...)
}
