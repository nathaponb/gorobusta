package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/nathaponb/robusta-gosrv/internal/db"
	"github.com/nathaponb/robusta-gosrv/internal/repository/book"
	"github.com/nathaponb/robusta-gosrv/internal/repository/user"
	"gorm.io/gorm"
)

const (
	port = "8080"
)

type Config struct {
	UserRepo user.UserRepository
	BookRepo book.BookRepository
	Server   *http.Server
}

func NewServer() (*Config, error) {

	// instantiate http server
	app := Config{}

	srv := &http.Server{
		Addr:    fmt.Sprintf("127.0.0.1:%s", port),
		Handler: app.routes(),
	}

	app.Server = srv

	// connect to db
	conn := db.ConnectDB(os.Getenv("DSN"))
	if conn == nil {
		log.Println("Unable to establish connection to postgres!")
		os.Exit(1)
	}

	// instantiate repositories
	app.setupRepo(conn)

	return &app, nil
}

func (app *Config) start() {

	// start http server
	app.Server.ListenAndServe()
}

func (app *Config) setupRepo(connection *gorm.DB) {

	userRepo := user.NewPostgresRepository(connection)
	bookRepo := book.NewBookRepository(connection)

	app.UserRepo = userRepo
	app.BookRepo = bookRepo
}
