package handler

import (
	"log"
	"time"

	"encoding/json"
	"net/http"

	"go-jwt-auth/internal/app/auth/db"
	"go-jwt-auth/internal/app/auth/model"
	"go-jwt-auth/internal/app/common/utils"
	"go-jwt-auth/pkg/encrypt"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	AccessToken string `json:"access_token"`
}

// PostLogin : Login
func PostLogin(w http.ResponseWriter, r *http.Request) {
	login := Login{}
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	dbHandler := db.GetDB()
	user := model.User{}
	if result := dbHandler.Where("username = ? AND password = ?", login.Username, encrypt.EncryptSHA256(login.Password)).First(&user); result.RecordNotFound() {
		http.Error(w, "Not Athorized", http.StatusUnauthorized)
		return
	}

	accessToken, err := encrypt.EncryptJWT(user.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token := Token{AccessToken: accessToken}

	// response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

type Me struct {
	ID         uint      `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	DateJoined time.Time `json:"date_joined"`
	IssuedAt   int64     `json:"issued_at"`
	ExpiresAt  int64     `json:"expires_at"`
}

// GetMe : Get own information
func GetMe(w http.ResponseWriter, r *http.Request) {
	claims, err := utils.JWTAthentication(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "Not Athorized", http.StatusUnauthorized)
		return
	}

	dbHandler := db.GetDB()
	user := model.User{}
	if result := dbHandler.Where("username = ?", claims.Audience).First(&user); result.RecordNotFound() {
		http.Error(w, "Not Athorized", http.StatusUnauthorized)
		return
	}

	response := Me{
		ID:         user.ID,
		Username:   user.Username,
		Email:      user.Email,
		DateJoined: user.DateJoined,
		IssuedAt:   claims.IssuedAt,
		ExpiresAt:  claims.ExpiresAt,
	}

	// response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// GetRefresh : Refresh JWT token
func GetRefresh(w http.ResponseWriter, r *http.Request) {
	claims, err := utils.JWTAthentication(r)
	if err != nil {
		log.Println(err)
		http.Error(w, "Not Athorized", http.StatusUnauthorized)
		return
	}

	// give new access token
	accessToken, err := encrypt.EncryptJWT(claims.Audience)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	token := Token{AccessToken: accessToken}

	// response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}
