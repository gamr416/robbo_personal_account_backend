package auth

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

type Delegate interface {
	SignIn(email, password string, role uint) (accessToken string, refreshToken string, err error)
	SignUp(userHttp *models.User, role models.Role) (accessToken string, refreshToken string, err error)
	ParseToken(token string, key []byte) (claims *models.UserClaims, err error)
	RefreshToken(token string) (newAccessToken string, err error)
}
