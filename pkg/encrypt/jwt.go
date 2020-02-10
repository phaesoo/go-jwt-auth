package encrypt

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var expPeriod = 5 * time.Minute
var secretKey = []byte("Secret!")

type Claims struct {
	jwt.StandardClaims
}

func EncryptJWT(username string) (string, error) {
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(expPeriod)

	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			Audience:  username,
			IssuedAt:  issuedAt.Unix(),
			ExpiresAt: expiresAt.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("Token signed error")
	} else {
		return tokenString, nil
	}
}

func DecryptJWT(token string) (*Claims, error) {
	claims := &Claims{}

	// ParseWithClaims include expiration time checking
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return claims, err
	}

	return claims, nil
}
