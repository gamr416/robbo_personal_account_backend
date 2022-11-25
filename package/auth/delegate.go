package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
)

type Delegate interface {
	SignIn(email, password string, role uint) (accessToken string, refreshToken string, err error)
	SignUp(userHttp *models.UserHTTP) (accessToken string, refreshToken string, err error)
	//ParseToken(token string, key []byte) (claims *models.UserClaims, err error)
	UserIdentity(c *gin.Context) (id string, role models.Role, err error)
	UserAccess(currentRole models.Role, roles []models.Role) (err error)
	RefreshToken(refreshToken string) (newAccessToken string, err error)
}
