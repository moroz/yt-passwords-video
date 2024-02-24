package config

import (
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

var ListenOn = GetEnvWithDefault("LISTEN_ON", "0.0.0.0:3000")
var RootPath, _ = os.Getwd()
var defaultKeypairDirectory = filepath.Join(RootPath, "priv")

var JWTKeypairDirectory = GetEnvWithDefault("JWT_KEYPAIR_DIRECTORY", defaultKeypairDirectory)
var JWTPrivkeyFilename = GetEnvWithDefault("JWT_PRIVKEY_FILENAME", "access.key")
var JWTPubkeyFilename = GetEnvWithDefault("JWT_PRIVKEY_FILENAME", "access.pub")
