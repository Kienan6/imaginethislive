package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var HttpServerAddressKey = "HTTP_ADDRESS"
var HttpServerPortKey = "HTTP_PORT"

type HttpServerConfig struct {
	Address string
	Port    string
}

func NewHttpServerConfig() *HttpServerConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &HttpServerConfig{
		Address: os.Getenv(HttpServerAddressKey),
		Port:    os.Getenv(HttpServerPortKey),
	}
}
