package config

import (
	"fmt"
	"os"
)

var PostgresAddressKey = "ITL_POSTGRES_ADDRESS"
var PostgresPortKey = "ITL_POSTGRES_PORT"
var PostgresUserKey = "ITL_POSTGRES_USER"
var PostgresPasswordKey = "ITL_POSTGRES_PASSWORD"
var PostgresDbName = "ITL_POSTGRES_NAME"

type PostgresConfig struct {
	Dsn string
}

func NewPostgresConfig() *PostgresConfig {
	address := os.Getenv(PostgresAddressKey)
	port := os.Getenv(PostgresPortKey)
	name := os.Getenv(PostgresDbName)
	user := os.Getenv(PostgresUserKey)
	pass := os.Getenv(PostgresPasswordKey)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", address, user, pass, name, port)
	return &PostgresConfig{
		Dsn: dsn,
	}
}
