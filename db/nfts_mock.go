package db

import (
	"github.com/go-pg/pg/v10"
	"log"
)

const FirstNftAddress = "0xBAYC"
const FirstNftName = "bayc"
const FirstNftSupply = 10
const SecondNftAddress = "0xCryptoPunks"
const SecondNftName = "cryptopunks"
const SecondNftSupply = 5

func MockGetNft(db *pg.DB) bool {
	nft, err := GetNft(db, FirstNftAddress)
	if err != nil {
		log.Printf("[MockGetNft] GetNft err: %v", err)
	}

	if nft.CollectionAddress != FirstNftAddress {
		log.Printf("nft.CollectionAddress != FirstNftAddress")
		return false
	}
	if nft.CollectionName != FirstNftName {
		log.Printf("nft.CollectionName != FirstNftName")
		return false
	}
	if nft.TotalSupply != FirstNftSupply {
		log.Printf("nft.TotalSupply != FirstNftSupply")
		return false
	}

	log.Printf("MockGetNft passed")
	return true
}

func MockCreateNft(db *pg.DB) bool {
	_, err := CreateNft(db, Nfts{
		CollectionAddress: SecondNftAddress,
		CollectionName:    SecondNftName,
		TotalSupply:       SecondNftSupply,
	})
	if err != nil {
		log.Printf("[MockCreateNft] CreateNft err: %v", err)
	}

	nft, err := GetNft(db, SecondNftAddress)
	if err != nil {
		log.Printf("[MockCreateNft] GetNft: %v", err)
	}

	if nft.CollectionAddress != SecondNftAddress {
		log.Printf("nft.CollectionAddress != SecondNftAddress")
		return false
	}
	if nft.CollectionName != SecondNftName {
		log.Printf("nft.CollectionName != SecondNftName")
		return false
	}
	if nft.TotalSupply != SecondNftSupply {
		log.Printf("nft.TotalSupply != SecondNftSupply")
		return false
	}

	log.Printf("MockCreateNft passed")
	return true
}
