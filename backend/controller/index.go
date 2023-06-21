package controller

import "go.uber.org/fx"

func Index() fx.Option {
	return fx.Module("itl.controllers",
		fx.Provide(
			NewSimpleRoutesController,
			NewUserRoutesController,
			NewGroupRoutesController,
		))
}
