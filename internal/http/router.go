package http

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func NewRouter(petHandler *PetHandler, vaccineHandler *VaccineHandler, logger *zap.Logger) *chi.Mux {
	r := chi.NewRouter()

	r.Use(ZapMiddleware(logger))

	r.Route("/api", func(r chi.Router) {
		r.Mount("/pets", PetRouter(petHandler))
		r.Mount("/vaccines", VaccineRouter(vaccineHandler))
	})

	return r
}

func PetRouter(PetHandler *PetHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", PetHandler.List)
	r.Post("/create", PetHandler.Create)
	r.Put("/update/{id}", PetHandler.Update)
	r.Delete("/delete/{id}", PetHandler.Delete)
	return r
}

func VaccineRouter(VaccineHandler *VaccineHandler) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", VaccineHandler.List)
	r.Post("/create", VaccineHandler.Create)
	return r
}
