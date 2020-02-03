package main

import (
	// "log"
	"net/http"
	// "strconv"

	"go-jwt-auth/internal/app/auth"
	// "go-jwt-auth/internal/app/auth/api"
	// "github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func main() {
	auth.InitDB()
	auth.MigrateDB()

	auth.SetInitialData()

	app := &auth.App{}
	app.Initialize()
	app.Run(":8000")

	// r := mux.NewRouter()
	// api := r.PathPrefix("/api/v1").Subrouter()
	// api.HandleFunc("/", notFound)
	// api.HandleFunc("/user/{userID}/comment/{commentID}", params).Methods(http.MethodGet)

	// api2 := r.PathPrefix("/api").Subrouter()
	// api2.HandleFunc("/auth/login", handler.PostLogin).Methods(http.MethodPost)

	// log.Fatal(http.ListenAndServe(":8000", r))
}
