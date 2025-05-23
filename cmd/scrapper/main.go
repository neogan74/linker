package main

import (
	"github.com/neogan74/linker/internal/config"
	"go.uber.org/fx"
)

func main() {
	addOpts := fx.Options(
		fx.Provide(config.NewConfig),
	)
	fx.New(addOpts).Run()
}
