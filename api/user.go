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
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-pg/pg/v10"
)

type UserResponse struct {
	Success bool      `json:"success"`
	Error   string    `json:"err"`
	User    *db.Users `json:"user"`
}

type UsersResponse struct {
	Success bool        `json:"success"`
	Error   string      `json:"err"`
	User    []*db.Users `json:"user"`
}

type PostAddressBindResponse struct {
	Success bool      `json:"success"`
	Error   string    `json:"err"`
	User    *db.Users `json:"user"`
}

type GetNftsFromAccountResponse struct {
	Success bool    `json:"success"`
	Error   string  `json:"err"`
	tokens  []token `json:"tokens"`
}

type token struct {
	ContractAddress string `json:"contract_address"`
	TokenId         int    `json:"token_id"`
}

type TatumResponse []struct {
	ContractAddress string   `json:"contractAddress"`
	Balances        []string `json:"balances"`
	Metadata        []struct {
		URL      string `json:"url"`
		Metadata struct {
			Description string        `json:"description"`
			Image       string        `json:"image"`
			Name        string        `json:"name"`
			Attributes  []interface{} `json:"attributes"`
		} `json:"metadata"`
		TokenID string `json:"tokenId"`
	} `json:"metadata"`
}

func GetUserByAddressW3A(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	userAddress := chi.URLParam(r, "address_w3a")

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

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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
	// query for the users
	users, err := db.GetUsers(pgdb)
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
	res := &UsersResponse{
		Success: true,
		Error:   "",
		User:    users,
	}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

func DeleteUserByAddressW3A(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	userAddress := chi.URLParam(r, "address_w3a")

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
	err := db.DeleteUser(pgdb, userAddress)
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
		User:    nil,
	}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

func PostUpsertBinding(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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
	enableCors(&w)
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
	res, err := queryTatum(user.Address_B)
	if err != nil {
		w = userErrResponse(err, w)
		return
	}

	// return a response
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

func queryTatum(address_b string) ([]token, error) {
	log.Printf("@@@@@@@@@@@ address_b %s\n", address_b)
	// Call Tatum
	apiKey := os.Getenv("TATUM")
	url := "https://api-eu1.tatum.io/v3/nft/address/balance/MATIC/" + address_b

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-api-key", apiKey)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// snippet only
	var result TatumResponse
	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}

	log.Printf("@@@@@@@@@@@ result : %+v\n", result)
	if len(result) < 1 {
		return nil, errors.New("tatum result length is 0")
	}
	contractAddress := result[0].ContractAddress
	var nfts []token
	for _, meta := range result[0].Metadata {
		tokenId, err := strconv.Atoi(meta.TokenID)
		if err != nil {
			return nil, err
		}

		nfts = append(nfts, token{
			ContractAddress: contractAddress,
			TokenId:         tokenId,
		})
	}
	return nfts, nil
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
