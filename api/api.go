package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pg/pg/v10"
	"net/http"
)

func NewAPI(pgdb *pg.DB) *chi.Mux {
	// setup router
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.WithValue("DB", pgdb))
	r.Route("/users", func(r chi.Router) {
		r.Post("/", CreateUser)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from NiftyRewards get /"))
	})

	return r
}
