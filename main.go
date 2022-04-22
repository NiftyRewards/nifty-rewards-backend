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
	db, err := db.NewDB()
	if err != nil {
		panic(err)
	}

	log.Printf("[main] We're up and running!")
	port := os.Getenv("PORT")

	go func() {
		router := api.NewAPI(db)
		err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
		if err != nil {
			log.Printf("err from  router: %v\n", err)
		}
	}()
}
