package auth

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edxApi"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

type Delegate interface {
	SignIn(userHttp *models.UserHttp) (accessToken string, refreshToken string, err error)
	SignUp(userHttp *models.UserHttp) (accessToken string, refreshToken string, err error)
	ParseToken(token string, key []byte) (claims *models.UserClaims, err error)
	RefreshToken(token string) (newAccessToken string, err error)
	Registration(userForm *edxApi.RegistrationForm) (err error)
	Login(email, password string) (err error)
}
