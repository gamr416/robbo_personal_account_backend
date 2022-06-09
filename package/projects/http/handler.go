package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	authDelegate auth.Delegate
}

func NewProjectsHandler(authDelegate auth.Delegate) Handler {
	return Handler{
		authDelegate: authDelegate,
	}
}

func (h *Handler) InitProjectRoutes(router *gin.Engine) {
	project := router.Group("/project")
	{
		project.POST("/", h.CreateProject)
		project.GET("/", h.GetProject)
		project.PUT("/", h.UpdateProject)
		project.DELETE("/", h.DeleteProject)
	}
}

type testResponse struct {
	Id uint `json:"id"`
}

type testRequest struct {
	Body string `json:"body"`
}

func (h *Handler) CreateProject(c *gin.Context) {

}

func (h *Handler) GetProject(c *gin.Context) {

}

func (h *Handler) UpdateProject(c *gin.Context) {
	jsonDataBytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fmt.Println(string(jsonDataBytes))

	c.JSON(http.StatusOK, testResponse{
		Id: '1',
	})
}

func (h *Handler) DeleteProject(c *gin.Context) {

}
