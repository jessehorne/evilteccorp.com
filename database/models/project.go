package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model

	Title       string
	Description string
	Answer      string
	Tags        string
	Reward      int
}
