package api

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-pg/pg/v10"
	"golang-server/db"
	"log"
	"net/http"
)

type RewardResponse struct {
	Success bool        `json:"success"`
	Error   string      `json:"err"`
	Reward  *db.Rewards `json:"reward"`
}

func GetRewardsByMerchantId(w http.ResponseWriter, r *http.Request) {
	merchantId := chi.URLParam(r, "merchant_id")

	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		w = rewardErrResponse(errors.New("could not get database from context"), w)
		return
	}
	// query for the reward
	reward, err := db.GetRewards(pgdb, &db.Rewards{
		AddressW3a: addressW3a,
		Address_B:  addressB,
	})
	if err != nil {
		w = rewardErrResponse(err, w)
		return
	}

	// return a response
	res := &RewardResponse{
		Success: true,
		Error:   "",
		Reward:  reward,
	}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

func rewardErrResponse(err error, w http.ResponseWriter) http.ResponseWriter {
	res := &RewardResponse{
		Success: false,
		Error:   err.Error(),
		Reward:  nil,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("err sending response: %v\n", err)
	}
	w.WriteHeader(http.StatusBadRequest)

	return w
}
