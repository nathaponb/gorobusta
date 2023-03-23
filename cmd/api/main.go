package main

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbAttemptCount int

type Config struct{}

func main() {
	// load local env to sys

	// connect db

	// populate server

	// start serrver
}

func connectDB(dsn string) *gorm.DB {

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("Connot connect to postgres", err)
		} else {
			return connection
		}

		// linit to 10 attempts
		if dbAttemptCount > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Wait for two seconds and try again!")
		time.Sleep(2 * time.Second)
		continue
	}

}

func openDB(dsn string) (*gorm.DB, error) {

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return conn, nil
}
