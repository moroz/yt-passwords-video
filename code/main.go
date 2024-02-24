package main

import (
	"log"
	"net/http"

	"github.com/moroz/yt-passwords-video/code/config"
)

func main() {
	log.Printf("Listening on %s...", config.ListenOn)
	log.Fatal(http.ListenAndServe(config.ListenOn, nil))
}
