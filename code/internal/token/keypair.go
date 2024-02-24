package token

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/moroz/yt-passwords-video/code/config"
)

func readPublicKey() (ed25519.PublicKey, error) {
	filename := filepath.Join(config.JWTKeypairDirectory, config.JWTPubkeyFilename)
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("readPublicKey: %w", err)
	}
	block, _ := pem.Decode(bytes)
	parsed, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("readPublicKey: %w", err)
	}
	if key, ok := parsed.(ed25519.PublicKey); ok {
		return key, nil
	}
	return nil, errors.New("invalid public key format")
}

func readPrivateKey() (ed25519.PrivateKey, error) {
	filename := filepath.Join(config.JWTKeypairDirectory, config.JWTPrivkeyFilename)
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(bytes)
	parsed, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	if key, ok := parsed.(ed25519.PrivateKey); ok {
		return key, nil
	}
	return nil, errors.New("invalid private key format")
}

func ReadED25519Keypair() (pub ed25519.PublicKey, priv ed25519.PrivateKey, err error) {
	pub, err = readPublicKey()
	if err != nil {
		return
	}

	priv, err = readPrivateKey()
	return
}
