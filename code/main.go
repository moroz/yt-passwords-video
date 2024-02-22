package main

import (
	"log"
	"net/http"
	"os"
)

func getEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func main() {
	listenOn := getEnvWithDefault("LISTEN_ON", "0.0.0.0:3000")

	log.Printf("Listening on %s...", listenOn)
	log.Fatal(http.ListenAndServe(listenOn, nil))
}
