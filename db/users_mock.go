package db

import (
	"github.com/go-pg/pg/v10"
	"log"
)

var err error

const FirstAddressW3A = "0xUser1_w3a"
const FirstAddressB = "0xUser1_b"

const SecondAddressW3A = "0xUser2_w3a"
const SecondAddressB = "0xUser2_b"

const ThirdAddressW3A = "0xUser3_w3a"
const ThirdAddressB = "0xUser3_b"

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
		log.Printf("user.AddressW3a != SecondAddressW3A")
		return false
	}
	if user.Address_B != SecondAddressB {
		log.Printf("user.Address_B != SecondAddressB")
		return false
	}

	log.Printf("MockCreateUser passed")
	return true
}

func MockUpdateUser(db *pg.DB) bool {
	_, err := UpdateUser(db, &Users{
		AddressW3a: FirstAddressW3A,
		Address_B:  SecondAddressB,
	})

	if err != nil {
		log.Printf("[MockUpdateUser] UpdateUser err: %v", err)
	}

	user, err := GetUser(db, FirstAddressW3A)
	if err != nil {
		log.Printf("[MockUpdateUser] GetUser: %v", err)
	}

	if user.AddressW3a != FirstAddressW3A {
		log.Printf("user.AddressW3a != FirstAddressW3A")
		return false
	}
	if user.Address_B != SecondAddressB {
		log.Printf("user.Address_B != SecondAddressB")
		return false
	}

	log.Printf("MockUpdateUser passed")
	return true
}

func MockUpsertUserExists(db *pg.DB) bool {
	_, err := UpsertUser(db, &Users{
		AddressW3a: FirstAddressW3A,
		Address_B:  FirstAddressB,
	})

	if err != nil {
		log.Printf("[MockUpsertUserExists] UpdateUser err: %v", err)
	}

	user, err := GetUser(db, FirstAddressW3A)
	if err != nil {
		log.Printf("[MockUpsertUserExists] GetUser: %v", err)
	}

	if user.AddressW3a != FirstAddressW3A {
		log.Printf("user.AddressW3a != FirstAddressW3A")
		return false
	}
	if user.Address_B != FirstAddressB {
		log.Printf("user.Address_B != FirstAddressB")
		return false
	}

	log.Printf("MockUpsertUserExists passed")
	return true
}

func MockUpsertUserDoesNotExists(db *pg.DB) bool {
	_, err := UpsertUser(db, &Users{
		AddressW3a: ThirdAddressW3A,
		Address_B:  ThirdAddressB,
	})

	if err != nil {
		log.Printf("[MockUpsertUserDoesNotExists] UpdateUser err: %v", err)
	}

	user, err := GetUser(db, ThirdAddressW3A)
	if err != nil {
		log.Printf("[MockUpsertUserDoesNotExists] GetUser: %v", err)
	}

	if user.AddressW3a != ThirdAddressW3A {
		log.Printf("user.AddressW3a != ThirdAddressW3A")
		return false
	}
	if user.Address_B != ThirdAddressB {
		log.Printf("user.Address_B != ThirdAddressB")
		return false
	}

	log.Printf("MockUpsertUserDoesNotExists passed")
	return true
}
