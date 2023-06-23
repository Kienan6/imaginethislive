package fixtures

import (
	"context"
	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"
	"itl/config"
	serverfx "itl/fx"
	"itl/runner"
	"log"
	"time"
)

func NewClient() *resty.Client {
	config := config.NewHttpServerConfig()
	return resty.New().SetBaseURL("http://" + config.Address + ":" + config.Port)
}

func NewSession() {
	app := fx.New(
		serverfx.Index(false),
		fx.Invoke(func(server *runner.HttpServer) {}),
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatalln("bad")
	}
}
