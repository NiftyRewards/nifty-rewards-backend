package db

import (
	"github.com/go-pg/pg/v10"
)

type Merchants struct {
	MerchantId   int    `pg:",pk" json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
}

func GetMerchant(db *pg.DB, merchantName string) (*Merchants, error) {
	merchant := &Merchants{}
	err := db.Model(merchant).
		Where("merchants.merchant_name = ?", merchantName).
		Select()

	return merchant, err
}

func GetMerchants(db *pg.DB) ([]*Merchants, error) {
	merchants := make([]*Merchants, 0)
	err := db.Model(&merchants).
		Select()

	return merchants, err
}

func CreateMerchant(db *pg.DB, req Merchants) (*Merchants, error) {
	_, err := db.Model(&req).Insert()
	if err != nil {
		return nil, err
	}

	merchant := &Merchants{}
	err = db.Model(merchant).
		Where("merchants.merchant_name = ?", req.MerchantName).
		Select()

	return merchant, err
}
