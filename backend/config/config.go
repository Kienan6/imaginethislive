package config

import (
	"os"
)

var HttpServerAddressKey = "HTTP_ADDRESS"
var HttpServerPortKey = "HTTP_PORT"

type HttpServerConfig struct {
	Address string
	Port    string
}

func NewHttpServerConfig() *HttpServerConfig {
	return &HttpServerConfig{
		Address: os.Getenv(HttpServerAddressKey),
		Port:    os.Getenv(HttpServerPortKey),
	}
}
