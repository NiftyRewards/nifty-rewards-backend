package db

import (
	"github.com/go-pg/pg/v10"
	"log"
)

var err error

const FirstAddressW3A = "0x123"
const FirstAddressB = "0x456"

func MockGetUser(db *pg.DB) bool {
	user, err := GetUser(db, FirstAddressW3A)
	if err != nil {
		log.Printf("[TestGetUser] err: %v", err)
	}

	if user.AddressW3a != FirstAddressW3A {
		log.Printf("user.AddressW3a != FirstAddressW3A")
		return false
	}
	if user.Address_B != FirstAddressB {
		log.Printf("user.Address_B != FirstAddressB")
		return false
	}

	log.Printf("MockGetUser passed")
	return true
}

func MockCreateUser(db *pg.DB) bool {
	_, err := CreateUser(db, FirstAddressW3A)
	log.Printf("[MockCreateUser] CreateUser err: %v", err)

	user, err := GetUser(db, FirstAddressW3A)
	if err != nil {
		log.Printf("[MockCreateUser] GetUser: %v", err)
	}

	if user.AddressW3a != FirstAddressW3A {
		log.Printf("user.AddressW3a != FirstAddressW3A")
		return false
	}
	if user.Address_B != FirstAddressB {
		log.Printf("user.AddressW3a != FirstAddressW3A")
		return false
	}

	log.Printf("MockCreateUser passed")
	return true
}
