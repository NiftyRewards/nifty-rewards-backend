package api

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-pg/pg/v10"
	"golang-server/db"
	"log"
	"net/http"
)

type CreateUserRequest struct {
	ID       int64  `pg:",pk" json:"id"`
	Username string `json:"username"`
}

type UserResponse struct {
	Success bool      `json:"success"`
	Error   string    `json:"err"`
	User    *db.Users `json:"user"`
}

type CreateuserResponse struct {
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
