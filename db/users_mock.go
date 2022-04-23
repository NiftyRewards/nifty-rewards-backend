package db

import (
	"github.com/go-pg/pg/v10"
	"log"
)

var err error

const FirstAddressW3A = "0x123"
const FirstAddressB = "0x456"
const SecondAddressW3A = "0xabc"
const SecondAddressB = "0xdef"

func MockGetUser(db *pg.DB) bool {
	user, err := GetUser(db, FirstAddressW3A)
	if err != nil {
		log.Printf("[TestGetUser] err: %v", err)
	}

	if user.AddressW3a != FirstAddressW3A {
		log.Printf("user.CollectionAddress != SecondAddressW3a")
		return false
	}
	if user.Address_B != FirstAddressB {
		log.Printf("user.CollectionName != FirstMerchantID")
		return false
	}

	log.Printf("MockGetUser passed")
	return true
}

func MockCreateUser(db *pg.DB) bool {
	_, err := CreateUser(db, Users{
		AddressW3a: SecondAddressW3A,
		Address_B:  SecondAddressB,
	})
	if err != nil {
		log.Printf("[MockCreateUser] CreateUser err: %v", err)
	}

	user, err := GetUser(db, SecondAddressW3A)
	if err != nil {
		log.Printf("[MockCreateUser] GetUser: %v", err)
	}

	if user.AddressW3a != SecondAddressW3A {
		log.Printf("user.CollectionAddress != SecondAddressW3a")
		return false
	}
	if user.Address_B != SecondAddressB {
		log.Printf("user.CollectionName != SecondMerchantID")
		log.Printf(user.Address_B)
		return false
	}

	log.Printf("MockCreateUser passed")
	return true
}
