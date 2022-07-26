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
}
