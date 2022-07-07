package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projects"
	"net/http"
)

type Handler struct {
	authDelegate        auth.Delegate
	projectsDelegate    projects.Delegate
	projectPageDelegate projectPage.Delegate
}

type updateProjectPageInput struct {
	ProjectPage models.ProjectPageHTTP `json:"project_page"`
}

func NewProjectPageHandler(authDelegate auth.Delegate, projectsDelegate projects.Delegate, projectPageDelegate projectPage.Delegate) Handler {
	return Handler{
		authDelegate:        authDelegate,
		projectsDelegate:    projectsDelegate,
		projectPageDelegate: projectPageDelegate,
	}
}

func (h *Handler) InitProjectRoutes(router *gin.Engine) {
	projectPage := router.Group("/projectPage")
	{
		projectPage.POST("/", h.CreateProjectPage)
		projectPage.GET("/:projectPageId", h.GetProjectPageByID)
    projectPage.GET("/", h.GetAllProjectPageByUserId)
		projectPage.PUT("/", h.UpdateProjectPage)
		projectPage.DELETE("/:projectPageID", h.DeleteProjectPage)
	}
}

type createProjectPageResponse struct {
	ProjectId string `json:"projectId"`
}

func (h *Handler) CreateProjectPage(c *gin.Context) {
	fmt.Println("CreateProjectPage")
	userId := h.userIdentity(c)

	projectId, err := h.projectPageDelegate.CreateProjectPage(userId)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, createProjectPageResponse{
		projectId,
	})
}

func (h *Handler) GetProjectPageById(c *gin.Context) {

}


func (h *Handler) GetAllProjectPageByUserId(c *gin.Context) {

}

func (h *Handler) UpdateProjectPage(c *gin.Context) {
	fmt.Println("Update Project Page")
	inp := new(updateInput)
	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
}

func (h *Handler) DeleteProjectPage(c *gin.Context) {
	fmt.Println("Delete Project Page")

	projectId := c.Param("projectId")

	projectHTTP := models.ProjectHTTP{}
	projectHTTP.ID = projectId

}
