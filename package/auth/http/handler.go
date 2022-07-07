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
		auth.GET("/refresh", h.Refresh)
		auth.POST("/sign-out", h.SignOut)
	}
}

type signInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type signInResponse struct {
	AccessToken string `json:"accessToken"`
}

func (h *Handler) SignIn(c *gin.Context) {
	inp := new(signInput)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, err := h.delegate.SignIn(inp.Email, inp.Password)
	if err != nil {
		if err == auth.ErrUserNotFound {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    refreshToken,
		MaxAge:   30 * 24 * 60 * 60,
		HttpOnly: false,
	})
	c.JSON(http.StatusOK, signInResponse{
		AccessToken: accessToken,
	})
}

func (h *Handler) SignUp(c *gin.Context) {
	fmt.Println("SignUp")
	inp := new(signInput)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, err := h.delegate.SignUp(inp.Email, inp.Password)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    refreshToken,
		MaxAge:   30 * 24 * 60 * 60,
		HttpOnly: false,
	})

	c.JSON(http.StatusOK, signInResponse{
		AccessToken: accessToken,
	})
}

func (h *Handler) Refresh(c *gin.Context) {
	fmt.Println("Refresh")
	tokenStr, err := c.Cookie("token")

	fmt.Println(tokenStr)
	if err != nil {
		if err == http.ErrNoCookie {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	newAccessToken, newRefreshToken, err := h.delegate.RefreshToken(tokenStr)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    newRefreshToken,
		MaxAge:   30 * 24 * 60 * 60,
		HttpOnly: false,
	})

	c.JSON(http.StatusOK, signInResponse{
		AccessToken: newAccessToken,
	})
}

func (h *Handler) SignOut(c *gin.Context) {
	fmt.Println("SignOut")

	c.SetCookie("token", "", -1, ":3030/", "localhost", false, false)

	tokenStr, _ := c.Cookie("token")
	fmt.Println("выход")
	fmt.Println(tokenStr)

	c.Status(http.StatusOK)
}
