package main

import (
	"github.com/go-jwt-auth/internal/app/auth"
	"github.com/go-jwt-auth/internal/app/auth/db"
)

func main() {
	db.InitDB()
	db.MigrateDB()
	db.SetInitialData()

	app := &auth.App{}
	app.Initialize()
	app.Run(":8000")
}
