package security

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT generates a new JWT
func GenerateJWT() (string, error) {

	// We are happy with the credentials, so build a token. We've given an expiry of 30 seconds
	// Using RSA256 Algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"user": "Christy Joseph Anoop",
		"exp":  time.Now().Add(time.Minute * time.Duration(5)).Unix(),
		"iat":  time.Now().Unix(),
	})

	// Read from RSA file
	privateKey, err := ioutil.ReadFile("keys/private.pem")
	if err != nil {
		return "", fmt.Errorf(`{"error":"Error reading from RSA file"}`)
	}

	// Generate Private Key
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		return "", fmt.Errorf(`{"error":"Unable to parse token"}`)
	}

	tokenString, err := token.SignedString(signKey)
	if err != nil {
		return "", fmt.Errorf(`{"error":"token generation failed"}`)
	}

	return tokenString, nil
}
