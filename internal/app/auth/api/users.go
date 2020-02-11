package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-jwt-auth/internal/app/auth/db"
	"github.com/go-jwt-auth/internal/app/auth/model"
	"github.com/go-jwt-auth/internal/app/common/utils"
)

// Get : Get all user information
func Get(w http.ResponseWriter, r *http.Request) {
	if _, err := utils.JWTAthentication(w, r); err != nil {
		log.Println(err)
		return
	}

	// get all user info from db
	dbHandler := db.GetDB()
	users := []model.User{}
	dbHandler = dbHandler.Select("id, username, email").Find(&users)
	if dbHandler.RecordNotFound() {
		http.NotFound(w, r)
		return
	}

	// response
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

// Post : Add new user
func Post(w http.ResponseWriter, r *http.Request) {
	_, err := utils.JWTAthentication(w, r)
	if err != nil {
		log.Println(err)
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
}

func PutUser(w http.ResponseWriter, r *http.Request) {
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
}
