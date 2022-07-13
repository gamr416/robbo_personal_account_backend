package auth

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

type Delegate interface {
	SignIn(email, password string) (string, string, error)
	SignUp(email, password string) (string, string, error)
	ParseToken(token string, key []byte) (claims *models.UserClaims, err error)
	RefreshToken(token string) (newAccessToken string, err error)
}
