package models

import "gorm.io/gorm"

type Solution struct {
	gorm.Model

	ProjectID uint `gorm:"column=project_id"`
	UserID    uint `gorm:"column=user_id"`
}
