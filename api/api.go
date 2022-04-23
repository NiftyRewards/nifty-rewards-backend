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
		r.Get("/", GetAllUsers)
		r.Get("/{address_w3a}", GetUserByAddressW3A)
		r.Post("/delete/{address_w3a}", DeleteUserByAddressW3A)
		r.Post("/bind/{address_w3a}/{address_to_bind}", PostUpsertBinding)
		r.Get("/rewards/{merchant_id}", GetRewardsByMerchantId)
		r.Put("/rewards/{reward_id}", PutRewardByRewardId)
		r.Get("/nfts/{address_w3a}", GetNftsOfAccount)
	})

	r.Route("/merchants", func(r chi.Router) {
		r.Get("/", GetAllMerchants)
		r.Get("/{merchant_id}", GetMerchantById)
		r.Post("/delete/{merchant_name}", DeleteMerchantByName)
	})

	r.Route("/merchant", func(r chi.Router) {
		r.Get("/campaigns/{merchant_id}", GetAllCampaignsByMerchantId)
		r.Post("/startcampaign", PostRewards)
	})

	r.Route("/collectionowner", func(r chi.Router) {
		r.Post("/add", ApproveCampaigns)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from NiftyRewards get /"))
	})

	return r
}
