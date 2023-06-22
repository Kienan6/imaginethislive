package service

import "go.uber.org/fx"

func Index() fx.Option {
	return fx.Module("itl.service",
		fx.Provide(
			NewUserService,
			NewGroupService,
			NewPostService,
		),
	)
}
