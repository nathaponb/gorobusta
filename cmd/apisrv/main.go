package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/nathaponb/robusta-gosrv/internal/server"
)

func main() {

	// load local env to sys
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// instantiate http server
	server, err := server.NewServer()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// start the server
	server.Start()
}
