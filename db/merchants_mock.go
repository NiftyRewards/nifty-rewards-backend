package db

import (
	"github.com/go-pg/pg/v10"
	"log"
)

const FirstMerchantName = "merchant1"
const FirstMerchantID = 1
const SecondMerchantName = "merchant2"
const SecondMerchantID = 2

func MockGetMerchant(db *pg.DB) bool {
	merchant, err := GetMerchant(db, FirstMerchantName)
	if err != nil {
		log.Printf("[MockGetMerchant] GetMerchant err: %v", err)
	}

	if merchant.MerchantName != FirstMerchantName {
		log.Printf("merchant.MerchantName != FirstMerchantName")
		return false
	}
	if merchant.MerchantId != FirstMerchantID {
		log.Printf("merchant.MerchantId != FirstMerchantID")
		return false
	}

	log.Printf("MockGetMerchant passed")
	return true
}

func MockCreateMerchant(db *pg.DB) bool {
	_, err := CreateMerchant(db, Merchants{
		MerchantName: SecondMerchantName,
	})
	if err != nil {
		log.Printf("[MockCreateMerchant] CreateUser err: %v", err)
	}

	merchant, err := GetMerchant(db, SecondMerchantName)
	if err != nil {
		log.Printf("[MockCreateMerchant] GetMerchant: %v", err)
	}

	if merchant.MerchantName != SecondMerchantName {
		log.Printf("merchant.MerchantName != SecondMerchantName")
		return false
	}
	if merchant.MerchantId != SecondMerchantID {
		log.Printf("merchant.MerchantId != SecondMerchantID")
		return false
	}

	log.Printf("MockCreateMerchant passed")
	return true
}
