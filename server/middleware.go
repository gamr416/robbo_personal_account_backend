package server

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/spf13/viper"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"strings"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		cookie, gerTokenErr := c.Cookie("refresh_token")
		if gerTokenErr == nil {
			c.Set("refresh_token", cookie)
		} else {
			c.Set("refresh_token", "")
		}
		if header == "" {
			//c.JSON(http.StatusUnauthorized, gin.H{
			//	"error": "token not found",
			//})
			//c.Abort()
			c.Set("user_id", "0")
			c.Set("user_role", models.Anonymous)
			c.Next()
			return
		}
		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 {
			//c.JSON(http.StatusUnauthorized, gin.H{
			//	"error": "invalid authorization header format",
			//})
			graphql.AddError(c, &gqlerror.Error{
				Path:    graphql.GetPath(c),
				Message: "invalid authorization header format",
				Extensions: map[string]interface{}{
					"code": "401",
				},
			})
			c.Abort()
			return
		}
		data, err := jwt.ParseWithClaims(headerParts[1], &models.UserClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte(viper.GetString("auth.access_signing_key")), nil
			})

		if err != nil {
			c.AbortWithStatusJSON(401, err)
			return
		}

		claims, ok := data.Claims.(*models.UserClaims)
		if !ok {
			//c.JSON(http.StatusUnauthorized, gin.H{
			//	"error": "token claims are not of type *StandardClaims",
			//})
			graphql.AddError(c, &gqlerror.Error{
				Path:    graphql.GetPath(c),
				Message: "token claims are not of type *StandardClaims",
				Extensions: map[string]interface{}{
					"code": "401",
				},
			})
			c.Abort()
			return
		}
		c.Set("user_id", claims.Id)
		c.Set("user_role", claims.Role)
		c.Next()
	}
}
