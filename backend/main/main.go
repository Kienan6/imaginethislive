package main

import (
	"go.uber.org/fx"
	serverfx "itl/fx"
)

// @title           Imagine This Live
// @version         1.0
// @description     Imagine this live backend applications

// @contact.name   Kienan O'Brien

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.basic BasicAuth

// @host      localhost:9005
// @BasePath  /v1

func main() {
	fx.New(
		serverfx.Index(false),
	).Run()
}
