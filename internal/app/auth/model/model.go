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
	Username    string `gorm:"size:150`
	Email       string `gorm:"size:254"`
	IsStaff     bool   `gorm:"not null"`
	IsActive    bool   `gorm:"not null"`
}
