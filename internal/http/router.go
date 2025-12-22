package http

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func NewRouter(clientHandler *ClientHandler, logger *zap.Logger) *chi.Mux {
	r := chi.NewRouter()

	r.Use(ZapMiddleware(logger))

	r.Route("/api", func(r chi.Router) {
		r.Mount("/pets", clientRouter(clientHandler))
	})

	return r
}

func clientRouter(clientHandler *ClientHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", clientHandler.List)
	r.Post("/create", clientHandler.Create)
	r.Put("/update/{id}", clientHandler.Update)
	return r
}
