package models

import (
	"gorm.io/gorm"
	"strconv"
)

type UserCore struct {
	ID       string
	Email    string
	Password string
}

type UserDB struct {
	gorm.Model

	Email    string `gorm:"not null;size:256"`
	Password string `gorm:"not null;size:256"`
}

func (em *UserDB) ToCore() *UserCore {
	return &UserCore{
		ID:       strconv.FormatUint(uint64(em.ID), 10),
		Email:    em.Email,
		Password: em.Password,
	}
}

func (em *UserDB) FromCore(user *UserCore) {
	id, _ := strconv.ParseUint(user.ID, 10, 64)
	em.ID = uint(id)
	em.Email = user.Email
	em.Password = user.Password
}
