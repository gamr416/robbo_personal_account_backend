package models

import "gorm.io/gorm"

type UserDB struct {
	gorm.Model
	email string `gorm:"not null;size:256"`
}
