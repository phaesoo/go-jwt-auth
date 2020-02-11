package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-jwt-auth/internal/app/auth/db"
	"github.com/go-jwt-auth/internal/app/auth/model"
	"github.com/go-jwt-auth/internal/app/common/utils"
	"github.com/go-jwt-auth/pkg/account"
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

type postUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Post : Add new user
func Post(w http.ResponseWriter, r *http.Request) {
	_, err := utils.JWTAthentication(w, r)
	if err != nil {
		log.Println(err)
		return
	}

	user := postUser{}
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// validate request body
	if ok, err := account.IsValidUsername(user.Username); !ok {
		if err != nil {
			log.Println(err)
		}
		http.Error(w, "Invalid username", http.StatusBadRequest)
		return
	}

	if ok, err := account.IsValidPassword(user.Password); !ok {
		if err != nil {
			log.Println(err)
		}
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}

	if ok, err := account.IsValidEmail(user.Email); !ok {
		if err != nil {
			log.Println(err)
		}
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}

	// update db
	dbHandler := db.GetDB()
	newUser := model.User{
		Username:    user.Username,
		Password:    user.Password,
		Email:       user.Email,
		IsSuperuser: false,
		IsStaff:     false,
		IsActive:    true,
		DateJoined:  time.Now(),
	}

	if err := dbHandler.Create(&newUser).Error; err != nil {
		log.Println(err.Error())
		http.Error(w, "Already existed username: "+user.Username, http.StatusBadRequest)
		return
	}

	// response
	w.WriteHeader(http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
}

func PutUser(w http.ResponseWriter, r *http.Request) {
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
}
