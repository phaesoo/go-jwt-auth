package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Password    string `gorm:"size:128;not null"`
	LastLogin   time.Time
	IsSuperuser bool   `gorm:"not null"`
	Username    string `gorm:"unique;not null;size:150"`
	Email       string `gorm:"size:254"`
	IsStaff     bool   `gorm:"not null"`
	IsActive    bool   `gorm:"not null"`
	DateJoined  time.Time   `gorm:"not null"`
}

func (u *User) UpdateUser(dbHandler *gorm.DB) error {
	tx := dbHandler.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Save(u).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}