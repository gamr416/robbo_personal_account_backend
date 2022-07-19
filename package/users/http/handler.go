package http

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/edxApi"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/users"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	authDelegate  auth.Delegate
	usersDelegate users.Delegate
}

func NewUsersHandler(authDelegate auth.Delegate, usersDelegate users.Delegate) Handler {
	return Handler{
		authDelegate:  authDelegate,
		usersDelegate: usersDelegate,
	}
}

type getUsersByRoleResponse struct {
	Users *[]models.UserHttp `json:"users"`
}

type getUserByIdResponse struct {
	Users models.UserHttp `json:"user"`
}

type loginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) InitUsersRoutes(router *gin.Engine) {
	users := router.Group("/user")
	{
		users.GET("/getUsersByRole/:role", h.GetUsersByRole)
		users.POST("/postEnrollment", h.PostEnrollment)
		users.POST("/postUnenroll", h.PostUnenroll)
		users.PUT("/updateUser", h.UpdateUser)
		users.DELETE("/deleteUser/:userId", h.DeleteUser)
		users.POST("/registrationUser", h.RegistrationUser)
		users.POST("/loginUser", h.LoginUser)
		users.POST("/createUser", h.CreateUser)
		users.GET("/getUserById/:userId", h.GetUsersById)
	}
}

func (h *Handler) GetUsersById(c *gin.Context) {
	fmt.Println("Get Users By Id")
	param := c.Param("userId")
	userId, _ := strconv.Atoi(param)

	user, err := h.usersDelegate.GetUserById(userId)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getUserByIdResponse{
		user,
	})
}

func (h *Handler) CreateUser(c *gin.Context) {
	fmt.Println("Create User")

	userHttp := &models.UserHttp{}

	if err := c.BindJSON(userHttp); err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	userId, err := h.usersDelegate.CreateUser(userHttp)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": userId,
	})

}

func (h *Handler) LoginUser(c *gin.Context) {
	fmt.Println("Login User")
	user := loginUser{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &user)
	fmt.Println(user)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = h.usersDelegate.LoginUser(user.Email, user.Password)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) RegistrationUser(c *gin.Context) {
	fmt.Println("Registration User")
	userForm := edxApi.RegistrationForm{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &userForm)
	fmt.Println(userForm)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = h.usersDelegate.RegistrationUser(&userForm)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetUsersByRole(c *gin.Context) {
	fmt.Println("Get Users By Role")
	roleParam := c.Param("role")
	roleId, _ := strconv.Atoi(roleParam)
	role := models.Role(roleId)
	users, err := h.usersDelegate.GetUsersByRole(&role)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getUsersByRoleResponse{
		users,
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	fmt.Println("Update User")
	userHTTP := models.UserHttp{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &userHTTP)
	fmt.Println(userHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = h.usersDelegate.UpdateUser(&userHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)

}

func (h *Handler) PostEnrollment(c *gin.Context) {
	fmt.Println("Post Enrollment")

	postEnrollmentHTTP := models.PostEnrollmentHTTP{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &postEnrollmentHTTP)
	fmt.Println(postEnrollmentHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = h.usersDelegate.PostEnrollment(&postEnrollmentHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) PostUnenroll(c *gin.Context) {
	fmt.Println("Post Unenroll")

	postUnenrollHTTP := models.PostEnrollmentHTTP{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &postUnenrollHTTP)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = h.usersDelegate.PostUnenroll(&postUnenrollHTTP)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	fmt.Println("Delete User")

	userId := c.Param("userId")
	id, _ := strconv.Atoi(userId)
	err := h.usersDelegate.DeleteUser(id)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
}
