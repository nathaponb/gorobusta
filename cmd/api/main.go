package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/nathaponb/robusta-gosrv/data"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	webPort = "8080"
)

var dbAttemptCount int

type Config struct {
	Repo data.Repository
}

func main() {

	// load local env to sys
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Errpr loading .env file")
	}

	// create app instance
	app := Config{}

	conn := connectDB(os.Getenv("DSN"))
	if conn == nil {
		log.Println("No connection with db is established!")
		os.Exit(1)
	}

	app.setupRepo(conn)

	// create http server
	srv := &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%s", webPort),
		Handler: app.routes(),
	}
	// start serrver
	srv.ListenAndServe()
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

func (app *Config) setupRepo(conn *gorm.DB) {
	db := data.NewPostgresRepository(conn)
	app.Repo = db
}
