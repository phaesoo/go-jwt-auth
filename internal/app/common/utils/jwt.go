package utils

import (
	"fmt"
	"go-jwt-auth/pkg/encrypt"
	"net/http"
)

// JWTAthentication : Parse Claims from http.Request
func JWTAthentication(r *http.Request) (*encrypt.Claims, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		return nil, fmt.Errorf("Empty access token")
	}

	claims, err := encrypt.DecryptJWT(token)
	if err != nil {
		return nil, fmt.Errorf("Error while decrypt token")
	}

	return claims, nil
}
