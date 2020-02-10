package utils

import (
	"fmt"
	"go-jwt-auth/pkg/encrypt"
	"net/http"
)

// JWTAthentication : Parse Claims from http.Request
func JWTAthentication(w http.ResponseWriter, r *http.Request) (*encrypt.Claims, error) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Not Authorized", http.StatusUnauthorized)
		return nil, fmt.Errorf("Empty access token")
	}

	claims, err := encrypt.DecryptJWT(token)
	if err != nil {
		http.Error(w, "Not Authorized", http.StatusUnauthorized)
		return nil, fmt.Errorf(err.Error())
	}

	return claims, nil
}
