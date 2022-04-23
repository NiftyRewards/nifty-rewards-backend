package db

import (
	"github.com/go-pg/pg/v10"
	"log"
)

const Reward1Id = 1
const Reward1Merchant = 1
const Reward1Address = "0x9999999"
const Reward1TokenId = 555
const Reward1Desc = "rewards1_desc"
const Reward1MaxQuantity = 4
const Reward1QuantityUsed = 0

const Reward2Id = 2
const Reward2Merchant = 2
const Reward2Address = "0x111111"
const Reward2TokenId = 111
const Reward2Desc = "rewards2_desc"
const Reward2MaxQuantity = 2
const Reward2QuantityUsed = 1

func MockGetReward(db *pg.DB) bool {
	reward, err := GetReward(db, Reward1Id)
	if err != nil {
		log.Printf("[MockGetReward] GetReward err: %v", err)
	}

	if reward.RewardId != Reward1Id {
		log.Printf("reward.RewardId != Reward1Id")
		return false
	}
	if reward.MerchantId != Reward1Merchant {
		log.Printf("reward.MerchantId != Reward1Merchant")
		return false
	}
	if reward.CollectionAddress != Reward1Address {
		log.Printf("reward.CollectionAddress != Reward1Address")
		return false
	}
	if reward.TokenId != Reward1TokenId {
		log.Printf("reward.TokenId != Reward1TokenId")
		return false
	}
	if reward.Description != Reward1Desc {
		log.Printf("reward.Description != Reward1Desc")
		return false
	}
	if reward.MaxQuantity != Reward1MaxQuantity {
		log.Printf("reward.MaxQuantity != Reward1MaxQuantity")
		return false
	}
	if reward.QuantityUsed != Reward1QuantityUsed {
		log.Printf("reward.QuantityUsed != Reward1QuantityUsed")
		return false
	}

	log.Printf("MockGetReward passed")
	return true
}

func MockCreateReward(db *pg.DB) bool {
	_, err := CreateReward(db, Rewards{
		MerchantId:        Reward2Merchant,
		CollectionAddress: Reward2Address,
		TokenId:           Reward2TokenId,
		Description:       Reward2Desc,
		MaxQuantity:       Reward2MaxQuantity,
		QuantityUsed:      Reward2QuantityUsed,
	})
	if err != nil {
		log.Printf("[MockCreateReward] CreateUser err: %v", err)
	}

	reward, err := GetReward(db, Reward2Id)
	if err != nil {
		log.Printf("[MockCreateReward] GetReward: %v", err)
	}

	if reward.RewardId != Reward2Id {
		log.Printf("reward.RewardId != Reward1Id")
		return false
	}
	if reward.MerchantId != Reward2Merchant {
		log.Printf("reward.MerchantId != Reward2Merchant")
		return false
	}
	if reward.CollectionAddress != Reward2Address {
		log.Printf("reward.CollectionAddress != Reward2Address")
		return false
	}
	if reward.TokenId != Reward2TokenId {
		log.Printf("reward.TokenId != Reward2TokenId")
		return false
	}
	if reward.Description != Reward2Desc {
		log.Printf("reward.Description != Reward2Desc")
		return false
	}
	if reward.MaxQuantity != Reward2MaxQuantity {
		log.Printf("reward.MaxQuantity != Reward2MaxQuantity")
		return false
	}
	if reward.QuantityUsed != Reward2QuantityUsed {
		log.Printf("reward.QuantityUsed != Reward2QuantityUsed")
		return false
	}

	log.Printf("MockCreateReward passed")
	return true
}
