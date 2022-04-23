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
		r.Get("/rewards/{merchant_id}", GetRewardsByMerchantId)
		r.Put("/rewards/{reward_id}", PutRewardByRewardId)
		// TODO:
		r.Get("/nfts/{address_w3a}", GetNftsOfAccount)
		// r.Put("/rewards/redeem", RedeemReward")
	})


	r.Route("/merchants", func(r chi.Router) {
		// TODO:
		// r.Get("", GetMerchantList)
	})

	r.Route("/merchant", func(r chi.Router) {
		// TODO:
		// r.Post("/startcampaign", StartCampaign)
		// r.Get("/campaigns", GetAllCampaignsByMerchant)
	})


	r.Route("/collectionowner", func(r chi.Router) {
		// TODO:
		// r.Post("/add", ApproveCampaign)
		// r.Get("/campaigns", GetAllCampaignsByMerchant)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from NiftyRewards get /"))
	})

	return r
}
