package main

import (
	"net/http"

	"github.com/neogan74/linker/internal/api/bot_api"
	"github.com/neogan74/linker/internal/api/bot_api/codegen"
	"github.com/neogan74/linker/internal/application/bot_app"
	"github.com/neogan74/linker/internal/bot_server"
	"github.com/neogan74/linker/internal/config"
	"github.com/neogan74/linker/internal/infra/clients"
	"go.uber.org/fx"
)

func main() {
	addOpts := fx.Options(
		fx.Provide(config.NewConfig),
		fx.Provide(bot_app.NewBot),
		fx.Provide(clients.NewBotClient),
		fx.Provide(bot_app.NewBotApp),
		fx.Provide(bot_api.NewBotHandler),
		fx.Provide(
			func(h codegen.ServerInterface) http.Handler {
				return codegen.Handler(h)
			},
		),
		fx.Provide(bot_server.NewServer),
		fx.Invoke(bot_server.RunServer),
		fx.Invoke(bot_app.RunBot),
	)
	fx.New(addOpts).Run()
}
