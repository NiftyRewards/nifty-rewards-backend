package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"golang-server/api"
	"golang-server/db"

	"log"
	"net/http"
	"os"
)

func main() {
	// If env is not prod, use read env from local.env
	if os.Getenv("ENV") != "PROD" {
		log.Printf("[env] Reading from local.env")
		err := godotenv.Load("local.env")
		if err != nil {
			log.Fatalf("Some err occured. Err: %s", err)
		}
	}

	// Init API and DB
	dbInst, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	log.Printf("[main] We're up and running!")

	go func() {
		router := api.NewAPI(dbInst)
		log.Printf("HERE1 %+v", router)
		err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
		log.Printf("HERE2")
		if err != nil {
			log.Printf("err from  router: %v\n", err)
		}
		log.Printf("HERE3")
	}()

	// USERS TEST
	if !db.MockGetUser(dbInst) {
		panic(err)
	}
	if !db.MockCreateUser(dbInst) {
		panic(err)
	}
	if !db.MockUpdateUser(dbInst) {
		panic(err)
	}
	if !db.MockUpsertUserExists(dbInst) {
		panic(err)
	}
	if !db.MockUpsertUserDoesNotExists(dbInst) {
		panic(err)
	}

	// MERCHANTS TEST
	if !db.MockGetMerchant(dbInst) {
		panic(err)
	}
	if !db.MockCreateMerchant(dbInst) {
		panic(err)
	}

	// NFTS TEST
	if !db.MockGetNft(dbInst) {
		panic(err)
	}
	if !db.MockCreateNft(dbInst) {
		panic(err)
	}

	// CAMPAIGNS TEST
	if !db.MockGetCampaign(dbInst) {
		panic(err)
	}
	if !db.MockCreateCampaign(dbInst) {
		panic(err)
	}

	// REWARSDS TEST
	if !db.MockGetReward(dbInst) {
		panic(err)
	}
	if !db.MockCreateReward(dbInst) {
		panic(err)
	}
}
