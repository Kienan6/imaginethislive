package interceptor

import "go.uber.org/fx"

func AsInterceptor(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Middleware)),
		fx.ResultTags(`group:"interceptors"`))
}

func Index() fx.Option {
	return fx.Module("itl.interceptor",
		fx.Provide(AsInterceptor(NewSimpleMiddleware)),
		fx.Provide(AsInterceptor(NewUserMiddleware)),
	)
}
