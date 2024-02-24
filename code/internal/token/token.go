package token

import (
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/moroz/yt-passwords-video/code/types"
	"golang.org/x/crypto/ed25519"
)

var (
	pubkey  ed25519.PublicKey
	privkey ed25519.PrivateKey
	err     error
)

func init() {
	pubkey, privkey, err = ReadED25519Keypair()
	if err != nil {
		log.Fatal(err)
	}
}

func keyfunc(token *jwt.Token) (any, error) {
	return pubkey, nil
}

const JWTValidity = 3600 * time.Second
const issuer = "ACME Corp."

func buildClaimsForUser(user *types.User) jwt.RegisteredClaims {
	now := time.Now()

	return jwt.RegisteredClaims{
		Issuer:    issuer,
		Subject:   strconv.Itoa(user.ID),
		IssuedAt:  jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(JWTValidity)),
	}
}

func IssueTokenForUser(user *types.User) (string, error) {
	claims := buildClaimsForUser(user)
	token := jwt.NewWithClaims(&jwt.SigningMethodEd25519{}, claims)
	return token.SignedString(privkey)
}

func DecodeToken(tokenString string) (claims jwt.RegisteredClaims, err error) {
	_, err = jwt.ParseWithClaims(tokenString, &claims, keyfunc, jwt.WithIssuer(issuer))
	return
}
