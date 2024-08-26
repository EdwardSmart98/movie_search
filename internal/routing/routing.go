package routing

import (
	"github.com/go-chi/chi/v5"
)

func AddRoutes(r *chi.Mux) {
	AddMovieRoutes(r)
	AddActorRoutes(r)
}
