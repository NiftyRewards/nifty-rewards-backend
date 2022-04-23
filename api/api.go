package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pg/pg/v10"
	"log"
	"net/http"
)

func NewAPI(pgdb *pg.DB) *chi.Mux {
	log.Printf("IN NewAPI")
	// setup router
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.WithValue("DB", pgdb))
	r.Route("/users", func(r chi.Router) {
		r.Post("/bind/{address_w3a}/{address_to_bind}", PostUpsertBinding)
		//r.Get("/rewards/{merchant_id}}", GetRewardsByMerchantId)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from NiftyRewards get /"))
	})

	return r
}
