package server

import (
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token not found",
			})
			c.Abort()
			return
		}
		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid authorization header format",
			})
			c.Abort()
			return
		}
		data, err := jwt.ParseWithClaims(headerParts[1], &models.UserClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(viper.GetString("auth.access_signing_key")), nil
			})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		claims, ok := data.Claims.(*models.UserClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "token claims are not of type *StandardClaims",
			})
			c.Abort()
			return
		}
		c.Set("user_id", claims.Id)
		c.Set("user_role", claims.Role)
		c.Next()
	}
}
