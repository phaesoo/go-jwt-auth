package db

import (
	"go-jwt-auth/internal/app/auth/model"
	"go-jwt-auth/pkg/encrypt"
	"log"
	"os"
	"time"

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

// InitTestDB : Prepare for test DB
func InitTestDB() {
	var err error
	db, err = gorm.Open("sqlite3", "/tmp/users.db")
	if err != nil {
		log.Fatal("Could not connect database")
	}
}

// RemoveTestDB : Remove Test DB
func RemoveTestDB() {
	err := os.Remove("/tmp/users.db")
	if err != nil {
		log.Fatal("Failed to remove Test DB")
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

	db.Create(&model.User{
		Password:    encrypt.EncryptSHA256("password"),
		IsSuperuser: false,
		Username:    "test",
		Email:       "test@test.com",
		IsStaff:     false,
		IsActive:    true,
		DateJoined:  time.Now(),
	})
}
