package config

import (
	"log"
	"os"
	"path/filepath"
)

func GetEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func MustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("FATAL: Environment variable %s is not set!", key)
	}
	return value
}

var ListenOn = GetEnvWithDefault("LISTEN_ON", "0.0.0.0:3000")
var RootPath, _ = os.Getwd()
var defaultKeypairDirectory = filepath.Join(RootPath, "priv")

var JWTKeypairDirectory = GetEnvWithDefault("JWT_KEYPAIR_DIRECTORY", defaultKeypairDirectory)
var JWTPrivkeyFilename = GetEnvWithDefault("JWT_PRIVKEY_FILENAME", "access.key")
var JWTPubkeyFilename = GetEnvWithDefault("JWT_PUBKEY_FILENAME", "access.pub")

var DatabaseURL = MustGetEnv("DATABASE_URL")
