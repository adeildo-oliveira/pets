package main

import (
	_ "github.com/denisenkom/go-mssqldb"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"pets/internal/config"
	"pets/internal/domain/interfaces"
	"pets/internal/domain/services"
	"pets/internal/http"
	"pets/internal/repository"
)

func main() {
	app := fx.New(
		fx.Provide(
			NewLogger,
			config.LoadConfig,
			http.NewClientHandler,
			http.NewRouter,
			http.NewHTTPServer,
		),
		fx.Provide(
			fx.Annotate(repository.NewPetDb, fx.As(new(interfaces.ISqlDB))),
			fx.Annotate(repository.NewPetRepository, fx.As(new(interfaces.IPetRepository))),
			fx.Annotate(services.NewPetService, fx.As(new(interfaces.IPetService))),
		),
		fx.Invoke(http.StartServer),
	)

	app.Run()
}

func NewLogger() (*zap.Logger, error) {
	return zap.NewProduction()
}
