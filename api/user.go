package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang-server/db"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-pg/pg/v10"
)

type UserResponse struct {
	Success bool      `json:"success"`
	Error   string    `json:"err"`
	User    *db.Users `json:"user"`
}

type PostAddressBindResponse struct {
	Success bool      `json:"success"`
	Error   string    `json:"err"`
	User    *db.Users `json:"user"`
}

func GetUserByAddressW3A(w http.ResponseWriter, r *http.Request) {
	userAddress := chi.URLParam(r, "addressW3A")

	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		res := &UserResponse{
			Success: false,
			Error:   "could not get database from context",
			User:    nil,
		}
		err := json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("err sending resopnse: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// query for the user
	user, err := db.GetUser(pgdb, userAddress)
	if err != nil {
		res := &UserResponse{
			Success: false,
			Error:   err.Error(),
			User:    nil,
		}
		err = json.NewEncoder(w).Encode(res)
		if err != nil {
			log.Printf("err sending response: %v\n", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// return a response
	res := &UserResponse{
		Success: true,
		Error:   "",
		User:    user,
	}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

func PostUpsertBinding(w http.ResponseWriter, r *http.Request) {
	addressW3a := chi.URLParam(r, "address_w3a")
	addressB := chi.URLParam(r, "address_to_bind")

	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		w = userErrResponse(errors.New("could not get database from context"), w)
		return
	}
	// query for the user
	user, err := db.UpsertUser(pgdb, &db.Users{
		AddressW3a: addressW3a,
		Address_B:  addressB,
	})
	if err != nil {
		w = userErrResponse(err, w)
		return
	}

	// return a response
	res := &UserResponse{
		Success: true,
		Error:   "",
		User:    user,
	}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

func GetNftsOfAccount(w http.ResponseWriter, r *http.Request) {
	// Get User Query Param
	addressW3a := chi.URLParam(r, "address_w3a")

	// Address Sanitation??uu

	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		w = userErrResponse(errors.New("could not get database from context"), w)
		return
	}

	// query for the user
	user, err := db.GetUser(pgdb, addressW3a)
	if err != nil {
		w = userErrResponse(err, w)
		return
	}

	address_b := user.Address_B

	// Call Tatum
	apiKey := os.Getenv("TATUM")
	url := "https://api-eu1.tatum.io/v3/nft/address/balance/MATIC/" + address_b

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-api-key", apiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	// return a response
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

func userErrResponse(err error, w http.ResponseWriter) http.ResponseWriter {
	res := &UserResponse{
		Success: false,
		Error:   err.Error(),
		User:    nil,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("err sending response: %v\n", err)
	}
	w.WriteHeader(http.StatusBadRequest)

	return w
}
