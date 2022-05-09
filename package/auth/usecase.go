package auth

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	SignIn(email, password string) (string, error)
	SignUp(email, password string) error
	ParseToken(accessToken string) (*models.UserCore, error)
}
