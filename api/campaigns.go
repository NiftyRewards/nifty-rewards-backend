package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-pg/pg/v10"
	"golang-server/db"
	"log"
	"net/http"
	"strconv"
	"time"
)

type CampaignsResponse struct {
	Success   bool            `json:"success"`
	Error     string          `json:"err"`
	Campaigns []*db.Campaigns `json:"rewards"`
}

type ApproveCampaignRequest struct {
	MerchantId        int    `json:"merchant_id"`
	CollectionAddress string `json:"collection_address"`
	StartTime         int64  `json:"start_time"`
	EndTime           int64  `json:"end_time"`
}

type ApproveCampaignsResponse struct {
	Success  bool          `json:"success"`
	Error    string        `json:"err"`
	Campaign *db.Campaigns `json:"campaign_id"`
}

func ApproveCampaigns(w http.ResponseWriter, r *http.Request) {
	var req ApproveCampaignRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error while decoding ApproveCampaignRequest %v\n", err)
		return
	}

	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		w = campaignErrResponse(errors.New("could not get database from context"), w)
		return
	}

	log.Printf(fmt.Sprintf("HERE IS req %+v", req))

	// Create campaign
	newCampaign := db.Campaigns{
		MerchantId:        req.MerchantId,
		CollectionAddress: req.CollectionAddress,
		StartTime:         time.Unix(req.StartTime, 0),
		EndTime:           time.Unix(req.EndTime, 0),
	}
	createdCampaign, err := db.CreateCampaign(pgdb, newCampaign)

	// return a response
	res := &ApproveCampaignsResponse{
		Success:  true,
		Error:    "",
		Campaign: createdCampaign,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("GetCampaignsByMerchantId err3: %v\n", err)
		w = rewardErrResponse(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)

	// Update rewards to approved
	rewards, err := db.GetAllRewardsByMerchantIdCollectionAddress(pgdb, req.MerchantId, req.CollectionAddress)
	for _, rew := range rewards {
		newRew := &db.Rewards{
			RewardId:          rew.RewardId,
			MerchantId:        rew.MerchantId,
			CollectionAddress: rew.CollectionAddress,
			TokenId:           rew.TokenId,
			Description:       rew.Description,
			MaxQuantity:       rew.MaxQuantity,
			QuantityUsed:      rew.QuantityUsed,
			Approved:          true,
		}

		_, err := db.UpdateReward(pgdb, newRew)
		if err != nil {
			log.Printf("GetCampaignsByMerchantId err3: %v\n", err)
			w = rewardErrResponse(err, w)
			return
		}
	}
}

func GetAllCampaignsByMerchantId(w http.ResponseWriter, r *http.Request) {
	merchantId, err := strconv.Atoi(chi.URLParam(r, "merchant_id"))
	if err != nil {
		log.Printf("GetCampaignsByMerchantId err1: %v\n", err)
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
	campaigns, err := db.GetCampaignsByMerchantId(pgdb, merchantId)
	if err != nil {
		log.Printf("GetCampaignsByMerchantId err2: %v\n", err)
		w = rewardErrResponse(err, w)
		return
	}

	// return a response
	res := &CampaignsResponse{
		Success:   true,
		Error:     "",
		Campaigns: campaigns,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("GetCampaignsByMerchantId err3: %v\n", err)
		w = rewardErrResponse(err, w)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func campaignErrResponse(err error, w http.ResponseWriter) http.ResponseWriter {
	res := &ApproveCampaignsResponse{
		Success:  false,
		Error:    err.Error(),
		Campaign: nil,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("err sending response: %v\n", err)
	}
	w.WriteHeader(http.StatusBadRequest)

	return w
}
