package fx

import (
	"go.uber.org/fx"
	"itl/config"
	"itl/controller"
	"itl/gin"
	"itl/interceptor"
	"itl/runner"
	"itl/service"
)

func Index() fx.Option {

	options := []fx.Option{
		fx.Provide(config.NewHttpServerConfig),
		interceptor.Index(),
		service.Index(),
		controller.Index(),
		fx.Provide(gin.NewEngine),
		runner.Index(),
	}

	return fx.Module("itl", options...)
}
