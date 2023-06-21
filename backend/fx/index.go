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

func Index() fx.Option {

	options := []fx.Option{
		fx.Provide(config.NewHttpServerConfig),
		fx.Provide(config.NewPostgresConfig),
		db.Index(),
		repository.Index(),
		interceptor.Index(),
		service.Index(),
		controller.Index(),
		fx.Provide(gin.NewEngine),
		runner.Index(),
	}

	return fx.Module("itl", options...)
}
