package runner

import (
	"context"
	"github.com/fvbock/endless"
	"go.uber.org/fx"
	"itl/config"
	"itl/gin"
	"log"
)

type HttpServer struct {
}

type HttpServerParams struct {
	fx.In
	Config *config.HttpServerConfig
	Engine *gin.Engine
}

// NewHttpServer
// Run http server with configured gin instance
func NewHttpServer(lc fx.Lifecycle, params HttpServerParams) *HttpServer {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				err := endless.ListenAndServe(params.Config.Address+":"+params.Config.Port, params.Engine.Router)
				if err != nil {
					log.Fatalln("Error starting server")
				}
			}()
			return nil
		},
	})
	return &HttpServer{}
}
