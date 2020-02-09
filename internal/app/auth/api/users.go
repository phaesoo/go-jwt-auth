package handler

import (
	"encoding/json"
	"net/http"

	"go-jwt-auth/internal/app/auth/db"
	"go-jwt-auth/internal/app/auth/model"
)

func Get(w http.ResponseWriter, r *http.Request) {
	dbHandler := db.GetDB()

	users := []model.User{}

	result := dbHandler.Select("id, username, email").Find(&users)
	if result.RecordNotFound() {
		http.NotFound(w, r)
		return
	}

	// response
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func Post(w http.ResponseWriter, r *http.Request) {
}

func GetUser(w http.ResponseWriter, r *http.Request) {
}

func PutUser(w http.ResponseWriter, r *http.Request) {
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
}
