package runner

import (
	"go.uber.org/fx"
)

func Index() fx.Option {
	return fx.Module("itl.runner",
		fx.Provide(NewHttpServer),
		fx.Invoke(func(server *HttpServer) {}),
	)
}
