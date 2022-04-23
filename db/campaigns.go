package db

import (
	"github.com/go-pg/pg/v10"
	"time"
)

type Campaigns struct {
	CampaignId        int       `pg:",pk" json:"campaign_id"`
	MerchantId        int       `json:"merchant_id"`
	CollectionAddress string    `json:"collection_address"`
	StartTime         time.Time `json:"start_time"`
	EndTime           time.Time `json:"end_time"`
}

func GetCampaign(db *pg.DB, CampaignId int) (*Campaigns, error) {
	campaign := &Campaigns{}
	err := db.Model(campaign).
		Where("campaigns.campaign_id = ?", CampaignId).
		Select()

	return campaign, err
}

func GetCampaigns(db *pg.DB) ([]*Campaigns, error) {
	campaigns := make([]*Campaigns, 0)
	err := db.Model(&campaigns).
		Select()

	return campaigns, err
}

func CreateCampaign(db *pg.DB, req Campaigns) (*Campaigns, error) {
	_, err := db.Model(&req).Insert()
	if err != nil {
		return nil, err
	}

	campaign := &Campaigns{}
	err = db.Model(campaign).
		Where("campaigns.collection_address = ?", req.CollectionAddress).
		Where("campaigns.merchant_id = ?", req.MerchantId).
		Where("campaigns.start_time = ?", req.StartTime).
		Where("campaigns.end_time = ?", req.EndTime).
		Select()
	if err != nil {
		return nil, err
	}

	return campaign, nil
}
