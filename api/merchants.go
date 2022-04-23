package api

import (
	"encoding/json"
	"errors"
	"github.com/go-pg/pg/v10"
	"golang-server/db"
	"log"
	"net/http"
)

type MerchantResponse struct {
	Success  bool          `json:"success"`
	Error    string        `json:"err"`
	Merchant *db.Merchants `json:"reward"`
}

type MerchantsResponse struct {
	Success   bool            `json:"success"`
	Error     string          `json:"err"`
	Merchants []*db.Merchants `json:"rewards"`
}

func GetAllMerchants(w http.ResponseWriter, r *http.Request) {
	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		w = rewardErrResponse(errors.New("could not get database from context"), w)
		return
	}
	// get all merchants
	merchants, err := db.GetMerchants(pgdb)
	if err != nil {
		log.Printf("GetAllMerchants err2: %v\n", err)
		w = rewardErrResponse(err, w)
		return
	}

	// return a response
	res := &MerchantsResponse{
		Success:   true,
		Error:     "",
		Merchants: merchants,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("GetRewardsByMerchantId err3: %v\n", err)
	}
	w.WriteHeader(http.StatusOK)
}
