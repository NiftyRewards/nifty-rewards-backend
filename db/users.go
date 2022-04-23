package db

import (
	"errors"
	"github.com/go-pg/pg/v10"
	"log"
)

type Users struct {
	AddressW3a string `pg:",pk" json:"address_w3a"`
	Address_B  string `json:"address_b"`
}

func GetUser(db *pg.DB, addressW3a string) (*Users, error) {
	user := &Users{}
	err := db.Model(user).
		Where("users.address_w3a = ?", addressW3a).
		Select()

	return user, err
}

func GetUsers(db *pg.DB) ([]*Users, error) {
	users := make([]*Users, 0)
	err := db.Model(&users).
		Select()

	return users, err
}

func CreateUser(db *pg.DB, req Users) (*Users, error) {
	_, err := db.Model(&req).Insert()
	if err != nil {
		return nil, err
	}

	user := &Users{}
	err = db.Model(user).
		Where("users.address_w3a = ?", req.AddressW3a).
		Select()

	return user, err
}

func UpdateUser(db *pg.DB, req *Users) (*Users, error) {
	_, err := db.Model(req).WherePK().Update()
	if err != nil {
		return nil, err
	}

	user := &Users{}
	err = db.Model(user).
		Where("users.address_w3a = ?", req.AddressW3a).
		Select()

	return nil, err
}

func UpsertUser(db *pg.DB, req *Users) (*Users, error) {
	// Try to get user
	user := &Users{}
	err := db.Model(user).
		Where("users.address_w3a = ?", req.AddressW3a).
		Select()

	// If does not exist
	if errors.Is(err, pg.ErrNoRows) {
		createUser, err := CreateUser(db, *req)
		if err != nil {
			log.Printf("[UpsertUser] UpsertUser err: %v", err)
		}
		return createUser, nil
	}

	// If exists, update user
	_, err = db.Model(req).WherePK().Update()
	if err != nil {
		return nil, err
	}

	err = db.Model(user).
		Where("users.address_w3a = ?", req.AddressW3a).
		Select()

	return nil, err
}
