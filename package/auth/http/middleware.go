package http

import (
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
)

type AuthMiddleware struct {
	delegate auth.Delegate
}

//func NewAuthMiddleware(delegate auth.Delegate) gin.HandlerFunc {
//	return (&AuthMiddleware{
//		delegate: delegate,
//	}).Handle
//}

//func (m *AuthMiddleware) Handle(c *gin.Context) {
//	authHeader := c.GetHeader("Authorization")
//	if authHeader == "" {
//		c.AbortWithStatus(http.StatusUnauthorized)
//		return
//	}
//
//	headerParts := strings.Split(authHeader, " ")
//	if len(headerParts) != 2 {
//		c.AbortWithStatus(http.StatusUnauthorized)
//		return
//	}
//
//	if headerParts[0] != "Bearer" {
//		c.AbortWithStatus(http.StatusUnauthorized)
//		return
//	}
//
//	user, err := m.delegate.ParseToken(c.Request.Context(), headerParts[1])
//	if err != nil {
//		status := http.StatusInternalServerError
//		if err == auth.ErrInvalidAccessToken {
//			status = http.StatusUnauthorized
//		}
//
//		c.AbortWithStatus(status)
//		return
//	}
//
//	c.Set(auth.CtxUserKey, user)
//}
