package db

import (
	"log"
	"time"
	"go-jwt-auth/internal/app/auth/model"
	"go-jwt-auth/pkg/encrypt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func InitDB() {
	var err error
	db, err = gorm.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal("Could not connect database")
	}
}

func GetDB() *gorm.DB {
	return db
}

func MigrateDB() {
	db.AutoMigrate(&model.User{})
}

func SetInitialData() {
	db.Create(&model.User{
		Password:    encrypt.EncryptSHA256("password"),
		IsSuperuser: true,
		Username:    "admin",
		Email:       "admin@admin.com",
		IsStaff:     true,
		IsActive:    true,
		DateJoined:  time.Now(),
	})
}