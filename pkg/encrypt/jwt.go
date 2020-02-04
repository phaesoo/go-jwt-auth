package encrypt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var expPeriod = 5 * time.Minute
var secretKey = []byte("Secret!")

func EncryptJWT(username string) string {
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(expPeriod)

	claims := jwt.StandardClaims{
		Audience:  username,
		IssuedAt:  issuedAt.Unix(),
		ExpiresAt: expiresAt.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		panic("asdsad")
	} else {
		return tokenString
	}
}
