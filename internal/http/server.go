package http

import (
	"context"
	lib "net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"pets/internal/config"
)

func NewHTTPServer(router *chi.Mux, cfg *config.Server) *lib.Server {
	return &lib.Server{
		Addr:    cfg.Config.ServerPort,
		Handler: router,
	}
}

func StartServer(lc fx.Lifecycle, server *lib.Server, logger *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info("server started", zap.String("addr", server.Addr))
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("server stopping")
			return server.Shutdown(ctx)
		},
	})
}
