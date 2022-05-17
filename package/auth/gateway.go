package auth

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	GetUser(email, password string) (user *models.UserCore, err error)
	CreateUser(user *models.UserCore) (id string, err error)
}
