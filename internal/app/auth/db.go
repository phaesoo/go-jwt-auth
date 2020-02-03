package auth

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DBWrapper struct {
	DB *gorm.DB
}

var MyDB = &DBWrapper{}

func InitDB() *DBWrapper {
	db, err := gorm.Open("sqlite3", "auth.db")
	if err != nil {
		log.Fatal("Could not connect database")
	}

	MyDB.DB = db
	return MyDB
}

func GetDB() *DBWrapper {
	return MyDB
}

func (db *DBWrapper) MigrateDB() *DBWrapper {
	return db
}
