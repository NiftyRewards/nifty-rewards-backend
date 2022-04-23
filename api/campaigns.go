package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-pg/pg/v10"
	"golang-server/db"
	"log"
	"net/http"
	"time"
)

type StartCampaignsRequest struct {
	MerchantId        int             `json:"merchant_id"`
	CollectionAddress string          `json:"collection_address"`
	StartTime         int64           `json:"start_time"`
	EndTime           int64           `json:"end_time"`
	RewardDescs       []rewardDescReq `json:"rewards"`
}

type rewardDescReq struct {
	Description string `json:"description"`
	MaxQuantity int    `json:"quantity"`
}

type StartCampaignsResponse struct {
	Success     bool             `json:"success"`
	Error       string           `json:"err"`
	Campaign    db.Campaigns     `json:"campaign_id"`
	RewardDescs []rewardDescResp `json:"rewards"`
}

type rewardDescResp struct {
	Description string `json:"description"`
	RowsAdded   int    `json:"rows_added"`
}

func PostCampaigns(w http.ResponseWriter, r *http.Request) {
	var req StartCampaignsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error whilde decoding StartCampaignsRequest %v\n", err)
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
	gettedCampaign, err := db.CreateCampaign(pgdb, newCampaign)

	//Create rewards for all tokens in a collection
	campaignNft, err := db.GetNft(pgdb, req.CollectionAddress)
	if err != nil {
		w = campaignErrResponse(errors.New(fmt.Sprintf("error getting nft from collection_address(%s)", req.CollectionAddress)), w)
		return
	}

	// Prepare response
	var res StartCampaignsResponse

	// For each reward description, populate reward for all tokens in NFT collection
	for _, rewardDesc := range req.RewardDescs {
		tokenIdCounter := 1
		for tokenIdCounter = 1; tokenIdCounter < campaignNft.TotalSupply; tokenIdCounter++ {
			newReward := db.Rewards{
				MerchantId:        req.MerchantId,
				CollectionAddress: req.CollectionAddress,
				TokenId:           tokenIdCounter,
				Description:       rewardDesc.Description,
				MaxQuantity:       rewardDesc.MaxQuantity,
				QuantityUsed:      0,
			}
			err := db.CreateReward(pgdb, newReward)
			if err != nil {
				w = campaignErrResponse(errors.New(fmt.Sprintf("error while populating rewards, reward:(%+v)", newReward)), w)
				return
			}
		}

		// Prepare response
		res.RewardDescs = append(
			res.RewardDescs,
			rewardDescResp{
				Description: rewardDesc.Description,
				RowsAdded:   tokenIdCounter,
			})
	}

	// return a response
	res.Campaign = *gettedCampaign
	res.Success = true
	res.Error = ""
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("GetRewardsByCampaignId err3: %v\n", err)
	}
	w.WriteHeader(http.StatusOK)
}

func campaignErrResponse(err error, w http.ResponseWriter) http.ResponseWriter {
	res := &StartCampaignsResponse{
		Success:     false,
		Error:       err.Error(),
		RewardDescs: nil,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("err sending response: %v\n", err)
	}
	w.WriteHeader(http.StatusBadRequest)

	return w
}
