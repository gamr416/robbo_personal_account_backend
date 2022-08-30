package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/spf13/viper"
	"strings"
)

type Middleware struct {
	authDelegate auth.Delegate
}

const (
	authorizationHeader = "Authorization"
)

func (h *Middleware) UserIdentity(c *gin.Context) (id string, role models.Role, err error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", models.Anonymous, auth.ErrTokenNotFound
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return "", models.Anonymous, auth.ErrTokenNotFound
		return
	}

	claims, err := h.authDelegate.ParseToken(headerParts[1], []byte(viper.GetString("auth.access_signing_key")))
	if err != nil {
		return "", models.Anonymous, auth.ErrInvalidAccessToken
	}
	return claims.Id, claims.Role, nil
}
