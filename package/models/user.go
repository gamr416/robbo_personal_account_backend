package models

import "gorm.io/gorm"

type UserCore struct {
	ID       string
	Username string
	Password string
}

type UserDB struct {
	gorm.Model
	email    string `gorm:"not null;size:256"`
	password string `gorm:"not null;size:256"`
}
