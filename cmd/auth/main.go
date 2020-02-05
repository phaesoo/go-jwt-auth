package main

import (
	"net/http"

	"go-jwt-auth/internal/app/auth"
	"go-jwt-auth/internal/app/auth/db"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(`{"message": "not found"}`))
}

func main() {
	db.InitDB()
	db.MigrateDB()
	db.SetInitialData()

	app := &auth.App{}
	app.Initialize()
	app.Run(":8000")
}
