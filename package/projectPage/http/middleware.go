package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strings"
)

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

	return h.authDelegate.ParseToken(headerParts[1], []byte(viper.GetString("auth.access_signing_key")))
}
