package api

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-pg/pg/v10"
	"golang-server/db"
	"log"
	"net/http"
	"strconv"
)

type RewardResponse struct {
	Success bool        `json:"success"`
	Error   string      `json:"err"`
	Reward  *db.Rewards `json:"reward"`
}

type RewardsResponse struct {
	Success bool          `json:"success"`
	Error   string        `json:"err"`
	Rewards []*db.Rewards `json:"rewards"`
}

func GetRewardsByMerchantId(w http.ResponseWriter, r *http.Request) {
	merchantId, err := strconv.Atoi(chi.URLParam(r, "merchant_id"))
	if err != nil {
		log.Printf("GetRewardsByMerchantId err1: %v\n", err)
		w = rewardErrResponse(err, w)
		return
	}

	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		w = rewardErrResponse(errors.New("could not get database from context"), w)
		return
	}
	// query for the reward
	rewards, err := db.GetRewardsByMerchantId(pgdb, merchantId)
	if err != nil {
		log.Printf("GetRewardsByMerchantId err2: %v\n", err)
		w = rewardErrResponse(err, w)
		return
	}

	// return a response
	res := &RewardsResponse{
		Success: true,
		Error:   "",
		Rewards: rewards,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("GetRewardsByMerchantId err3: %v\n", err)
	}
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
