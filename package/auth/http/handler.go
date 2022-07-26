package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
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
		auth.GET("/check-auth", h.CheckAuth)
	}
}

type signInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     uint   `json:"role"`
}

type signInResponse struct {
	AccessToken string `json:"accessToken"`
}

func (h *Handler) SignIn(c *gin.Context) {
	fmt.Println("SignIn")

	userHttp := &models.UserHttp{}

	if err := c.BindJSON(userHttp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, err := h.delegate.SignIn(userHttp)
	if err != nil {
		ErrorHandling(err, c)
		return
	}

	cookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		MaxAge:   60 * 60 * 24 * 7,
		HttpOnly: true,
	}

	http.SetCookie(c.Writer, cookie)

	c.JSON(http.StatusOK, signInResponse{
		AccessToken: accessToken,
	})
}

func (h *Handler) SignUp(c *gin.Context) {
	fmt.Println("SignUp")

	userHttp := &models.UserHttp{}

	if err := c.BindJSON(userHttp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, err := h.delegate.SignUp(userHttp)
	if err != nil {
		ErrorHandling(err, c)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    refreshToken,
		MaxAge:   60 * 60 * 24 * 7,
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, signInResponse{
		AccessToken: accessToken,
	})
}

func (h *Handler) Refresh(c *gin.Context) {
	fmt.Println("Refresh")

	tokenStr, err := c.Cookie("refreshToken")

	if err != nil {
		if err == http.ErrNoCookie {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	newAccessToken, err := h.delegate.RefreshToken(tokenStr)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, signInResponse{
		AccessToken: newAccessToken,
	})
}

func (h *Handler) SignOut(c *gin.Context) {
	fmt.Println("SignOut")

	cookie := &http.Cookie{
		Name:     "refreshToken",
		Value:    "",
		MaxAge:   0,
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)

	c.Status(http.StatusOK)
}

type userIdentity struct {
	Id   string `json:"id"`
	Role uint   `json:"role"`
}

func (h *Handler) CheckAuth(c *gin.Context) {
	fmt.Println("CheckAuth")
	userId, role, err := h.userIdentity(c)
	if err != nil {
		ErrorHandling(err, c)
		return
	}
	c.JSON(http.StatusOK, &userIdentity{
		userId,
		uint(role),
	})
}

func ErrorHandling(err error, c *gin.Context) {
	switch err {
	case auth.ErrUserAlreadyExist:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	case auth.ErrInvalidAccessToken:
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	case auth.ErrInvalidTypeClaims:
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	case auth.ErrUserNotFound:
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	case auth.ErrTokenNotFound:
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
		return
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}
