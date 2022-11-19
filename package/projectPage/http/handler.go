package http

import (
	"github.com/gin-gonic/gin"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/auth"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/models"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projectPage"
	"github.com/skinnykaen/robbo_student_personal_account.git/package/projects"
	"log"
	"net/http"
)

type Handler struct {
	authDelegate        auth.Delegate
	projectsDelegate    projects.Delegate
	projectPageDelegate projectPage.Delegate
}

func NewProjectPageHandler(
	authDelegate auth.Delegate,
	projectsDelegate projects.Delegate,
	projectPageDelegate projectPage.Delegate,
) Handler {
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
		projectPage.DELETE("/:projectId", h.DeleteProjectPage)
	}
}

type createProjectPageResponse struct {
	ProjectId string `json:"projectId"`
}

func (h *Handler) CreateProjectPage(c *gin.Context) {
	log.Println("Create Project Page")
	userId, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.Student, models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}

	projectId, err := h.projectPageDelegate.CreateProjectPage(userId)

	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
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
	log.Println("Get Project Page By ID")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.Student, models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	projectPageId := c.Param("projectPageId")
	projectPage, err := h.projectPageDelegate.GetProjectPageById(projectPageId)
	if err != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}

	c.JSON(http.StatusOK, getProjectPageResponse{
		&projectPage,
	})
}

type getAllProjectPageResponse struct {
	ProjectPages []*models.ProjectPageHTTP `json:"projectPages"`
}

func (h *Handler) GetAllProjectPageByUserId(c *gin.Context) {
	log.Println("Get All Project Page By User ID")
	userId, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.Student, models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}

	projectPages, err := h.projectPageDelegate.GetAllProjectPagesByUserId(userId)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	c.JSON(http.StatusOK, getAllProjectPageResponse{
		ProjectPages: projectPages,
	})
}

type updateProjectPageInput struct {
	ProjectPage *models.ProjectPageHTTP `json:"projectPage"`
}

func (h *Handler) UpdateProjectPage(c *gin.Context) {
	log.Println("Update Project Page")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.Student, models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}

	inp := new(updateProjectPageInput)
	if err := c.BindJSON(&inp); err != nil {
		err = projectPage.ErrBadRequestBody
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	log.Println(inp)
	err := h.projectPageDelegate.UpdateProjectPage(inp.ProjectPage)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) DeleteProjectPage(c *gin.Context) {
	log.Println("Delete Project Page")
	_, role, userIdentityErr := h.authDelegate.UserIdentity(c)
	if userIdentityErr != nil {
		log.Println(userIdentityErr)
		ErrorHandling(userIdentityErr, c)
		return
	}
	allowedRoles := []models.Role{models.Student, models.UnitAdmin, models.SuperAdmin}
	accessErr := h.authDelegate.UserAccess(role, allowedRoles)
	if accessErr != nil {
		log.Println(accessErr)
		ErrorHandling(accessErr, c)
		return
	}
	projectId := c.Param("projectId")

	err := h.projectPageDelegate.DeleteProjectPage(projectId)
	if err != nil {
		log.Println(err)
		ErrorHandling(err, c)
		return
	}
	c.Status(http.StatusOK)
}

func ErrorHandling(err error, c *gin.Context) {
	switch err {
	case projectPage.ErrBadRequest:
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	case projectPage.ErrInternalServerLevel:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	case projectPage.ErrPageNotFound:
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
	case projectPage.ErrBadRequestBody:
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
	case auth.ErrInvalidAccessToken:
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
	case auth.ErrTokenNotFound:
		c.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
	case auth.ErrNotAccess:
		c.AbortWithStatusJSON(http.StatusForbidden, err.Error())
	default:
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
}
