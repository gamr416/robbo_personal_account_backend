package auth

import "github.com/skinnykaen/robbo_student_personal_account.git/package/models"

type UseCase interface {
	SignIn(email, password string) (err error)
	SignUp(email, password string) (err error)
	ParseToken(accessToken string) (*models.UserCore, error)
}
