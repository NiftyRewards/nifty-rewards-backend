package db

import (
	"github.com/go-pg/pg/v10"
)

type Rewards struct {
	RewardId          int    `pg:",pk" json:"reward_id"`
	MerchantId        int    `json:"merchant_id"`
	CollectionAddress string `json:"collection_address"`
	TokenId           *int   `json:"token_id"`
	Description       string `json:"Description"`
	MaxQuantity       int    `json:"max_quantity"`
	QuantityUsed      int    `json:"quantity_used"`
	Approved          bool   `json:"approved"`
}

func GetReward(db *pg.DB, RewardId int) (*Rewards, error) {
	reward := &Rewards{}
	err := db.Model(reward).
		Where("rewards.reward_id = ?", RewardId).
		Select()

	return reward, err
}

func GetRewards(db *pg.DB) ([]*Rewards, error) {
	rewards := make([]*Rewards, 0)
	err := db.Model(&rewards).
		Select()

	return rewards, err
}

func GetRewardsByMerchantId(db *pg.DB, merchantId int) ([]*Rewards, error) {
	rewards := make([]*Rewards, 0)
	err := db.Model(&rewards).
		Where("rewards.merchant_id = ?", merchantId).
		Select()

	return rewards, err
}

func CreateReward(db *pg.DB, req Rewards) error {
	_, err := db.Model(&req).Insert()
	if err != nil {
		return err
	}

	return nil
}

func UpdateReward(db *pg.DB, req *Rewards) (*Rewards, error) {
	_, err := db.Model(req).WherePK().Update()
	if err != nil {
		return nil, err
	}

	reward := &Rewards{}
	err = db.Model(reward).
		Where("rewards.reward_id = ?", req.RewardId).
		Select()

	return reward, err
}

func GetAllRewardsByMerchantIdCollectionAddress(db *pg.DB, merchantId int, collectionAddress string) ([]*Rewards, error) {
	rewards := make([]*Rewards, 0)
	err := db.Model(&rewards).
		Where("rewards.merchant_id = ?", merchantId).
		Where("rewards.collection_address = ?", collectionAddress).
		Select()

	return rewards, err
}

func GetAllRewardsByMerchantIdCollectionAddressTokenId(db *pg.DB, merchantId, tokenId int, collectionAddress string) ([]*Rewards, error) {
	rewards := make([]*Rewards, 0)
	err := db.Model(&rewards).
		Where("rewards.merchant_id = ?", merchantId).
		Where("rewards.token_id = ?", tokenId).
		Where("rewards.collection_address = ?", collectionAddress).
		Select()

	return rewards, err
}
