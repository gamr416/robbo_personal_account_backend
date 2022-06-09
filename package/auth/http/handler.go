package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"io/ioutil"
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
		auth.POST("/test", h.Test)
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
		HttpOnly: true,
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
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, signInResponse{
		AccessToken: accessToken,
	})
}

func (h *Handler) Refresh(c *gin.Context) {
	fmt.Println("Refresh")
	tokenStr, err := c.Cookie("token")

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
		HttpOnly: true,
	})

	c.JSON(http.StatusOK, signInResponse{
		AccessToken: newAccessToken,
	})
}

func (h *Handler) SignOut(c *gin.Context) {
	fmt.Println("SignOut")

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "token",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
	})
}

type testResponse struct {
	Id uint `json:"id"`
}

type testRequest struct {
	Body string `json:"body"`
}

func (h *Handler) Test(c *gin.Context) {
	//inp := new(testRequest)
	jsonDataBytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println(string(jsonDataBytes))

	//if err := c.BindJSON(inp); err != nil {
	//	fmt.Println(err)
	//	c.AbortWithStatus(http.StatusBadRequest)
	//	return
	//}
	//
	//fmt.Println(inp)

	c.JSON(http.StatusOK, testResponse{
		Id: '1',
	})
}
