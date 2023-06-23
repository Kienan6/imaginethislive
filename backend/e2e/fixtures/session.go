package fixtures

import (
	"context"
	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"
	"itl/config"
	serverfx "itl/fx"
	"log"
)

func NewClient() *resty.Client {
	config := config.NewHttpServerConfig()
	return resty.New().SetBaseURL("http://" + config.Address + ":" + config.Port)
}

func NewSession(options ...interface{}) {

	app := fx.New(
		serverfx.Index(false),
		fx.Populate(options...),
	)

	if err := app.Start(context.Background()); err != nil {
		log.Fatalln("bad")
	}
	defer app.Stop(context.Background())
}
