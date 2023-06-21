package main

import (
	"go.uber.org/fx"
	serverfx "itl/fx"
)

func main() {
	fx.New(
		serverfx.Index(),
	).Run()
}
