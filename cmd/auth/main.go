package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct{
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "get called"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "post called"}`))
	case "PUT":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "put called"}`))
	case "DELETE":
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message": "delete called"}`))
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"message": "not found"}`))
	}
}

// https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj
func main() {

	mux.NewRouter()
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}