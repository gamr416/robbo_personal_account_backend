package http

import (
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
)

func RegisterHTTPEndpoints(router *gin.Engine, delegate auth.Delegate) {
	h := NewAuthHandler(delegate)

	authEndpoints := router.Group("/auth")
	{
		authEndpoints.POST("/sign-up", h.SignUp)
		authEndpoints.POST("/sign-in", h.SignIn)
	}
}
