package api

import (
	"encoding/json"
	"errors"
	"github.com/go-pg/pg/v10"
	"golang-server/db"
	"log"
	"net/http"
)

type NftResponse struct {
	Success bool     `json:"success"`
	Error   string   `json:"err"`
	Nft     *db.Nfts `json:"nft"`
}

type NftsResponse struct {
	Success bool       `json:"success"`
	Error   string     `json:"err"`
	Nft     []*db.Nfts `json:"nft"`
}

func GetAllNfts(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	// get the database from context
	pgdb, ok := r.Context().Value("DB").(*pg.DB)
	if !ok {
		w = nftErrResponse(errors.New("could not get database from context"), w)
		return
	}
	// query for the nfts
	nfts, err := db.GetNfts(pgdb)
	if err != nil {
		w = nftErrResponse(err, w)
		return
	}

	// return a response
	res := &NftsResponse{
		Success: true,
		Error:   "",
		Nft:     nfts,
	}
	_ = json.NewEncoder(w).Encode(res)
	w.WriteHeader(http.StatusOK)
}

//func DeleteNftByAddressW3A(w http.ResponseWriter, r *http.Request) {
//	enableCors(&w)
//	nftAddress := chi.URLParam(r, "address_w3a")
//
//	// get the database from context
//	pgdb, ok := r.Context().Value("DB").(*pg.DB)
//	if !ok {
//		res := &NftResponse{
//			Success: false,
//			Error:   "could not get database from context",
//			Nft:    nil,
//		}
//		err := json.NewEncoder(w).Encode(res)
//		if err != nil {
//			log.Printf("err sending resopnse: %v\n", err)
//		}
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//	// query for the nft
//	err := db.DeleteNft(pgdb, nftAddress)
//	if err != nil {
//		res := &NftResponse{
//			Success: false,
//			Error:   err.Error(),
//			Nft:    nil,
//		}
//		err = json.NewEncoder(w).Encode(res)
//		if err != nil {
//			log.Printf("err sending response: %v\n", err)
//		}
//		w.WriteHeader(http.StatusBadRequest)
//		return
//	}
//
//	// return a response
//	res := &NftResponse{
//		Success: true,
//		Error:   "",
//		Nft:    nil,
//	}
//	_ = json.NewEncoder(w).Encode(res)
//	w.WriteHeader(http.StatusOK)
//}

func nftErrResponse(err error, w http.ResponseWriter) http.ResponseWriter {
	res := &NftResponse{
		Success: false,
		Error:   err.Error(),
		Nft:     nil,
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("err sending response: %v\n", err)
	}
	w.WriteHeader(http.StatusBadRequest)

	return w
}
