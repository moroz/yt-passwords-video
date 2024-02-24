package main

import (
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/moroz/yt-passwords-video/code/config"
	"github.com/moroz/yt-passwords-video/code/handler"

	_ "github.com/lib/pq"
)

func initDatabase() *sqlx.DB {
	return sqlx.MustConnect("postgres", config.DatabaseURL)
}

func main() {
	router := handler.Router(initDatabase())

	log.Printf("Listening on %s...", config.ListenOn)
	log.Fatal(http.ListenAndServe(config.ListenOn, router))
}
