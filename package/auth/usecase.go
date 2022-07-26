package auth

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"time"
)

type UseCase interface {
	SignIn(email, password string, role uint) (accessToken string, refreshToken string, err error)
	SignUp(userCore *models.User, role models.Role) (accessToken string, refreshToken string, err error)
	ParseToken(token string, key []byte) (claims *models.UserClaims, err error)
	RefreshToken(refreshToken string) (newAccessToken string, err error)
	GenerateToken(user *models.UserToken, duration time.Duration, signingKey []byte) (token string, err error)
}
