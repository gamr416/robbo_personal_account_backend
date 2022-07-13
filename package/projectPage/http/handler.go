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
		projectPage.GET("/:projectPageId", h.GetProjectPageById)
		projectPage.GET("/", h.GetAllProjectPageByUserId)
		projectPage.PUT("/", h.UpdateProjectPage)
		projectPage.DELETE("/:projectPageId", h.DeleteProjectPage)
	}
}

type createProjectPageResponse struct {
	ProjectId string `json:"projectId"`
}

func (h *Handler) CreateProjectPage(c *gin.Context) {
	fmt.Println("CreateProjectPage")

	userId, _, userIdentityErr := h.userIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	projectId, err := h.projectPageDelegate.CreateProjectPage(userId)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, createProjectPageResponse{
		projectId,
	})
}

type getProjectPageResponse struct {
	ProjectPage *models.ProjectPageHTTP `json:"projectPage"`
}

func (h *Handler) GetProjectPageById(c *gin.Context) {
	fmt.Println("Get Project Page By ID")
	projectId := c.Param("projectPageId")
	project_page, err := h.projectPageDelegate.GetProjectPageById(projectId)
	switch err {
	case projectPage.ErrPageNotFound:
		c.AbortWithStatus(http.StatusNotFound)
		return
	case projectPage.ErrInternalServerLevel:
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	case projectPage.ErrBadRequest:
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, getProjectPageResponse{
		&project_page,
	})
}

type getAllProjectPageResponse struct {
	ProjectPages []*models.ProjectPageHTTP `json:"projectPages"`
}

func (h *Handler) GetAllProjectPageByUserId(c *gin.Context) {
	fmt.Println("GetAllProjectPageByUserId")

	userId, _, userIdentityErr := h.userIdentity(c)
	if userIdentityErr != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	projectPages, err := h.projectPageDelegate.GetAllProjectPages(userId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, getAllProjectPageResponse{
		ProjectPages: projectPages,
	})
}

type updateProjectPageInput struct {
	ProjectPage *models.ProjectPageHTTP `json:"projectPage"`
}

func (h *Handler) UpdateProjectPage(c *gin.Context) {
	fmt.Println("Update Project Page")
	inp := new(updateProjectPageInput)
	if err := c.BindJSON(&inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err := h.projectPageDelegate.UpdateProjectPage(inp.ProjectPage)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

func (h *Handler) DeleteProjectPage(c *gin.Context) {
	fmt.Println("Delete Project Page")

	projectId := c.Param("projectPageId")

	err := h.projectPageDelegate.DeleteProjectPage(projectId)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}
