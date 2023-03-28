package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var attemptCount int

func openDB(dsn string) (*gorm.DB, error) {

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func ConnectDB(dsn string) *gorm.DB {
	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Cannot connect to Postgres", err)
		} else {
			return connection
		}

		if attemptCount > 10 {
			return nil
		}

		log.Println("Wait for two seconds and try again!")
		time.Sleep(2 * time.Second)
		continue

	}
}
