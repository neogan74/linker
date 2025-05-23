package bot_server

import (
	"context"
	"net/http"

	"github.com/neogan74/linker/internal/config"
	"go.uber.org/fx"
)

func NewServer(handler http.Handler, cfg *config.Config) *http.Server {
	return &http.Server{
		Addr:    cfg.BotServerPort,
		Handler: handler,
	}
}

func RunServer(lc fx.Lifecycle, server *http.Server) error {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
	return nil
}
