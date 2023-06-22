package repository

import "go.uber.org/fx"

func Index() fx.Option {
	return fx.Module("itl.repository",
		fx.Provide(
			NewUserRepository,
			NewPostRepository,
			NewGroupRepository,
		),
	)
}
