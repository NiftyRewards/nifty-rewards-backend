package db

import (
	"github.com/go-pg/pg/v10"
	"log"
	"time"
)

const Campaign1Id = 1
const Campaign1Merchant = 1
const Campaign1Address = "0x9999999"

const Campaign2Id = 2
const Campaign2Merchant = 2
const Campaign2Address = "0x111111"

var StartTime = time.Now().Round(time.Microsecond)
var EndTime = time.Now().Round(time.Microsecond)

func MockGetCampaign(db *pg.DB) bool {
	campaign, err := GetCampaign(db, Campaign1Id)
	if err != nil {
		log.Printf("[MockGetCampaign] GetCampaign err: %v", err)
	}

	if campaign.CampaignId != Campaign1Id {
		log.Printf("campaign.CampaignId != Campaign1Id")
		return false
	}
	if campaign.CampaignId != Campaign1Merchant {
		log.Printf("campaign.CampaignId != Campaign1Merchant")
		return false
	}
	if campaign.CollectionAddress != Campaign1Address {
		log.Printf("campaign.CollectionAddress != Campaign1Address")
		return false
	}

	log.Printf("MockGetCampaign passed")
	return true
}

func MockCreateCampaign(db *pg.DB) bool {
	err = CreateCampaign(db, Campaigns{
		MerchantId:        Campaign2Merchant,
		CollectionAddress: Campaign2Address,
		StartTime:         StartTime,
		EndTime:           EndTime,
	})
	if err != nil {
		log.Printf("[MockCreateCampaign] CreateUser err: %v", err)
	}

	campaign, err := GetCampaign(db, Campaign2Id)
	if err != nil {
		log.Printf("[MockCreateCampaign] GetCampaign: %v", err)
	}

	if campaign.CampaignId != Campaign2Id {
		log.Printf("campaign.CampaignId != Campaign2Id")
		return false
	}
	if campaign.CampaignId != Campaign2Merchant {
		log.Printf("campaign.CampaignId != Campaign2Merchant")
		return false
	}
	if campaign.CollectionAddress != Campaign2Address {
		log.Printf("campaign.CollectionAddress != Campaign2Address")
		return false
	}
	if !campaign.StartTime.Equal(StartTime) {
		log.Printf("campaign.StartTime(%v) != StartTime(%v)", campaign.StartTime, StartTime)
		return false
	}
	if !campaign.EndTime.Equal(EndTime) {
		log.Printf("campaign.EndTime(%v) != EndTime(%v)", campaign.EndTime, EndTime)
		return false
	}

	log.Printf("MockCreateCampaign passed")
	return true
}
