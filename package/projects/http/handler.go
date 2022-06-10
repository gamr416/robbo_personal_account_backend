package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projects"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	authDelegate     auth.Delegate
	projectsDelegate projects.Delegate
}

func NewProjectsHandler(authDelegate auth.Delegate, projectsDelegate projects.Delegate) Handler {
	return Handler{
		authDelegate:     authDelegate,
		projectsDelegate: projectsDelegate,
	}
}

func (h *Handler) InitProjectRoutes(router *gin.Engine) {
	project := router.Group("/project")
	{
		project.POST("/", h.CreateProject)
		project.GET("/:projectId", h.GetProject)
		project.PUT("/:projectId", h.UpdateProject)
		project.DELETE("/", h.DeleteProject)
	}
}

type testResponse struct {
	Id string `json:"id"`
}

type testRequest struct {
	Body string `json:"body"`
}

func (h *Handler) CreateProject(c *gin.Context) {
	fmt.Println("Create Project")
	jsonDataBytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	projectHTTP := models.ProjectHTTP{}
	projectHTTP.Json = string(jsonDataBytes)

	projectId, err := h.projectsDelegate.CreateProject(&projectHTTP)
	fmt.Println(projectId)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, testResponse{
		Id: projectId,
	})
}

func (h *Handler) GetProject(c *gin.Context) {

}

func (h *Handler) UpdateProject(c *gin.Context) {
	fmt.Println("Update Project")
	jsonDataBytes, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	projectId := c.Param("projectId")

	projectHTTP := models.ProjectHTTP{}
	projectHTTP.ID = projectId
	projectHTTP.Json = string(jsonDataBytes)

	err = h.projectsDelegate.UpdateProject(&projectHTTP)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, testResponse{
		Id: "1",
	})
}

func (h *Handler) DeleteProject(c *gin.Context) {

}
