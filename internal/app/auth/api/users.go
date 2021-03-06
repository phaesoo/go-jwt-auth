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
	"github.com/go-jwt-auth/pkg/encrypt"
	"github.com/gorilla/mux"
)

// Get : Get all user information
func Get(w http.ResponseWriter, r *http.Request) {
	if _, err := utils.JWTAthentication(w, r); err != nil {
		log.Println(err.Error())
		return
	}

	// get all user info from db
	dbHandler := db.GetDB()
	users := []model.User{}
	if result := dbHandler.Select("id, username, email").Find(&users); result.RecordNotFound() {
		log.Println("Record not found")
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
	if _, err := utils.JWTAthentication(w, r); err != nil {
		log.Println(err.Error())
		return
	}

	user := postUser{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println(err.Error())
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// validate request body
	if ok, err := account.IsValidUsername(user.Username); !ok {
		if err != nil {
			log.Println(err.Error())
		}
		http.Error(w, "Invalid username", http.StatusBadRequest)
		return
	}

	if ok, err := account.IsValidPassword(user.Password); !ok {
		if err != nil {
			log.Println(err.Error())
		}
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}

	if ok, err := account.IsValidEmail(user.Email); !ok {
		if err != nil {
			log.Println(err.Error())
		}
		http.Error(w, "Invalid email", http.StatusBadRequest)
		return
	}

	// update db
	newUser := model.User{
		Username:    user.Username,
		Password:    user.Password,
		Email:       user.Email,
		IsSuperuser: false,
		IsStaff:     false,
		IsActive:    true,
		DateJoined:  time.Now(),
	}

	dbHandler := db.GetDB()
	if err := dbHandler.Create(&newUser).Error; err != nil {
		log.Println(err.Error())
		http.Error(w, "Already existed username: "+user.Username, http.StatusBadRequest)
		return
	}

	// response
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// GetUser : Get target user info.
func GetUser(w http.ResponseWriter, r *http.Request) {
	if _, err := utils.JWTAthentication(w, r); err != nil {
		log.Println(err.Error())
		return
	}

	vars := mux.Vars(r)
	username := vars["username"]

	dbHandler := db.GetDB()
	user := model.User{}
	if result := dbHandler.Select("id, username, email").Where("username = ?", username).First(&user); result.RecordNotFound() {
		log.Println("Not found username: "+username)
		http.Error(w, "Not found username: "+username, http.StatusBadRequest)
		return
	}

	// response
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

type parseUpdate struct {
	Password    string `json"password"`
	NewPassword string `json"new_password"`
	Email       string `json"email"`
}

// PutUser : update user info
func PutUser(w http.ResponseWriter, r *http.Request) {
	if _, err := utils.JWTAthentication(w, r); err != nil {
		log.Println(err.Error())
		return
	}

	// might have to add additional logic for user validation for update permission
	vars := mux.Vars(r)
	username := vars["username"]

	parsed := parseUpdate{}
	err := json.NewDecoder(r.Body).Decode(&parsed)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// validate request body
	var newPassword string
	var email string

	if parsed.Password == "" {
		log.Println("Empty request body: password")
		http.Error(w, "Empty request body: password", http.StatusBadRequest)
		return
	}

	if parsed.NewPassword != "" {
		if ok, err := account.IsValidPassword(parsed.NewPassword); !ok {
			if err != nil {
				log.Println(err.Error())
			}
			log.Println("Invalid password format: new_password")
			http.Error(w, "Invalid password format: new_password", http.StatusBadRequest)
			return
		}
		newPassword = parsed.NewPassword
	}

	if parsed.Email != "" {
		if ok, err := account.IsValidEmail(parsed.Email); !ok {
			if err != nil {
				log.Println(err.Error())
			}
			log.Println("Invalid email format: email")
			http.Error(w, "Invalid email format: email", http.StatusBadRequest)
			return
		}
		email = parsed.Email
	}

	if (newPassword == "") && (email == "") {
		log.Println("Empty new password and email")
		http.Error(w, "No data for updating", http.StatusBadRequest)
		return
	}

	// check username and password from DB
	user := model.User{}

	dbHandler := db.GetDB()
	if result := dbHandler.Where("username = ? AND password = ?", username, encrypt.EncryptSHA256(parseUpdate.Password)).First(&user); result.RecordNotFound() {
		log.Println("Invalid username and password")
		http.Error(w, "Invalid username and password", http.StatusUnauthorized)
		return
	}

	// update password wth new password
	// update db with transaction
	user.Password = encrypt.EncryptSHA256(newPassword)
	if err := user.UpdateUser(dbHandler); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// response
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// DeleteUser : delete user from db
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if _, err := utils.JWTAthentication(w, r); err != nil {
		log.Println(err.Error())
		return
	}

	vars := mux.Vars(r)
	username := vars["username"]

	if err := model.DeleteUser(db.GetDB(), username); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// response
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
