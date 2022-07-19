package users

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type Gateway interface {
	GetUsersByRole(role *models.Role) (users *[]models.UserCore, err error)
	UpdateUser(userCore *models.UserCore) (err error)
	DeleteUser(userId int) (err error)
	GetUserById(userId int) (user *models.UserCore, err error)
	CreateUser(user *models.UserCore) (id string, err error)
}
