package security

import (
	"fmt"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
)

// ParseTokenFromSignedTokenString parses JWT token
func ParseTokenFromSignedTokenString(tokenString string) (*jwt.Token, error) {

	publicKey, err := ioutil.ReadFile("keys/public.pem")
	if err != nil {
		return nil, fmt.Errorf("error reading public key file %v", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, fmt.Errorf("error parsing RSA public key %v", err)
	}

	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}

	return parsedToken, nil
}
