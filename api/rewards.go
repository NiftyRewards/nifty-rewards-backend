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
	enableCors(&w)
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

func PutRewardByRewardId(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	rewardId, err := strconv.Atoi(chi.URLParam(r, "reward_id"))
	if err != nil {
		log.Printf("PutRewardByRewardId err: %v\n", err)
		w = rewardErrResponse(err, w)
		return
	}

	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		w = rewardErrResponse(errors.New("could not get database from context"), w)
		return
	}

	// Get reward to determine quantityUsed
	reward, err := db.GetReward(pgdb, rewardId)

	// Reward does not exist
	if errors.Is(err, pg.ErrNoRows) {
		w = rewardErrResponse(errors.New(fmt.Sprintf("reward of reward_id(%d) does not exists", rewardId)), w)
		return
	}
	if err != nil {
		w = rewardErrResponse(errors.New(fmt.Sprintf("reward of reward_id(%d) does not exists", rewardId)), w)
		return
	}

	// Reward fully redeemed
	if reward.QuantityUsed >= reward.MaxQuantity {
		w = rewardErrResponse(errors.New("reward is fully redeemed"), w)
		return
	}

	// Reward can be redeemed
	updateReward, err := db.UpdateReward(pgdb, &db.Rewards{
		RewardId:          rewardId,
		MerchantId:        reward.MerchantId,
		CollectionAddress: reward.CollectionAddress,
		TokenId:           reward.TokenId,
		Description:       reward.Description,
		MaxQuantity:       reward.MaxQuantity,
		QuantityUsed:      reward.QuantityUsed + 1,
	})
	if err != nil {
		log.Printf("PutRewardByRewardId err2: %v\n", err)
		w = rewardErrResponse(err, w)
		return
	}

	// return a response
	res := &RewardResponse{
		Success: true,
		Error:   "",
		Reward:  updateReward,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("PutRewardByRewardId err3: %v\n", err)
	}
	w.WriteHeader(http.StatusOK)
}

type PostRewardsRequest struct {
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

type PostRewardResponse struct {
	Success     bool             `json:"success"`
	Error       string           `json:"err"`
	Campaign    db.Campaigns     `json:"campaign_id"`
	RewardDescs []rewardDescResp `json:"rewards"`
}

type rewardDescResp struct {
	Description string `json:"description"`
	RowsAdded   int    `json:"rows_added"`
}

func PostRewards(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var req PostRewardsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error whilde decoding PostRewardsRequest %v\n", err)
		return
	}

	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		w = campaignErrResponse(errors.New("could not get database from context"), w)
		return
	}

	log.Printf(fmt.Sprintf("HERE IS req %+v", req))

	//Create rewards for all tokens in a collection
	campaignNft, err := db.GetNft(pgdb, req.CollectionAddress)
	if err != nil {
		w = campaignErrResponse(errors.New(fmt.Sprintf("error getting nft from collection_address(%s)", req.CollectionAddress)), w)
		return
	}

	// Prepare response
	var res PostRewardResponse

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
				Approved:          false,
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
	res.Success = true
	res.Error = ""
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("GetRewardsByCampaignId err3: %v\n", err)
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
