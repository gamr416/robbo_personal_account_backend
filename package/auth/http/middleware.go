package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	authDelegate auth.Delegate
}

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) (id string) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	id, err := h.delegate.ParseToken(headerParts[1], []byte(viper.GetString("auth.access_signing_key")))
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	return
}
