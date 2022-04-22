package db

import (
	"github.com/go-pg/pg/v10"
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

func CreateUser(db *pg.DB, addressW3a string) (*Users, error) {
	req := Users{
		AddressW3a: addressW3a,
	}

	_, err := db.Model(&req).Insert()
	if err != nil {
		return nil, err
	}

	user := &Users{}
	err = db.Model(user).
		Where("users.address_w3a = ?", addressW3a).
		Select()

	return user, err
}
