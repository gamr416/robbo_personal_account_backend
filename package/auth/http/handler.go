package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"net/http"
)

type Handler struct {
	delegate auth.Delegate
}

func NewAuthHandler(authDelegate auth.Delegate) Handler {
	return Handler{
		delegate: authDelegate,
	}
}

func (h *Handler) InitAuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}
}

type signInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) SignIn(c *gin.Context) {
	fmt.Println("SignIn")
	c.Status(http.StatusOK)
}

type SignUpResponse struct {
	Succes bool `json:"succes"`
}

func (h *Handler) SignUp(c *gin.Context) {
	fmt.Println("SignUp")
	c.JSON(http.StatusOK, &SignUpResponse{
		Succes: true,
	})
}
