package db

import (
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"itl/config"
	"log"
	"os"
	"time"
)

type PostgresConnection struct {
	Db *gorm.DB
}

type PostgresConnectionParams struct {
	fx.In
	Config *config.PostgresConfig
}

func NewPostgresConnection(params PostgresConnectionParams) *PostgresConnection {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	db, err := gorm.Open(postgres.Open(params.Config.Dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalln("Error connecting to db")
	}
	return &PostgresConnection{
		Db: db,
	}
}
