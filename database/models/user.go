package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model

	Email          string `gorm:"unique"`
	Password       string
	Token          *string
	TokenUpdatedAt *time.Time
	Coins          int
}
