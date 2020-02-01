package main

import (
	"log"
	"fmt"
	"net/http"
	"strconv"

	"go-jwt-auth/internal/app/auth/api"
	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func params(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)

	fmt.Println("pathParams", pathParams)

	w.Header().Set("Content-Type", "application/json")

	userID := -1
	var err error
	if val, ok := pathParams["userID"]; ok {
		if val == "" {
			fmt.Println("nil userID")
		}

		userID, err = strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message": "need a number"}`))
			return
		}
	}

	commentID := -1
    if val, ok := pathParams["commentID"]; ok {
        commentID, err = strconv.Atoi(val)
        if err != nil {
            w.WriteHeader(http.StatusInternalServerError)
            w.Write([]byte(`{"message": "need a number"}`))
            return
        }
    }

    query := r.URL.Query()
	location := query.Get("location")


	fmt.Println(query, location)
	
	w.Write([]byte(fmt.Sprintf(`{"userID": %d, "commentID": %d, "location": "%s" }`, userID, commentID, location)))
}

func main() {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/", notFound)
	api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)

	api2 := r.PathPrefix("/api").Subrouter()
	api2.HandleFunc("/", handler.PostLogin).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8000", r))
}