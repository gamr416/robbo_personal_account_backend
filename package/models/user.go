package models

import (
	"github.com/dgrijalva/jwt-go/v4"
	"gorm.io/gorm"
	"strconv"
)

type Role int

const (
	student Role = iota
	teacher
	parent
	freeListener
	unitAdmin
	superAdmin
)

type UserClaims struct {
	jwt.StandardClaims

	Id   string
	Role Role
}

type UserHttp struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     uint   `json:"role"`
}

type UserCore struct {
	ID       string
	Email    string
	Password string
	Role     Role
}

type UserDB struct {
	gorm.Model

	Email    string `gorm:"not null;size:256"`
	Password string `gorm:"not null;size:256"`
	Role     uint   `gorm:"not null"`
}

func (em *UserDB) ToCore() *UserCore {
	return &UserCore{
		ID:       strconv.FormatUint(uint64(em.ID), 10),
		Email:    em.Email,
		Password: em.Password,
		Role:     Role(em.Role),
	}
}

func (em *UserDB) FromCore(user *UserCore) {
	id, _ := strconv.ParseUint(user.ID, 10, 64)
	em.ID = uint(id)
	em.Email = user.Email
	em.Password = user.Password
	em.Role = uint(user.Role)
}

func (em *UserHttp) ToCore() *UserCore {
	return &UserCore{
		Email:    em.Email,
		Password: em.Password,
		Role:     Role(em.Role),
	}
}

func (em *UserHttp) FromCore(user *UserCore) {
	em.Email = user.Email
	em.Password = user.Password
	em.Role = uint(user.Role)
}
