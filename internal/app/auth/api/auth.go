package handler

import (
	"time"
	"log"
	//"io/ioutil"
	"encoding/json"
	"net/http"

	"go-jwt-auth/internal/app/auth/db"

	//"github.com/gorilla/mux"
	"go-jwt-auth/internal/app/auth/model"
	"go-jwt-auth/pkg/encrypt"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	AccessToken string `json:"access_token"`
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	login := Login{}
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	dbHandler := db.GetDB()
	user := model.User{}
	if result := dbHandler.Where("username = ? AND password = ?", login.Username, encrypt.EncryptSHA256(login.Password)).First(&user); result.RecordNotFound() {
		http.Error(w, "Not Athorized", http.StatusUnauthorized)
		return
	}

	access_token, err := encrypt.EncryptJWT(user.Username)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	token := Token{}
	token.AccessToken = access_token

	// Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

type RespMe struct {
	Id uint `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	DateJoined time.Time `json:"date_joined"`
	IssuedAt int64 `json:"issued_at"`
	ExpiresAt int64 `json:"expires_at"`
}

func GetMe(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		http.Error(w, "Not Athorized", http.StatusUnauthorized)
		return
	}
	log.Println(token)

	claims, err := encrypt.DecryptJWT(token)
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

	response := RespMe{
		Id: user.ID,
		Username: user.Username,
		Email: user.Email,
		DateJoined: user.DateJoined,
		IssuedAt: claims.IssuedAt,
		ExpiresAt: claims.ExpiresAt,
	}

	// Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func GetRefresh(w http.ResponseWriter, r *http.Request) {
}
