package db

import (
	"github.com/go-pg/migrations/v8"
	"github.com/go-pg/pg/v10"
	"log"
	"os"
)

func NewDB() (*pg.DB, error) {
	var opts *pg.Options
	var err error

	if os.Getenv("ENV") == "PROD" {
		opts, err = pg.ParseURL(os.Getenv("DATABASE_URL"))
		if err != nil {
			return nil, err
		}
	} else {
		opts = &pg.Options{
			Addr:     "db:5432",
			User:     "postgres",
			Password: "admin",
		}
	}

	// connect to db
	db := pg.Connect(opts)

	// run migrations
	log.Printf("@@@@@ run migrations")
	collection := migrations.NewCollection()
	err = collection.DiscoverSQLMigrations("migrations")

	log.Printf("@@@@@ collections: %v", collection)
	if err != nil {
		return nil, err
	}

	// init db
	log.Printf("@@@@@ init")
	_, _, err = collection.Run(db, "init")
	if err != nil {
		return nil, err
	}

	log.Printf("@@@@@ up")
	oldVersion, newVersion, err := collection.Run(db, "up")
	log.Printf("@@@@@ oldVersion %v, newVersion %v", oldVersion, newVersion)
	if err != nil {
		return nil, err
	}
	if newVersion != oldVersion {
		log.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		log.Printf("version is %d\n", oldVersion)
	}

	//return the db connections
	return db, nil
}
