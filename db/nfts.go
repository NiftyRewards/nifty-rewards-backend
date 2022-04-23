package db

import (
	"github.com/go-pg/pg/v10"
)

type Nfts struct {
	CollectionAddress string `pg:",pk" json:"collection_address"`
	CollectionName    string `json:"collection_name"`
	TotalSupply       int    `json:"total_supply"`
}

func GetNft(db *pg.DB, collectionAddress string) (*Nfts, error) {
	nft := &Nfts{}
	err := db.Model(nft).
		Where("nfts.collection_address = ?", collectionAddress).
		Select()

	return nft, err
}

func GetNfts(db *pg.DB) ([]*Nfts, error) {
	nfts := make([]*Nfts, 0)
	err := db.Model(&nfts).
		Select()

	return nfts, err
}

func CreateNft(db *pg.DB, req Nfts) (*Nfts, error) {
	_, err := db.Model(&req).Insert()
	if err != nil {
		return nil, err
	}

	nft := &Nfts{}
	err = db.Model(nft).
		Where("nfts.collection_address = ?", req.CollectionAddress).
		Select()

	return nft, err
}
