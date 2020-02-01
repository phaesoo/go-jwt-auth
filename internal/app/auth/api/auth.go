package handler

import (
	"fmt"
	"net/http"
	//"github.com/gorilla/mux"
)


func PostLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)

	decoder := json.Decoder(r.Body)
	fmt.Println(decoder)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func GetMe(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func GetRefresh(w http.ResponseWriter, r *http.Request) {
}