package handler

import (
	"log"
	"time"

	"encoding/json"
	"net/http"

	"github.com/go-jwt-auth/internal/app/auth/db"
	"github.com/go-jwt-auth/internal/app/auth/model"
	"github.com/go-jwt-auth/internal/app/common/utils"
	"github.com/go-jwt-auth/pkg/encrypt"
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
	// parse login from request body
	login := Login{}
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// select user from db
	dbHandler := db.GetDB()
	user := model.User{}
	if result := dbHandler.Where(
		"username = ? AND password = ?", login.Username, encrypt.EncryptSHA256(login.Password)
		).First(&user); result.RecordNotFound() {
		log.Println("Record not found from db: username, password")
		http.Error(w, "Not Athorized", http.StatusUnauthorized)
		return
	}

	// generate token
	var token Token
	if accessToken, err := encrypt.EncryptJWT(user.Username); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	} else {
		token = Token{AccessToken: accessToken}
	}

	// response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

type RespMe struct {
	ID         uint      `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	DateJoined time.Time `json:"date_joined"`
	IssuedAt   int64     `json:"issued_at"`
	ExpiresAt  int64     `json:"expires_at"`
}

// GetMe : Get own information
func GetMe(w http.ResponseWriter, r *http.Request) {
	// jwt authentication
	claims, err := utils.JWTAthentication(w, r)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// select user from db
	dbHandler := db.GetDB()
	user := model.User{}
	if result := dbHandler.Where("username = ?", claims.Audience).First(&user); result.RecordNotFound() {
		http.Error(w, "Not Athorized", http.StatusUnauthorized)
		return
	}

	response := RespMe{
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
	// jwt authentication
	claims, err := utils.JWTAthentication(w, r)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// generate new access token
	var token Token
	if accessToken, err := encrypt.EncryptJWT(claims.Audience); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	} else {
		token = Token{AccessToken: accessToken}
	}

	// response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}
