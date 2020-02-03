package handler

import (
	//"io/ioutil"
	"encoding/json"
	"net/http"
	//"github.com/gorilla/mux"
)

type Login struct {
	Username string `json:"username`
	Password string `json:"password"`
}

func PostLogin(w http.ResponseWriter, r *http.Request) {
	var login Login
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

func GetMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func GetRefresh(w http.ResponseWriter, r *http.Request) {
}