package models

import (
	"github.com/dgrijalva/jwt-go/v4"
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
<<<<<<< HEAD

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
=======
type UserToken struct {
	ID   string
	Role Role
}

type User struct {
	Email      string `json:"email"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Password   string `json:"password"`
	Role       Role   `json:"role"`
>>>>>>> b51413c19b53a2b40776b2746be8d694d6f8e40e
}
