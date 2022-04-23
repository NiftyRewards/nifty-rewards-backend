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

func GetMerchantById(w http.ResponseWriter, r *http.Request) {
	merchantId, err := strconv.Atoi(chi.URLParam(r, "merchant_id"))
	if err != nil {
		log.Printf("GetRewardsByMerchantId err1: %v\n", err)
		w = rewardMerchantResponse(err, w)
		return
	}

	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &MerchantResponse{
			Success:  false,
			Error:    "could not get database from context",
			Merchant: nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("err sending resopnse: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// query for the merchant
	merchant, err := db.GetMerchantById(pgdb, merchantId)
	if err != nil {
		res := &MerchantResponse{
			Success:  false,
			Error:    err.Error(),
			Merchant: nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("err sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// return a response
	res := &MerchantResponse{
		Success:  true,
		Error:    "",
		Merchant: merchant,
	}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

func DeleteMerchantByName(w http.ResponseWriter, r *http.Request) {
	merchantName := chi.URLParam(r, "merchant_name")

	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &MerchantResponse{
			Success:  false,
			Error:    "could not get database from context",
			Merchant: nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("err sending resopnse: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// query for the merchant
	err := db.DeleteMerchant(pgdb, merchantName)
	if err != nil {
		res := &MerchantResponse{
			Success:  false,
			Error:    err.Error(),
			Merchant: nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("err sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// return a response
	res := &MerchantResponse{
		Success:  true,
		Error:    "",
		Merchant: nil,
	}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

func GetAllMerchants(w http.ResponseWriter, r *http.Request) {
	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		w = rewardMerchantResponse(errors.New("could not get database from context"), w)
		return
	}
	// get all merchants
	merchants, err := db.GetMerchants(pgdb)
	if err != nil {
		log.Printf("GetAllMerchants err2: %v\n", err)
		w = rewardMerchantResponse(err, w)
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

func rewardMerchantResponse(err error, w http.ResponseWriter) http.ResponseWriter {
	res := &MerchantResponse{
		Success:  false,
		Error:    err.Error(),
		Merchant: nil,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("err sending response: %v\n", err)
	}
	w.WriteHeader(http.StatusBadRequest)

	return w
}
