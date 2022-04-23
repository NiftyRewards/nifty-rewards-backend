package main

import (
	"fmt"
	"golang-server/api"
	"golang-server/db"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf("HELLO WORLD")

	// Init API and DB
	dbInst, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	log.Printf("[main] We're up and running!")

	go func() {
		router := api.NewAPI(dbInst)
		err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
		if err != nil {
			log.Printf("err from  router: %v\n", err)
		}
	}()
	// USERS TEST
	if !db.MockGetUser(dbInst) {
		panic(err)
	}
	if !db.MockCreateUser(dbInst) {
		panic(err)
	}
	if !db.MockGetMerchant(dbInst) {
		panic(err)
	}
	if !db.MockCreateMerchant(dbInst) {
		panic(err)
	}
	if !db.MockGetNft(dbInst) {
		panic(err)
	}
	if !db.MockCreateNft(dbInst) {
		panic(err)
	}
}
