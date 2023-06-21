package db

import "go.uber.org/fx"

func Index() fx.Option {
	return fx.Module("itl.db",
		fx.Provide(NewPostgresConnection),
		fx.Invoke(func(db *PostgresConnection) {}),
	)
}
