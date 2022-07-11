package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/spf13/viper"
	"strings"
)

type AuthMiddleware struct {
	authDelegate auth.Delegate
}

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) (id string, err error) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		return "", errors.New("token not found")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		return "", errors.New("token not found")
		return
	}

	return h.delegate.ParseToken(headerParts[1], []byte(viper.GetString("auth.access_signing_key")))
}
