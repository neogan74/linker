package main

import (
	"github.com/neogan74/linker/internal/config"
	"go.uber.org/fx"
)

func main() {
	addOpts := fx.Options(
		fx.Provide(config.NewConfig),
		fx.Provide(scrapper_app.NewScheduler),
	)
	fx.New(addOpts).Run()
}
