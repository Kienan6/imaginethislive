package repository

import "go.uber.org/fx"

func Index() fx.Option {
	return fx.Module("itl.domain",
		fx.Provide(
			NewUserRepository,
			NewPostRepository,
			NewGroupRepository,
			NewCommentRepository,
		),
	)
}
